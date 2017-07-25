package classpath

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}
