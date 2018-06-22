package util

type InitUtil struct {
	log LogStruct
	// db  iface.DataBase
}

func (iu *InitUtil) LogInit() LogStruct {
	log := LogStruct{}
	log.Config("127.0.0.1:5000")
	defer log.Flush()

	return log
}
