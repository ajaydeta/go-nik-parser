# go-nik-parser

This is **NIK** (nomor induk kependudukan) parser in `GoLang`

This module is in developing, do not use in porduction

## Example

```go
package main

import (
	"fmt"
	gonikparser "github.com/ajaydeta/go-nik-parser"
	"log"
)

func main() {
	nik, err := gonikparser.NikParse("3671012001010001")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(nik.GetProvinsi())
	fmt.Println(nik.GetKabKot())
	fmt.Println(nik.GetKecamatan())
	fmt.Println(nik.GetUnicode())
	fmt.Println(nik.GetPostalCode())
	fmt.Println(nik.GetBornDay())
	fmt.Println(nik.GetBornMonth())
	fmt.Println(nik.GetBornYear())
	fmt.Println(nik.GetBirdDay())
}

```

