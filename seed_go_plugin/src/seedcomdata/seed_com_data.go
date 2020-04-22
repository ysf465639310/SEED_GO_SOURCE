package seedcomdata

type CommInterFace interface {
	GetVersion() string
}

type  ServerRegister struct {
	Name string `json:"Name"`
	Type string `json:"Type"`
	Config string `json:"Config"`
	ListenPort int `json:"ListenPort"`
	ListerPath string `json:"ListerPath"`
}

type SeedEtcdResp struct {
	Key string
	Value string
}

type PROTOCOL_ID int32

const (
	REQ_ONLINE  PROTOCOL_ID =0
	REP_ONLINE  PROTOCOL_ID =1
	REQ_OFFLINE PROTOCOL_ID =2
	REP_OFFLINE PROTOCOL_ID =3

	/*-----------------------*/
	REQ_SENDMS PROTOCOL_ID =1000
	REP_SENDMS PROTOCOL_ID =1001
	REQ_CRHOME PROTOCOL_ID =1002
	REP_CRHOME PROTOCOL_ID =1003
)

type ReqSend struct {

	HomeID   string `json:"home_id"`
	ToUser   string `json:"to_user"`
	FromUser string `json:"from_user"`
	Message  string `json:"message"`
}

type RepSend struct {
	HomeID   string `json:"home_id"`
	ToUser   string `json:"to_user"`
	FromUser string `json:"from_user"`
	Message  string `json:"message"`
	Err  int32
}

type Login struct {
	Username string `json:"username"`
	RoomID string `json:"room_id"`
}

type Rep struct {
	Message  string `json:"message"`
	Err  int32
}







