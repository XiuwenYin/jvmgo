package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

/* newCompositeEntry 构造函数把参数(路 径列表)按分隔符分成小路径，然后把每个小路径都转换成具体的 Entry实例 */
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

/*依次调用每一个子路径的readClass()方法，
如果成功读取到class数据，返回数据即可;
如果收到错误信息，则继续;
如果遍历完所有的子路径还没有找到class文件，则返回错误*/
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, form, err := entry.readClass(className)
		if err == nil {
			return data, form, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

/*调用每一个子路径的String()方法，然 后把得到的字符串用路径分隔符拼接起来即可 */
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
