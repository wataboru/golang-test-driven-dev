package convert

import (
	"errors"
	"strconv"
)

func ToString(v interface{}) (string, error) {
	switch vt := v.(type) {
	case string:
		return vt, nil
	case int:
		return strconv.Itoa(vt), nil
	case bool:
		return strconv.FormatBool(vt), nil
	default:
		return "", errors.New("Error")
	}

}
