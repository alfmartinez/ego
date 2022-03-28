/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/alfmartinez/ego/tools/spritesheet"
	"github.com/spf13/cobra"
)

// ssheetCmd represents the ssheet command
var ssheetCmd = &cobra.Command{
	Use:   "ssheet",
	Short: "A sprite sheet building tool",
	Long:  `ssheet takes a folder with images, and creates a composite image to be used as a sprite sheet.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var inPath = args[0]
		var outPath = ""
		if len(args) > 1 {
			outPath += "/" + args[1]
		} else {
			outPath += "/" + args[0] + ".png"
		}
		sheet := spritesheet.New()

		defer func() {
			r := recover()
			if r != nil {
				fmt.Printf("Error : %s", r)
			}
		}()

		sheet.Load(inPath)
		sheet.Export(outPath)
		fmt.Printf("File created : %s", outPath)
	},
}

func init() {
	rootCmd.AddCommand(ssheetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ssheetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ssheetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
