package ipnet

type Ver uint8

const (
	v4 Ver = 4
	v6 Ver = 6
)

func (v Ver) String() string {
	switch v {
	case v4:
		return "v4"
	case v6:
		return "v6"
	default:
		return "UNKNOWN"
	}
}
