package converterhelper

import (
	"fmt"
	"strconv"
	"unicode"
)

func BoolToInt(arg bool) int {
	if arg {
		return 1
	}
	return 0
}

func InterfaceToString(v interface{}) string {
	if val, ok := v.(string); ok {
		return val
	}

	return fmt.Sprintf("%v", v)
}

func InterfaceToInt64(v interface{}) int64 {
	if val, ok := v.(int64); ok {
		return val
	}

	i, err := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func InterfaceToInt(v interface{}) int {
	if val, ok := v.(int); ok {
		return val
	}

	i, err := strconv.Atoi(fmt.Sprintf("%v", v))
	if err != nil {
		return 0
	}
	return i

}

func IntToBool(arg int) bool {
	return arg == 1
}

func StringToInt64(arg string) int64 {
	intArg, _ := strconv.Atoi(arg)
	return int64(intArg)
}

func Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}
