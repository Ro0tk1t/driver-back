package public

import "os"

const TmpDir = "/tmp/driver-back"

var StoreType = "minio"

func init() {
	if _, err := os.Stat(TmpDir); err != nil {
		os.Mkdir(TmpDir, os.ModePerm)
	}
}
