package ipnet

import (
	"fmt"
	"strconv"
)

type v6IP struct {
	address
}

func (i v6IP) addr() address {
	return i.address
}

func (i v6IP) bytes() [8]uint64 {
	b := i.address.bytes()
	var c [8]uint64
	for i := range c {
		x := i * 2
		c[i] = uint64(b[x])<<8 | uint64(b[x+1])
	}
	return c
}

func (i v6IP) String() string {
	b := i.bytes()
	return fmt.Sprintf("%x:%x:%x:%x:%x:%x:%x:%x",
		b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7])
}

func (i v6IP) Version() Ver {
	return v6
}

func splitV6(s string) ([]string, bool) {
	if len(s) == 0 {
		return []string{}, false
	}

	var o []string
	var b int
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == v6sep {
			if len(s[b:i]) == 0 || len(s[b:i]) > 4 {
				return []string{}, false
			}
			o = append(o, s[b:i])
			b = i + 1
		}
	}
	if len(s[b:]) == 0 || len(s[b:]) > 4 {
		return []string{}, false
	}
	o = append(o, s[b:])

	if len(o) > 8 {
		return []string{}, false
	}

	return o, true
}

func splitShortV6(s string) ([]string, bool) {
	var f bool
	var o = make([][]string, 2)
	l := len(s) - 1
	for i := 0; i < l; i++ {
		if s[i] == v6sep && s[i+1] == v6sep {
			for j, v := range []string{s[:i], s[i+2:]} {
				if len(v) == 0 {
					continue
				}
				u, ok := splitV6(v)
				if !ok {
					return []string{}, ok
				}
				o[j] = u
			}
			f = true
			break
		}
	}
	if f {
		return stretchV6(o)
	}
	return nil, false
}

func stretchV6(o [][]string) ([]string, bool) {
	var f []string

	h, hl := o[0], len(o[0])
	l, ll := o[1], len(o[1])

	if hl+ll > 7 {
		return []string{}, false
	}

	ml := 8 - (hl + ll)

	for i := 0; i < hl; i++ {
		f = append(f, h[i])
	}
	for i := 0; i < ml; i++ {
		f = append(f, "0")
	}
	for i := 0; i < ll; i++ {
		f = append(f, l[i])
	}
	return f, true
}

func parseV6(s string) ([8]uint64, bool) {
	sg, ok := splitShortV6(s)
	if ok {
		return a16toUint64(sg)
	}
	sg, ok = splitV6(s)
	if !ok {
		return [8]uint64{}, false
	}
	return a16toUint64(sg)
}

func a16toUint64(s []string) ([8]uint64, bool) {
	var r [8]uint64
	if len(s) != 8 {
		return [8]uint64{}, false
	}
	for idx, v := range s {
		b, err := strconv.ParseUint(v, 16, 64)
		if err != nil {
			return [8]uint64{}, false
		}
		r[idx] = b
	}
	return r, true
}

func parseIPv6(addrStr string) (v6IP, error) {
	if len(addrStr) <= 1 {
		return v6IP{}, fmt.Errorf("%s", Invalid)
	}
	var a, ok = parseV6(addrStr)
	if !ok {
		return v6IP{}, fmt.Errorf("%s", Invalid)
	}
	var ipv6 v6IP
	ipv6.address.set(a[0]<<48|a[1]<<32|a[2]<<16|a[3], a[4]<<48|a[5]<<32|a[6]<<16|a[7])
	return ipv6, nil
}
