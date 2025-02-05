package util

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

func ConvertToStringArray(array []interface{}) []string {
	var values []string
	for _, v := range array {
		if str, ok := v.(string); ok {
			values = append(values, str)
		} else {
			logrus.Warnf("not a string type: %v", v)
		}
	}
	return values
}
func ConvertToIntArray(array []interface{}) []int32 {
	var values []int32
	for _, v := range array {
		if value, ok := v.(int64); ok {
			values = append(values, int32(value))
		} else {
			logrus.Warnf("not a int type: %v", v)
		}
	}
	return values
}

func ConvertIntArrayToString(array []int32, separator string) string {
	str := ""
	for _, v := range array {
		str += separator + strconv.Itoa(int(v))
	}
	return str[1:]
}
