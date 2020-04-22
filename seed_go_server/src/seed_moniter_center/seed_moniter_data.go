package main

import "encoding/json"

const ENV_HOST  = "ENV_HOST"

type SeedMonConfig struct {

	Uid  string    `json:"uid"`
	Name string    `json:"name"`
	Type string    `json:"type"`
	Host string    `json:"host"`
	Port  int	   `json:"port"`
	VirHost string `json:"vir_host"`

}

func (base* SeedMonConfig) CreateJson() (string, error){

	data, err := json.Marshal(base)

	return string(data), err
}


func (base* SeedMonConfig) ParseJson(data string) error {
	return  json.Unmarshal([]byte(data), base)
}