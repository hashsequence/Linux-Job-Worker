package utils 

import (
	"strings"
	"time"
	"os"
	guuid "github.com/google/uuid"
)

func GetTimeStamp() string {
	return strings.ReplaceAll(time.Now().Format("20060102150405.000000.000000000"),".","")
}

func CreateFile(name string) error {
    file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    return file.Close()
}

func GetNewUUID() string {
	return guuid.New().String()
}

type LogWriter struct {
	file *os.File
}

func NewLogWriter(file *os.File) *LogWriter {
    lw := &LogWriter{
		file : file,
	}
    return lw
}

func (lw *LogWriter) Write (p []byte) (n int, err error) {
	lw.file.Write(p)
    return len(p), nil
}
