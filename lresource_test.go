package lresource

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.004
// @date    2019-11-09

import (
	"testing"
)

func TestResources(t *testing.T) {

	src := `MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OD
kwMTIzNDU2Nzg5MDEyMzQ1Njc4OTA=`

	Add("test", "text/plain", false, src)

	wait := "12345678901234567890123456789012345678901234567890"

	res := Get("test")
	if res == nil {
		t.Fatal("Get not work")
	}

	if string(res.Data) != wait {
		t.Fatal("Get return wrong data")
	}

	res = Get("test1")
	if res != nil {
		t.Fatal("return value for unknown source")
	}

	Delete("test1")

	if Get("test1") != nil {
		t.Fatal("Delete not work")
	}

	Delete("test1")
}
