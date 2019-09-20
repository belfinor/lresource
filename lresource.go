package lresource

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-09-20

import (
	"encoding/base64"
)

type FileData struct {
	Name        string
	ContentType string
	ModifyTime  int64
	Data        []byte
}

var (
	data map[string]*FileData
)

func init() {
	data = map[string]*FileData{}
}

func Add(name, contentType string, ts int64, content string) {

	cont, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		panic(err)
	}

	data[name] = &FileData{
		Name:        name,
		ContentType: contentType,
		ModifyTime:  ts,
		Data:        cont,
	}
}

func Get(name string) *FileData {

	val, _ := data[name]
	return val
}
