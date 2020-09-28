package inverter

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestGetBatInfo(t *testing.T) {

	bi := batInfo{
		ModelType:       "C",
		FirmwareVersion: "1619",
		SerialNumber:    "66",
	}

	answer := []byte{0x43, 0x06, 0x53, 0x00, 0x00, 0x00, 0x42}

	x, _ := bi.GetInfo()

	for i, v := range x {
		if v != answer[i] {
			t.Error("got", x, "want", answer)
			break
		}
	}
}

func BenchmarkGetBatInfo(b *testing.B) {

	bi := batInfo{
		ModelType:       "C",
		FirmwareVersion: "1619",
		SerialNumber:    "66",
	}

	for j := 0; j < b.N; j++ {
		bi.GetInfo()
	}
}

func TestGetBatData(t *testing.T) {

	bd := batData{
		HighestCellVoltage: "3844.5",
		LowestCellVoltage:  "3840",
		SysMaxTemperature:  "32.5",
		SysAvgTemperature:  "30.6",
		SysMinTemperature:  "28.6",
	}

	answer := []byte{0x0A, 0x03, 0x0A, 0x00, 0x01, 0x45, 0x01, 0x32, 0x01, 0x1E}

	x, _ := bd.GetData()

	for i, v := range x {
		if v != answer[i] {
			t.Error("got", x, "want", answer)
			break
		}
	}
}

func BenchmarkGetBatData(b *testing.B) {

	bd := batData{
		HighestCellVoltage: "3844.5",
		LowestCellVoltage:  "3840",
		SysMaxTemperature:  "32.5",
		SysAvgTemperature:  "30.6",
		SysMinTemperature:  "28.6",
	}

	for j := 0; j < b.N; j++ {
		bd.GetData()
	}
}

