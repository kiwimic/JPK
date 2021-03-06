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
	vatCmd.PersistentFlags().String("file", "", "file jpk to read.")
	vatCmd.PersistentFlags().String("dir", "C:/jpk", "dir to export result tables")
}

var vatCmd = &cobra.Command{
	Use:   "vat",
	Short: "convert jpk_vat.xml to csv tables exported to dir you choose",
	//Long: `A Fast and Flexible Static Site Generator built with
	//              love by spf13 and friends in Go.
	//              Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {

		st := time.Now()
		fmt.Println("Let's get started: ", st)
		//FilePath := "C:\\jpk\\xml_duzy.xml"
		//ExportDir := "C:\\jpk"
		//FilePath to JPK file
		FilePath := cmd.Flag("file").Value.String()
		fmt.Println("File path: ", FilePath)
		//ExportDir where you save files
		ExportDir := cmd.Flag("dir").Value.String()
		fmt.Println("export dir: ", ExportDir)
		SprzedazWierszName := ExportDir + "\\" + "SprzedazWiersz.csv"
		ZakupWierszName := ExportDir + "\\" + "ZakupWiersz.csv"
		SprzedazCtrlName := ExportDir + "\\" + "SprzedazCtrl.csv"
		ZakupCtrlName := ExportDir + "\\" + "ZakupCtrl.csv"
		NaglowekName := ExportDir + "\\" + "Naglowek.csv"

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
		fmt.Println("File loaded: ", v2)
		var jpk_vat utils.JPK_VAT
		xml.Unmarshal(byteValue, &jpk_vat)
		v3 := time.Now()
		fmt.Println("Parsing ended: ", v3)
		//str := createCSVFaktura(JPK_FA.Faktura)
		jpk_vat.CreateCSVFromRowsSprzedazWiersz(200, SprzedazWierszName)
		//exportCSV(csvFaktura, "JPK_FA_Faktura.csv")

		fmt.Println("File created: SprzedazWiersz.csv:", time.Now())

		jpk_vat.CreateCSVFromRowsZakupWiersz(200, ZakupWierszName)

		fmt.Println("File created: ZakupWiersz.csv: ", time.Now())
		end := time.Now()

		test := len(jpk_vat.SprzedazWiersz)
		fmt.Println("Count of SprzedazWiersz: ", test)
		fmt.Println("Liczba wierszy zakupu to: ", len(jpk_vat.ZakupWiersz))
		//exportCSV(csvFakturaWiersz, "JPK_FA_FakturaWiersz.csv")

		jpk_vat.CreateCSVSprzedazCtrl(SprzedazCtrlName)
		jpk_vat.CreateCSVZakupCtrl(ZakupCtrlName)
		jpk_vat.CreateCSVNaglowek(NaglowekName)
		fmt.Println("Stworzono sumy kontrolne, oraz nagłowek", time.Now())

		fmt.Println("Start: ", st, "\nEnd: ", end)
		fmt.Println("Wszystko trwało: ", time.Since(st))

	},
}

func init() {
	RootCmd.AddCommand(vatCmd)
}
