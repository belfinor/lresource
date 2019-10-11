package lresource

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
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
	Data        []byte
}

var (
	data map[string]*FileData
)

func init() {
	data = map[string]*FileData{}
}

func Add(name, contentType string, arch bool, content string) {

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
		Data:        cont,
	}
}

func Get(name string) *FileData {

	val, _ := data[name]
	return val
}
