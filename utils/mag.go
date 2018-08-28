package utils

import "strings"

// JPK_MAG cały nosiciel JPK to do niego wkładamy zprasowany plik
type JPK_MAG struct {
	Naglowek NaglowekMAG `xml:"Naglowek"`
	Magazyn  string      `xml:"Magazyn"`
	PZ       []PZ        `xml:"PZ"`
	WZ       []WZ        `xml:"WZ"`
	RW       []RW        `xml:"RW"`
	MM       []MM        `xml:"MM"`
}

type NaglowekMAG struct {
	KodFormularza      string `xml:"KodFormularza"`
	WariantFormularza  string `xml:"WariantFormularza"`
	CelZlozenia        string `xml:"CelZlozenia"`
	DataWytworzeniaJPK string `xml:"DataWytworzeniaJPK"`
	DataOd             string `xml:"DataOd"`
	DataDo             string `xml:"DataDo"`
	DomyslnyKodWaluty  string `xml:"DomyslnyKodWaluty"`
	KodUrzedu          string `xml:"KodUrzedu"`
}

type PZ struct {
	PZWartosc []PZWartosc `xml:"PZWartosc"`
	PZWiersz  []PZWiersz  `xml:"PZWiersz"`
	PZCtrl    PZCtrl      `xml:"PZCtrl"`
}

type PZWartosc struct {
	NumerPZ          string `xml:"NumerPZ"`
	DataPZ           string `xml:"DataPZ"`
	WartoscPZ        string `xml:"WartoscPZ"`
	DataOtrzymaniaPZ string `xml:"DataOtrzymaniaPZ"`
	Dostawca         string `xml:"Dostawca"`
	NumerFaPZ        string `xml:"NumerFaPZ"`
	DataFaPZ         string `xml:"DataFaPZ"`
}

type PZWiersz struct {
	Numer2PZ         string `xml:"Numer2PZ"`
	KodTowaruPZ      string `xml:"KodTowaruPZ"`
	NazwaTowaruPZ    string `xml:"NazwaTowaruPZ"`
	IloscPrzyjetaPZ  string `xml:"IloscPrzyjetaPZ"`
	JednostkaMiaryPZ string `xml:"JednostkaMiaryPZ"`
	CenaJednPZ       string `xml:"CenaJednPZ"`
	WartoscPozycjiPZ string `xml:"WartoscPozycjiPZ"`
}

type PZCtrl struct {
	LiczbaPZ string `xml:"LiczbaPZ"`
	SumaPZ   string `xml:"SumaPZ"`
}

type WZ struct {
	WZWartosc []WZWartosc `xml:"WZWartosc"`
	WZWiersz  []WZWiersz  `xml:"WZWiersz"`
	WZCtrl    WZCtrl      `xml:"WZCtrl"`
}

type WZWartosc struct {
	NumerWZ       string `xml:"NumerWZ"`
	DataWZ        string `xml:"DataWZ"`
	WartoscWZ     string `xml:"WartoscWZ"`
	DataWydaniaWZ string `xml:"DataWydaniaWZ"`
	OdbiorcaWZ    string `xml:"OdbiorcaWZ"`
	NumerFaWZ     string `xml:"NumerFaWZ"`
	DataFaWZ      string `xml:"DataFaWZ"`
}

type WZWiersz struct {
	Numer2WZ         string `xml:"Numer2WZ"`
	KodTowaruWZ      string `xml:"KodTowaruWZ"`
	NazwaTowaruWZ    string `xml:"NazwaTowaruWZ"`
	IloscWydanaWZ    string `xml:"IloscWydanaWZ"`
	JednostkaMiaryWZ string `xml:"JednostkaMiaryWZ"`
	CenaJednWZ       string `xml:"CenaJednWZ"`
	WartoscPozycjiWZ string `xml:"WartoscPozycjiWZ"`
}

type WZCtrl struct {
	LiczbaWZ string `xml:"LiczbaWZ"`
	SumaWZ   string `xml:"SumaWZ"`
}

type RW struct {
	RWWartosc []string `xml:"RWWartosc"`
	RWWiersz  []string `xml:"RWWiersz"`
	RWCtrl    string   `xml:"RWCtrl"`
}

type RWWartosc struct {
	NumerRW       string `xml:"NumerRW"`
	DataRW        string `xml:"DataRW"`
	WartoscRW     string `xml:"WartoscRW"`
	DataWydaniaRW string `xml:"DataWydaniaRW"`
	SkadRW        string `xml:"SkadRW"`
	DokadRW       string `xml:"DokadRW"`
}

