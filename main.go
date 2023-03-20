package main

import (
	"learn_mongo/config"
)

func main() {
	config.Init()
	config.InitMongoClient()

}
