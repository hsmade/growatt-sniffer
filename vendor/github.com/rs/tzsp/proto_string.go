// Code generated by "stringer -type=Proto"; DO NOT EDIT

package tzsp

import "fmt"

const (
	_Proto_name_0 = "ProtoEthernet"
	_Proto_name_1 = "ProtoIEEE80211"
	_Proto_name_2 = "ProtoPrismHeader"
	_Proto_name_3 = "ProtoWLANAVS"
)

var (
	_Proto_index_0 = [...]uint8{0, 13}
	_Proto_index_1 = [...]uint8{0, 14}
	_Proto_index_2 = [...]uint8{0, 16}
	_Proto_index_3 = [...]uint8{0, 12}
)

func (i Proto) String() string {
	switch {
	case i == 1:
		return _Proto_name_0
	case i == 18:
		return _Proto_name_1
	case i == 119:
		return _Proto_name_2
	case i == 127:
		return _Proto_name_3
	default:
		return fmt.Sprintf("Proto(%d)", i)
	}
}