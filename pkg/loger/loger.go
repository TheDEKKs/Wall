package pkg

import (
	"log"
	"os"
)

var Log *log.Logger
var file *os.File

func InitLoger(file_log string) error{

	var err error
	file, err = os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil{
		log.Println("Error create file ", err)
		return err
	}
	defer file.Close()

	Logs := log.New(file, "ERROR: ", log.LstdFlags)


	Logs.Println("Good create")

	
	return nil
}