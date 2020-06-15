package test

func Finalize() {
	DBpool.Close()
	PQconn.Close()
	if Cleaner != nil {
		Cleaner.Close()
	}
}
