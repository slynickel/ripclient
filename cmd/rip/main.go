package main

import (
	"fmt"
	"os"

	"github.com/slynickel/ripclient"
)

func main() {
	fmt.Println("Test")
	auth := os.Getenv("SNAUTH")
	if auth == "" {
		fmt.Println("SNAUTH must be set")
		return
	}
	host := os.Getenv("SNHOST")
	if auth == "" {
		fmt.Println("SNHOST must be set")
		return
	}
	c := ripclient.Config{
		Hostname: host,
		Auth:     auth,
	}
	resp, err := c.Post()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", resp)
}
