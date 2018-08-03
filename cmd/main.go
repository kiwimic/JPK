package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type JPK struct {
	Naglowek          Naglowek          `xml:"Naglowek"`
	Faktura           []Faktura         `xml:"Faktura"`
	FakturaCtrl       FakturaWierszCtrl `xml:"FakturaCtrl"`
	FakturaWiersz     []FakturaWiersz   `xml:"FakturaWiersz"`
	FakturaWierszCtrl FakturaWierszCtrl `xml:"FakturaWierszCtrl"`
}

type Naglowek struct {
	KodFormularza      string `xml:"KodFormularza"`
	WariantFormularza  string `xml:"WariantFormularza"`
	CelZlozenia        string `xml:"CelZlozenia"`
	DataWytworzeniaJPK string `xml:"DataWytworzeniaJPK"`
	DataOd             string `xml:"DataOd"`
	DataDo             string `xml:"DataDo"`
	DomyslnyKodWaluty  string `xml:"DomyslnyKodWaluty"`
	KodUrzedu          string `xml:"KodUrzedu"`
}

type Faktura struct {
	P_1                string `xml:"P_1"`
	P_2A               string `xml:"P_2A"`
	P_3A               string `xml:"P_3A"`
	P_3B               string `xml:"P_3B"`
	P_3C               string `xml:"P_3C"`
	P_3D               string `xml:"P_3D"`
	P_4A               string `xml:"P_4A"`
	P_4B               string `xml:"P_4B"`
	P_5A               string `xml:"P_5A"`
	P_5B               string `xml:"P_5B"`
	P_6                string `xml:"P_6"`
	P_13_1             string `xml:"P_13_1"`
	P_14_1             string `xml:"P_14_1"`
	P_13_2             string `xml:"P_13_2"`
	P_14_2             string `xml:"P_14_2"`
	P_13_3             string `xml:"P_13_3"`
	P_14_3             string `xml:"P_14_3"`
	P_13_4             string `xml:"P_13_4"`
	P_14_4             string `xml:"P_14_4"`
	P_13_5             string `xml:"P_13_5"`
	P_14_5             string `xml:"P_14_5"`
	P_13_6             string `xml:"P_13_6"`
	P_13_7             string `xml:"P_13_7"`
	P_15               string `xml:"P_15"`
	P_16               string `xml:"P_16"`
	P_17               string `xml:"P_17"`
	P_18               string `xml:"P_18"`
	P_19               string `xml:"P_19"`
	P_19A              string `xml:"P_19A"`
	P_19B              string `xml:"P_19B"`
	P_19C              string `xml:"P_19C"`
	P_20               string `xml:"P_20"`
	P_20A              string `xml:"P_20A"`
	P_20B              string `xml:"P_20B"`
	P_21               string `xml:"P_21"`
	P_21A              string `xml:"P_21A"`
	P_21B              string `xml:"P_21B"`
	P_21C              string `xml:"P_21C"`
	P_22A              string `xml:"P_22A"`
	P_22B              string `xml:"P_22B"`
	P_22C              string `xml:"P_22C"`
	P_23               string `xml:"P_23"`
	P_106E_2           string `xml:"P_106E_2"`
	P_106E_3           string `xml:"P_106E_3"`
	P_106E_3A          string `xml:"P_106E_3A"`
	RodzajFaktury      string `xml:"RodzajFaktury"`
	PrzyczynaKorekty   string `xml:"PrzyczynaKorekty"`
	NrFaKorygowanej    string `xml:"NrFaKorygowanej"`
	OkresFaKorygowanej string `xml:"OkresFaKorygowanej"`
	ZALZaplata         string `xml:"ZALZaplata"`
	ZALPodatek         string `xml:"ZALPodatek"`
}

type FakturaCtrl struct {
	LiczbaFaktur  string `xml:"LiczbaFaktur"`
	WartoscFaktur string `xml:"WartoscFaktur"`
}

type FakturaWiersz struct {
	P_2B  string `xml:"P_2B"`
	P_7   string `xml:"P_7"`
	P_8A  string `xml:"P_8A"`
	P_8B  string `xml:"P_8B"`
	P_9A  string `xml:"P_9A"`
	P_9B  string `xml:"P_9B"`
	P_10  string `xml:"P_10"`
	P_11  string `xml:"P_11"`
	P_11A string `xml:"P_11A"`
	P_12  string `xml:"P_12"`
}

type FakturaWierszCtrl struct {
	LiczbaWierszyFaktur  string `xml:"LiczbaWierszyFaktur"`
	WartoscWierszyFaktur string `xml:"WartoscWierszyFaktur"`
}

