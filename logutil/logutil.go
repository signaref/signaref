// logutil project logutil.go
package logutil

import (
	"io"
	"log"
	"os"
	_ "runtime"
	"strconv"
	"time"
)

type Stringer interface {
	String() string
}

type LoggerUtil struct {
	*log.Logger
}

func NewLogger(out io.Writer, prefix string, flag int) *LoggerUtil {
	lu := new(LoggerUtil)
	lu.Logger = log.New(out, prefix, flag)
	return lu
}

func (lu *LoggerUtil) WriteLog(v ...interface{}) {
	y, m, d := time.Now().Date()
	file, err := os.OpenFile("logs/"+strconv.Itoa(y)+"-"+m.String()+"-"+strconv.Itoa(d), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf(err.Error())
	}
	lu.SetOutput(file)
	lu.Println(v)
}

func (lu *LoggerUtil) LogString(s Stringer) {
	if s.String() != "" {
		y, m, d := time.Now().Date()
		file, err := os.OpenFile("logs/"+strconv.Itoa(y)+"-"+m.String()+"-"+strconv.Itoa(d), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Printf(err.Error())
		}
		lu.SetOutput(file)
		lu.Println(s.String())
	}
}
