package util

import (
	"strings"
	"fmt"
)

func FromIso8859_1toUtf8(iso8859_1_buf []byte) string {
	buf := make([]rune, len(iso8859_1_buf))
	for i, b := range iso8859_1_buf {
		buf[i] = rune(b)
	}
	return string(buf)
}

func Empty(value string) bool {
	return "" == strings.TrimSpace(value)
}

func FromGeneric(generic interface{}) string {
	to := fmt.Sprintf("%v", generic)

	return strings.TrimSpace(to)
}