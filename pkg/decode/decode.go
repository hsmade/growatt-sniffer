package decode

import (
	"encoding/binary"
	"log/slog"
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
	Date                                  time.Time
	DataLoggerSerial                      string // string
	Serial                                string // string
	Status                                Status
	PowerIn                               float64
	TotalPowerOut                         float64
	Frequency                             float64 // in or out?
	PV1Voltage                            float64
	PV1Current                            float64
	PV1Power                              float64
	PV2Voltage                            float64
	PV2Current                            float64
	PV2Power                              float64
	GridFase1Voltage                      float64
	GridFase1Current                      float64
	GridFase1Power                        float64
	GridFase2Voltage                      float64
	GridFase2Current                      float64
	GridFase2Power                        float64
	GridFase3Voltage                      float64
	GridFase3Current                      float64
	GridFase3Power                        float64
	EnergyToday                           float64
	EnergyTotal                           float64
	PV1EnergyToday                        float64
	PV1EnergyTotal                        float64
	PV2EnergyToday                        float64
	PV2EnergyTotal                        float64
	PVEnergyTotal                         float64
	TotalWorkingTime                      time.Duration
	InverterTemperature                   float64
	IntelligentPowerManagementTemperature float64
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
	// T06NNNNX
	s.DataLoggerSerial = string(data[8:18])
	s.Serial = string(data[38:48])
	year := binary.BigEndian.Uint16([]byte{0x00, data[68:69][0]}) + 2000
	month := binary.BigEndian.Uint16([]byte{0x00, data[69:70][0]})
	day := binary.BigEndian.Uint16([]byte{0x00, data[70:71][0]})
	hour := binary.BigEndian.Uint16([]byte{0x00, data[71:72][0]})
	minute := binary.BigEndian.Uint16([]byte{0x00, data[72:73][0]})
	second := binary.BigEndian.Uint16([]byte{0x00, data[73:74][0]})
	s.Date = time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, time.Local)
	_ = float64(binary.BigEndian.Uint16(data[75:77])) // record type 1? (3000)
	_ = float64(binary.BigEndian.Uint16(data[77:79])) // record type 2? (3124)
	s.Status = Status(binary.BigEndian.Uint16(data[79:81]))
	s.PowerIn = float64(binary.BigEndian.Uint32(data[81:85])) / 10
	s.PV1Voltage = float64(binary.BigEndian.Uint16(data[85:87])) / 10
	s.PV1Current = float64(binary.BigEndian.Uint16(data[87:89])) / 10
	s.PV1Power = float64(binary.BigEndian.Uint32(data[89:93])) / 10
	s.PV2Voltage = float64(binary.BigEndian.Uint16(data[93:95])) / 10
	s.PV2Current = float64(binary.BigEndian.Uint16(data[95:97])) / 10
	s.PV2Power = float64(binary.BigEndian.Uint32(data[97:101])) / 10
	s.TotalPowerOut = float64(int(data[128])+int(data[127])*256+int(data[126])*256*256+int(data[125])*256*256*256) / 10
	s.Frequency = float64(binary.BigEndian.Uint16(data[129:131])) / 100
	s.GridFase1Voltage = float64(binary.BigEndian.Uint16(data[131:133])) / 10
	s.GridFase1Current = float64(binary.BigEndian.Uint16(data[133:135])) / 10
	s.GridFase1Power = float64(binary.BigEndian.Uint32(data[135:139])) / 10
	s.GridFase2Voltage = float64(binary.BigEndian.Uint16(data[139:141])) / 10
	s.GridFase2Current = float64(binary.BigEndian.Uint16(data[141:143])) / 10
	s.GridFase2Power = float64(binary.BigEndian.Uint32(data[143:147])) / 10
	s.GridFase3Voltage = float64(binary.BigEndian.Uint16(data[147:149])) / 10
	s.GridFase3Current = float64(binary.BigEndian.Uint16(data[149:151])) / 10
	s.GridFase3Power = float64(binary.BigEndian.Uint32(data[151:155])) / 10
	s.TotalWorkingTime = time.Duration(binary.BigEndian.Uint32(data[173:177])/2) * time.Second
	s.EnergyToday = float64(binary.BigEndian.Uint32(data[177:181])) / 10
	s.EnergyTotal = float64(binary.BigEndian.Uint32(data[181:185])) / 10
	s.PVEnergyTotal = float64(binary.BigEndian.Uint32(data[185:189])) / 10
	s.PV1EnergyToday = float64(binary.BigEndian.Uint32(data[189:193])) / 10
	s.PV1EnergyTotal = float64(binary.BigEndian.Uint32(data[193:197])) / 10
	s.PV2EnergyToday = float64(binary.BigEndian.Uint32(data[197:201])) / 10
	s.PV2EnergyTotal = float64(binary.BigEndian.Uint32(data[201:205])) / 10
	s.InverterTemperature = float64(binary.BigEndian.Uint16(data[265:267])) / 10
	s.IntelligentPowerManagementTemperature = float64(binary.BigEndian.Uint16(data[273:275])) / 10
	_ = float64(binary.BigEndian.Uint16(data[275:277])) / 10          // pbusvolt, inverter bus? (361.1)
	_ = float64(binary.BigEndian.Uint16(data[277:279])) / 10          // nbusvolt, battery bus? (0)
	_ = float64(binary.BigEndian.Uint16([]byte{0x00, data[277]})) / 1 // battery1soc (0)
	defer func() {
		if r := recover(); r != nil {
			slog.Error("recovered from panic", "r", r)
		}
	}()
	return nil
}
