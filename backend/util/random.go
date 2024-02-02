package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomTime(from time.Time, to time.Time) time.Time {
	if to.Before(from) {
		panic("endTime should be after startTime")
	}

	diff := to.Sub(from)
	randomDiff := time.Duration(rand.Int63n(int64(diff)))

	return from.Add(randomDiff)
}
