package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil { // 首先打开ZIP文件，如果这一步出错的话，直接返回
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File { // 遍历 ZIP压缩包里的文件，看能否找到class文件
		if f.Name == className {
			rc, err := f.Open() // 1. 打开 class文件
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close() // 2.把内容读取出来
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil // 3. 返回
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
