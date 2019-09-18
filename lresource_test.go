package lresource

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-09-18

import (
	"testing"
)

func TestResources(t *testing.T) {

	src := `MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OD
kwMTIzNDU2Nzg5MDEyMzQ1Njc4OTA=`

	Add("test", src)

	wait := "12345678901234567890123456789012345678901234567890"

	res, err := Get("test")
	if err != nil {
		t.Fatal("Get not work")
	}

	if string(res) != wait {
		t.Fatal("Get return wrong data")
	}

	_, err = Get("test1")
	if err == nil {
		t.Fatal("return value for unknown source")
	}
}
