# lresource

lresource is a simple solution for bundling static assets inside of Go binaries.

# Install

```
go get github.com/belfinor/lresource
```

# Embed resource

You have directory *data*, that contains *text.txt* file (may be binary) of the following contents:

```
English texts for beginners to practice reading and comprehension online and for free. Practicing your comprehension of written English will both improve your vocabulary and understanding of grammar and word order. The texts below are designed to help you develop while giving you an instant evaluation of your progress.

Prepared by experienced English teachers, the texts, articles and conversations are brief and appropriate to your level of proficiency. Take the multiple-choice quiz following each text, and you'll get the results immediately. You will feel both challenged and accomplished! You can even download (as PDF) and print the texts and exercises. It's enjoyable, fun and free. Good luck!
```

Create *generator.go* into *data*:

```go
// +build ignore

package main

import (

	"github.com/belfinor/lresource"
)

func main() {

	err := lresource.Convert("text.txt", "text.txt.go", "data", "data/text.txt", false)
	if err != nil {
		panic(err)
	}
}
```

Create *data.go* with generate instruction:

```go
package data

//go:generate go run generator.go
```

Go to data and run generate:

```
cd data
go generate
```

And you have file *text.txt.go* like this:

```go
package data

// @comment this file was generated by lresource

import (
  "github.com/belfinor/lresource"
)

func init() {

  lresource.Add("data/text.txt", "text/plain; charset=utf-8", true, `
H4sIAAAAAAAA/1ySPY7cMAyF+z0Ft9oEmPgU+UG6KbZJKUvPNrM0qVCyPc7pA2lm
sUBa0w/fRz5901m4LFRxq4Umcxoxsyq8UDXKHmLlCHKExDpT0ETR1uxYoIVNyVRY
0QctPjkw0PWea4nTNv8/MtHhXCuU3vkHi9BodSFes9uOe263GMZNgp8dsGmClxq0
u9hEs4d1Dd6Hh3ki8wQf6HXBY6URYgcFByUUnhWp7bVAciNQwg6xTMfCApp5fyhT
UGJtqErYg2yhPsy7V3abHaUMT09XRw6ORONJuGU4QyMSfVw2xAVeLlTfpS4UvHIU
lMc9dYeXTihddXTG1GchZ7fsHCqad4dLc24q2W3i2HjnQK/hDR2xblI5C77ExVp1
fzb+S5OJ2NGWazpd49IBp20vIjSj9rCjbFIL8boiNaqcA/2y7V7QBDxaiksQgc5I
d83YGm77Ij33/2NQwg6lZIeKhUSfQqHr1++feyA7a/24SP+GGzxyQRnoZ30pBP1t
ZxgFF5o2vb+w/rp+mCWSLb49P/0LAAD///xIsyHAAgAA`)

}
```

# Access resource from code

```go
import (

	"github.com/belfinor/lresource"
	_ "github.com/login/proj/data" // load data

)

func Get() []byte {

  data := lresource.Get("data/text.txt")

	if data != nil {
		return data.Data
	}

  return nil
}
```

# If resource changed

If the *data.txt* file has changed, you go to the data directory and call *go generate*.
