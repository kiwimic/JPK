package utils

import "strings"

// JPK_MAG cały nosiciel JPK to do niego wkładamy zprasowany plik
type JPK_MAG struct {
	Naglowek NaglowekMAG `xml:"Naglowek"`
	Magazyn  string      `xml:"Magazyn"`
	PZ       PZ          `xml:"PZ"`
	WZ       WZ          `xml:"WZ"`
	RW       RW          `xml:"RW"`
	MM       MM          `xml:"MM"`
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
	RWWartosc []RWWartosc `xml:"RWWartosc"`
	RWWiersz  []RWWiersz  `xml:"RWWiersz"`
	RWCtrl    RWCtrl      `xml:"RWCtrl"`
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
	MMWartosc []MMWartosc `xml:"MMWartosc"`
	MMWiersz  []MMWiersz  `xml:"MMWiersz"`
	MMCtrl    MMCtrl      `xml:"MMCtrl"`
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
		t.PZ.PZWartosc[i].NumerPZ,
		t.PZ.PZWartosc[i].DataPZ,
		t.PZ.PZWartosc[i].WartoscPZ,
		t.PZ.PZWartosc[i].DataOtrzymaniaPZ,
		t.PZ.PZWartosc[i].Dostawca,
		t.PZ.PZWartosc[i].NumerFaPZ,
		t.PZ.PZWartosc[i].DataFaPZ}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowWZWartosc(i int) string {
	str := ""

	strSlice := []string{
		t.WZ.WZWartosc[i].NumerWZ,
		t.WZ.WZWartosc[i].DataWZ,
		t.WZ.WZWartosc[i].WartoscWZ,
		t.WZ.WZWartosc[i].DataWydaniaWZ,
		t.WZ.WZWartosc[i].OdbiorcaWZ,
		t.WZ.WZWartosc[i].NumerFaWZ,
		t.WZ.WZWartosc[i].DataFaWZ,
	}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowRWWartosc(i int) string {
	str := ""

	strSlice := []string{
		t.RW.RWWartosc[i].NumerRW,
		t.RW.RWWartosc[i].DataRW,
		t.RW.RWWartosc[i].WartoscRW,
		t.RW.RWWartosc[i].DataWydaniaRW,
		t.RW.RWWartosc[i].SkadRW,
		t.RW.RWWartosc[i].DokadRW,
	}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowMMWartosc(i int) string {
	str := ""

	strSlice := []string{
		t.MM.MMWartosc[i].NumerMM,
		t.MM.MMWartosc[i].DataMM,
		t.MM.MMWartosc[i].WartoscMM,
		t.MM.MMWartosc[i].DataWydaniaMM,
		t.MM.MMWartosc[i].SkadMM,
		t.MM.MMWartosc[i].DokadMM,
	}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowPZWiersz(i int) string {
	str := ""

	strSlice := []string{
		t.PZ.PZWiersz[i].Numer2PZ,
		t.PZ.PZWiersz[i].KodTowaruPZ,
		t.PZ.PZWiersz[i].NazwaTowaruPZ,
		t.PZ.PZWiersz[i].IloscPrzyjetaPZ,
		t.PZ.PZWiersz[i].JednostkaMiaryPZ,
		t.PZ.PZWiersz[i].CenaJednPZ,
		t.PZ.PZWiersz[i].WartoscPozycjiPZ}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowWZWiersz(i int) string {
	str := ""

	strSlice := []string{
		t.WZ.WZWiersz[i].Numer2WZ,
		t.WZ.WZWiersz[i].KodTowaruWZ,
		t.WZ.WZWiersz[i].NazwaTowaruWZ,
		t.WZ.WZWiersz[i].IloscWydanaWZ,
		t.WZ.WZWiersz[i].JednostkaMiaryWZ,
		t.WZ.WZWiersz[i].CenaJednWZ,
		t.WZ.WZWiersz[i].WartoscPozycjiWZ,
	}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowRWWiersz(i int) string {
	str := ""

	strSlice := []string{
		t.RW.RWWiersz[i].Numer2RW,
		t.RW.RWWiersz[i].KodTowaruRW,
		t.RW.RWWiersz[i].NazwaTowaruRW,
		t.RW.RWWiersz[i].IloscWydanaRW,
		t.RW.RWWiersz[i].JednostkaMiaryRW,
		t.RW.RWWiersz[i].CenaJednRW,
		t.RW.RWWiersz[i].WartoscPozycjiRW}

	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + joined + "\n"

	return str
}

func (t JPK_MAG) createRowMMWiersz(i int) string {
	str := ""

	strSlice := []string{
		t.MM.MMWiersz[i].Numer2MM,
		t.MM.MMWiersz[i].KodTowaruMM,
		t.MM.MMWiersz[i].NazwaTowaruMM,
		t.MM.MMWiersz[i].IloscWydanaMM,
		t.MM.MMWiersz[i].JednostkaMiaryMM,
		t.MM.MMWiersz[i].CenaJednMM,
		t.MM.MMWiersz[i].WartoscPozycjiMM,
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

	colNames := "NumerWZ;DataWZ;WartoscWZ;DataWydaniaWZ;OdbiorcaWZ;NumerFaWZ;DataFaWZ\n"

	for i := range t.WZ.WZWartosc {
		rowPaste := t.createRowWZWartosc(i)

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

func (t JPK_MAG) CreateCSVFromRowsWZWiersz(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}

	colNames := "Numer2WZ;KodTowaruWZ;NazwaTowaruWZ;IloscWydanaWZ;JednostkaMiaryWZ;CenaJednWZ;WartoscPozycjiWZ\n"

	for i := range t.WZ.WZWiersz {
		rowPaste := t.createRowWZWiersz(i)

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

func (t JPK_MAG) CreateCSVFromRowsPZWartosc(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}

	colNames := "NumerPZ;WartoscPZ;DataOtrzymaniaPZ;Dostawca;NumerFaPZ;DataFaPZ\n"
	for i := range t.PZ.PZWartosc {
		rowPaste := t.createRowPZWartosc(i)

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

func (t JPK_MAG) CreateCSVFromRowsPZWiersz(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}

	colNames := "Numer2PZ;KodTowaruPZ;NazwaTowaruPZ;IloscPrzyjetaPZ;JednostkaMiaryPZ;CenaJednPZ;WartoscPozycjiPZ\n"
	for i := range t.PZ.PZWiersz {
		rowPaste := t.createRowPZWiersz(i)

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

func (t JPK_MAG) CreateCSVFromRowsRWWartosc(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}

	colNames := "NumerRW;DataRW;WartoscRW;DataWydaniaRW;SkadRW;DokadRW\n"
	for i := range t.RW.RWWartosc {
		rowPaste := t.createRowRWWartosc(i)

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

func (t JPK_MAG) CreateCSVFromRowsRWWiersz(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}

	colNames := "Numer2RW;KodTowaruRW;NazwaTowaruRW;IloscWydanaRW;JednostkaMiaryRW;CenaJednRW;WartoscPozycjiRW\n"
	for i := range t.RW.RWWiersz {
		rowPaste := t.createRowRWWiersz(i)

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

func (t JPK_MAG) CreateCSVFromRowsMMWartosc(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}

	colNames := "NumerMM;DataMM;WartoscMM;DataWydaniaMM;SkadMM;DokadMM\n"
	for i := range t.MM.MMWartosc {
		rowPaste := t.createRowMMWartosc(i)

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

func (t JPK_MAG) CreateCSVFromRowsMMWiersz(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}

	colNames := "Numer2MM;KodTowaruMM;NazwaTowaruMM;IloscWydanaMM;JednostkaMiaryMM;CenaJednMM;WartoscPozycjiMM\n"
	for i := range t.MM.MMWiersz {
		rowPaste := t.createRowMMWiersz(i)

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

func (t JPK_MAG) CreateCSVPZCtrl(filename string) {
	str := ""

	strSlice := []string{
		t.PZ.PZCtrl.LiczbaPZ,
		t.PZ.PZCtrl.SumaPZ}

	colNames := "LiczbaPZ;SumaPZ\n"
	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + colNames + joined + "\n"

	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVWZCtrl(filename string) {
	str := ""

	strSlice := []string{
		t.WZ.WZCtrl.LiczbaWZ,
		t.WZ.WZCtrl.SumaWZ}

	colNames := "LiczbaWZ;SumaWZ\n"
	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + colNames + joined + "\n"

	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVRWCtrl(filename string) {
	str := ""

	strSlice := []string{
		t.RW.RWCtrl.LiczbaRW,
		t.RW.RWCtrl.SumaRW}

	colNames := "LiczbaRW;SumaRW\n"
	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + colNames + joined + "\n"

	WriteToCSV(str, filename)
}

func (t JPK_MAG) CreateCSVMMCtrl(filename string) {
	str := ""

	strSlice := []string{
		t.MM.MMCtrl.LiczbaMM,
		t.MM.MMCtrl.SumaMM}

	colNames := "LiczbaMM;SumaMM\n"
	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + colNames + joined + "\n"

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

	colNames := "KodFormularza;WariantFormularza;CelZlozenia;DataWytworzeniaJPK;DataOd;DataDo;DomyslnyKodWaluty;KodUrzedu\n"
	strSlice = RemoveStringFromSliceOfString(strSlice, ";", "", -1)
	joined := strings.Join(strSlice, ";")
	str = str + colNames + joined + "\n"

	WriteToCSV(str, filename)

}
