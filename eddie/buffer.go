package eddie

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/theapemachine/boogie/apeterm"
	"golang.org/x/crypto/ssh/terminal"
)

type Buffer struct {
	ui       *apeterm.UI
	filePath *string
	data     []string
}

func NewBuffer(ui *apeterm.UI, filePath *string) *Buffer {
	return &Buffer{
		ui:       ui,
		filePath: filePath,
		data:     make([]string, 0),
	}
}

func (buffer *Buffer) Load() {
	fh, err := os.Open(*buffer.filePath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		buffer.data = append(buffer.data, scanner.Text())
	}

	fmt.Print(buffer.GetData())
}

func (buffer *Buffer) SetFocus() {
	oldState, err := terminal.MakeRaw(0)

	if err != nil {
		panic(err)
	}

	defer terminal.Restore(0, oldState)

	for {
		str, _ := buffer.ui.ReadInputByte()

		if str == "q" {
			break
		}

		fmt.Print(str)
	}
}

func rawReadInput(f *os.File) []byte {
	var buf []byte

	for {
		_, err := f.Read(buf[:])

		if err != nil && err != io.EOF {
			break
		}

		if len(buf) > 0 {
			break
		}
	}

	return buf
}

func (buffer *Buffer) GetData() string {
	return strings.Join(buffer.data, "\n")
}
