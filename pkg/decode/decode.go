package decode

import (
	"encoding/binary"
	"time"
)

type Status int

const (
	StatusStandby Status = iota
	StatusNoUse
	StatusDischarge
	StatusFault
	StatusFlash
	StatusPvCharge
	StatusAcCharge
	StatusCombineCharge
	StatusCombineChargeAndBypass
	StatusPvChargeAndBypass
	StatusAcChargeAndBypass
	StatusBypass
	StatusPvChargeAndDischarge
)

type Data struct {
	Date             time.Time
	DataLoggerSerial [10]byte // string
	Serial           [10]byte // string
	Status           Status
	PowerIn          float64
	PowerOut         float64
	Frequency        float64 // in or out?
	PV1Voltage       float64
	PV1Current       float64
	PV1Watt          float64
	PV2Voltage       float64
	PV2Current       float64
	PV2Watt          float64
	GridVoltage      float64
	GridCurrent      float64
	GridPower        float64 // difference from PowerOut?
	EnergyToday      float64
	EnergyTotal      float64
	PV1EnergyToday   float64
	PV1EnergyTotal   float64
	PV2EnergyToday   float64
	PV2EnergyTotal   float64
	PVEnergyTotal    float64
	TotalWorkingTime time.Duration
	Temperature1     float64
	Temperature2     float64
}

func (S Status) String() string {
	switch S {
	case 0:
		return "Standby"
	case 1:
		return "NoUse"
	case 2:
		return "Discharge"
	case 3:
		return "Fault"
	case 4:
		return "Flash"
	case 5:
		return "PV Charge"
	case 6:
		return "AC Charge"
	case 7:
		return "Combine Charge"
	case 8:
		return "Combine Charge and Bypass"
	case 9:
		return "PV Charge and Bypass"
	case 10:
		return "AC Charge and Bypass"
	case 11:
		return "Bypass"
	case 12:
		return "PC Charge and Discharge"
	}
	return "unknown status"
}

func Decode(data []byte) (Data, error) {

	result := Data{}
	err := UnmarshalBinary(data, &result)
	return result, err
}

func UnmarshalBinary(data []byte, s *Data) error {
	//s.Date = float64(binary.BigEndian.Uint32(data[68:???])) / 10
	copy(s.DataLoggerSerial[:], data[8:18])
	copy(s.Serial[:], data[38:48])
	s.Status = Status(binary.BigEndian.Uint32(data[79:81]))
	s.PowerIn = float64(binary.BigEndian.Uint32(data[81:85])) / 10
	s.PV1Voltage = float64(binary.BigEndian.Uint16(data[85:87])) / 10
	s.PV1Current = float64(binary.BigEndian.Uint16(data[87:89])) / 10
	s.PV1Watt = float64(binary.BigEndian.Uint32(data[89:93])) / 10
	s.PV2Voltage = float64(binary.BigEndian.Uint16(data[93:95])) / 10
	s.PV2Current = float64(binary.BigEndian.Uint16(data[95:97])) / 10
	s.PV2Watt = float64(binary.BigEndian.Uint32(data[97:101])) / 10
	s.PowerOut = float64(binary.BigEndian.Uint32(data[101:105])) / 10 // fixme: is signed int!
	s.Frequency = float64(binary.BigEndian.Uint16(data[105:107])) / 100
	s.GridVoltage = float64(binary.BigEndian.Uint16(data[107:109])) / 10
	s.GridCurrent = float64(binary.BigEndian.Uint16(data[109:111])) / 10
	s.GridPower = float64(binary.BigEndian.Uint32(data[111:115])) / 10
	// grid 2 and grid 3 volt/current/power
	s.EnergyToday = float64(binary.BigEndian.Uint32(data[131:135])) / 10
	s.EnergyTotal = float64(binary.BigEndian.Uint32(data[135:139])) / 10
	s.TotalWorkingTime = float64(binary.BigEndian.Uint32(data[139:143])) / 7200
	s.Temperature1 = float64(binary.BigEndian.Uint16(data[143:145])) / 10
	// fault codes
	s.Temperature2 = float64(binary.BigEndian.Uint16(data[161:163])) / 10
	// bus volt
	s.PV1EnergyToday = float64(binary.BigEndian.Uint32(data[179:183])) / 10
	s.PV1EnergyTotal = float64(binary.BigEndian.Uint32(data[183:187])) / 10
	s.PV2EnergyToday = float64(binary.BigEndian.Uint32(data[187:191])) / 10
	s.PV2EnergyTotal = float64(binary.BigEndian.Uint32(data[191:195])) / 10
	s.PVEnergyTotal = float64(binary.BigEndian.Uint32(data[195:199])) / 10
	return nil
}
