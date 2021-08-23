package singlee_sdk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func GetSign(data map[string]string, key string) string {
	keys := make([]string, 0, len(data))
	for k := range data {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		if data[k] == "" {
			continue
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(fmt.Sprintf("%v", data[k]))
		buf.WriteByte('&')
	}
	buf.WriteString("key=" + key)

	hash := hmac.New(sha256.New, []byte(key))
	hash.Write(buf.Bytes())
	ret := hash.Sum(nil)

	return strings.ToUpper(hex.EncodeToString(ret[:]))
}
