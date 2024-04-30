package main

import (
	"flag"
	"fmt"

	"go.jetify.com/typeid"
)

func main() {
	var prefix string
	flag.StringVar(&prefix, "p", "", "Prefix")
	flag.Parse()

	var tid = typeid.Must(typeid.WithPrefix(prefix))
	fmt.Println(tid)
}
