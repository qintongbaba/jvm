//解析classpath为zip

package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

//实现接口Entry
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath) //打开zip
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()

	//遍历zip中的文件
	for _, f := range r.File {
		if f.Name == className { //如果改文件为需要读取的类
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)

}

func (self *ZipEntry) String() string {
	return self.absPath
}
