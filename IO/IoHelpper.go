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

func (f *FRead) Rread(Fstr string) FRead {
	f1, _ := FileHelpper.RFile(Fstr)
	f.Fi = f1

	f.Read = bufio.NewReaderSize(f1, 4096*3)
	return *f
}
