//Package config represents the configuration for each communication interface
package config

import (
	"fmt"
	"os"
)

//Config defines the default structure of the configuration file
type Config struct {
	Name         string `json:"name"`
	Path         string `json:"path"`
	DataFilePath string `json:"dataFilePath"`
	Protocol     string `json:"protocol"`
	BaudRate     int32  `json:"baudRate"` //optional, defaults to 1200
}

//communication protocols (enum)
const (
	inv string = "INVERTER"
)

//Supported communication protocols
var protocols = []string{inv}

//Supported Baud Rates
var baudRates = []int32{1200, 2400, 4800, 9600, 14400, 19200, 28800, 38400, 57600, 76800, 115200, 230400}

//configMap stores loaded configs
var configMap = make(map[string]*Config)

//Active represents the active configuartion (last loaded)
var Active *Config

//DataFile points to file open on load from current config
var DataFile *os.File

//print prints config details to the console - helper function
func (c *Config) toString() string {
	s := fmt.Sprintf("Loaded configuration:\n\tName: %v\n\tPath: %v\n\tDataFilePath: %v\n\tProtocol: %v\n\tBaudrate: %v\n",
		c.Name, c.Path, c.DataFilePath, c.Protocol, c.BaudRate)
	return s
}
