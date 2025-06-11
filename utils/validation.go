package utils

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strconv"
)

func ValidationRequied(fieldName, value string) error {
	if value == "" {
		return fmt.Errorf("%s bat buoc phai nhap", fieldName)
	}
	return nil
}
func ValidationStringLength(fieldName, value string, min, max int) error {
	if len(value) < min || len(value) > max {
		return fmt.Errorf("%s phai co do dai tu %d toi %d", fieldName, min, max)
	}
	return nil
}
func ValidationRegex(fieldName, value, errorMessage string, regex *regexp.Regexp) error {
	if !regex.MatchString(value) {
		return fmt.Errorf("%s: %s ", fieldName, errorMessage)
	}
	return nil
}
func ValidationPositiveInteger(fieldName string, value string) (int, error) {
	if value == "" {
		return 0, fmt.Errorf("%s bat buoc phai nhap", fieldName)
	}
	ValueInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s phai la so ", fieldName)
	}
	if ValueInt <= 0 {
		return 0, fmt.Errorf("%s phai la so nguyen duong", fieldName)
	}
	return ValueInt, nil
}
func ValidationUUID(fieldName, value string) (uuid.UUID, error) {
	uid, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s khong hop le: %v", fieldName, err)
	}
	return uid, nil
}
func ValidationInList(fieldName, value string, allowed map[string]bool) error {
	if _, ok := allowed[value]; !ok {
		return fmt.Errorf("%s khong hop le, phai la mot trong %v", fieldName, keys(allowed))
	}
	return nil
}
func keys(m map[string]bool) []string {
	var keys []string // tạo một slice để chứa các khóa
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
