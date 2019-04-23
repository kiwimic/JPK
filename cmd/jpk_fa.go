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
	faCmd.PersistentFlags().String("file", "", "file jpk to read.")
	faCmd.PersistentFlags().String("dir", "C:/jpk", "dir to export result tables")
}

var faCmd = &cobra.Command{
	Use:   "fa",
	Short: "convert jpk_fa.xml to csv tables exported to dir you choose",
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
		fakturaName := ExportDir + "\\" + "Faktura.csv"
		fakturawierszName := ExportDir + "\\" + "FakturaWiersz.csv"
		fakturaCtrlName := ExportDir + "\\" + "FakturaCtrl.csv"
		fakturaWierszCtrlName := ExportDir + "\\" + "FakturaWierszCtrl.csv"
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
		fmt.Println("File loadeds: ", v2)
		var jpk_fa utils.JPK_FA
		xml.Unmarshal(byteValue, &jpk_fa)
		v3 := time.Now()
		fmt.Println("Parsing ended: ", v3)
		//str := createCSVFaktura(JPK_FA.Faktura)
		jpk_fa.CreateCSVFromRowsFaktura(200, fakturaName)
		fmt.Println("File created: faktura.csv: ", time.Now())
		//exportCSV(csvFaktura, "JPK_FA_Faktura.csv")

		jpk_fa.CreateCSVFromRowsFakturaWiersz(200, fakturawierszName)

		fmt.Println("File created: fakturawiersz.csv", time.Now())
		test := len(jpk_fa.Faktura)
		fmt.Println("Count of Faktura: ", test)
		fmt.Println("Count of FakturaWiersz: ", len(jpk_fa.FakturaWiersz))
		//exportCSV(csvFakturaWiersz, "JPK_FA_FakturaWiersz.csv")

		jpk_fa.CreateCSVFakturaCtrl(fakturaCtrlName)
		jpk_fa.CreateCSVFakturaWierszCtrl(fakturaWierszCtrlName)
		jpk_fa.CreateCSVNaglowek(NaglowekName)
		fmt.Println("File created: facturaCtrl, fakturaWierszCtrl, Naglowek", time.Now())

		end := time.Now()
		fmt.Println("Start: ", st, "\nEnd: ", end)
		fmt.Println("Everything lasted: ", time.Since(st))

	},
}

func init() {
	RootCmd.AddCommand(faCmd)
}
