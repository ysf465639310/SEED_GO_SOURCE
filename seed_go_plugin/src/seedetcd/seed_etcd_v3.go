package seedetcd

import (
	"../seedcomdata"
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
)

// cat /dev/null > log.txt

func (base*V3) Set (key string, json string) error {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   base.EPath,
		DialTimeout: base.DialTimeout,
	})
	if err != nil {
		//log.Fatal(err)
		base.Log.GetLogHandle().WithFields(logrus.Fields{"error":err}).Error("Fail connect ET server...")
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), base.RequestTimeout)

	_, err = cli.Put(ctx, key, json)
	cancel()
	if err != nil {
		//log.Fatal(err)
		base.Log.GetLogHandle().WithFields(logrus.Fields{"error":err}).Error("Fail PUT to ET server...")
		return err
	}
	return nil
}

func (base*V3) Gets (key string) ([] seedcomdata.SeedEtcdResp, error)  {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   base.EPath,
		DialTimeout: base.DialTimeout,
	})
	if err != nil {
		//log.Fatal(err)
		base.Log.GetLogHandle().WithFields(logrus.Fields{"error":err}).Error("Fail connect ET server...")
		return nil,err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), base.RequestTimeout)

	resp, err := cli.Get(ctx, key)
	cancel()

	if err != nil {
		//log.Fatal(err)
		base.Log.GetLogHandle().WithFields(logrus.Fields{"error":err}).Error("Fail GETS to ET server...")
		return nil ,err
	}

	var items [] seedcomdata.SeedEtcdResp

	for _, ev := range resp.Kvs {
		var item = seedcomdata.SeedEtcdResp { Key:string(ev.Key) , Value:string(ev.Value) }
		item.Value=string(ev.Value)
		item.Key = string(ev.Key)
		items = append(items,item)
		//fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
	return items, nil
}

func (base*V3) Del (key string) error {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   base.EPath,
		DialTimeout: base.DialTimeout,
	})

	if err != nil {
		//log.Fatal(err)
		base.Log.GetLogHandle().WithFields(logrus.Fields{"error":err}).Error("Fail connect ET server...")
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), base.RequestTimeout)
	defer cancel()

	// count keys about to be deleted
	gresp, err := cli.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		//log.Fatal(err)
		base.Log.GetLogHandle().WithFields(logrus.Fields{"error":err}).Error("Fail GETS to ET server...")
		return err
	}

	// delete the keys
	resp, err := cli.Delete(ctx, key, clientv3.WithPrefix())
	if err != nil {
		//log.Fatal(err)
		base.Log.GetLogHandle().WithFields(logrus.Fields{"error":err}).Error("Fail DELETE to ET server...")
		return err
	}

	//fmt.Println("Deleted all keys:", int64(len(gresp.Kvs)) == resp.Deleted)
	base.Log.GetLogHandle().WithFields(logrus.Fields{"Kvs":int64(len(gresp.Kvs)),
		"Delete":resp.Deleted}).Error("Deleted all keys ...")

	return nil
}

