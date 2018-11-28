package utils

import "strings"

// JPK_WB cały nosiciel JPK to do niego wkładamy zprasowany plik
type JPK_WB struct {
	Naglowek      NaglowekWB     `xml:"Naglowek"`
	NumerRachunku string         `xml:"NumerRachunku"`
	Salda         Salda          `xml:"Salda"`
	WyciagWiersz  []WyciagWiersz `xml:"WyciagWiersz"`
	WyciagCtrl    WyciagCtrl     `xml:"WyciagCtrl"`
}

type NaglowekWB struct {
	KodFormularza      string `xml:"KodFormularza"`
	WariantFormularza  string `xml:"WariantFormularza"`
	CelZlozenia        string `xml:"CelZlozenia"`
	DataWytworzeniaJPK string `xml:"DataWytworzeniaJPK"`
	DataOd             string `xml:"DataOd"`
	DataDo             string `xml:"DataDo"`
	DomyslnyKodWaluty  string `xml:"DomyslnyKodWaluty"`
	KodUrzedu          string `xml:"KodUrzedu"`
}

type PodmiotWB struct {
	IndetyfikatorPodmiotu string `xml:"IdentyfikatorPodmiotu"`
	AdresPodmiotu         string `xml:"AdresPodmiotu"`
}

type Salda struct {
	SaldoPoczatkowe string `xml:"SaldoPoczatkowe"`
	SaldoKoncowe    string `xml:"SaldoKoncowe"`
}

type WyciagWiersz struct {
	NumerWiersza  string `xml:"NumerWiersza"`
	DataOperacji  string `xml:"DataOperacji"`
	NazwaPodmiotu string `xml:"NazwaPodmiotu"`
	OpisOperacji  string `xml:"OpisOperacji"`
	KwotaOperacji string `xml:"KwotaOperacji"`
	SaldoOperacji string `xml:"SaldoOperacji"`
}

type WyciagCtrl struct {
	LiczbaWierszy string `xml:"LiczbaWierszy"`
	SumaObciazen  string `xml:"SumaObciazen"`
	SumaUznan     string `xml:"SumaUznan"`
}

func (t JPK_WB) createRowSalda(i int) string {
	str := ""

	strSlice := []string{
		t.Salda.SaldoPoczatkowe,
		t.Salda.SaldoKoncowe}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_WB) createRowWyciagWiersz(i int) string {
	str := ""

	strSlice := []string{
		t.WyciagWiersz[i].NumerWiersza,
		t.WyciagWiersz[i].DataOperacji,
		t.WyciagWiersz[i].NazwaPodmiotu,
		t.WyciagWiersz[i].OpisOperacji,
		t.WyciagWiersz[i].KwotaOperacji,
		t.WyciagWiersz[i].SaldoOperacji}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_WB) LiczbaWierszyWyciagow() int {
	return (len(t.WyciagWiersz))
}

func (t JPK_WB) CreateCSVFromRowsWyciagWiersz(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}

	colNames := "NumerWiersza;DataOperacji;NazwaPodmiotu;OpisOperacji;KwotaOperacji;SaldoOperacji\n"
	for i := range t.WyciagWiersz {
		rowPaste := t.createRowWyciagWiersz(i)

		if i == 0 {
			str = str + colNames
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

func (t JPK_WB) CreateCSVWyciagCtrl(filename string) {
	str := ""

	strSlice := []string{
		t.WyciagCtrl.LiczbaWierszy,
		t.WyciagCtrl.SumaObciazen,
		t.WyciagCtrl.SumaUznan}

	colNames := "LiczbaWierszy;SumaObciazen;SumaUznan\n"
	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + colNames + joined + "\n"

	WriteToCSV(str, filename)
}

func (t JPK_WB) CreateCSVSalda(filename string) {
	str := ""

	strSlice := []string{
		t.Salda.SaldoPoczatkowe,
		t.Salda.SaldoKoncowe}

	colNames := "SaldoPoczatkowe;SaldoKoncowe\n"
	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + colNames + joined + "\n"

	WriteToCSV(str, filename)
}

func (t JPK_WB) CreateCSVNumerRachunku(filename string) {
	str := ""
	strSlice := []string{
		t.NumerRachunku}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + "NumerRachunku\n" + joined + "\n"

	WriteToCSV(str, filename)

}

func (t JPK_WB) CreateCSVNaglowek(filename string) {
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
	str = str + colNames + joined

	WriteToCSV(str, filename)
}
