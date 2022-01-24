package ipnet

import (
	"fmt"
	"strconv"
)

type v4IP struct {
	address
}

func (i v4IP) addr() address {
	return i.address
}

func (i v4IP) String() string {
	b := i.address.bytes()
	return fmt.Sprintf("%d.%d.%d.%d", b[12], b[13], b[14], b[15])
}

func (i v4IP) Version() Ver {
	return v4
}

func splitV4(s string) ([]string, bool) {
	l := len(s)
	if l == 0 {
		return []string{}, false
	}
	var o []string
	var b int
	for i := 0; i < l; i++ {
		if s[i] == v4sep {
			if len(s[b:i]) == 0 || len(s[b:i]) > 3 {
				return []string{}, false
			}
			o = append(o, s[b:i])
			b = i + 1
		}
	}
	if len(s[b:]) == 0 || len(s[b:]) > 3 {
		return []string{}, false
	}
	o = append(o, s[b:])
	if len(o) != 4 {
		return []string{}, false
	}
	return o, true
}

func parseV4(s string) ([4]uint64, bool) {
	so, ok := splitV4(s)
	if !ok {
		return [4]uint64{}, false
	}
	return a10toUint64(so)
}

func a10toUint64(s []string) ([4]uint64, bool) {
	var r [4]uint64
	if len(s) != 4 {
		return [4]uint64{}, false
	}
	for idx, v := range s {
		b, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return [4]uint64{}, false
		}
		if b < 0 || b > 255 {
			return [4]uint64{}, false
		}
		r[idx] = b
	}
	return r, true
}

func parseIPv4(addr string) (v4IP, error) {
	if len(addr) < 7 {
		return v4IP{}, fmt.Errorf("%s", Invalid)
	}
	so, ok := parseV4(addr)
	if !ok {
		return v4IP{}, fmt.Errorf("%s", Invalid)
	}

	var ipv4 v4IP
	ipv4.address.set(0, so[0]<<24|(so[1]<<16)|(so[2]<<8)|so[3])

	return ipv4, nil
}
