package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
	zipRC   *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath, nil}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if self.zipRC == nil {
		err := self.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	classFile := self.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)
	return data, self, err
}

func (self *ZipEntry) openJar() error {
	r, err := zip.OpenReader(self.absPath)
	if err == nil {
		self.zipRC = r
	}
	return err
}

func (self *ZipEntry) findClass(className string) *zip.File {
	for _, f := range self.zipRC.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	// read class data
	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (self *ZipEntry) String() string {
	return self.absPath
}

//func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
//	r, err := zip.OpenReader(self.absPath)
//	if err != nil { // 首先打开ZIP文件，如果这一步出错的话，直接返回
//		return nil, nil, err
//	}
//	defer r.Close()
//	for _, f := range r.File { // 遍历 ZIP压缩包里的文件，看能否找到class文件
//		if f.Name == className {
//			rc, err := f.Open() // 1. 打开 class文件
//			if err != nil {
//				return nil, nil, err
//			}
//			defer rc.Close() // 2.把内容读取出来
//			data, err := ioutil.ReadAll(rc)
//			if err != nil {
//				return nil, nil, err
//			}
//			return data, self, nil // 3. 返回
//		}
//	}
//	return nil, nil, errors.New("class not found: " + className)
//}
