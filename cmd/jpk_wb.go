package cmd

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/jpk/utils"
	"github.com/spf13/cobra"
)

func init() {
	wbCmd.PersistentFlags().String("file", "", "file jpk to read.")
	wbCmd.PersistentFlags().String("dir", "C:/jpk", "dir to export result tables")
}

var wbCmd = &cobra.Command{
	Use:   "wb",
	Short: "convert jpk_wb.xml to csv tables exported to dir you choose",
	//Long: `A Fast and Flexible Static Site Generator built with
	//              love by spf13 and friends in Go.
	//              Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {

		st := time.Now()
		fmt.Println("Zaczynamy: ", st)
		//FilePath := "C:\\jpk\\xml_duzy.xml"
		//ExportDir := "C:\\jpk"
		//FilePath to JPK file
		FilePath := cmd.Flag("file").Value.String()
		fmt.Println("File path: ", FilePath)
		//ExportDir where you save files
		ExportDir := cmd.Flag("dir").Value.String()
		fmt.Println("export dir: ", ExportDir)
		NumerRachunkuName := ExportDir + "\\" + "NumerRachunku.txt"
		SaldaName := ExportDir + "\\" + "Salda.txt"
		WyciagCtrlName := ExportDir + "\\" + "WyciagCtrl.txt"
		WYciagWierszName := ExportDir + "\\" + "WyciagWiersz.txt"
		NaglowekName := ExportDir + "\\" + "Naglowek.txt"

		xmlFile, err := os.Open(FilePath)
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
		var jpk_wb utils.JPK_WB
		xml.Unmarshal(byteValue, &jpk_wb)
		v3 := time.Now()
		fmt.Println("Zakończono parsowanie: ", v3)
		//str := createCSVFaktura(JPK_FA.Faktura)
		jpk_wb.CreateCSVFromRowsWyciagWiersz(200, WYciagWierszName)
		fmt.Println("Stworzono CSV WyciagWiersz: ", time.Now())
		//exportCSV(csvFaktura, "JPK_FA_Faktura.txt")

		jpk_wb.CreateCSVSalda(SaldaName)
		fmt.Println("Stworzono CSV Salda: ", time.Now())
		jpk_wb.CreateCSVNumerRachunku(NumerRachunkuName)
		fmt.Println("Zapisano CSV NumerRachunku: ", time.Now())
		jpk_wb.CreateCSVNaglowek(NaglowekName)
		fmt.Println("Zapisano CSV Naglowek: ", time.Now())
		jpk_wb.CreateCSVWyciagCtrl(WyciagCtrlName)
		fmt.Println("Stworzono sumy kontrolne", time.Now())

		end := time.Now()
		fmt.Println("Start: ", st, "\nEnd: ", end)
		fmt.Println("Wszystko trwało: ", time.Since(st))

	},
}

func init() {
	RootCmd.AddCommand(wbCmd)
}
