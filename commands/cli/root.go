package cli

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Short: "Simple Archivator",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
