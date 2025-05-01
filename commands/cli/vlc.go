package cli

import (
	"archiver/pkg/archivers/vlc"
	"archiver/pkg/lib"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable length algorithm",
	Run:   handle,
}

const extension = ".vlc"

func handle(_ *cobra.Command, args []string) {
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

	packed := vlc.Encode(string(data))

	err = os.WriteFile(lib.FileName(filePath, extension), []byte(packed), 644)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
