package data_parser

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadCSVFile reads a CSV file and returns a slice of maps where each map represents a row in the CSV file.
func ReadCSVFile(filename string) ([]map[string]string, error) {
	file, err := os.Open(filename)
	if err!= nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err!= nil {
		return nil, err
	}

	var result []map[string]string
	for _, record := range records {
		result = append(result, map[string]string{})
		for i, field := range record {
			result[len(result)-1][strconv.Itoa(i)] = field
		}
	}

	return result, nil
}

// ReadCSVFileWithHeader reads a CSV file and returns a slice of maps where each map represents a row in the CSV file.
// The keys in each map are the column names.
func ReadCSVFileWithHeader(filename string) ([]map[string]string, error) {
	file, err := os.Open(filename)
	if err!= nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err!= nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	var result []map[string]string
	for _, record := range records[1:] {
		result = append(result, map[string]string{})
		for i, field := range record {
			result[len(result)-1][strconv.Itoa(i)] = field
		}
	}

	return result, nil
}

// SplitString splits a string into a slice of substrings based on a delimiter.
func SplitString(s string, delimiter string) []string {
	return strings.Split(s, delimiter)
}

// GetInt attempts to parse a string into an integer.
func GetInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// GetFloat attempts to parse a string into a float64.
func GetFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// GetBool attempts to parse a string into a boolean.
func GetBool(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean value: %s", s)
	}
}