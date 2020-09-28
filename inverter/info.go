package inverter

import (
	"fmt"
)

//info represents the inverter information
type info struct {
	PhaseNumber     string `json:"phaseNumber"`
	VARating        string `json:"vaRating"`
	FirmwareVersion string `json:"fwVersion"`
	ModelName       string `json:"modelName"`
	Manufacturer    string `json:"manufacturer"`
	SerialNumber    string `json:"serialNumber"`
	NominalPV       string `json:"nominalPv"`
}

//GetInfo returns a byte slice with inverter's info
func (i *info) GetInfo() ([]byte, error) {

	info := strToBytes(i.PhaseNumber, 1, paddingRight)

	info = append(info, strToBytes(i.VARating, 6, paddingLeft)...)
	info = append(info, strToBytes(i.FirmwareVersion, 10, paddingLeft)...)
	info = append(info, strToBytes(i.ModelName, 16, paddingRight)...)
	info = append(info, strToBytes(i.Manufacturer, 16, paddingRight)...)
	info = append(info, strToBytes(i.SerialNumber, 16, paddingRight)...)

	s, err := numerics(i.NominalPV, 10)
	if err != nil {
		return nil, err
	}
	info = append(info, strToBytes(s, 4, paddingRight)...)

	return info, nil
}

//validate does validation for the inverter's info fields
func (i *info) validate() error {
	if len(i.PhaseNumber) != 1 {
		return fmt.Errorf("Invalid Phase Number: %v", i.PhaseNumber)
	}

	if len(i.VARating) <= 0 || len(i.VARating) > 6 {
		return fmt.Errorf("Invalid VA Rating: %v", i.VARating)
	}

	if len(i.FirmwareVersion) <= 0 || len(i.FirmwareVersion) > 10 {
		return fmt.Errorf("Invalid Firmware Version: %v", i.FirmwareVersion)
	}

	if len(i.ModelName) <= 0 || len(i.ModelName) > 16 {
		return fmt.Errorf("Invalid Model Name: %v", i.ModelName)
	}

	if len(i.Manufacturer) <= 0 || len(i.Manufacturer) > 16 {
		return fmt.Errorf("Invalid Manufacturer: %v", i.Manufacturer)
	}

	if len(i.SerialNumber) <= 0 || len(i.SerialNumber) > 16 {
		return fmt.Errorf("Invalid Serial Number: %v", i.SerialNumber)
	}

	s, err := numerics(i.NominalPV, 10)
	if err != nil {
		return err
	}

	if len(s) <= 0 || len(s) > 4 {
		return fmt.Errorf("Invalid Nominal PV: %v", i.NominalPV)
	}

	return nil
}
