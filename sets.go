package main

import "strings"

const (
	Lower   = "abcdefghijklmnopqrstuvwxyz"
	Upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits  = "0123456789"
	Symbols = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

	Base58Num   = "123456789"
	Base58Upper = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	Base58Lower = "abcdefghijkmnopqrstuvwxyz"

	WordSymbols = "/-+\\~_."
)

var (
	Alpha    = strings.Join([]string{Lower, Upper}, "")
	AlphaNum = strings.Join([]string{Lower, Upper, Digits}, "")
	Any      = strings.Join([]string{Lower, Upper, Digits, Symbols}, "")

	Base58 = strings.Join([]string{Base58Num, Base58Upper, Base58Lower}, "")
)
