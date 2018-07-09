package main

import (
	"fmt"
	"io/ioutil"

	"github.com/philoticnet/go-multiple-apps/cmd/app1/utils"
)

func main() {
	fmt.Printf("start app2")

	b, err := ioutil.ReadFile("app1.txt")
	if err != nil {
		fmt.Print(err)
	}

	utils.WriteFile("app2.txt", string(b))
}
