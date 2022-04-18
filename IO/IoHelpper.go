package IO

import (
	"bufio"
	"github.com/tianfei212/units/IO/FileHelpper"
	"os"
)

//////////////////////////
//IO/IoHelpper.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 3/25/2022 15:03
// note =
/////////////////////////

type FRead struct {
	Fi   *os.File
	Read *bufio.Reader
}
type FWead struct {
	Fi    *os.File
	Write *bufio.Writer
}

var fh FileHelpper.FileH

func (f *FRead) Rread(fp string) FRead {
	fh = FileHelpper.FileH{}
	f1, _ := fh.Open(fp)
	f.Fi = f1

	f.Read = bufio.NewReaderSize(f1, 4096*3)
	return *f
}
func (f *FWead) Cread(fp string) FWead {
	fh = FileHelpper.FileH{}
	f1, _ := fh.OpenFile(fp)
	f.Fi = f1

	f.Write = bufio.NewWriterSize(f1, 4096*3)
	return *f
}
