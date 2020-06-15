package app

import "github.com/LiamYabou/top100-scrapy/v2/variable"

func Finalize() {
  if variable.Env == "development" {
    file.Close()
  }
  DBpool.Close()
  AMQPconn.Close()
}
