package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCommand = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

const packedExtention = "vlc"

func init() {
	rootCmd.AddCommand(vlcCommand)
}

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		HandleErr(errors.New("define file path"))
	}

	filepath := args[0]

	r, err := os.Open(filepath)
	if err != nil {
		HandleErr(err)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		HandleErr(err)
	}

	defer r.Close() // closed file

	//we should transfer this var Data to encode function

	packed := "" + string(data)

	err = os.WriteFile(packedFileName(filepath), []byte(packed), 0644)
	if err != nil {
		HandleErr(err)
	}

}

func packedFileName(path string) string {
	// path/to/my/file/my_file.txt -> my_file.vlc

	fileName := filepath.Base(path)

	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)

	return baseName + "." + packedExtention
}

func init() {
	packCmd.AddCommand(vlcCommand)
}
