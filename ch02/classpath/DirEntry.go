package classpath

/**
*entry 的实现类
 */
import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

// newDirEntry方法 先把参数转换成绝对路径
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil { // 转换出错
		panic(err) // 调用panic 函数终止程序执行
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className) // 将目录和class文件名拼成完整的路径
	data, err := ioutil.ReadFile(fileName)            // 读取class文件内容
	return data, self, err
}

// 返回目录
func (self *DirEntry) String() string {
	return self.absDir
}
