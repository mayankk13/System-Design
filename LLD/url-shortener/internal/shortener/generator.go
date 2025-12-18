package shortener

import (
	"crypto/rand"
	"math/big"
)

const (
	NumCharsShortLink = 7
	Alphabet          = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func generateRandomShortURL() string {
	result := make([]byte, NumCharsShortLink)

	for {
		for i := 0; i < NumCharsShortLink; i++ {
			n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(Alphabet))))
			result[i] = Alphabet[n.Int64()]
		}

		short := string(result)

		if !shortLinkExists(short) {
			return short
		}
	}
}
