package inverter

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"
)

func TestValidateInverter(t *testing.T) {
	type test struct {
		data   Inverter
		answer error
	}

	tests := []test{
		{
			data: Inverter{
				info{
					PhaseNumber:     "1",
					VARating:        "6000",
					FirmwareVersion: "00.01.0001",
					ModelName:       "ES 6096s UK",
					Manufacturer:    "EATON",
					SerialNumber:    "RB76H39016",
					NominalPV:       "360.0",
				},
				bms{
					Info: batInfo{
						ModelType:       "C",
						FirmwareVersion: "1619",
						SerialNumber:    "66",
					},
					Data: batData{
						HighestCellVoltage: "3844.5",
						LowestCellVoltage:  "3840",
						SysMaxTemperature:  "32.5",
						SysAvgTemperature:  "30.6",
						SysMinTemperature:  "28.6",
					},
				},
			},
			answer: nil,
		}, {
			data: Inverter{
				info{
					PhaseNumber:     "ZZZZZZZZ",
					VARating:        "6000",
					FirmwareVersion: "00.01.0001",
					ModelName:       "ES 6096s UK",
					Manufacturer:    "EATON",
					SerialNumber:    "RB76H39016",
					NominalPV:       "360.0",
				},
				bms{
					Info: batInfo{
						ModelType:       "C",
						FirmwareVersion: "1619",
						SerialNumber:    "66",
					},
					Data: batData{
						HighestCellVoltage: "3844.5",
						LowestCellVoltage:  "3840",
						SysMaxTemperature:  "32.5",
						SysAvgTemperature:  "30.6",
						SysMinTemperature:  "28.6",
					},
				},
			},
			answer: fmt.Errorf("error"),
		}, {
			data: Inverter{
				info{
					PhaseNumber:     "1",
					VARating:        "6000",
					FirmwareVersion: "00.01.0001",
					ModelName:       "ES 6096s UK",
					Manufacturer:    "EATON",
					SerialNumber:    "RB76H39016",
					NominalPV:       "360.0",
				},
				bms{
					Info: batInfo{
						ModelType:       "ZZZZZ",
						FirmwareVersion: "1619",
						SerialNumber:    "66",
					},
					Data: batData{
						HighestCellVoltage: "3844.5",
						LowestCellVoltage:  "3840",
						SysMaxTemperature:  "32.5",
						SysAvgTemperature:  "30.6",
						SysMinTemperature:  "28.6",
					},
				},
			},
			answer: fmt.Errorf("error"),
		}, {
			data: Inverter{
				info{
					PhaseNumber:     "1",
					VARating:        "6000",
					FirmwareVersion: "00.01.0001",
					ModelName:       "ES 6096s UK",
					Manufacturer:    "EATON",
					SerialNumber:    "RB76H39016",
					NominalPV:       "360.0",
				},
				bms{
					Info: batInfo{
						ModelType:       "C",
						FirmwareVersion: "1619",
						SerialNumber:    "66",
					},
					Data: batData{
						HighestCellVoltage: "3844.5",
						LowestCellVoltage:  "3840",
						SysMaxTemperature:  "32.5",
						SysAvgTemperature:  "30.6",
						SysMinTemperature:  "ZZZZZZZZZ",
					},
				},
			},
			answer: fmt.Errorf("error"),
		},
	}

	for _, v := range tests {
		x := v.data.Validate()
		if (x == nil && v.answer != nil) || (x != nil && v.answer == nil) {
			t.Error("got", x, "want", v.answer)
		}
	}
}

func BenchmarkValidateInverter(b *testing.B) {

	inv := Inverter{
		info{
			PhaseNumber:     "1",
			VARating:        "6000",
			FirmwareVersion: "00.01.0001",
			ModelName:       "ES 6096s UK",
			Manufacturer:    "EATON",
			SerialNumber:    "RB76H39016",
			NominalPV:       "360.0",
		},
		bms{
			Info: batInfo{
				ModelType:       "C",
				FirmwareVersion: "1619",
				SerialNumber:    "66",
			},
			Data: batData{
				HighestCellVoltage: "3844.5",
				LowestCellVoltage:  "3840",
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  "30.6",
				SysMinTemperature:  "28.6",
			},
		},
	}

	for j := 0; j < b.N; j++ {
		inv.Validate()
	}
}

func TestNumerics(t *testing.T) {
	type test struct {
		data       string
		multiplier float64
		result     string
	}

	tester := test{
		data:       "360.0",
		multiplier: 10,
		result:     "3600",
	}

	x, _ := numerics(tester.data, tester.multiplier)

	if tester.result != x {
		t.Error("got", x, "want", tester.result)
	}
}

func BenchmarkNumerics(b *testing.B) {

	type test struct {
		data       string
		multiplier float64
		result     string
	}

	tester := test{
		data:       "360.0",
		multiplier: 10,
		result:     "3600",
	}

	for i := 0; i < b.N; i++ {
		numerics(tester.data, tester.multiplier)
	}
}

