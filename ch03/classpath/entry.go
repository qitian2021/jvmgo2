package classpath

/**
* 类路径接口
**/

import (
	"os"
	"strings"
)

// 常量pathListSeparator 是string 类型，存放路径分隔符
const pathListSeparator = string(os.PathListSeparator)

// Entry接口中两个方法
type Entry interface {
	// param readClass方法参数-class文件的相对路径 路径之间用（/）例如：读取java.lang.Object传入参数：java/lang/Object.class
	// return: 读取到的字节数据、最终定位到class文件的Entry,以及错误信息
	readClass(className string) ([]byte, Entry, error) // readClass方法 负责寻找和加载class文件
	String() string                                    // 类似java 中的toString()
}

func newEntry(path string) Entry {

	//...
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
