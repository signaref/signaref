// utils project utils.go
package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//To validate file extensions.. only jpg,jpeg and png are allowd.
//resizer works only with jpeg and jpg at this point of time.
func ValidateFileTypes(ext string) (err error) {
	if strings.ToLower(ext) == ".jpg" || strings.ToLower(ext) == ".jpeg" || strings.ToLower(ext) == ".png" {
		return nil
	}
	return errors.New("invalid file type.only jpg,jpeg and png formats are allowed")
}

func GetMagicNumbers(max, min, currentCount, noToFetch int) ([]string, error) {
	if min > max {
		return nil, errors.New("min cannot be greater than max value")
	}
	if min > currentCount || max > currentCount {
		return nil, errors.New("min or max values cannot be greater than current file number")
	}
	nums := make([]string, noToFetch)
	j := 0
	for i := currentCount; i >= 1; i-- {
		if i < min || i > max {
			if j == noToFetch {
				break
			} else {
				nums[j] = strconv.Itoa(i)
			}
			j++
		}
	}
	return nums, nil
}

//get the requested ip address
func GetIpAddr(r *http.Request) string {
	ip := r.Header.Get("x-forwarded-for")
	if ip == "" || len(ip) == 0 {
		ip = r.Header.Get("Proxy-Client-IP")
	}
	if ip == "" || len(ip) == 0 {
		ip = r.Header.Get("WL-Proxy-Client-IP")
	}
	if ip == "" || len(ip) == 0 {
		ip = r.RemoteAddr
	}
	return ip
}

func trace(i int) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s\n", file, line, f.Name())
}

func WhereAmI(depthList ...int) {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, _, line, _ := runtime.Caller(depth)
	fmt.Printf(" Function: %s Line: %d", runtime.FuncForPC(function).Name(), line)
	//return fmt.Sprintf(" Function: %s Line: %d", runtime.FuncForPC(function).Name(), line)
}

//retuns guid
func GUID() string {
	// generate 32 bits timestamp
	unix32bits := uint32(time.Now().UTC().Unix())
	buff := make([]byte, 12)
	rand.Read(buff)
	return fmt.Sprintf("%x-%x-%x-%x-%x-%x", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8], buff[8:])
}
