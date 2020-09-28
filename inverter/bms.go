package inverter

import (
	"fmt"
)

//bms represents the battery
type bms struct {
	Info batInfo `json:"info"`
	Data batData `json:"data"`
}

//batInfo represents the battery description
type batInfo struct {
	ModelType       string `json:"modelType"`
	FirmwareVersion string `json:"fwVersion"`
	SerialNumber    string `json:"serialNumber"`
}

//batData represents the battery data
type batData struct {
	HighestCellVoltage string `json:"highestCellVoltage"`
	LowestCellVoltage  string `json:"lowestCellVoltage"`
	SysMaxTemperature  string `json:"sysMaxTemperature"`
	SysAvgTemperature  string `json:"sysAvgTemperature"`
	SysMinTemperature  string `json:"sysMinTemperature"`
}

const (
	voltFactor float64 = 1 / 1.5
	tempFactor float64 = 1 / 0.1
)

//GetInfo returns a byte slice with battery info
func (b *batInfo) GetInfo() ([]byte, error) {
	info := strToBytes(b.ModelType, 1, paddingRight)

	fw, err := numToBytes(b.FirmwareVersion, 1, 2)
	if err != nil {
		return nil, err
	}
	info = append(info, fw...)

	sn, err := numToBytes(b.SerialNumber, 1, 4)
	if err != nil {
		return nil, err
	}
	info = append(info, sn...)

	return info, nil
}

//GetData returns a byte slice with battery data
func (b *batData) GetData() ([]byte, error) {

	bs, err := numToBytes(b.HighestCellVoltage, voltFactor, 2)
	if err != nil {
		return nil, err
	}
	data := bs

	bs, err = numToBytes(b.LowestCellVoltage, voltFactor, 2)
	if err != nil {
		return nil, err
	}
	data = append(data, bs...)

	bs, err = numToBytes(b.SysMaxTemperature, tempFactor, 2)
	if err != nil {
		return nil, err
	}
	data = append(data, bs...)

	bs, err = numToBytes(b.SysAvgTemperature, tempFactor, 2)
	if err != nil {
		return nil, err
	}
	data = append(data, bs...)

	bs, err = numToBytes(b.SysMinTemperature, tempFactor, 2)
	if err != nil {
		return nil, err
	}
	data = append(data, bs...)

	return data, nil
}

//validate does battery validation (data and info)
func (b *bms) validate() error {
	err := b.Info.validate()
	if err != nil {
		return fmt.Errorf("Battery Info error: %v", err)
	}
	err = b.Data.validate()
	if err != nil {
		return fmt.Errorf("Battery Data error: %v", err)
	}
	return nil
}

//validate does validation for battery info
func (b *batInfo) validate() error {
	if b.ModelType != "C" && b.ModelType != "R" {
		return fmt.Errorf("Invalid Model Type: %v", b.ModelType)
	}

	if !checkSize(b.FirmwareVersion, 1, 2) {
		return fmt.Errorf("Invalid Firmware Version: %v", b.FirmwareVersion)
	}

	if !checkSize(b.SerialNumber, 1, 4) {
		return fmt.Errorf("Invalid Serial Number: %v", b.SerialNumber)
	}

	return nil
}

//validate does validation for battery data
func (b *batData) validate() error {
	if !checkSize(b.HighestCellVoltage, voltFactor, 2) {
		return fmt.Errorf("Invalid Highest Cell Voltage: %v", b.HighestCellVoltage)
	}

	if !checkSize(b.LowestCellVoltage, voltFactor, 2) {
		return fmt.Errorf("Invalid Lowest Cell Voltage: %v", b.LowestCellVoltage)
	}

	if !checkSize(b.SysMaxTemperature, tempFactor, 2) {
		return fmt.Errorf("Invalid System Max Temperature: %v", b.SysMaxTemperature)
	}

	if !checkSize(b.SysAvgTemperature, tempFactor, 2) {
		return fmt.Errorf("Invalid System Average Temperature: %v", b.SysAvgTemperature)
	}

	if !checkSize(b.SysMinTemperature, tempFactor, 2) {
		return fmt.Errorf("Invalid System Min Temperature: %v", b.SysMinTemperature)
	}

	return nil
}
