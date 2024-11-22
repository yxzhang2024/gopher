package main

import (
	"field_con/library3"
)

func main() {

	Conf := library3.MyTLSConfig{InsecureSkipVerify: true}
	Conf.Conn()

}
