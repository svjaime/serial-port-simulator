package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/viper"
	"ubiwhere.com/serial-port-simulator/inverter"
)

//Load will load configuration from file and return a config instance
func Load() (*Config, error) {

	viper.AddConfigPath("files")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config: %v", err)
	}

	c := &Config{}

	err = viper.Unmarshal(c)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	err = c.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	err = save(c)
	if err != nil {
		return nil, fmt.Errorf("problem saving config: %v", err)
	}

	err = loadDataFile(c)
	if err != nil {
		return nil, fmt.Errorf("problem storing data file: %v", err)
	}

	fmt.Print(c.toString())
	return c, nil
}

//save stores config in memory
func save(c *Config) error {

	if _, found := configMap[c.Name]; found {
		return fmt.Errorf("found duplicate: %v", c.Name)
	}

	configMap[c.Name] = c
	Active = c

	return nil
}

//loadDataFile stores a *os.File in memory, from the loaded config
func loadDataFile(c *Config) error {

	file, err := os.OpenFile(c.DataFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("Failed to open datafile: %v", err)
	}

	DataFile = file

	content := make([]byte, 1024)

	readBytes, err := DataFile.Read(content)
	if err != nil {
		if err != io.EOF {
			return fmt.Errorf("error reading data file: %v", err)
		}
	}

	err = json.Unmarshal(content[:readBytes], &inverter.Active)
	if err != nil {
		return fmt.Errorf("json unmarshal error: %v", err)
	}

	err = inverter.Active.Validate()
	if err != nil {
		return fmt.Errorf("inverter validation error: %v", err)
	}

	//debug
	fmt.Println(inverter.Active)

	return nil
}
