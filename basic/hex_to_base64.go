package basic

import (
	"encoding/base64"
	"encoding/hex"
)

// ByteStrTobase64 convert hex string into base64
func ByteStrTobase64(input string) string {
	decoded, _ := hex.DecodeString(input)
	return base64.StdEncoding.EncodeToString(decoded)
}
