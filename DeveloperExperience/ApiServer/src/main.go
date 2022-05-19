package main

import (
	"apiserver/handler"
)

func run() {
	handler.Start("127.0.0.1:8080")
}

func init() {
	handler.InitConfig()
	handler.InitRedisClient()
}

func main() {
	run()
}
