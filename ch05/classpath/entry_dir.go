package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry 只有一个字段，用于存放目录的绝对路径
type DirEntry struct {
	absDir string
}

/* newDirEntry 命名规范表示这是属于DirEntry的构造函数（也可去掉new，非强制要求）
先把参数转换成绝对路径，如果转换过程出现错误，则调用panic（）函数终止程序执行，否则创建DirEntry实例并返回*/
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

/* readClass 先把目录和class文件名拼成一个完整的路径，
然后调用ioutil包提供的ReadFile（）函数读取class文件内容，最后返回 */
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
