package ipnet

import "fmt"

type IP interface {
	addr() address
	String() string
	Version() Ver
}

func ParseIP(addr string) (IP, error) {
	for i := 0; i < len(addr); i++ {
		switch addr[i] {
		case v4sep:
			ip, err := parseIPv4(addr)
			if err != nil {
				return nil, err
			}
			return ip, nil
		case v6sep:
			ip, err := parseIPv6(addr)
			if err != nil {
				return nil, err
			}
			return ip, nil
		}
	}
	return nil, fmt.Errorf("%s", Invalid)
}
