package main

import (
	"flag"
	"fmt"

	"github.com/FlyingRadish/rcong"
)


func main() {
	// 定义命令行参数
	ip := flag.String("ip", "", "server ip")
	port := flag.Int("p", "", "server port")
	pwd := flag.String("pwd", "", "auth password")
	command := flag.String("c", "", "rcon command")
	flag.Parse()
	conn := rcong.NewRCONConnection(ip, port, pwd, 3, 10)
	conn.Connect()
	defer conn.Close()

	response, err := conn.ExecCommand(command)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(response)
}