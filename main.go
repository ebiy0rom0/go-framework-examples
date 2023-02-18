package main

import (
	"examples/framework"
	"examples/infra"
	"flag"
	"fmt"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	var fw framework.Framework
	flag.StringVar((*string)(&fw), "framework", "", "exec framework")
	flag.StringVar((*string)(&fw), "f", "gin", "exec framework")

	if err := infra.InitializeDb("db/example.db"); err != nil {
		return err
	}
	return framework.Start(fw)
}
