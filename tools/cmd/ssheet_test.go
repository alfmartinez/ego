package cmd

import "github.com/spf13/cobra"

func Example_ssheet_nooutput() {
	args := []string{"testdata/self"}
	defer func() {

	}()
	ssheetCmd.Run(&cobra.Command{}, args)
	// Output:
	// File created : /testdata/self.png
}

func Example_ssheet_output() {
	args := []string{"testdata/self", "testdata/output/other.png"}
	defer func() {

	}()
	ssheetCmd.Run(&cobra.Command{}, args)
	// Output:
	// File created : /testdata/output/other.png
}
