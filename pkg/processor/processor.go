package processor

type ALFolderProcessor interface {
	ProcessAllConfigsForALFolder(alCode string, configs []interface{}) error
	Accepts(config interface{}) bool
}

type ConfigProcessor interface {
	Process(config interface{}) error
}
