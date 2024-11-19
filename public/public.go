package public

import "os"

const TmpDir = "/tmp/driver-back"

var StoreType = "minio"

func init() {
	if _, err := os.Stat(TmpDir); err != nil {
		os.Mkdir(TmpDir, os.ModePerm)
	}
}

type FileType int

const (
	Directory FileType = iota
	Image
	Audio
	Text
)

var ftypeStr = []string{"directory", "image", "audio", "text"}

func (f FileType) String() string {
	return ftypeStr[f]
}

func (f FileType) Index() int {
	return int(f)
}
