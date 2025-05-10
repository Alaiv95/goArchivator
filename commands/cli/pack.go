package cli

import (
	chunks "archiver/pkg/archivers"
	"archiver/pkg/archivers/vlc"
	"archiver/pkg/archivers/vlc/table/shannon_fano"
	"archiver/pkg/lib"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   handlePack,
}

func init() {
	rootCmd.AddCommand(packCmd)
}

const extension = ".fano"

func handlePack(_ *cobra.Command, args []string) {
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

	enc := SelectEncoder(args[0])
	packed := enc.Encode(string(data))

	err = os.WriteFile(lib.FileName(filePath, enc.GetExt()), packed, 644)
	if err != nil {
		log.Fatal(err)
	}
}

func SelectEncoder(alg string) chunks.Encoder {
	var enc chunks.Encoder

	switch alg {
	case "vlc":
		enc = vlc.NewEncoder(shannon_fano.NewGenerator())
	default:
		log.Fatal("Unsupported extension")
	}

	return enc
}
