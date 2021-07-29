package go_nik_parser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Nik struct {
	nikNum       string
	provinsi     string
	kota         string
	kecamatan    string
	tanggalLahir string
	isFemale     bool
	unicode      string

	provinsiMap  map[string]interface{}
	kabkotMap    map[string]interface{}
	kecamatanMap map[string]interface{}
}

func NikParse(nik string) (Nik, error) {
	n := &Nik{}
	n.nikNum = nik

	err := n.readFileWilayah()
	if err != nil {
		return Nik{}, err
	}

	if !n.validate() {
		err = errors.New("NIK number not valid")
		return Nik{}, err
	}

	n.provinsi = n.provinsiMap[nik[0:2]].(string)
	n.kota = n.kabkotMap[nik[0:4]].(string)
	n.kecamatan = n.kecamatanMap[nik[0:6]].(string)

	n.unicode = nik[12:16]
	n.tanggalLahir = nik[6:12]

	n.isFemale = len(nik[6:8]) > 40

	return *n, nil
}

func (n Nik) GetProvinsi() string {
	return n.provinsi
}

func (n Nik) GetKabKot() string {
	return n.kota
}

func (n Nik) GetKecamatan() string {
	a := strings.Split(n.kecamatan, "--")
	return strings.Trim(a[0], " ")
}

func (n Nik) GetPostalCode() string {
	a := strings.Split(n.kecamatan, "--")
	return strings.Trim(a[1], " ")
}

func (n Nik) GetBornDay() int {
	atoi, _ := strconv.Atoi(n.tanggalLahir[0:2])
	if n.isFemale {
		atoi -= 40
	}
	return atoi
}

func (n Nik) GetBornMonth() int {
	atoi, _ := strconv.Atoi(n.tanggalLahir[2:4])
	return atoi
}

func (n Nik) GetBornYear() int {
	//t, _ := strconv.Atoi( strings.Split(time.Now().Format("01/02/06"), "/")[2])
	t, _ := strconv.Atoi(time.Now().Format("06"))
	nikT, _ := strconv.Atoi(n.tanggalLahir[4:6])
	if nikT < t {
		nikT += 2000
	} else {
		nikT += 1900
	}
	return nikT
}

func (n Nik) GetBirdDay() time.Time {
	l := strconv.Itoa(n.GetBornDay()) + " " + convBulan(n.GetBornMonth()) + " " + strconv.Itoa(n.GetBornYear())

	parse, err := time.Parse("02 January 2006", l)
	if err != nil {
		log.Println(err)
		return time.Time{}
	}
	return parse
}

func (n Nik) GetUnicode() string {
	return n.unicode
}

func (n *Nik) validate() bool {
	if len(n.nikNum) != 16 {
		return false
	}

	provinsi := n.provinsiMap[n.nikNum[0:2]]
	kabkot := n.kabkotMap[n.nikNum[0:4]]
	kec := n.kecamatanMap[n.nikNum[0:6]]

	return provinsi != nil && kabkot != nil && kec != nil
}

func (n *Nik) readFileWilayah() error {
	file, err := ioutil.ReadFile("wilayah.json")
	if err != nil {
		return err
	}
	var a map[string]interface{}
	err = json.Unmarshal(file, &a)
	if err != nil {
		return err
	}

	prov, err := marshaler(a["provinsi"])
	if err != nil {
		return err
	}
	n.provinsiMap = prov

	kabkot, err := marshaler(a["kabkot"])
	if err != nil {
		return err
	}
	n.kabkotMap = kabkot

	kec, err := marshaler(a["kecamatan"])
	if err != nil {
		return err
	}
	n.kecamatanMap = kec

	return nil
}
