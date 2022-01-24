package ipnet

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

type v4Prefix struct {
	first       address
	cursor      address
	inverseMask address
	last        address
	len         int
	mask        address
}

func (p v4Prefix) Member(in IP) bool {
	if in.Version() == p.Version() {
		addr := in.addr()
		if addr.ge(p.first) && addr.le(p.last) {
			return true
		}
	}
	return false
}

func (p v4Prefix) Version() Ver {
	return v4
}

func (p v4Prefix) First() IP {
	return v4IP{address: p.first}
}

func (p v4Prefix) Last() IP {
	return v4IP{address: p.last}
}

func (p v4Prefix) Mask() IP {
	return v4IP{address: p.mask}
}

func (p v4Prefix) Length() int {
	return p.len
}

func (p v4Prefix) Cursor() IP {
	return v4IP{address: p.cursor}
}

func (p *v4Prefix) Next() (IP, bool) {
	// address overflow shouldn't be possible with IPv4, so we're ignoring the ok from address.add()
	next, _ := p.cursor.add(1)
	ip := v4IP{address: next}
	if !p.Member(ip) {
		return v4IP{address: p.cursor}, false
	}
	p.cursor = next
	return ip, true
}

func (p *v4Prefix) Prev() (IP, bool) {
	prev, ok := p.cursor.subtract(1)
	if !ok {
		return v4IP{address: p.cursor}, false
	}
	ip := v4IP{address: prev}
	if !p.Member(ip) {
		return v4IP{address: p.cursor}, false
	}
	p.cursor = prev
	return ip, true
}

func (p v4Prefix) Random() IP {
	var rb = [4]byte{}
	if _, err := rand.Read(rb[:]); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	var ra address
	ra.set(0, uint64(rb[0])<<24|uint64(rb[1])<<16|uint64(rb[2])<<8|uint64(rb[3]))
	first, mask := p.First(), p.Mask()
	ra = ra.or(mask.addr()).xor(mask.addr()).or(first.addr())
	return v4IP{address: ra}
}

func (p v4Prefix) String() string {
	return fmt.Sprintf("%s/%d", v4IP{address: p.first}.String(), p.len)
}

func (p v4Prefix) IsInterface() bool {
	if p.first != p.cursor {
		return true
	}
	return false
}

func v4SplitLen(pfx string) ([2]string, bool) {
	l := len(pfx)
	s := l - 4
	var r [2]string
	for i := l - 1; i > s; i-- {
		if pfx[i] == 47 {
			r[0], r[1] = pfx[:i], pfx[i+1:]
			return r, true
		}
	}
	return [2]string{}, false
}

func parseIPv4Prefix(pfx string) (*v4Prefix, error) {
	var p = &v4Prefix{}
	x, ok := v4SplitLen(pfx)
	if !ok {
		return nil, fmt.Errorf("%s", Invalid)
	}
	ipString, lengthString := x[0], x[1]

	p.len, ok = parseIPv4Length(lengthString)
	if !ok {
		return nil, fmt.Errorf("%s", Invalid)
	}

	// masks
	p.mask, p.inverseMask = ipv4Mask(p.len)

	ip, err := parseIPv4(ipString)
	if err != nil {
		return nil, err
	}

	p.first = ip.address.or(p.inverseMask).xor(p.inverseMask)
	p.last = ip.address.or(p.inverseMask)

	p.cursor = ip.address

	return p, nil
}

func ipv4Mask(length int) (mask address, inverseMask address) {
	for i := 0; i < length; i++ {
		mask[1] |= 1 << (32 - 1 - i)
	}
	inverseMask[1] = 0xffffffff ^ mask[1]
	return
}

func parseIPv4Length(length string) (int, bool) {
	l, err := strconv.Atoi(length)
	if err != nil {
		return 0, false
	}
	if !(0 <= l && l <= 32) {
		return 0, false
	}
	return l, true
}
