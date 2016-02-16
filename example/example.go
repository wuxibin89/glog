package main

import (
    "time"
    "math/rand"
    "github.com/chasex/glog"
)

func main() {
    options := glog.LogOptions{
	File: "./abc.log",
	Flag: glog.LstdFlags | glog.Lmicroseconds | glog.Lshortfile,
	Level: glog.Ldebug,
	Mode: glog.R_Size,
	Maxsize: 1024 * 1024 * 16,
    }
    logger, err := glog.New(options)
    if err != nil {
	panic(err)
    }

    con := 100
    done := make(chan int, con)
    for k := 0; k < con; k++ {
	go func() {
	    for i := 0; i < 30000; i++ {
		s := RandStringBytesRmndr(rand.Int() % 1024)
		logger.Info(s)
		time.Sleep(time.Duration(rand.Int() % 10) * time.Millisecond)
	    }
	    done <- 1
	}()
    }

    for i := 0; i < con; i++ {
	<-done
    }

    logger.Flush()
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func RandStringBytesRmndr(n int) string {
    b := make([]byte, n)
    for i := range b {
	b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}
