package chain

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "0123456789abcdef"

const (
	bits = 6           // 6 bits to represent a letter index
	mask = 1<<bits - 1 // All 1-bits, as many as bits
	max  = 63 / bits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randomHash(len int) string {
	return "0x" + randHexString(len)
}

func randHexString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for max characters!
	for i, cache, remain := n-1, src.Int63(), max; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), max
		}
		if idx := int(cache & mask); idx < len(charset) {
			sb.WriteByte(charset[idx])
			i--
		}
		cache >>= bits
		remain--
	}

	return sb.String()
}
