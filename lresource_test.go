package lresource

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-09-20

import (
	"testing"
)

func TestResources(t *testing.T) {

	src := `MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OD
kwMTIzNDU2Nzg5MDEyMzQ1Njc4OTA=`

	Add("test", "text/plain", 12312313, src)

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
}
