package utils

import "strings"

// JPK_VAT cały nosiciel JPK to do niego wkładamy zprasowany plik
type JPK_VAT struct {
	Naglowek       NaglowekVAT      `xml:"Naglowek"`
	SprzedazWiersz []SprzedazWiersz `xml:"SprzedazWiersz"`
	SprzedazCtrl   SprzedazCtrl     `xml:"SprzedazCtrl"`
	ZakupWiersz    []ZakupWiersz    `xml:"ZakupWiersz"`
	ZakupCtrl      ZakupCtrl        `xml:"ZakupCtrl"`
}

type NaglowekVAT struct {
	KodFormularza      string `xml:"KodFormularza"`
	WariantFormularza  string `xml:"WariantFormularza"`
	CelZlozenia        string `xml:"CelZlozenia"`
	DataWytworzeniaJPK string `xml:"DataWytworzeniaJPK"`
	DataOd             string `xml:"DataOd"`
	DataDo             string `xml:"DataDo"`
	DomyslnyKodWaluty  string `xml:"DomyslnyKodWaluty"`
	KodUrzedu          string `xml:"KodUrzedu"`
}

type SprzedazWiersz struct {
	LpSprzedazy      string `xml:"LpSprzedazy"`
	NrKontrahenta    string `xml:"NrKontrahenta"`
	NazwaKontrahenta string `xml:"NazwaKontrahenta"`
	AdresKontrahenta string `xml:"AdresKontrahenta"`
	DowodSprzedazy   string `xml:"DowodSprzedazy"`
	DataWystawienia  string `xml:"DataWystawienia"`
	DataSprzedazy    string `xml:"DataSprzedazy"`
	K_10             string `xml:"K_10"`
	K_11             string `xml:"K_11"`
	K_12             string `xml:"K_12"`
	K_13             string `xml:"K_13"`
	K_14             string `xml:"K_14"`
	K_15             string `xml:"K_15"`
	K_16             string `xml:"K_16"`
	K_17             string `xml:"K_17"`
	K_18             string `xml:"K_18"`
	K_19             string `xml:"K_19"`
	K_20             string `xml:"K_20"`
	K_21             string `xml:"K_21"`
	K_22             string `xml:"K_22"`
	K_23             string `xml:"K_23"`
	K_24             string `xml:"K_24"`
	K_25             string `xml:"K_25"`
	K_26             string `xml:"K_26"`
	K_27             string `xml:"K_27"`
	K_28             string `xml:"K_28"`
	K_29             string `xml:"K_29"`
	K_30             string `xml:"K_30"`
	K_31             string `xml:"K_31"`
	K_32             string `xml:"K_32"`
	K_33             string `xml:"K_33"`
	K_34             string `xml:"K_34"`
	K_35             string `xml:"K_35"`
	K_36             string `xml:"K_36"`
	K_37             string `xml:"K_37"`
	K_38             string `xml:"K_38"`
	K_39             string `xml:"K_39"`
}

type SprzedazCtrl struct {
	LiczbaWierszySprzedazy string `xml:"LiczbaWierszySprzedazy"`
	PodatekNalezny         string `xml:"PodatekNalezny"`
}

type ZakupWiersz struct {
	LpZakupu      string `xml:"LpZakupu"`
	NrDostawcy    string `xml:"NrDostawcy"`
	NazwaDostawcy string `xml:"NazwaDostawcy"`
	AdresDostawcy string `xml:"AdresDostawcy"`
	DowodZakupu   string `xml:"DowodZakupu"`
	DataZakupu    string `xml:"DataZakupu"`
	DataWplywu    string `xml:"DataWplywu"`
	K_43          string `xml:"K_43"`
	K_44          string `xml:"K_44"`
	K_45          string `xml:"K_45"`
	K_46          string `xml:"K_46"`
	K_47          string `xml:"K_47"`
	K_48          string `xml:"K_48"`
	K_49          string `xml:"K_49"`
	K_50          string `xml:"K_50"`
}

type ZakupCtrl struct {
	LiczbaWierszyZakupoow string `xml:"LiczbaWierszyZakupow"`
	PodatekNaliczony      string `xml:"PodatekNaliczony"`
}

