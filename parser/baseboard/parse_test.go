package baseboard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/0x1un/dmidecode/parser/baseboard"
	"github.com/0x1un/dmidecode/smbios"
)

var (
	s = &smbios.Structure{
		Header: smbios.Header{
			Type:   2,
			Length: 8,
			Handle: 512,
		},
		Formatted: []byte{0x1, 0x2, 0x3, 0x4},
		Strings:   []string{"Dell Inc.", "0CNCJW", "A05", ".HVSRF52.CN7475154A0700."},
	}
)

func TestParse(t *testing.T) {
	should := assert.New(t)

	bios, err := baseboard.Parse(s)
	if assert.NoError(t, err) {
		should.Equal("Dell Inc.", bios.Manufacturer)
		should.Equal("0CNCJW", bios.ProductName)
		should.Equal("A05", bios.Version)
		should.Equal(".HVSRF52.CN7475154A0700.", bios.SerialNumber)
	}
}
