package random

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	crand "crypto/rand"
	mrand "math/rand"
)

var once sync.Once

func RandUint64() uint64 {
	once.Do(func() {
		mrand.Seed(time.Now().UnixNano())
		logrus.Debugf("Seeded random with current time!")
	})
	return mrand.Uint64()
}

func RandomString(stringLength int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, stringLength)
	if _, err := crand.Read(bytes); err != nil {
		return "", err
	}
	// Run through bytes; replacing each with the equivalent random char.
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}
