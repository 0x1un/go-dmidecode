package slot_test

import (
	"testing"

	"github.com/0x1un/go-dmidecode/parser/slot"
	"github.com/0x1un/go-dmidecode/smbios"

	"github.com/stretchr/testify/assert"
)

var (
	s = &smbios.Structure{
		Header: smbios.Header{
			Type:   9,
			Length: 17,
			Handle: 2304,
		},
		Formatted: []byte{0x1, 0xb1, 0xd, 0x3, 0x4, 0x1, 0x0, 0x4, 0x1, 0xff, 0xff, 0xff, 0xff},
		Strings:   []string{"PCIe Slot 1"},
	}
)

func TestParse(t *testing.T) {
	bios, err := slot.Parse(s)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}