func TestStrToBytes(t *testing.T) {
	type test struct {
		data    string
		size    byte
		padding func([]byte, byte) []byte
		result  []byte
	}

	tests := []test{
		{
			data:    "6000",
			size:    6,
			padding: paddingLeft,
			result:  []byte{0x20, 0x20, 0x36, 0x30, 0x30, 0x30},
		}, {
			data:    "ES 6096s UK",
			size:    16,
			padding: paddingRight,
			result:  []byte{0x45, 0x53, 0x20, 0x36, 0x30, 0x39, 0x36, 0x73, 0x20, 0x55, 0x4B, 0x20, 0x20, 0x20, 0x20, 0x20},
		},
	}

	for _, v := range tests {
		x := strToBytes(v.data, v.size, v.padding)

		for i, val := range x {
			if val != v.result[i] {
				t.Error("got", x, "want", v.result)
			}
		}
	}
}

func BenchmarkStrToBytes(b *testing.B) {

	type test struct {
		data    string
		size    byte
		padding func([]byte, byte) []byte
	}

	tester := test{
		data:    "6000",
		size:    6,
		padding: paddingLeft,
	}

	for i := 0; i < b.N; i++ {
		strToBytes(tester.data, tester.size, tester.padding)
	}
}

func TestNumToBytes(t *testing.T) {
	type test struct {
		data       string
		multiplier float64
		size       byte
		result     []byte
	}

	tests := []test{
		{
			data:       "3844.5",
			multiplier: voltFactor,
			size:       2,
			result:     []byte{0x0A, 0x03},
		}, {
			data:       "30.6",
			multiplier: tempFactor,
			size:       2,
			result:     []byte{0x01, 0x32},
		}, {
			data:       "66",
			multiplier: 1,
			size:       4,
			result:     []byte{0x00, 0x00, 0x00, 0x42},
		},
	}

	for _, v := range tests {
		x, _ := numToBytes(v.data, v.multiplier, v.size)

		for i, val := range x {
			if val != v.result[i] {
				t.Error("got", x, "want", v.result)
			}
		}
	}
}

func BenchmarkNumToBytes(b *testing.B) {

	type test struct {
		data       string
		multiplier float64
		size       byte
	}

	tester := test{
		data:       "3844.5",
		multiplier: voltFactor,
		size:       2,
	}

	for i := 0; i < b.N; i++ {
		numToBytes(tester.data, tester.multiplier, tester.size)
	}
}

func TestPaddingRight(t *testing.T) {

	test := []byte("test")
	size := byte(7)

	answer := make([]byte, size)

	for i := range answer {
		if i < len(test) {
			answer[i] = test[i]
			continue
		}
		answer[i] = 0x20
	}

	x := paddingRight(test, size)

	for i, v := range x {
		if v != answer[i] {
			t.Error("got", x, "want", answer)
			break
		}
	}
}

func BenchmarkPaddingRight(b *testing.B) {

	test := []byte("test")
	size := byte(7)

	for i := 0; i < b.N; i++ {
		paddingRight(test, size)
	}
}

func TestPaddingLeft(t *testing.T) {

	test := []byte("test")
	size := byte(7)

	answer := make([]byte, size)

	for i := range answer {
		if i < int(size)-len(test) {
			answer[i] = 0x20
			continue
		}
		answer[i] = test[i-(int(size)-len(test))]
	}

	x := paddingLeft(test, size)

	for i, v := range x {
		if v != answer[i] {
			t.Error("got", x, "want", answer)
			break
		}
	}
}

func BenchmarkPaddingLeft(b *testing.B) {

	test := []byte("test")
	size := byte(7)

	for i := 0; i < b.N; i++ {
		paddingLeft(test, size)
	}
}

func TestCheckSize(t *testing.T) {

	type test struct {
		data       string
		multiplier float64
		maxsize    byte
		answer     bool
	}

	tests := []test{
		{
			data:       strconv.Itoa(rand.Intn(100)),
			multiplier: rand.Float64() * 100,
			maxsize:    2,
			answer:     true,
		}, {
			data:       strconv.Itoa(rand.Intn(100)),
			multiplier: rand.Float64() * 100,
			maxsize:    4,
			answer:     true,
		},
		{
			data:       strconv.Itoa(rand.Intn(100)),
			multiplier: rand.Float64() * 100,
			maxsize:    6,
			answer:     false,
		}, {
			data:       strconv.Itoa(math.MaxUint16 + 1),
			multiplier: 1,
			maxsize:    2,
			answer:     false,
		},
		{
			data:       strconv.Itoa(math.MaxUint32 + 1),
			multiplier: 1,
			maxsize:    4,
			answer:     false,
		},
		{
			data:       strconv.Itoa(math.MaxInt32 + 1),
			multiplier: tempFactor,
			maxsize:    4,
			answer:     false,
		},
		{
			data:       strconv.Itoa(math.MaxInt16 + 1),
			multiplier: tempFactor,
			maxsize:    2,
			answer:     false,
		},
		{
			data:       "-30.2",
			multiplier: voltFactor,
			maxsize:    4,
			answer:     false,
		},
		{
			data:       "-30.2",
			multiplier: tempFactor,
			maxsize:    4,
			answer:     true,
		},
	}

	for _, v := range tests {
		x := checkSize(v.data, v.multiplier, v.maxsize)
		if x != v.answer {
			t.Error("got", x, "want", v.answer)
		}
	}
}

func BenchmarkCheckSize(b *testing.B) {

	type test struct {
		data       string
		multiplier float64
		maxsize    byte
	}

	tester := test{
		data:       strconv.Itoa(rand.Intn(100)),
		multiplier: rand.Float64() * 100,
		maxsize:    2,
	}

	for i := 0; i < b.N; i++ {
		checkSize(tester.data, tester.multiplier, tester.maxsize)
	}
}
