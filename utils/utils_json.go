package utils

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strings"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(in interface{}) (str string, err error) {
	var (
		buf []byte
	)
	buf, err = json.Marshal(in)
	if err != nil {
		return
	}
	str = string(buf)
	return
}

func ObjToByte(in interface{}) (buf []byte, err error) {
	buf, err = json.Marshal(in)
	return
}

func ByteToObj(buf []byte, out interface{}) (err error) {
	dc := json.NewDecoder(bytes.NewReader(buf))
	dc.UseNumber()
	return dc.Decode(out)
}

func Unmarshal(in string, out interface{}) error {
	//return json.Unmarshal([]byte(in), out)
	dc := json.NewDecoder(strings.NewReader(in))
	dc.UseNumber()
	return dc.Decode(out)
}

func ObjToMap(in interface{}) map[string]interface{} {
	var (
		maps map[string]interface{}
		buf  []byte
		err  error
	)
	if buf, err = json.Marshal(in); err != nil {
		//fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(buf))
		d.UseNumber()
		if err = d.Decode(&maps); err != nil {
			//fmt.Println(err)
		} else {
			for k, v := range maps {
				maps[k] = v
			}
		}
	}
	return maps
}

func MapToObj(maps map[string]interface{}, out interface{}) error {
	buf, err := json.Marshal(maps)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, out)
	if err != nil {
		return err
	}

	return nil
}

func ObjToJsonStr(obj interface{}) (str string) {
	str = ""
	var err error
	if obj == nil {
		_ = fmt.Errorf("obj is nil")
		return
	}
	str, err = Marshal(obj)
	if err != nil {
		_ = fmt.Errorf(err.Error())
		return
	}
	return
}
