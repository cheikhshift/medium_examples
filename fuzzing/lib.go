package fuzzing

import (
	"strings"
	"strconv"
	"errors"
)

func ParseString(s string) error {
	parts := strings.Split(s, ",")

	if(len(parts) < 3){
		return errors.New("Invalid string")
	}

	if parts[0] == "" {
		parts[0] = "-1"
	}

	_, err := strconv.Atoi(parts[0])

	if err != nil {
		return err
	}

	parts[2] += ", Senegal"

	return nil
}