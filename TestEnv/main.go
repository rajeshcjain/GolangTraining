package main

import (
	"github.com/alecthomas/kingpin"
	"fmt"
)

var(
	check = kingpin.Command("check","check the network")
	appServerIP = check.Arg("appserverIP","enter the AppServer IP").Required().String()
        concentratorIP = check.Arg("concentratorIP","enter the Concentrator IP").Required().String()
	egmIP = check.Arg("egmIP","enter the EGM IP").Required().String()
)

func main(){
	switch kingpin.Parse(){

	case "check":
		fmt.Println(appServerIP , concentratorIP, egmIP)

	}
}
