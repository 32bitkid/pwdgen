package main

import "strings"

const (
	Lower = "abcdefghijklmnopqrstuvwxyz"
	Upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digit = "0123456789"

	Base58Lower = "abcdefghijkmnopqrstuvwxyz"
	Base58Upper = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	Base58Digit = "123456789"

	Symbol     = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	WordSymbol = "/-+\\~_."
)

var (
	Alpha    = strings.Join([]string{Lower, Upper}, "")
	AlphaNum = strings.Join([]string{Alpha, Digit}, "")

	Base58Alpha    = strings.Join([]string{Base58Lower, Base58Upper}, "")
	Base58AlphaNum = strings.Join([]string{Base58Alpha, Base58Digit}, "")
)
