// ref from https://github.com/spiegel-im-spiegel/logf

package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger struct {
	mu  sync.Mutex
	lg  *log.Logger
	min Level
}

func new(opts ...OptFunc) *Logger {
	l := &Logger{lg: log.New(os.Stderr, "", log.LstdFlags)}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *Logger) SetMinLevel(lv Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.min = lv
}

func (l *Logger) lprintf(lv Level, format string, v ...interface{}) {
	_ = l.output(lv, 4, fmt.Sprintf(format+"\n", v...))
}

func (l *Logger) output(lv Level, calldepth int, s string) error {
	if lv >= l.min {
		return l.lg.Output(calldepth, fmt.Sprintf("[%v] %s", lv, s))
	}
	return nil
}

func (l *Logger) Printf(format string, v ...interface{}) { l.lprintf(INFO, format, v...) }

type OptFunc func(*Logger)

var logger = new(func(l *Logger) {
	// common settings
	l.lg.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
	l.lg.SetOutput(os.Stdout)
	l.lg.SetPrefix(fmt.Sprintf("[Proc: %d]", os.Getpid()))
	l.lg.Println("Initialized Logger!")
})

//Tracef calls std.Tracef() to print to the logger.
func Tracef(format string, v ...interface{}) { logger.lprintf(TRACE, format, v...) }

//Debugf calls logger.Debugf() to print to the logger.
func Debugf(format string, v ...interface{}) { logger.lprintf(DEBUG, format, v...) }

//Printf calls logger.Printf() to print to the logger.
func Printf(format string, v ...interface{}) { logger.lprintf(INFO, format, v...) }

//Warnf calls logger.Warnf() to print to the logger.
func Warnf(format string, v ...interface{}) { logger.lprintf(WARN, format, v...) }

//Errorf calls logger.Errorf() to print to the logger.
func Errorf(format string, v ...interface{}) { logger.lprintf(ERROR, format, v...) }

//Fatalf calls logger.Fatalf() to print to the logger.
func Fatalf(format string, v ...interface{}) { logger.lprintf(FATAL, format, v...) }

//Panicf is equivalent() to logger.Output() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	_ = logger.output(FATAL, 4, s)
	panic(s)
}

//Panic is equivalent() to logger.output() followed by a call to panic().
func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	_ = logger.output(FATAL, 4, s)
	panic(s)
}

//Panicln is equivalent() to logger.output() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	_ = logger.output(FATAL, 4, s)
	panic(s)
}

func init() {
	if os.Getenv("USER") != "takeru" {
		logger.SetMinLevel(INFO)
	} else {
		logger.SetMinLevel(DEBUG)
	}
}

/* Copyright 2018,2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
