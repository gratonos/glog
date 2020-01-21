package file

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gratonos/glog/internal/util"
	"github.com/gratonos/glog/pkg/glog/iface"

	gos "github.com/gratonos/goutil/os"
)

const (
	checkInterval = time.Second * 5
	dateFormat    = "%04d_%02d%02d"
	timeFormat    = "%02d%02d%02d.%09d"
	dirPerm       = 0770
)

var extensions = [...]string{
	iface.Binary: ".log.bin",
	iface.Text:   ".log.txt",
}

type Writer struct {
	config iface.FileWriter

	writer    io.WriteCloser
	nextDay   time.Time
	checkTime time.Time
	path      string
	fileSize  int64
}

func (this *Writer) Write(log []byte, tm time.Time) {
	err := this.checkFile(tm)
	if err == nil {
		var n int
		n, err = this.writer.Write(this.convert(log))
		this.fileSize += int64(n)
	}

	if err != nil && this.config.ErrorHandler != nil {
		this.config.ErrorHandler(tm, err)
	}
}

func (this *Writer) SetConfig(config iface.FileWriter) error {
	if !config.Enable {
		return this.closeFile()
	}
	if !config.Format.Legal() {
		return fmt.Errorf("illegal Format '%d'", config.Format)
	}
	if config.MaxFileSize <= 0 {
		return errors.New("MaxFileSize must be positive")
	}
	if err := this.checkDir(config.Dir); err != nil {
		return err
	}
	this.config = config
	return nil
}

func (this *Writer) checkDir(dir string) error {
	if dir == "" {
		return errors.New("Dir is empty")
	}
	if dir == this.config.Dir {
		return nil
	}
	if err := mkdir(dir); err != nil {
		return err
	}
	if err := this.closeFile(); err != nil {
		return err
	}
	return nil
}

func (this *Writer) convert(log []byte) []byte {
	switch this.config.Format {
	case iface.Binary:
		return log
	case iface.Text:
		text, err := util.BinaryToText(log, this.config.TextConfig.Coloring)
		if err != nil {
			panic(fmt.Sprintf("glog: corrupted log: %v", err))
		}
		return text
	default:
		panic(fmt.Sprintf("glog: illegal format '%d'", this.config.Format))
	}
}

func (this *Writer) checkFile(tm time.Time) error {
	if this.writer == nil ||
		this.fileSize >= this.config.MaxFileSize ||
		tm.Sub(this.nextDay) >= 0 {
		return this.createFile(tm)
	} else if tm.Sub(this.checkTime) >= checkInterval {
		this.checkTime = tm
		ok, err := gos.FileExists(this.path)
		if err != nil {
			return err
		}
		if !ok {
			return this.createFile(tm)
		}
	}
	return nil
}

func (this *Writer) createFile(tm time.Time) error {
	if err := this.closeFile(); err != nil {
		return err
	}

	dir := filepath.Join(this.config.Dir, dateStr(tm))
	if err := mkdir(dir); err != nil {
		return err
	}

	filename := clockStr(tm) + extensions[this.config.Format]
	path := filepath.Join(dir, filename)

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	this.writer = file
	this.nextDay = nextDay(tm)
	this.path = path
	this.fileSize = 0

	return nil
}

func (this *Writer) closeFile() error {
	if this.writer != nil {
		if err := this.writer.Close(); err != nil {
			return err
		}
		this.writer = nil
	}
	return nil
}

func nextDay(tm time.Time) time.Time {
	year, month, day := tm.Date()
	dayBegin := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return dayBegin.AddDate(0, 0, 1)
}

func dateStr(tm time.Time) string {
	year, month, day := tm.Date()
	return fmt.Sprintf(dateFormat, year, month, day)
}

func clockStr(tm time.Time) string {
	hour, minute, second := tm.Clock()
	return fmt.Sprintf(timeFormat, hour, minute, second, tm.Nanosecond())
}

func mkdir(dir string) error {
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return err
	}
	return syscall.Access(dir, 7 /* R_OK | W_OK | X_OK */)
}
