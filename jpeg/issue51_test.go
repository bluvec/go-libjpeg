package jpeg

import (
	"bytes"
	"testing"
)

// https://github.com/bluvec/go-libjpeg/issues/51
func TestIssue51(t *testing.T) {
	data := []byte("\xff\xd8\xff\xdb\x00C\x000000000000000" +
		"00000000000000000000" +
		"00000000000000000000" +
		"00000000000\xff\xc9\x00\v\b00\x000" +
		"\x01\x01\x14\x00\xff\xda\x00\b\x01\x010\x00?\x0000")

	Decode(bytes.NewReader(data), &DecoderOptions{})
}
