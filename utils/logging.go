package utils

//ログファイルにはシステムの状態やイベント、エラー、アクティビティなどの情報が記録される
//webapp.log

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer logfile.Close()

	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//設定するフラグの定数はあらかじめ決められている, 他にもいくつかある.
	log.SetOutput(multiLogFile)
	//コンソールへのlogの出力時にファイルにも出力する(出力先を決定)
}