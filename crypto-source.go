package main

import (
	"crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

type cryptoSource struct{}

func (s *cryptoSource) Seed(seed int64) { /*no-op*/ }
func (s *cryptoSource) Uint64() (value uint64) {
	err := binary.Read(rand.Reader, binary.BigEndian, &value)
	if err != nil {
		panic(err)
	}
	return
}
func (s *cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func newCryptoSeededSource() mrand.Source {
	var seed int64
	binary.Read(rand.Reader, binary.BigEndian, &seed)
	return mrand.NewSource(seed)
}
