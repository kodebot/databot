package logger

import (
	"fmt"

	"github.com/golang/glog"
)

// Tracef records trace level log
func Tracef(format string, args ...interface{}) {
	// glog.Infof("TRACE>>>>"+format, args...)
}

// Infof records info level log
func Infof(format string, args ...interface{}) {
	//	glog.Infof(format, args...)
}

// Warnf records warning level log
func Warnf(format string, args ...interface{}) {
	//	glog.Warningf(format, args...)
}

// Errorf records errors level log
func Errorf(format string, args ...interface{}) {
	glog.ErrorDepth(1, fmt.Sprintf(format, args...))
}

// Fatalf records fatal level log
func Fatalf(format string, args ...interface{}) {
	glog.FatalDepth(1, fmt.Sprintf(format, args...))
}