type RWWiersz struct {
	Numer2RW         string `xml:"Numer2RW"`
	KodTowaruRW      string `xml:"KodTowaruRW"`
	NazwaTowaruRW    string `xml:"NazwaTowaruRW"`
	IloscWydanaRW    string `xml:"IloscWydanaRW"`
	JednostkaMiaryRW string `xml:"JednostkaMiaryRW"`
	CenaJednRW       string `xml:"CenaJednRW"`
	WartoscPozycjiRW string `xml:"WartoscPozycjiRW"`
}

type RWCtrl struct {
	LiczbaRW string `xml:"LiczbaRW"`
	SumaRW   string `xml:"SumaRW"`
}

type MM struct {
	MMWartosc []string `xml:"MMWartosc"`
	MMWiersz  []string `xml:"MMWiersz"`
	MMCtrl    string   `xml:"MMCtrl"`
}

type MMWartosc struct {
	NumerMM       string `xml:"NumerMM"`
	DataMM        string `xml:"DataMM"`
	WartoscMM     string `xml:"WartoscMM"`
	DataWydaniaMM string `xml:"DataWydaniaMM"`
	SkadMM        string `xml:"SkadMM"`
	DokadMM       string `xml:"DokadMM"`
}

type MMWiersz struct {
	Numer2MM         string `xml:"Numer2MM"`
	KodTowaruMM      string `xml:"KodTowaruMM"`
	NazwaTowaruMM    string `xml:"NazwaTowaruMM"`
	IloscWydanaMM    string `xml:"IloscWydanaMM"`
	JednostkaMiaryMM string `xml:"JednostkaMiaryMM"`
	CenaJednMM       string `xml:"CenaJednMM"`
	WartoscPozycjiMM string `xml:"WartoscPozycjiMM"`
}

type MMCtrl struct {
	LiczbaMM string `xml:"LiczbaMM"`
	SumaMM   string `xml:"SumaMM"`
}

