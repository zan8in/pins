package cryptoutil

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexDecode(str string) ([]byte, error) {
	return hex.DecodeString(fmt.Sprintf("%x", str))
}

func Base64Decode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
