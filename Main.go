package main

import (
	services "github.com/ourcolour/aliacm/services"
	"io"
	"log"
	"os"
)

func initLogger(logFilePath string) *log.Logger {
	loggerFileHandler, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if nil != err {
		log.Fatal("Failed to create file.", err)
	}
	defer loggerFileHandler.Close()

	writerArray := []io.Writer{
		loggerFileHandler,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writerArray...)
	logger = log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime|log.Lshortfile)

	return logger
}

/**
 * Constants
 */
//time.Now().Format("20060102")
const LOG_FILE_PATH = "./log/" + "log" + ".log"

/**
 * Variables
 */
var logger *log.Logger

func main() {
	logger = initLogger(LOG_FILE_PATH)

	var endpoint = "acm.aliyun.com"
	var namespaceId = ""
	var accessKey = ""
	var secretKey = ""
	var group = "prd"

	client := services.NewAliACMSvs(
		endpoint,
		namespaceId,
		accessKey,
		secretKey,
		group,
	)

	var content = client.Load("last-fetch")

	logger.Println(content)
}
