package main

import (
	"fmt"

	"github.com/SavenkoArtem/pshelp/configs"
)

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf)
}
