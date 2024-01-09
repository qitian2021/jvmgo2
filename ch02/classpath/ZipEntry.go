package classpath

/**
*entry 的实现类 ZipEntry
 */
import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string // 存放ZIP或JAR的绝对路径
}

// newDirEntry方法 先把参数转换成绝对路径
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil { // 转换出错
		panic(err) // 调用panic 函数终止程序执行
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) String() string {
	return self.absPath
}

// 如何从ZIP文件中提取class文件
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath) // 首先打开ZIP文件
	if err != nil {
		return nil, nil, err // 出错直接返回
	}
	defer r.Close()            // 确保打开的文件关闭
	for _, f := range r.File { // 遍历ZIP压缩包里的文件，看能否找到class文件
		if f.Name == className { //
			rc, err := f.Open() // 找到文件打开
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc) // 读取文件
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil // 返回读取的文件
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
