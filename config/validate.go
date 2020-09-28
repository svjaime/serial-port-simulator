package config

import (
	"fmt"
	"os"
	"regexp/syntax"
	"strings"
)

//Validate does config validation, and stores it in memory
func (c *Config) Validate() error {

	err := validateName(c)
	if err != nil {
		return fmt.Errorf("invalid name: %v", err)
	}

	err = validatePath(c)
	if err != nil {
		return fmt.Errorf("invalid path: %v", err)
	}

	err = validateDataFilePath(c)
	if err != nil {
		return fmt.Errorf("invalid data file path: %v", err)
	}

	err = validateProtocol(c)
	if err != nil {
		return fmt.Errorf("invalid protocol: %v", err)
	}

	err = validateBaudRate(c)
	if err != nil {
		return fmt.Errorf("invalid baudrate: %v", err)
	}

	return nil
}

//validateName does name validation
func validateName(c *Config) error {

	if len(c.Name) <= 3 || len(c.Name) >= 20 {
		return fmt.Errorf("%v", c.Name)
	}

	for _, char := range c.Name {
		if !syntax.IsWordChar(char) {
			return fmt.Errorf("%v", c.Name)
		}
	}

	c.Name = strings.ToLower(c.Name)
	return nil
}

//validatePath does path validation
func validatePath(c *Config) error {

	if len(c.Path) == 0 {
		return fmt.Errorf("%v", c.Path)
	}
	return nil
}

//validateDataFilePath does data file path validation
func validateDataFilePath(c *Config) error {

	if len(c.DataFilePath) == 0 {
		return fmt.Errorf("data file path is empty")
	}

	//Checking if the given file exists or not
	file, err := os.Stat(c.DataFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}

	if file.IsDir() {
		return fmt.Errorf("Path points to a directory: %v", c.DataFilePath)
	}

	return nil
}

//validateProtocol does protocol validation
func validateProtocol(c *Config) error {

	for _, val := range protocols {
		if val == c.Protocol {
			return nil
		}
	}
	return fmt.Errorf("%v", c.Protocol)
}

//validateBaudRate does baudrate validation
func validateBaudRate(c *Config) error {

	for _, val := range baudRates {
		if val == c.BaudRate {
			return nil
		}
	}

	//defaults to 1200
	if c.BaudRate == 0 {
		c.BaudRate = baudRates[0]
		return nil
	}

	return fmt.Errorf("%v", c.BaudRate)
}
