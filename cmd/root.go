package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "jpk",
	Short: "'jpk' is an CLI app which helps you to parse JPK.xml files really fast :)",
	//Long: `This is long description. Now i don't know what to type here.
	//Maybe there will be something interesting here later.`,
}
