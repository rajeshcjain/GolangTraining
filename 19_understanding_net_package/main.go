package main

import (
	"net"
	"fmt"
)

func main(){
	inter,_ := net.Interfaces()

	for _, val := range inter{
		fmt.Println(val.Name,"->",val.HardwareAddr.String())
		fmt.Println(val.Addrs())
	}
}
