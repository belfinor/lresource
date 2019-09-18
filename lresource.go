package lresource

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-09-18

import (
	"encoding/base64"
	"errors"
)

var (
	data map[string][]byte
)

func init() {
	data = map[string][]byte{}
}

func Add(name, content string) {

	cont, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		panic(err)
	}

	data[name] = cont
}

func Get(name string) ([]byte, error) {

	val, has := data[name]
	if !has {
		return nil, errors.New("data not found")
	}

	return val, nil
}
