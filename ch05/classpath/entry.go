package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator) //常量pathListSeparator是string类型，存放路径分隔符

// Entry
// readClass 方法负责寻找和加载class文件；
// String 方法的作用相当于Java中的toString（），用于返回变量的字符串表示
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

// newEntry 根据参数创建不同类型的Entry实例
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
