package main

import (
	"fmt"
	"log"
	"os"
)

func readEnv() string {
	os.Setenv("GOLANG_PROJECT", "初學GOlang")
	project := os.Getenv("GOLANG_PROJECT")
	fmt.Println(project)
	content, err := os.ReadFile(".env")
	if err != nil {
		log.Fatal(err)
	}
	str := string(content)
	return str
}
func main() {

	fmt.Println(readEnv())
}
