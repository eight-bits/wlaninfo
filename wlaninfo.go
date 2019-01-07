package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	iconv "github.com/djimenez/iconv-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No parametrs")
		fmt.Println("============")
		fmt.Println("all - display all networks with saved parameters")
		fmt.Println("<name> - display the data of the specified network")
		os.Exit(0)
	}
	param := os.Args[1]
	if string(param) == "all" {
		cm := exec.Command("netsh", "wlan", "show", "profiles")
		var buf bytes.Buffer
		cm.Stdout = &buf
		err := cm.Start()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		err = cm.Wait()
		str, errConv := iconv.ConvertString(buf.String(), "866", "UTF-8")
		if errConv != nil {
			fmt.Printf("error: %v\n", errConv)
			os.Exit(1)
		}
		fmt.Println(str)
		os.Exit(0)
	} else {
		namewlan := os.Args[1]
		cm := exec.Command("netsh", "wlan", "show", "profiles", string(namewlan), "key=clear")
		var buf bytes.Buffer
		cm.Stdout = &buf
		err := cm.Start()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		err = cm.Wait()
		str, errConv := iconv.ConvertString(buf.String(), "866", "UTF-8")
		if errConv != nil {
			fmt.Printf("error: %v\n", errConv)
			os.Exit(1)
		}
		fmt.Println(str)
	}
}
