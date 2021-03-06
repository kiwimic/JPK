package utils

import "strings"

//  JPK_FA ...
type JPK_FA struct {
	Naglowek          Naglowek          `xml:"Naglowek"`
	Faktura           []Faktura         `xml:"Faktura"`
	FakturaCtrl       FakturaCtrl       `xml:"FakturaCtrl"`
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

func (t JPK_FA) CreateRowFaktura(i int) string {
	str := ""

	strSlice := []string{
		t.Faktura[i].P_1,
		t.Faktura[i].P_2A,
		t.Faktura[i].P_3A,
		t.Faktura[i].P_3B,
		t.Faktura[i].P_3C,
		t.Faktura[i].P_3D,
		t.Faktura[i].P_4A,
		t.Faktura[i].P_4B,
		t.Faktura[i].P_5A,
		t.Faktura[i].P_5B,
		t.Faktura[i].P_6,
		t.Faktura[i].P_13_1,
		t.Faktura[i].P_14_1,
		t.Faktura[i].P_13_2,
		t.Faktura[i].P_14_2,
		t.Faktura[i].P_13_3,
		t.Faktura[i].P_14_3,
		t.Faktura[i].P_13_4,
		t.Faktura[i].P_14_4,
		t.Faktura[i].P_13_5,
		t.Faktura[i].P_14_5,
		t.Faktura[i].P_13_6,
		t.Faktura[i].P_13_7,
		t.Faktura[i].P_15,
		t.Faktura[i].P_16,
		t.Faktura[i].P_17,
		t.Faktura[i].P_18,
		t.Faktura[i].P_19,
		t.Faktura[i].P_19A,
		t.Faktura[i].P_19B,
		t.Faktura[i].P_19C,
		t.Faktura[i].P_20,
		t.Faktura[i].P_20A,
		t.Faktura[i].P_20B,
		t.Faktura[i].P_21,
		t.Faktura[i].P_21A,
		t.Faktura[i].P_21B,
		t.Faktura[i].P_21C,
		t.Faktura[i].P_22A,
		t.Faktura[i].P_22B,
		t.Faktura[i].P_22C,
		t.Faktura[i].P_23,
		t.Faktura[i].P_106E_2,
		t.Faktura[i].P_106E_3,
		t.Faktura[i].P_106E_3A,
		t.Faktura[i].RodzajFaktury,
		t.Faktura[i].PrzyczynaKorekty,
		t.Faktura[i].NrFaKorygowanej,
		t.Faktura[i].OkresFaKorygowanej,
		t.Faktura[i].ZALZaplata,
		t.Faktura[i].ZALPodatek}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_FA) CreateRowFakturaWiersz(i int) string {
	str := ""

	strSlice := []string{
		t.FakturaWiersz[i].P_2B,
		t.FakturaWiersz[i].P_7,
		t.FakturaWiersz[i].P_8A,
		t.FakturaWiersz[i].P_8B,
		t.FakturaWiersz[i].P_9A,
		t.FakturaWiersz[i].P_9B,
		t.FakturaWiersz[i].P_10,
		t.FakturaWiersz[i].P_11,
		t.FakturaWiersz[i].P_11A,
		t.FakturaWiersz[i].P_12}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_FA) LiczbaFaktur() int {
	return (len(t.Faktura))
}

func (t JPK_FA) LiczbaWierszyFaktur() int {
	return (len(t.FakturaWiersz))
}

func (t JPK_FA) CreateCSVFromRowsFaktura(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}

	colNamesToBind := []string{"P_1",
		"P_2A",
		"P_3A",
		"P_3B",
		"P_3C",
		"P_3D",
		"P_4A",
		"P_4B",
		"P_5A",
		"P_5B",
		"P_6",
		"P_13_1",
		"P_14_1",
		"P_13_2",
		"P_14_2",
		"P_13_3",
		"P_14_3",
		"P_13_4",
		"P_14_4",
		"P_13_5",
		"P_14_5",
		"P_13_6",
		"P_13_7",
		"P_15",
		"P_16",
		"P_17",
		"P_18",
		"P_19",
		"P_19A",
		"P_19B",
		"P_19C",
		"P_20",
		"P_20A",
		"P_20B",
		"P_21",
		"P_21A",
		"P_21B",
		"P_21C",
		"P_22A",
		"P_22B",
		"P_22C",
		"P_23",
		"P_106E_2",
		"P_106E_3",
		"P_106E_3A",
		"RodzajFaktury",
		"PrzyczynaKorekty",
		"NrFaKorygowanej",
		"OkresFaKorygowanej",
		"ZALZaplata",
		"ZALPodatek"}

	colNames := strings.Join(colNamesToBind, ";")

	for i := range t.Faktura {
		rowPaste := t.CreateRowFaktura(i)

		if i == 0 {
			str = str + colNames + "\n"
		}

		str = str + rowPaste
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			WriteToCSV(str, filename)
			str = ""
		}
	}
	WriteToCSV(str, filename)
}

func (t JPK_FA) CreateCSVFromRowsFakturaWiersz(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}
	colNamesToBind := []string{"P_2B",
		"P_7",
		"P_8A",
		"P_8B",
		"P_9A",
		"P_9B",
		"P_10",
		"P_11",
		"P_11A",
		"P_12"}

	colNames := strings.Join(colNamesToBind, ";")

	for i := range t.FakturaWiersz {
		rowPaste := t.CreateRowFakturaWiersz(i)

		if i == 0 {
			str = str + colNames + "\n"
		}
		str = str + rowPaste
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			WriteToCSV(str, filename)
			str = ""
		}
	}
	WriteToCSV(str, filename)
}

func (t JPK_FA) CreateCSVFakturaCtrl(filename string) {
	str := ""

	strSlice := []string{
		t.FakturaCtrl.LiczbaFaktur,
		t.FakturaCtrl.WartoscFaktur}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	colNames := "LiczbaFaktur;WartoscFaktur\n"
	str = str + colNames + joined + "\n"

	WriteToCSV(str, filename)
}

func (t JPK_FA) CreateCSVFakturaWierszCtrl(filename string) {
	str := ""

	strSlice := []string{
		t.FakturaWierszCtrl.LiczbaWierszyFaktur,
		t.FakturaWierszCtrl.WartoscWierszyFaktur}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")

	colNames := "LiczbaWierszyFaktur;WartoscWierszyFaktur\n"
	str = str + colNames + joined + "\n"

	WriteToCSV(str, filename)
}

func (t JPK_FA) CreateCSVNaglowek(filename string) {
	str := ""

	strSlice := []string{
		t.Naglowek.KodFormularza,
		t.Naglowek.WariantFormularza,
		t.Naglowek.CelZlozenia,
		t.Naglowek.DataWytworzeniaJPK,
		t.Naglowek.DataOd,
		t.Naglowek.DataDo,
		t.Naglowek.DomyslnyKodWaluty,
		t.Naglowek.KodUrzedu}
	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	colNames := "KodFormularza;WariantFormularza;CelZlozenia;DataWytworzeniaJPK;DataOd;DataDo;DomyslnyKodWaluty;KodUrzedu\n"
	str = str + colNames + joined + "\n"

	WriteToCSV(str, filename)
}
