package cmd

import "github.com/spf13/cobra"

func Example_ssheet_nooutput() {
	args := []string{"self"}
	defer func() {

	}()
	ssheetCmd.Run(&cobra.Command{}, args)
	// Output:
	// Error : cannot read path /
}

func Example_ssheet_output() {
	args := []string{"self", "other"}
	defer func() {

	}()
	ssheetCmd.Run(&cobra.Command{}, args)
	// Output:
	// Error : cannot read path /*
}
