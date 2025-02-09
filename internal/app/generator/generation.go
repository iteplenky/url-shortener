package generator

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789"

type Generator interface {
	RandomString(length int) string
}
type Generate struct{}

func New() *Generate {
	return &Generate{}
}

func (Generate) RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
