package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator) // 路径分隔符

type Entry interface {
	readClass(className string) ([]byte, Entry, error) //寻找和加载class文件
	String() string                                    //可以理解为java中toString()方法
}

//根据传入的path创建不同的Entry
func newEntry(path string) {

	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(pathList)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
