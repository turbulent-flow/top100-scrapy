package test

func Finalize() {
	DBconn.Close()
	if Cleaner != nil {
		Cleaner.Close()
	}
}
