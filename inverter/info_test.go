package inverter

import (
	"fmt"
	"testing"
)

func TestGetInverterInfo(t *testing.T) {

	i := info{
		PhaseNumber:     "1",
		VARating:        "6000",
		FirmwareVersion: "00.01.0001",
		ModelName:       "ES 6096s UK",
		Manufacturer:    "EATON",
		SerialNumber:    "RB76H39016",
		NominalPV:       "360.0",
	}
	/*
		A := i.PhaseNumber
		B := strings.Repeat(" ", 6-len(i.VARating)) + i.VARating
		C := strings.Repeat(" ", 10-len(i.FirmwareVersion)) + i.FirmwareVersion
		D := i.ModelName + strings.Repeat(" ", 16-len(i.ModelName))
		E := i.Manufacturer + strings.Repeat(" ", 16-len(i.Manufacturer))
		F := i.SerialNumber + strings.Repeat(" ", 16-len(i.SerialNumber))
		G := "3600"

		s := A + B + C + D + E + F + G
	*/

	// "1  600000.01.0001ES 6096s UK     EATON           RB76H39016      3600"
	answer := []byte{0x31, 0x20, 0x20, 0x36, 0x30, 0x30, 0x30, 0x30, 0x30, 0x2e, 0x30, 0x31, 0x2e, 0x30, 0x30, 0x30, 0x31, 0x45, 0x53, 0x20, 0x36, 0x30, 0x39, 0x36, 0x73, 0x20, 0x55, 0x4b, 0x20, 0x20, 0x20, 0x20, 0x20, 0x45, 0x41, 0x54, 0x4f, 0x4e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x52, 0x42, 0x37, 0x36, 0x48, 0x33, 0x39, 0x30, 0x31, 0x36, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x33, 0x36, 0x30, 0x30}

	x, _ := i.GetInfo()

	for i, v := range x {
		if v != answer[i] {
			t.Error("got", x, "want", answer)
			break
		}
	}
}

func BenchmarkGetInverterInfo(b *testing.B) {

	i := info{
		PhaseNumber:     "1",
		VARating:        "6000",
		FirmwareVersion: "00.01.0001",
		ModelName:       "ES 6096s UK",
		Manufacturer:    "EATON",
		SerialNumber:    "RB76H39016",
		NominalPV:       "360.0",
	}

	for j := 0; j < b.N; j++ {
		i.GetInfo()
	}
}

func TestValidateInverterInfo(t *testing.T) {

	type test struct {
		data   info
		answer error
	}

	tests := []test{
		{
			data: info{
				PhaseNumber:     "1",
				VARating:        "6000",
				FirmwareVersion: "00.01.0001",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "EATON",
				SerialNumber:    "RB76H39016",
				NominalPV:       "360.0",
			},

			answer: nil,
		}, {
			data: info{
				PhaseNumber:     "",
				VARating:        "6000",
				FirmwareVersion: "00.01.0001",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "EATON",
				SerialNumber:    "RB76H39016",
				NominalPV:       "360.0",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "1",
				VARating:        "",
				FirmwareVersion: "00.01.0001",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "EATON",
				SerialNumber:    "RB76H39016",
				NominalPV:       "360.0",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "1",
				VARating:        "6000",
				FirmwareVersion: "",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "EATON",
				SerialNumber:    "RB76H39016",
				NominalPV:       "360.0",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "1",
				VARating:        "6000",
				FirmwareVersion: "00.01.0001",
				ModelName:       "",
				Manufacturer:    "EATON",
				SerialNumber:    "RB76H39016",
				NominalPV:       "360.0",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "1",
				VARating:        "6000",
				FirmwareVersion: "00.01.0001",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "",
				SerialNumber:    "RB76H39016",
				NominalPV:       "360.0",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "1",
				VARating:        "6000",
				FirmwareVersion: "00.01.0001",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "EATON",
				SerialNumber:    "",
				NominalPV:       "360.0",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "1",
				VARating:        "6000",
				FirmwareVersion: "00.01.0001",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "EATON",
				SerialNumber:    "RB76H39016",
				NominalPV:       "",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "11",
				VARating:        "6000",
				FirmwareVersion: "00.01.0001",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "EATON",
				SerialNumber:    "RB76H39016",
				NominalPV:       "360.0",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "1",
				VARating:        "6000000",
				FirmwareVersion: "00.01.0001",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "EATON",
				SerialNumber:    "RB76H39016",
				NominalPV:       "360.0",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "1",
				VARating:        "6000",
				FirmwareVersion: "00.01.00011",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "EATON",
				SerialNumber:    "RB76H39016",
				NominalPV:       "360.0",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "1",
				VARating:        "6000",
				FirmwareVersion: "00.01.0001",
				ModelName:       "ES 6096s UKZZZZZZ",
				Manufacturer:    "EATON",
				SerialNumber:    "RB76H39016",
				NominalPV:       "360.0",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "1",
				VARating:        "6000",
				FirmwareVersion: "00.01.0001",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "EATONZZZZZZZZZZZZ",
				SerialNumber:    "RB76H39016",
				NominalPV:       "360.0",
			},

			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "1",
				VARating:        "6000",
				FirmwareVersion: "00.01.0001",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "EATON",
				SerialNumber:    "RB76H39016ZZZZZZZ",
				NominalPV:       "360.0",
			},
			answer: fmt.Errorf("error"),
		}, {
			data: info{
				PhaseNumber:     "1",
				VARating:        "6000",
				FirmwareVersion: "00.01.0001",
				ModelName:       "ES 6096s UK",
				Manufacturer:    "EATON",
				SerialNumber:    "RB76H39016",
				NominalPV:       "3600",
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

func BenchmarkValidateInverterInfo(b *testing.B) {

	i := info{
		PhaseNumber:     "1",
		VARating:        "6000",
		FirmwareVersion: "00.01.0001",
		ModelName:       "ES 6096s UK",
		Manufacturer:    "EATON",
		SerialNumber:    "RB76H39016",
		NominalPV:       "360.0",
	}

	for j := 0; j < b.N; j++ {
		i.validate()
	}
}
