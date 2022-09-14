package share

import (
	"bytes"
	"sort"
)

func Signature(params map[string]string) []byte {
	var b bytes.Buffer
	keys := make([]string, len(params))
	for k, _ := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		if key == "sign" {
			continue
		}

		val := params[key]
		if key != "" && val != "" {
			b.WriteString(key)
			b.WriteString(val)
		}
	}
	return b.Bytes()
}
