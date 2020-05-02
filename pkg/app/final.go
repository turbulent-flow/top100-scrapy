package app

func Finalize() {
  if env == "development" {
    file.Close()
  }
  DBconn.Close()
}