func createRowFaktura(FakturaStruct Faktura) string {
	str := ""

	strSlice := []string{FakturaStruct.P_1,
		FakturaStruct.P_2A,
		FakturaStruct.P_3A,
		FakturaStruct.P_3B,
		FakturaStruct.P_3C,
		FakturaStruct.P_3D,
		FakturaStruct.P_4A,
		FakturaStruct.P_4B,
		FakturaStruct.P_5A,
		FakturaStruct.P_5B,
		FakturaStruct.P_6,
		FakturaStruct.P_13_1,
		FakturaStruct.P_14_1,
		FakturaStruct.P_13_2,
		FakturaStruct.P_14_2,
		FakturaStruct.P_13_3,
		FakturaStruct.P_14_3,
		FakturaStruct.P_13_4,
		FakturaStruct.P_14_4,
		FakturaStruct.P_13_5,
		FakturaStruct.P_14_5,
		FakturaStruct.P_13_6,
		FakturaStruct.P_13_7,
		FakturaStruct.P_15,
		FakturaStruct.P_16,
		FakturaStruct.P_17,
		FakturaStruct.P_18,
		FakturaStruct.P_19,
		FakturaStruct.P_19A,
		FakturaStruct.P_19B,
		FakturaStruct.P_19C,
		FakturaStruct.P_20,
		FakturaStruct.P_20A,
		FakturaStruct.P_20B,
		FakturaStruct.P_21,
		FakturaStruct.P_21A,
		FakturaStruct.P_21B,
		FakturaStruct.P_21C,
		FakturaStruct.P_22A,
		FakturaStruct.P_22B,
		FakturaStruct.P_22C,
		FakturaStruct.P_23,
		FakturaStruct.P_106E_2,
		FakturaStruct.P_106E_3,
		FakturaStruct.P_106E_3A,
		FakturaStruct.RodzajFaktury,
		FakturaStruct.PrzyczynaKorekty,
		FakturaStruct.NrFaKorygowanej,
		FakturaStruct.OkresFaKorygowanej,
		FakturaStruct.ZALZaplata,
		FakturaStruct.ZALPodatek}

	joined := strings.Join(strSlice, "\";\"")
	joined = "\"" + joined + "\""
	str = str + joined + "\n"

	return str
}

func createRowFakturaWiersz(FakturaStruct FakturaWiersz) string {
	str := ""

	strSlice := []string{FakturaStruct.P_2B,
		FakturaStruct.P_7,
		FakturaStruct.P_8A,
		FakturaStruct.P_8B,
		FakturaStruct.P_9A,
		FakturaStruct.P_9B,
		FakturaStruct.P_10,
		FakturaStruct.P_11,
		FakturaStruct.P_11A,
		FakturaStruct.P_12}

	joined := strings.Join(strSlice, "\";\"")
	joined = "\"" + joined + "\""
	str = str + joined + "\n"

	return str
}

func (t JPK) createCSVFromRowsFaktura() {
	str := ""
	for i := range t.Faktura {
		ss := createRowFaktura(t.Faktura[i])
		str = str + ss + "\n"
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			writeToCSV(str, "JPK_FA_Faktura.txt")
			str = ""
		}
	}
	writeToCSV(str, "JPK_FA_Faktura.txt")
}

func (t JPK) createCSVFromRowsFakturaWiersz() {
	str := ""
	for i := range t.FakturaWiersz {
		ss := createRowFakturaWiersz(t.FakturaWiersz[i])
		str = str + ss + "\n"
		if i%500 == 0 {
			//fmt.Println("FakturaWiersz: ", i)
			writeToCSV(str, "C:\\Users\\msiwik\\Documents\\JPK_FA_FakturaWiersz.txt")
			str = ""
		}
	}
	writeToCSV(str, "C:\\Users\\msiwik\\Documents\\JPK_FA_FakturaWiersz.txt")
}

func exportCSV(st, name string) {
	dataBytes := []byte(st)
	// Use WriteFile to create a file with byte data.
	ioutil.WriteFile(name, dataBytes, 0)
}

func writeToCSV(st, filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(f, "%s", st)

	f.Close()
}
func main() {
	pathToFile := flag.String("path", "data/xml_duzy.xml", "a string with path to JPK xml file")
	flag.Parse()
	st := time.Now()
	fmt.Println("Zaczynami: ", st)
	xmlFile, err := os.Open(*pathToFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)
	v2 := time.Now()
	fmt.Println("Wczytano plik: ", v2)
	var JPK_FA JPK
	xml.Unmarshal(byteValue, &JPK_FA)
	v3 := time.Now()
	fmt.Println("Zakończono parsowanie: ", v3)
	//str := createCSVFaktura(JPK_FA.Faktura)
	JPK_FA.createCSVFromRowsFaktura()

	fmt.Println("Stworzono CSV faktura: ", time.Now())
	//exportCSV(csvFaktura, "JPK_FA_Faktura.txt")

	fmt.Println("Zapisano CSV faktura: ", time.Now())

	JPK_FA.createCSVFromRowsFakturaWiersz()

	fmt.Println("Stworzono CSV fakturaWiersz: ", time.Now())

	//exportCSV(csvFakturaWiersz, "JPK_FA_FakturaWiersz.txt")

	fmt.Println("Zapisano CSV fakturaWiersz: ", time.Now())
	end := time.Now()
	fmt.Println("Start: ", st, "\nEnd: ", end)
	fmt.Println("Wszystko trwało: ", time.Since(st))
}
