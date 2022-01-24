package ipnet

import (
	"fmt"
)

type Prefix interface {
	First() IP

	// Cursor represents the IP of the CIDR string that was parsed, but may also
	// be used to track a location within the prefix.
	//
	//  Example:
	//
	//  p, err := ipnet.ParsePrefix("192.168.127.240/10")
	//	if err != nil {
	//		//
	//	}
	//  for ip, ok := p.Cursor(), true; ok; ip, ok = p.Next() {
	//		//
	//  }
	Cursor() IP
	IsInterface() bool
	Last() IP
	Length() int
	Mask() IP
	Member(ip IP) bool
	Next() (IP, bool)
	Prev() (IP, bool)
	Random() IP
	String() string
	Version() Ver
}

func ParsePrefix(p string) (Prefix, error) {
	for i := 0; i < len(p); i++ {
		switch p[i] {
		case v4sep:
			ip, err := parseIPv4Prefix(p)
			if err != nil {
				return nil, err
			}
			return ip, nil
		case v6sep:
			ip, err := parseIPv6Prefix(p)
			if err != nil {
				return nil, err
			}
			return ip, nil
		}
	}
	return nil, fmt.Errorf("%s", Invalid)
}
