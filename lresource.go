package lresource

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2019-10-11

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
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

func Add(name, contentType string, ts int64, arch bool, content string) {

	cont, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		panic(err)
	}

	if arch {
		gz, _ := gzip.NewReader(bytes.NewReader(cont))

		data, err := ioutil.ReadAll(gz)
		if err != nil {
			panic(err)
		}

		cont = data
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