func TestValidateBatInfo(t *testing.T) {

	type test struct {
		data   batInfo
		answer error
	}

	tests := []test{
		{
			data: batInfo{
				ModelType:       "C",
				FirmwareVersion: "1619",
				SerialNumber:    "66",
			},

			answer: nil,
		}, {
			data: batInfo{
				ModelType:       "R",
				FirmwareVersion: "1619",
				SerialNumber:    "66",
			},

			answer: nil,
		}, {
			data: batInfo{
				ModelType:       "J",
				FirmwareVersion: "1619",
				SerialNumber:    "66",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: batInfo{
				ModelType:       "",
				FirmwareVersion: "1619",
				SerialNumber:    "66",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: batInfo{
				ModelType:       "CC",
				FirmwareVersion: "1619",
				SerialNumber:    "66",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: batInfo{
				ModelType:       "C",
				FirmwareVersion: "",
				SerialNumber:    "66",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: batInfo{
				ModelType:       "C",
				FirmwareVersion: strconv.Itoa(math.MaxUint16 + 1),
				SerialNumber:    "66",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: batInfo{
				ModelType:       "C",
				FirmwareVersion: "1619",
				SerialNumber:    "",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: batInfo{
				ModelType:       "C",
				FirmwareVersion: "1619",
				SerialNumber:    strconv.Itoa(math.MaxUint32 + 1),
			},
			answer: fmt.Errorf("error"),
		},
		{
			data: batInfo{
				ModelType:       "C",
				FirmwareVersion: "asdasds",
				SerialNumber:    "66",
			},
			answer: fmt.Errorf("error"),
		},
	}

	for _, v := range tests {
		x := v.data.validate()
		if (x == nil && v.answer != nil) || (x != nil && v.answer == nil) {
			t.Error("got", x, "want", v.answer)
		}
	}
}

func BenchmarkValidateBatInfo(b *testing.B) {

	bi := batInfo{
		ModelType:       "C",
		FirmwareVersion: "1619",
		SerialNumber:    "66",
	}

	for j := 0; j < b.N; j++ {
		bi.validate()
	}
}

func TestValidateBatData(t *testing.T) {

	type test struct {
		data   batData
		answer error
	}

	tests := []test{
		{
			data: batData{
				HighestCellVoltage: "3844.5",
				LowestCellVoltage:  "3840",
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  "30.6",
				SysMinTemperature:  "28.6",
			},

			answer: nil,
		}, {
			data: batData{
				HighestCellVoltage: "3844.5",
				LowestCellVoltage:  "3840",
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  "-30.6",
				SysMinTemperature:  "28.6",
			},
			answer: nil,
		}, {
			data: batData{
				HighestCellVoltage: "",
				LowestCellVoltage:  "3840",
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  "30.6",
				SysMinTemperature:  "28.6",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: batData{
				HighestCellVoltage: "3844.5",
				LowestCellVoltage:  "",
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  "30.6",
				SysMinTemperature:  "28.6",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: batData{
				HighestCellVoltage: "3844.5",
				LowestCellVoltage:  "3840",
				SysMaxTemperature:  "",
				SysAvgTemperature:  "30.6",
				SysMinTemperature:  "28.6",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: batData{
				HighestCellVoltage: "3844.5",
				LowestCellVoltage:  "3840",
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  "",
				SysMinTemperature:  "28.6",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: batData{
				HighestCellVoltage: "3844.5",
				LowestCellVoltage:  "3840",
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  "30.6",
				SysMinTemperature:  "",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: batData{
				HighestCellVoltage: strconv.FormatFloat((math.MaxUint16+1)/voltFactor, 'f', 1, 32),
				LowestCellVoltage:  "3840",
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  "30.6",
				SysMinTemperature:  "28.6",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: batData{
				HighestCellVoltage: "3844.5",
				LowestCellVoltage:  strconv.FormatFloat((math.MaxUint16+1)/voltFactor, 'f', 1, 32),
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  "30.6",
				SysMinTemperature:  "28.6",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: batData{
				HighestCellVoltage: "3844.5",
				LowestCellVoltage:  "3840",
				SysMaxTemperature:  strconv.FormatFloat((math.MaxInt16+1)/tempFactor, 'f', 1, 32),
				SysAvgTemperature:  "30.6",
				SysMinTemperature:  "28.6",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: batData{
				HighestCellVoltage: "3844.5",
				LowestCellVoltage:  "3840",
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  strconv.FormatFloat((math.MaxInt16+1)/tempFactor, 'f', 1, 32),
				SysMinTemperature:  "28.6",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: batData{
				HighestCellVoltage: "3844.5",
				LowestCellVoltage:  "3840",
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  "30.6",
				SysMinTemperature:  strconv.FormatFloat((math.MaxInt16+1)/tempFactor, 'f', 1, 32),
			},
			answer: fmt.Errorf("error"),
		}, {
			data: batData{
				HighestCellVoltage: "asdasdsad",
				LowestCellVoltage:  "3840",
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  "30.6",
				SysMinTemperature:  "28.6",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: batData{
				HighestCellVoltage: "3844.5",
				LowestCellVoltage:  "-3840",
				SysMaxTemperature:  "32.5",
				SysAvgTemperature:  "30.6",
				SysMinTemperature:  "28.6",
			},
			answer: fmt.Errorf("error"),
		},
	}

	for _, v := range tests {
		x := v.data.validate()
		if (x == nil && v.answer != nil) || (x != nil && v.answer == nil) {
			t.Error("got", x, "want", v.answer)
		}
	}
}

func BenchmarkValidateBatData(b *testing.B) {

	bd := batData{
		HighestCellVoltage: "3844.5",
		LowestCellVoltage:  "3840",
		SysMaxTemperature:  "32.5",
		SysAvgTemperature:  "30.6",
		SysMinTemperature:  "28.6",
	}

	for j := 0; j < b.N; j++ {
		bd.validate()
	}
}

func TestValidateBms(t *testing.T) {

	type test struct {
		data   bms
		answer error
	}

	tests := []test{
		{
			data: bms{
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

			answer: nil,
		}, {
			data: bms{
				Info: batInfo{
					ModelType:       "C",
					FirmwareVersion: "qweqweqw19",
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

			answer: fmt.Errorf("error"),
		}, {
			data: bms{
				Info: batInfo{
					ModelType:       "C",
					FirmwareVersion: "1619",
					SerialNumber:    "6qweqwewqe6",
				},
				Data: batData{
					HighestCellVoltage: "3844.5",
					LowestCellVoltage:  "3840",
					SysMaxTemperature:  "32.5",
					SysAvgTemperature:  "30.6",
					SysMinTemperature:  "28.6",
				},
			},

			answer: fmt.Errorf("error"),
		}, {
			data: bms{
				Info: batInfo{
					ModelType:       "C",
					FirmwareVersion: "1619",
					SerialNumber:    "66",
				},
				Data: batData{
					HighestCellVoltage: "3qweqwe",
					LowestCellVoltage:  "3840",
					SysMaxTemperature:  "32.5",
					SysAvgTemperature:  "30.6",
					SysMinTemperature:  "28.6",
				},
			},
			answer: fmt.Errorf("error"),
		}, {
			data: bms{
				Info: batInfo{
					ModelType:       "C",
					FirmwareVersion: "1619",
					SerialNumber:    "66",
				},
				Data: batData{
					HighestCellVoltage: "3844.5",
					LowestCellVoltage:  "sdaasdd",
					SysMaxTemperature:  "32.5",
					SysAvgTemperature:  "30.6",
					SysMinTemperature:  "28.6",
				},
			},

			answer: fmt.Errorf("error"),
		}, {
			data: bms{
				Info: batInfo{
					ModelType:       "C",
					FirmwareVersion: "1619",
					SerialNumber:    "66",
				},
				Data: batData{
					HighestCellVoltage: "3844.5",
					LowestCellVoltage:  "3840",
					SysMaxTemperature:  "sadadasdasd",
					SysAvgTemperature:  "30.6",
					SysMinTemperature:  "28.6",
				},
			},

			answer: fmt.Errorf("error"),
		}, {
			data: bms{
				Info: batInfo{
					ModelType:       "C",
					FirmwareVersion: "1619",
					SerialNumber:    "66",
				},
				Data: batData{
					HighestCellVoltage: "3844.5",
					LowestCellVoltage:  "3840",
					SysMaxTemperature:  "32.5",
					SysAvgTemperature:  "asdasdasd",
					SysMinTemperature:  "28.6",
				},
			},

			answer: fmt.Errorf("error"),
		}, {
			data: bms{
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
					SysMinTemperature:  "aaaaaaaaa",
				},
			},

			answer: fmt.Errorf("error"),
		},
	}

	for _, v := range tests {
		x := v.data.validate()
		if (x == nil && v.answer != nil) || (x != nil && v.answer == nil) {
			t.Error("got", x, "want", v.answer)
		}
	}
}

func BenchmarkValidateBms(b *testing.B) {

	bms := bms{
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
	}

	for j := 0; j < b.N; j++ {
		bms.validate()
	}
}
