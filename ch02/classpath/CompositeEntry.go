package classpath

/**
*entry 的实现类 CompositeEntry
 */
import (
	"errors"
)

type CompositeEntry []Entry

// newDirEntry方法 先把参数转换成绝对路径
func newCompositeEntry(pathList string) CompositeEntry {

}

//
func (self CompositeEntry) String() string {
	return self.String()
}

// 如何从ZIP文件中提取class文件
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	return nil, nil, errors.New("class not found: " + className)
}
