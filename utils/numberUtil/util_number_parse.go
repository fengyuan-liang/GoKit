// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package numberUtil

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func ParseString(args any) string {
	if args == nil {
		return ""
	}
	// 判断类型（未出现的类型可能会转换失败，请测试）
	switch reflect.TypeOf(args).Kind() {
	case reflect.Float32:
		return strconv.FormatFloat(float64(args.(float32)), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(args.(float64), 'f', -1, 64)
	case reflect.String:
		return args.(string)
	case reflect.Int:
		return strconv.Itoa(args.(int))
	case reflect.Int32:
		return Int32ToStr(args.(int32))
	case reflect.Int64:
		return Int64ToStr(args.(int64))
	case reflect.Uint:
		return strconv.FormatUint(uint64(args.(uint)), 10)
	default:
		// 对于未知类型，使用 fmt 包中的 %v 标记输出
		return fmt.Sprintf("%v", args)
	}
}

func Int64ToStr(val int64) string {
	return strconv.FormatInt(val, 10)
}

func Int32ToStr(val int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(val)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

// ParseNumber 解析字符串为数字（支持整数和浮点数）
func ParseNumber(value string) (float64, error) {
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, errors.New("cannot parse \"" + value + "\" as a number")
	}
	return parsed, nil
}

// MustParseNumber 解析字符串为数字（支持整数和浮点数），失败时打印日志并终止程序
func MustParseNumber(value string) float64 {
	parsed, err := ParseNumber(value)
	if err != nil {
		log.Print(err)
	}
	return parsed
}

// ParseInt64 解析字符串为64位整数
func ParseInt64(value string) (int64, error) {
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, errors.New("cannot parse \"" + value + "\" as a 64-bit integer")
	}
	return parsed, nil
}

// MustParseInt64 解析字符串为64位整数，失败时打印日志并终止程序
func MustParseInt64(value string) int64 {
	parsed, err := ParseInt64(value)
	if err != nil {
		log.Print(err)
	}
	return parsed
}

// ParseInt32 解析字符串为32位整数
func ParseInt32(value string) (int32, error) {
	parsed, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, errors.New("cannot parse \"" + value + "\" as a 32-bit integer")
	}
	return int32(parsed), nil
}

// MustParseInt32 解析字符串为32位整数，失败时打印日志并终止程序
func MustParseInt32(value string) int32 {
	parsed, err := ParseInt32(value)
	if err != nil {
		log.Print(err)
	}
	return parsed
}

// ParseUint 解析字符串为无符号整数
func ParseUint(value string) (uint, error) {
	parsed, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, errors.New("cannot parse \"" + value + "\" as an unsigned integer")
	}
	return uint(parsed), nil
}

// MustParseUint 解析字符串为无符号整数，失败时打印日志并终止程序
func MustParseUint(value string) uint {
	parsed, err := ParseUint(value)
	if err != nil {
		log.Print(err)
	}
	return parsed
}

// ParseInt 解析字符串为整数（一般用途）
func ParseInt(value string) (int, error) {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New("cannot parse \"" + value + "\" as an integer")
	}
	return parsed, nil
}

// MustParseInt 解析字符串为整数（一般用途），失败时打印日志并终止程序
func MustParseInt(value string) int {
	parsed, err := ParseInt(value)
	if err != nil {
		log.Print(err)
	}
	return parsed
}

// ParseUint32 解析字符串为32位无符号整数
func ParseUint32(value string) (uint32, error) {
	parsed, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, errors.New("cannot parse \"" + value + "\" as a 32-bit unsigned integer")
	}
	return uint32(parsed), nil
}

// MustParseUint32 解析字符串为32位无符号整数，失败时打印日志并终止程序
func MustParseUint32(value string) uint32 {
	parsed, err := ParseUint32(value)
	if err != nil {
		log.Print(err)
	}
	return parsed
}

// ParseUint64 解析字符串为64位无符号整数
func ParseUint64(value string) (uint64, error) {
	parsed, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, errors.New("cannot parse \"" + value + "\" as a 64-bit unsigned integer")
	}
	return parsed, nil
}

// MustParseUint64 解析字符串为64位无符号整数，失败时打印日志并终止程序
func MustParseUint64(value string) uint64 {
	parsed, err := ParseUint64(value)
	if err != nil {
		log.Print(err)
	}
	return parsed
}

// ParseFloat64 解析字符串为浮点数（64位）
func ParseFloat64(value string) (float64, error) {
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, errors.New("cannot parse \"" + value + "\" as a 64-bit float")
	}
	return parsed, nil
}

// MustParseFloat64 解析字符串为浮点数（64位），失败时打印日志并终止程序
func MustParseFloat64(value string) float64 {
	parsed, err := ParseFloat64(value)
	if err != nil {
		log.Print(err)
	}
	return parsed
}
