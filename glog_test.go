package glog

import (
	"math/rand"
	"testing"
	"time"
)

func TestRotateNone(t *testing.T) {
	options := LogOptions{
		File:  "./abc.log",
		Flag:  LstdFlags,
		Level: Ldebug,
		Mode:  R_None,
	}
	logger, err := New(options)
	if err != nil {
		t.Error(err)
	}
	name := "Mr Poirot"
	logger.Debugf("hello, %s", name)
	logger.Info("testing message")
	logger.Warn("testing message")
	logger.Error("testing message")
	logger.Fatal("oops, fatal error!")
}

func TestRotateSize(t *testing.T) {
	options := LogOptions{
		File:    "./abc.log",
		Flag:    LstdFlags | Lmicroseconds | Lshortfile,
		Level:   Ldebug,
		Mode:    R_Size,
		Maxsize: 1024 * 1024,
	}
	logger, err := New(options)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 10000; i++ {
		s := RandStringBytesRmndr(rand.Int() % 1024)
		logger.Info(s)
	}

	logger.Flush()
}

func TestRotateHour(t *testing.T) {
	options := LogOptions{
		File:    "./abc.log",
		Flag:    LstdFlags | Lmicroseconds | Lshortfile,
		Level:   Ldebug,
		Mode:    R_Hour,
		Maxsize: 1024 * 1024,
	}
	logger, err := New(options)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 10000; i++ {
		s := RandStringBytesRmndr(rand.Int() % 1024)
		logger.Info(s)
		time.Sleep(3 * time.Millisecond)
	}

	logger.Flush()
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
