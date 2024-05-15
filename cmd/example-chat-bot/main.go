package main

import (
	"fmt"

	"github.com/sarthakjdev/wapi.go/pkg/client"
)

func main() {
	fmt.Println("This is the main package entry point of my golang file")
	whatsappClient := client.NewClient()
	fmt.Print(whatsappClient)
}