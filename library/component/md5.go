package component

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	path = string(path[0:strings.LastIndex(path, "/")])
	return path
}

func Md5sum(path string) (err error, md5str string) {
	file, err := os.Open(path)
	if err == nil {
		md5h := md5.New()
		_, _ = io.Copy(md5h, file)
		fmt.Printf("%x", md5h.Sum([]byte(""))) //md5
		md5str = fmt.Sprintf("%x", md5h.Sum([]byte("")))
		return nil, md5str
	}
	return err, "111"
}
