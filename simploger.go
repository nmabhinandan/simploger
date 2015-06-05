// Package simploger provides a simple-to-use, level based logging functionality to go programs.
// This package uses http://github.com/fatih/color package to display colored output in consoles.
// This is a simplest example of simploger
//
//   func main() {
// 	  sl := &simploger.Simplogger {
// 		  Verbosity: 1,
// 		  Logfile: simploger.Logfile{
// 			  Win: "C:\\MyApp\\logs",
// 			  Nix: "var/log/myapp/logs",
// 		  },
// 	  },
//   }
//
package simploger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"

	"gopkg.in/fatih/color.v0"
)

// Logfile is the structure of path of the logfile in different os environments
type Logfile struct {
	// Win is the log file location for Windows based os'
	Win string
	// Nix is the log file location for Unix based os'
	Nix string
}

// Simplogger is the main struct of simploger package
type Simplogger struct {
	// Verbosity is the level of logs to be printed on console
	// -1: Quiet mode
	// 1: Prints only Error and ForceInfo level logs
	// 2: Prints only Error, Warn and ForceInfo level logs
	// 3: Prints logs of all levels
	Verbosity int
	Logfile
}

// Info logs information level logs
func (l *Simplogger) Info(msgs ...string) {

	raw := make([]string, len(msgs))
	copy(raw, msgs)

	for key, msg := range msgs {
		msgs[key] = color.WhiteString("[INFO] ") + color.WhiteString(msg)
		raw[key] = "[INFO] " + msg
	}

	go l.writeToFile(raw...)
	if l.Verbosity > 1 {
		l.handle(os.Stdout, msgs...)
	}
}

// ForceInfo logs information level logs unless in quite mode
func (l *Simplogger) ForceInfo(msgs ...string) {

	raw := make([]string, len(msgs))
	copy(raw, msgs)

	for key, msg := range msgs {
		msgs[key] = color.WhiteString("[INFO] ") + color.WhiteString(msg)
		raw[key] = "[INFO] " + msg
	}

	go l.writeToFile(raw...)
	if l.Verbosity > -1 {
		l.handle(os.Stdout, msgs...)
	}
}

// Warn logs waning level logs
func (l *Simplogger) Warn(msgs ...string) {

	raw := make([]string, len(msgs))
	copy(raw, msgs)

	for key, msg := range msgs {
		msgs[key] = color.YellowString("[WARN] ") + color.WhiteString(msg)
		raw[key] = "[WARN] " + msg
	}

	go l.writeToFile(raw...)
	if l.Verbosity > 0 {
		l.handle(os.Stdout, msgs...)
	}
}

// Err logs error level logs
func (l *Simplogger) Err(msgs ...string) {

	raw := make([]string, len(msgs))
	copy(raw, msgs)

	for key, msg := range msgs {
		msgs[key] = color.RedString("[ERROR] ") + color.WhiteString(msg)
		raw[key] = "[ERROR] " + msg
	}

	go l.writeToFile(raw...)
	if l.Verbosity > -1 {
		l.handle(os.Stderr, msgs...)
	}
}

// prints the logs to the console
func (l *Simplogger) handle(w io.Writer, msgs ...string) {

	for _, msg := range msgs {
		fmt.Fprintf(w, "[%d-%d-%d] [%d:%d:%d] %s\n", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), msg)
	}

}

// writes the logs to the logfile
func (l *Simplogger) writeToFile(msgs ...string) {
	var p string
	if runtime.GOOS == "windows" {
		p = path.Join(l.Logfile.Win + strconv.Itoa(time.Now().Year()) + ".log")
	} else {
		p = path.Join(l.Logfile.Nix + strconv.Itoa(time.Now().Year()) + ".log")
	}

	f, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	for _, msg := range msgs {
		fmt.Fprintf(f, "[%d-%d-%d] [%d:%d:%d] %s\n", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), msg)
	}
}
