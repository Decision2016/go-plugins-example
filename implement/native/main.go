package main

import "encoding/hex"

type HexConverter struct{}

func (c *HexConverter) Run(s string) interface{} {
	return hex.EncodeToString([]byte(s))
}

var Converter HexConverter
