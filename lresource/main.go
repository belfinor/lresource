package main

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.004
// @date    2019-10-16

import (
	"flag"
	"fmt"
	"os"

	"github.com/belfinor/lresource"
)

func usage() {

	fmt.Println(`
Use lresource:

  -name string
        internal resource name
  -noarch
        gzip data
  -output string
        output file name
  -package string
        package name
  -source string
        source file
 `)

	os.Exit(1)
}

func main() {

	var pkg string
	var src string
	var name string
	var dest string
	var noarch bool

	flag.StringVar(&pkg, "package", "", "package name")
	flag.StringVar(&src, "source", "", "source file")
	flag.StringVar(&name, "name", "", "internal resource name")
	flag.StringVar(&dest, "output", "", "output file name")
	flag.BoolVar(&noarch, "noarch", false, "don't gzip data")

	flag.Parse()

	err := lresource.Convert(src, dest, pkg, name, noarch)
	if err != nil {
		fmt.Println(err.Error())
		usage()
	}
}
