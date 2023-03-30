package utils

import (
	"os"
	"log"
	"fmt"
	"net"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetNetworkAddress() {
	PORT := os.Getenv("PORT")
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Printf("You can access via the same network with http://%s:%s\n", localAddr.IP, PORT)
}
