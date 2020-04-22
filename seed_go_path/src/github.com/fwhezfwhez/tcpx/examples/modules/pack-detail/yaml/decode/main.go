package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

func main() {
	var buf = []byte{
		0, 0, 0, 42, 0, 0, 0, 1, 0, 0, 0, 14, 0, 0, 0, 16, 123, 34, 97, 117, 116, 104, 34, 58, 34, 97, 98, 99, 34, 125,
		117, 115, 101, 114, 95, 110, 97, 109, 101, 58, 32, 116, 99, 112, 120, 10,
	}
	length := binary.BigEndian.Uint32(buf[0:4])
	messageID := binary.BigEndian.Uint32(buf[4:8])
	headerLength := binary.BigEndian.Uint32(buf[8:12])
	bodyLength := binary.BigEndian.Uint32(buf[12:16])

	var header map[string]interface{}
	var body struct {
		Username string `yaml:"user_name"`
	}

	json.Unmarshal(buf[16:16+headerLength], &header)
	yaml.Unmarshal(buf[16+headerLength:16+headerLength+bodyLength], &body)

	fmt.Println(length)
	fmt.Println(messageID)
	fmt.Println(headerLength)
	fmt.Println(bodyLength)

	fmt.Println(header)
	fmt.Println(body)
}
