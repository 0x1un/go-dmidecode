package oem

import "github.com/0x1un/dmidecode/smbios"

// Parse 解析smbios struct数据
func Parse(s *smbios.Structure) (info *OEM, err error) {
	defer smbios.ParseRecovery(s, &err)

	info = &OEM{
		Header:  s.Header,
		Strings: s.Strings,
	}

	return info, nil
}
