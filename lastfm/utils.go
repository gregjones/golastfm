package lastfm

import (
	"bytes"
	"net/url"
)

func urlencode(data map[string]string) string {
	var buf bytes.Buffer
	for k, v := range data {
		buf.WriteString(url.QueryEscape(k))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(v))
		buf.WriteByte('&')
	}
	s := buf.String()
	return s[0 : len(s)-1]
}

func boolToString(b bool) string {
	var boolString = "0"
	if b {
		boolString = "1"
	}
	return boolString
}