//
func (t JPK_MAG) createRowPZWartosc(i int) string {
	str := ""

	strSlice := []string{
		t.PZ[i].PZWartosc[i].NumerPZ,
		t.PZ[i].PZWartosc[i].DataPZ,
		t.PZ[i].PZWartosc[i].WartoscPZ,
		t.PZ[i].PZWartosc[i].DataOtrzymaniaPZ,
		t.PZ[i].PZWartosc[i].Dostawca,
		t.PZ[i].PZWartosc[i].NumerFaPZ,
		t.PZ[i].PZWartosc[i].DataFaPZ}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowWZWartosc(i int) string {
	str := ""

	strSlice := []string{
		t.WZ[i].WZWartosc[i].NumerWZ,
		t.WZ[i].WZWartosc[i].DataWZ,
		t.WZ[i].WZWartosc[i].WartoscWZ,
		t.WZ[i].WZWartosc[i].DataWydaniaWZ,
		t.WZ[i].WZWartosc[i].OdbiorcaWZ,
		t.WZ[i].WZWartosc[i].NumerFaWZ,
		t.WZ[i].WZWartosc[i].DataFaWZ,
	}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowRWWartosc(i int) string {
	str := ""

	strSlice := []string{
		t.RW[i].RWWartosc[i].NumerRW,
		t.RW[i].RWWartosc[i].DataRW,
		t.RW[i].RWWartosc[i].WartoscRW,
		t.RW[i].RWWartosc[i].DataWydaniaRW,
		t.RW[i].RWWartosc[i].SkadRW,
		t.RW[i].RWWartosc[i].DokadRW,
	}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowMMWartosc(i int) string {
	str := ""

	strSlice := []string{
		t.MM[i].MMWartosc[i].NumerMM,
		t.MM[i].MMWartosc[i].DataMM,
		t.MM[i].MMWartosc[i].WartoscMM,
		t.MM[i].MMWartosc[i].DataWydaniaMM,
		t.MM[i].MMWartosc[i].SkadMM,
		t.MM[i].MMWartosc[i].DokadMM,
	}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowPZWiersz(i int) string {
	str := ""

	strSlice := []string{
		t.PZ[i].PZWiersz[i].Numer2PZ,
		t.PZ[i].PZWiersz[i].KodTowaruPZ,
		t.PZ[i].PZWiersz[i].NazwaTowaruPZ,
		t.PZ[i].PZWiersz[i].IloscPrzyjetaPZ,
		t.PZ[i].PZWiersz[i].JednostkaMiaryPZ,
		t.PZ[i].PZWiersz[i].CenaJednPZ,
		t.PZ[i].PZWiersz[i].WartoscPozycjiPZ}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowWZWiersz(i int) string {
	str := ""

	strSlice := []string{
		t.WZ[i].WZWiersz[i].Numer2WZ,
		t.WZ[i].WZWiersz[i].KodTowaruWZ,
		t.WZ[i].WZWiersz[i].NazwaTowaruWZ,
		t.WZ[i].WZWiersz[i].IloscWydanaWZ,
		t.WZ[i].WZWiersz[i].JednostkaMiaryWZ,
		t.WZ[i].WZWiersz[i].CenaJednWZ,
		t.WZ[i].WZWiersz[i].WartoscPozycjiWZ,
	}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowRWWiersz(i int) string {
	str := ""

	strSlice := []string{
		t.RW[i].RWWartosc[i].Numer2RW,
		t.RW[i].RWWartosc[i].KodTowaruRW,
		t.RW[i].RWWartosc[i].NazwaTowaruRW,
		t.RW[i].RWWartosc[i].IloscWydanaRW,
		t.RW[i].RWWartosc[i].JednostkaMiaryRW,
		t.RW[i].RWWartosc[i].CenaJednRW,
		t.RW[i].RWWartosc[i].WartoscPozycjiRW}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowMMWiersz(i int) string {
	str := ""

	strSlice := []string{
		t.MM[i].MMWartosc[i].Numer2MM,
		t.MM[i].MMWartosc[i].KodTowaruMM,
		t.MM[i].MMWartosc[i].NazwaTowaruMM,
		t.MM[i].MMWartosc[i].IloscWydanaMM,
		t.MM[i].MMWartosc[i].JednostkaMiaryMM,
		t.MM[i].MMWartosc[i].CenaJednMM,
		t.MM[i].MMWartosc[i].WartoscPozycjiMM,
	}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) CreateCSVFromRowsWZWartosc(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}
	for i := range t.WyciagWiersz {
		rowPaste := t.createRowWZWartosc(i)
		str = str + rowPaste
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			WriteToCSV(str, filename)
			str = ""
		}
	}
	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVFromRowsWZWiersz(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}
	for i := range t.WyciagWiersz {
		rowPaste := t.createRowWZWiersz(i)
		str = str + rowPaste
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			WriteToCSV(str, filename)
			str = ""
		}
	}
	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVFromRowsPZWartosc(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}
	for i := range t.WyciagWiersz {
		rowPaste := t.createRowPZWartosc(i)
		str = str + rowPaste
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			WriteToCSV(str, filename)
			str = ""
		}
	}
	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVFromRowsPZWiersz(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}
	for i := range t.WyciagWiersz {
		rowPaste := t.createRowPZWiersz(i)
		str = str + rowPaste
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			WriteToCSV(str, filename)
			str = ""
		}
	}
	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVFromRowsRWWartosc(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}
	for i := range t.WyciagWiersz {
		rowPaste := t.createRowRWWartosc(i)
		str = str + rowPaste
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			WriteToCSV(str, filename)
			str = ""
		}
	}
	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVFromRowsRWWiersz(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}
	for i := range t.WyciagWiersz {
		rowPaste := t.createRowRWWiersz(i)
		str = str + rowPaste
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			WriteToCSV(str, filename)
			str = ""
		}
	}
	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVFromRowsMMWartosc(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}
	for i := range t.WyciagWiersz {
		rowPaste := t.createRowMMWartosc(i)
		str = str + rowPaste
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			WriteToCSV(str, filename)
			str = ""
		}
	}
	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVFromRowsMMWiersz(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}
	for i := range t.WyciagWiersz {
		rowPaste := t.createRowMMWiersz(i)
		str = str + rowPaste
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			WriteToCSV(str, filename)
			str = ""
		}
	}
	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVPZCtrl(filename string) {
	str := ""

	strSlice := []string{
		t.PZ.PZCtrl.LiczbaPZ,
		t.PZ.PZCtrl.SumaPZ}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVWZCtrl(filename string) {
	str := ""

	strSlice := []string{
		t.WZ.WZCtrl.LiczbaWZ,
		t.WZ.WZCtrl.SumaWZ}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVRWCtrl(filename string) {
	str := ""

	strSlice := []string{
		t.RW.RWCtrl.LiczbaRW,
		t.RW.RWCtrl.SumaRW}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVMMCtrl(filename string) {
	str := ""

	strSlice := []string{
		t.MM.MMCtrl.LiczbaMM,
		t.MM.MMCtrl.SumaMM}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVNaglowek(filename string) {
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
	str = str + joined + "\n"

	WriteToCSV(str, filename)
}
