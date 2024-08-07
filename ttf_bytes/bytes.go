package ttf_bytes

import (
	"os"
	"path"
	"runtime"
)

var (
	NotoSansTCRegular []byte
	NotoSansTCBold    []byte
	NotoSansSCRegular []byte
	NotoSansSCBold    []byte
)

var (
	filepath string
)

func init() {
	initFilepath()
	NotoSansTCRegular = getBytes("NotoSansTC-Regular.ttf")
	NotoSansTCBold = getBytes("NotoSansTC-Bold.ttf")
	NotoSansSCRegular = getBytes("NotoSansSC-Regular.ttf")
	NotoSansSCBold = getBytes("NotoSansSC-Bold.ttf")
}

func initFilepath() {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		filepath = path.Dir(filename) + "/"
	} else {
		panic("Cannot get current file path")
	}
}

func getBytes(filename string) []byte {
	b, err := os.ReadFile(filepath + filename)
	if err != nil {
		panic(err)
	}
	return b
}
