package classpath

import (
	"os"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath                         {}
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {}
func (self *Classpath) String() string                                    {}

//判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
