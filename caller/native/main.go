package main

import (
	"fmt"
	"github.io/decision2016/go-plugins-example/interfaces"
	"plugin"
)

func main() {
	p, err := plugin.Open("hex.so")
	if err != nil {
		panic(err)
	}

	symbol, err := p.Lookup("Converter")
	if err != nil {
		panic(err)
	}

	c := symbol.(interfaces.NativeConverter)
	hexString := c.Run("test plugin")
	fmt.Println(hexString)
}
