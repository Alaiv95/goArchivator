package cli

import (
	"archiver/pkg/archivers/vlc"
	"archiver/pkg/lib"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

var unpackVlcCmd = &cobra.Command{
	Use:   "unpack-vlc",
	Short: "Unpack file using variable length algorithm",
	Run:   handleUnpack,
}

const unpackExtension = ".txt"

func handleUnpack(_ *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("File path was not provided")
	}

	filePath := args[0]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	packed := vlc.Decode(data)

	err = os.WriteFile(lib.FileName(filePath, unpackExtension), []byte(packed), 644)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	unpackCmd.AddCommand(unpackVlcCmd)
}
