package main

import (
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"math/rand"
	"os"
	"path/filepath"
)

var pwdCount = pflag.IntP("count", "n", 1, "Total generated passwords")
var pwdLength = pflag.IntP("length", "l", 16, "Password length in characters")
var base58 = pflag.Bool("base58", false, "Use base-58 safe characters")

var nDigits = pflag.IntSliceP("digits", "D", nil, "Numeric character limits [min, max]")
var nLowers = pflag.IntSliceP("lowers", "L", nil, "Lower-case character limits [min, max]")
var nUppers = pflag.IntSliceP("uppers", "U", nil, "Upper-case character limits [min, max]")
var nSymbols = pflag.IntSliceP("symbols", "S", []int{1, 2}, "Symbolic character limits [min, max]")

var fast = pflag.Bool("fast", false, "use cytro-seeded PRNG. (not recommended!)")
var wordSafe = pflag.Bool("word-safe", false, "use word-safe symbols")


func init() {
	pflag.Usage = func() {
		cmd := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage: %s [ OPTIONS ]  \n\nOptions supported by %s:\n", cmd, cmd)
		pflag.PrintDefaults()
	}
	pflag.ErrHelp = errors.New("")
	pflag.Parse()
}

func lower(r []int) int {
	if len(r) < 1 {
		return 1
	}
	return r[0]
}

func upper(r []int) int {
	if len(r) < 1 {
		return -1
	}
	return r[len(r)-1]
}

func main() {
	var req []CharSet

	symbolSet := Symbol
	if *wordSafe {
		symbolSet = WordSymbol
	}

	if *base58 {
		req = append(
			req,
			CharSet{Base58Digit, lower(*nDigits), upper(*nDigits)},
			CharSet{Base58Lower, lower(*nLowers), upper(*nLowers)},
			CharSet{Base58Upper, lower(*nUppers), upper(*nUppers)},
			CharSet{symbolSet, lower(*nSymbols), upper(*nSymbols)},
		)
	} else {
		req = append(
			req,
			CharSet{Digit, lower(*nDigits), upper(*nDigits)},
			CharSet{Lower, lower(*nLowers), upper(*nLowers)},
			CharSet{Upper, lower(*nUppers), upper(*nUppers)},
			CharSet{symbolSet, lower(*nSymbols), upper(*nSymbols)},
		)
	}

	var source rand.Source
	if *fast {
		source = newCryptoSeededSource()
	}

	gen := Generator{
		Length:       *pwdLength,
		RequiredSets: req,
		Source:       source,
	}

	count := *pwdCount
	for count > 0 {
		fmt.Println(gen.MustGenerate())
		count--
	}

}
