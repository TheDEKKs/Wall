package pkg

import (
	"log"
	"os"

	"go.uber.org/zap"
)


//Нужно для вызова логирования
var Zap *zap.Logger



func InitLoger(file_log string) error{
	var err error
	//Оу я тут открываю фаял
	file, err := os.OpenFile(file_log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil{
		log.Println("Error create file ", err)
		return err
	}

	defer file.Close()
	

	//Здесь я настраиваю логирование 
	loger := zap.NewProductionConfig()
	loger.OutputPaths = []string{"stdout", "./log/.log"}

	//Дальше я уже фиг знаю что делал (я был слишком уставшим)
	Zap, err = loger.Build()

	if err != nil {
		log.Panic(err)
		return err
	}

	defer Zap.Sync()

	return nil

	/*path, _ := os.Executable()var err error

	log.Println(file, "Test", path+file_log)

	defer file.Close()

	log.SetOutput(file)

	Log = log.New(file, "ERROR: ", log.LstdFlags)

	
	Log.Println("Test")
	
	return nil
	*/
}