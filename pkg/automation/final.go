package automation

func Finalize() {
	DBpool.Close()
	SecondDBpool.Close()
}
