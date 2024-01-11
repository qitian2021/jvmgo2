package classpath

/**
*entry 的实现类 CompositeEntry
 */
import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

// newDirEntry方法 先把参数转换成绝对路径
// 构造函数把参数（路径列表）按分隔符分成小路径，然后把每个小路径都转换成具体的Entry实例
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

// 依次调用每一个子路径的readClass（）方法，如果成功读取到class数据，返回数据即可
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

// 调用每一个子路径的String（）方法，然后把得到的字符串路用路径分隔符拼接起来
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
