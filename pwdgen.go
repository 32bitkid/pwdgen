package main

import (
	"fmt"
)

func main() {
	gen := Generator{
		Length:     16,
		DefaultSet: Base58,
		RequiredSets: map[string]int{
			Base58Upper: 1,
			Base58Lower: 1,
			Base58Num:   1,
			WordSymbols: 1,
		},
	}

	fmt.Println(gen.MustGenerate())
}
