package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello Go")
	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		fmt.Println("Hello", name)
	} else {
		log.Fatal(err)
	}
}