func (t JPK_VAT) createRowSprzedazWiersz(i int) string {
	str := ""

	strSlice := []string{
		t.SprzedazWiersz[i].LpSprzedazy,
		t.SprzedazWiersz[i].NrKontrahenta,
		t.SprzedazWiersz[i].NazwaKontrahenta,
		t.SprzedazWiersz[i].AdresKontrahenta,
		t.SprzedazWiersz[i].DowodSprzedazy,
		t.SprzedazWiersz[i].DataWystawienia,
		t.SprzedazWiersz[i].DataSprzedazy,
		t.SprzedazWiersz[i].K_10,
		t.SprzedazWiersz[i].K_11,
		t.SprzedazWiersz[i].K_12,
		t.SprzedazWiersz[i].K_13,
		t.SprzedazWiersz[i].K_14,
		t.SprzedazWiersz[i].K_15,
		t.SprzedazWiersz[i].K_16,
		t.SprzedazWiersz[i].K_17,
		t.SprzedazWiersz[i].K_18,
		t.SprzedazWiersz[i].K_19,
		t.SprzedazWiersz[i].K_20,
		t.SprzedazWiersz[i].K_21,
		t.SprzedazWiersz[i].K_22,
		t.SprzedazWiersz[i].K_23,
		t.SprzedazWiersz[i].K_24,
		t.SprzedazWiersz[i].K_25,
		t.SprzedazWiersz[i].K_26,
		t.SprzedazWiersz[i].K_27,
		t.SprzedazWiersz[i].K_28,
		t.SprzedazWiersz[i].K_29,
		t.SprzedazWiersz[i].K_30,
		t.SprzedazWiersz[i].K_31,
		t.SprzedazWiersz[i].K_32,
		t.SprzedazWiersz[i].K_33,
		t.SprzedazWiersz[i].K_34,
		t.SprzedazWiersz[i].K_35,
		t.SprzedazWiersz[i].K_36,
		t.SprzedazWiersz[i].K_37,
		t.SprzedazWiersz[i].K_38,
		t.SprzedazWiersz[i].K_39}

	joined := strings.Join(strSlice, "\";\"")
	joined = "\"" + joined + "\""
	str = str + joined + "\n"

	return str
}

func (t JPK_VAT) createRowZakupWiersz(i int) string {
	str := ""

	strSlice := []string{
		t.ZakupWiersz[i].LpZakupu,
		t.ZakupWiersz[i].NrDostawcy,
		t.ZakupWiersz[i].NazwaDostawcy,
		t.ZakupWiersz[i].AdresDostawcy,
		t.ZakupWiersz[i].DowodZakupu,
		t.ZakupWiersz[i].DataZakupu,
		t.ZakupWiersz[i].DataWplywu,
		t.ZakupWiersz[i].K_43,
		t.ZakupWiersz[i].K_44,
		t.ZakupWiersz[i].K_45,
		t.ZakupWiersz[i].K_46,
		t.ZakupWiersz[i].K_47,
		t.ZakupWiersz[i].K_48,
		t.ZakupWiersz[i].K_49,
		t.ZakupWiersz[i].K_50}

	joined := strings.Join(strSlice, "\";\"")
	joined = "\"" + joined + "\""
	str = str + joined + "\n"

	return str
}

func (t JPK_VAT) LiczbaFakturSprzedazy() int {
	return (len(t.SprzedazWiersz))
}

func (t JPK_VAT) LiczbaFakturZakupu() int {
	return (len(t.ZakupWiersz))
}

func (t JPK_VAT) CreateCSVFromRowsSprzedazWiersz(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}
	for i := range t.SprzedazWiersz {
		rowPaste := t.createRowSprzedazWiersz(i)
		str = str + rowPaste
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			WriteToCSV(str, filename)
			str = ""
		}
	}
	WriteToCSV(str, filename)
}

func (t JPK_VAT) CreateCSVFromRowsZakupWiersz(buforSize int, filename string) {
	str := ""
	if buforSize == 0 {
		buforSize = 200
	}
	for i := range t.ZakupWiersz {
		rowPaste := t.createRowZakupWiersz(i)
		str = str + rowPaste
		if i%200 == 0 {
			//fmt.Println("Faktura: ", i)
			WriteToCSV(str, filename)
			str = ""
		}
	}
	WriteToCSV(str, filename)
}

func (t JPK_VAT) CreateCSVSprzedazCtrl(filename string) {
	str := ""

	strSlice := []string{
		t.SprzedazCtrl.LiczbaWierszySprzedazy,
		t.SprzedazCtrl.PodatekNalezny}

	joined := strings.Join(strSlice, "\";\"")
	joined = "\"" + joined + "\""
	str = str + joined + "\n"

	WriteToCSV(str, filename)
}
