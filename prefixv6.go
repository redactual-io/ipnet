package ipnet

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

type v6Prefix struct {
	first       address
	cursor      address
	inverseMask address
	last        address
	len         int
	mask        address
}

func (p v6Prefix) Member(in IP) bool {
	if in.Version() == p.Version() {
		addr := in.addr()
		if addr.ge(p.first) && addr.le(p.last) {
			return true
		}
	}
	return false
}

func (p v6Prefix) Version() Ver {
	return v6
}

func (p v6Prefix) First() IP {
	return v6IP{address: p.first}
}

func (p v6Prefix) Last() IP {
	return v6IP{address: p.last}
}

func (p v6Prefix) Mask() IP {
	return v6IP{address: p.mask}
}

func (p v6Prefix) Length() int {
	return p.len
}

func (p v6Prefix) Cursor() IP {
	return v6IP{address: p.cursor}
}

func (p *v6Prefix) Next() (IP, bool) {
	next, ok := p.cursor.add(1)
	if !ok {
		return v6IP{address: p.cursor}, false
	}
	ip := v6IP{address: next}
	if !p.Member(ip) {
		return v6IP{address: p.cursor}, false
	}
	p.cursor = next
	return ip, true
}

func (p *v6Prefix) Prev() (IP, bool) {
	prev, ok := p.cursor.subtract(1)
	if !ok {
		return v6IP{address: p.cursor}, false
	}
	ip := v6IP{address: prev}
	if !p.Member(ip) {
		return v6IP{address: p.cursor}, false
	}
	p.cursor = prev
	return ip, true
}

func (p v6Prefix) Random() IP {
	var rb = [16]byte{}
	if _, err := rand.Read(rb[:]); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	var ra address
	var shifts = [8]int{56, 48, 40, 32, 24, 16, 8, 0}
	for set, order := range []orders{higher, lower} {
		for idx, shift := range shifts {
			ra[order] = ra[order] | uint64(rb[idx+(set*8)])<<shift
		}
	}
	first, mask := p.First(), p.Mask()
	ra = ra.or(mask.addr()).xor(mask.addr()).or(first.addr())
	return v6IP{address: ra}
}

func (p v6Prefix) String() string {
	return fmt.Sprintf("%s/%d", v6IP{address: p.first}.String(), p.len)
}

func (p v6Prefix) IsInterface() bool {
	if p.first != p.cursor {
		return true
	}
	return false
}

func v6SplitLen(pfx string) ([2]string, bool) {
	l := len(pfx)
	s := l - 5
	var r [2]string
	for i := l - 1; i > s; i-- {
		if pfx[i] == 47 {
			r[0], r[1] = pfx[:i], pfx[i+1:]
			return r, true
		}
	}
	return [2]string{}, false
}

func parseIPv6Prefix(pfx string) (*v6Prefix, error) {
	var p = &v6Prefix{}
	x, ok := v6SplitLen(pfx)
	if !ok {
		return nil, fmt.Errorf("%s", Invalid)
	}
	ipString, lengthString := x[0], x[1]

	p.len, ok = parseIPv6Length(lengthString)
	if !ok {
		return nil, fmt.Errorf("%s", Invalid)
	}

	// masks
	p.mask, p.inverseMask = ipv6Mask(p.len)

	ip, err := parseIPv6(ipString)
	if err != nil {
		return nil, err
	}

	p.first = ip.address.or(p.inverseMask).xor(p.inverseMask)
	p.last = ip.address.or(p.inverseMask)

	p.cursor = ip.address

	return p, nil
}

func ipv6Mask(length int) (mask address, inverseMask address) {
	for i := 0; i < length; i++ {
		var o = i / 64
		mask[o] |= 1 << (64 - 1 - (i - 64*o))
	}
	inverseMask = address{0xffffffffffffffff ^ mask[0], 0xffffffffffffffff ^ mask[1]}
	return mask, inverseMask
}

func parseIPv6Length(length string) (int, bool) {
	l, err := strconv.Atoi(length)
	if err != nil {
		return 0, false
	}
	if !(0 <= l && l <= 128) {
		return 0, false
	}
	return l, true
}
