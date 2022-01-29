package main

import (
	"apiserver/handler"
	"encoding/json"
	"fmt"
)

func run() {
	handler.Start("127.0.0.1:8080")
}

func test() {
	a := handler.GetDeveloperIssueBehaviorRecord()
	ab, _ := json.Marshal(a)
	fmt.Println(string(ab))
}
func init() {
	handler.InitConfig()
}

func main() {
	run()
	// test()
}
