package shorter

import (
	"crypto/sha256"
	"hash/fnv"
	"math/rand"
)

const base63 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func getHash(original string) (int64, error) {
	shaHash := sha256.Sum256([]byte(original))
	fnvHash := fnv.New64a()
	_, err := fnvHash.Write(shaHash[:])
	if err != nil {
		return 0, err
	}
	return int64(fnvHash.Sum64()), nil
}
func RandString() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = base63[rand.Intn(len(base63))]
	}
	return string(b)
}

// GetShort - пока заглушка
func GetShort(original string) (string, error) {
	hash, err := getHash(original)
	if err != nil {
		return "", err
	}
	rand.Seed(hash)

	return RandString(), nil
}
