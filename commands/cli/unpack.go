package cli

import (
	chunks "archiver/pkg/archivers"
	"archiver/pkg/archivers/vlc"
	"archiver/pkg/lib"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   handleUnpack,
}

func init() {
	rootCmd.AddCommand(unpackCmd)
}

func handleUnpack(_ *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("Algoritm or File path was not provided")
	}

	filePath := args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	enc := SelectDecoder(args[0])
	packed := enc.Decode(data)

	err = os.WriteFile(lib.FileName(filePath, enc.GetExt()), []byte(packed), 644)
	if err != nil {
		log.Fatal(err)
	}
}

func SelectDecoder(alg string) chunks.Decoder {
	var enc chunks.Decoder

	switch alg {
	case "vlc":
		enc = vlc.NewDecoder()
	default:
		log.Fatal("Unsupported extension")
	}

	return enc
}
