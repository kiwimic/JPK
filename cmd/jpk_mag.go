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
	magCmd.PersistentFlags().String("file", "", "file jpk to read.")
	magCmd.PersistentFlags().String("dir", "C:/jpk", "dir to export result tables")
}

var magCmd = &cobra.Command{
	Use:   "mag",
	Short: "convert jpk_mag.xml to csv tables exported to dir you choose",
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

		PZWartosc := ExportDir + "\\" + "PZWartosc.csv"
		WZWartosc := ExportDir + "\\" + "WZWartosc.csv"
		RWWartosc := ExportDir + "\\" + "RWWartosc.csv"
		MMWartosc := ExportDir + "\\" + "MMWartosc.csv"

		PZWiersz := ExportDir + "\\" + "PZWiersz.csv"
		WZWiersz := ExportDir + "\\" + "WZWiersz.csv"
		RWWiersz := ExportDir + "\\" + "RWWiersz.csv"
		MMWiersz := ExportDir + "\\" + "MMWiersz.csv"

		PZCtrlName := ExportDir + "\\" + "PZCtrl.csv"
		WZCtrlName := ExportDir + "\\" + "WZCtrl.csv"
		RWCtrlName := ExportDir + "\\" + "RWCtrl.csv"
		MMCtrlName := ExportDir + "\\" + "MMCtrl.csv"

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
		fmt.Println("Wczytano plik: ", v2)
		var jpk_mag utils.JPK_MAG
		xml.Unmarshal(byteValue, &jpk_mag)
		v3 := time.Now()
		fmt.Println("Zakończono parsowanie: ", v3)

		jpk_mag.CreateCSVFromRowsPZWartosc(200, PZWartosc)
		jpk_mag.CreateCSVFromRowsWZWartosc(200, WZWartosc)
		jpk_mag.CreateCSVFromRowsRWWartosc(200, RWWartosc)
		jpk_mag.CreateCSVFromRowsMMWartosc(200, MMWartosc)

		v4 := time.Now()
		fmt.Println("Zakończono zapisywanie PZ, WZ, RW, MM Wartosc", v4)

		jpk_mag.CreateCSVFromRowsPZWiersz(200, PZWiersz)
		jpk_mag.CreateCSVFromRowsWZWiersz(200, WZWiersz)
		jpk_mag.CreateCSVFromRowsRWWiersz(200, RWWiersz)
		jpk_mag.CreateCSVFromRowsMMWiersz(200, MMWiersz)

		v5 := time.Now()
		fmt.Println("Zakończono zapisywanie PZ, WZ, RW, MM Wiersz", v5)

		jpk_mag.CreateCSVPZCtrl(PZCtrlName)
		jpk_mag.CreateCSVWZCtrl(WZCtrlName)
		jpk_mag.CreateCSVRWCtrl(RWCtrlName)
		jpk_mag.CreateCSVMMCtrl(MMCtrlName)
		jpk_mag.CreateCSVNaglowek(NaglowekName)

		v6 := time.Now()
		fmt.Println("Zakończono zapisywanie sum kontrolnych i nagłowka", v6)

		end := time.Now()
		fmt.Println("Start: ", st, "\nEnd: ", end)
		fmt.Println("Wszystko trwalo: ", time.Since(st))

	},
}

func init() {
	RootCmd.AddCommand(magCmd)
}
