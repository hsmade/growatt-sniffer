package decode

import "encoding/binary"

type Data struct {
	DataLoggerSerial [10]byte // string
	PVSerial         [10]byte // string
	PVPowerin        float64
	P1Voltage        float64
	PV1Current       float64
	PV1Watt          float64
}

func Decode(data []byte) (Data, error) {

	result := Data{}
	err := UnmarshalBinary(data, &result)
	return result, err
}

func UnmarshalBinary(data []byte, s *Data) error {
	copy(s.DataLoggerSerial[:], data[8:18])
	copy(s.PVSerial[:], data[38:48])
	s.PVPowerin = float64(binary.BigEndian.Uint32(data[81:85])) / 10
	s.P1Voltage = float64(binary.BigEndian.Uint16(data[85:87])) / 10
	s.PV1Current = float64(binary.BigEndian.Uint16(data[87:89])) / 10
	s.PV1Watt = float64(binary.BigEndian.Uint32(data[89:93])) / 10

	return nil
}
