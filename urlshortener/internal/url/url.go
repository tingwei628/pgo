package url

import (
	"crypto/sha256"
	"encoding/hex"
)

func Shorten(url string) string {
	h := sha256.New()
	h.Write([]byte(url))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash[:8]
}
