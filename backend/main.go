package main

import (
	"fmt"
)

func main(){
	fmt.Println("hello go!")
	loadConfigFile("./config.json")
	server := Server{}
	server.start(Config.Ipaddr, Config.Port)
}