package main

import (
	"fmt"

	toJson "encoding/json"

	example "github.com/frank269/protoc-gen-code/example/generated"
)

func main() {
	config := &example.RequestEnv{}
	err := config.FromEnv()
	if err != nil {
		panic(err)
	}
	json, err := toJson.Marshal(config)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))

	r, _ := config.AsProto()
	json, err = r.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))

	r2 := &example.Request{}
	err = r2.UnmarshalJSON(json)
	if err != nil {
		panic(err)
	}
	fmt.Println(r2)

}
