package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	classpath := &Classpath{}
	classpath.parseBootAndExtClasspath(jreOption)
	classpath.parseUserClasspath(cpOption)
	return classpath
}
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, nil
	}

	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, nil
	}

	return self.userClasspath.readClass(className)

}
func (self *Classpath) String() string {
	return self.userClasspath.String()
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {

	jreDir := getJreDir(jreOption)

	//jre/lib/*
	jreDirPath := filepath.Join(jreDir, "lib", "*")

	self.bootClasspath = newWildcardEntry(jreDirPath)

	//jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")

	self.extClasspath = newWildcardEntry(jreExtPath)

}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists(". /jre") {
		return ". /jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder!")
}

//判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
