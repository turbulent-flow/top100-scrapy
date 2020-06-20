package kit

import (
	"time"
	"math/rand"
	"strings"
)

// BuildRandomStrings generates the random string.
// length argument represents the length of the result you expect.
func BuildRandomStrings(length int) string {
	rand.Seed(time.Now().UnixNano())
    chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
                    "abcdefghijklmnopqrstuvwxyz" +
                    "0123456789")
    var b strings.Builder
    for i := 0; i < length; i++ {
        b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
