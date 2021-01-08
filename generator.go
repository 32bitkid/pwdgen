package main

import (
	"math/rand"
	"strings"
)

type CharSet struct {
	Set   string
	Count int
}

type Generator struct {
	Length       int
	DefaultSet   string
	RequiredSets []CharSet
	Source       rand.Source

	rng *rand.Rand
}

func (gen Generator) Generate() (string, error) {
	var sets []string

	// Generate required character sets
	for _, group := range gen.RequiredSets {
		c := group.Count
		if c == 0 {
			c = 1
		}
		for i := 0; i < c; i++ {
			sets = append(sets, group.Set)
		}
	}

	// Determine the default character set
	defaultChars := gen.DefaultSet
	if len(defaultChars) == 0 {
		defaultChars = gen.uniqueRequiredChars()
	}

	// Fill until the desired length
	for len(sets) < gen.Length {
		sets = append(sets, defaultChars)
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

func (gen Generator) uniqueRequiredChars() string {
	uniq := ""
	found := map[rune]struct{}{}
	for _, g := range gen.RequiredSets {
		for _, r := range g.Set {
			if _, ok := found[r]; !ok {
				uniq = uniq + string(r)
				found[r] = struct{}{}
			}
		}
	}
	return uniq
}
