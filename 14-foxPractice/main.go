package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readEnv() string {
	content, err := ioutil.ReadFile(".env")
	if err != nil {
		log.Fatal(err)
	}
	str := string(content)
	return str
}
func main() {

	fmt.Println(readEnv())
}
