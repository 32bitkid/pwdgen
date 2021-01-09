package main

import (
	"math/rand"
	"strings"
)

type CharSet struct {
	Chars string
	Min   int
	Max   int
}

type Generator struct {
	Length       int
	DefaultChars string
	RequiredSets []CharSet
	Source       rand.Source

	rng *rand.Rand
}

func (gen Generator) Generate() (string, error) {
	var sets []string

	// Generate required character sets
	for _, group := range gen.RequiredSets {
		for i := 0; i < group.Min; i++ {
			sets = append(sets, group.Chars)
		}
	}

	// Fill until the desired length
	for i := 0; len(sets) < gen.Length; i++ {
		chars := gen.DefaultChars

		if len(chars) == 0 {
			uniq := ""
			found := map[rune]struct{}{}
			for _, g := range gen.RequiredSets {
				if g.Min + i >= g.Max && g.Max >= 0 {
					continue
				}

				for _, r := range g.Chars {
					if _, ok := found[r]; ok {
						continue
					}

					uniq = uniq + string(r)
					found[r] = struct{}{}
				}
			}
			if len(uniq) == 0 {
				panic("out of available characters! check length and character-max settings!")
			}

			chars = uniq
		}

		sets = append(sets, chars)
	}

	// Shuffle
	for i := range sets {
		j := gen.next(i + 1)
		sets[i], sets[j] = sets[j], sets[i]
	}

	var pwd []string

	// Resolve
	for _, set := range sets {
		chars := strings.Split(set, "")
		idx := gen.next(len(chars))
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

func (gen Generator) next(n int) int {
	if gen.rng != nil {
		return gen.rng.Intn(n)
	}

	src := gen.Source
	if src == nil {
		src = &cryptoSource{}
	}

	r := rand.New(src)
	gen.rng = r

	return r.Intn(n)
}
