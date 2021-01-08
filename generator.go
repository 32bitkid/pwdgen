package main

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"strings"
)

type Generator struct {
	Length       int
	DefaultSet   string
	RequiredSets map[string]int
	Random       io.Reader
}

func (gen Generator) Generate() (string, error) {
	var sets []string

	// Generate required character sets
	for set, count := range gen.RequiredSets {
		for i := 0; i < count; i++ {
			sets = append(sets, set)
		}
	}

	// Determine the default character set
	defaultChars := gen.DefaultSet
	if len(defaultChars) == 0 {
		chars := ""
		for set, _ := range gen.RequiredSets {
			chars = chars + set
		}
		defaultChars = uniqueChars(chars)
	}

	// Fill until the desired length
	for len(sets) < gen.Length {
		sets = append(sets, defaultChars)
	}

	// Shuffle
	for i := range sets {
		j, err := gen.next(i + 1)
		if err != nil {
			return "", err
		}
		sets[i], sets[j] = sets[j], sets[i]
	}

	var pwd []string

	// Resolve
	for _, set := range sets {
		chars := strings.Split(set, "")
		idx, err := gen.next(len(chars))
		if err != nil {
			return "", err
		}
		pwd = append(pwd, chars[idx])
	}

	return strings.Join(pwd, ""), nil
}

func (gen Generator) MustGenerate() string {
	pwd, err := gen.Generate()
	if err != nil {
		panic(err)
	}
	return pwd
}

const maxRand = 65536

func (gen Generator) next(n int) (int, error) {
	reader := gen.Random
	if reader == nil {
		reader = rand.Reader
	}

	var next uint16
	var output int
	remainder := maxRand % n

	for {
		err := binary.Read(reader, binary.LittleEndian, &next)
		if err != nil {
			return 0, err
		}

		x := int(next)
		output = x % n
		if x < maxRand-remainder {
			return output, nil
		}
	}
}

func uniqueChars(input string) string {
	set := map[rune]struct{}{}
	for _, c := range input {
		set[c] = struct{}{}
	}
	uniq := ""
	for r := range set {
		uniq = uniq + string(r)
	}
	return uniq
}
