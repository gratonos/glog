package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gratonos/glog/internal/encoding/binary"
	"github.com/gratonos/glog/internal/encoding/text"
	"github.com/gratonos/glog/internal/writers/file"
	"github.com/gratonos/glog/pkg/glog/iface"
)

func processPath(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			errorf("processing %s: %v", path, err)
			return nil
		}

		name := info.Name()
		if info.IsDir() {
			if strings.HasPrefix(name, ".") && name != "." && name != ".." {
				return filepath.SkipDir
			}
		} else {
			ext := file.Extensions[iface.Binary]
			if !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ext) {
				taskChan <- conversionTask{
					Path:    path,
					ModTime: info.ModTime(),
				}
			}
		}

		return nil
	})
}

func processFile(path string, modTime time.Time) {
	outPath := toOutPath(path)
	outInfo, err := os.Stat(outPath)
	if err != nil || modTime.Sub(outInfo.ModTime()) > 0 {
		convertFile(path, outPath)
	}
}

func toOutPath(path string) string {
	inExt := file.Extensions[iface.Binary]
	outExt := file.Extensions[iface.Text]

	inBase := filepath.Base(path)
	var outBase string
	if strings.HasSuffix(inBase, inExt) {
		outBase = inBase[:len(inBase)-len(inExt)] + outExt
	} else {
		outBase = inBase + outExt
	}

	return filepath.Join(filepath.Dir(path), outBase)
}

func convertFile(inPath, outPath string) {
	inFile, err := os.Open(inPath)
	if err != nil {
		errorf("processing %s: %v", inPath, err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(outPath)
	if err != nil {
		errorf("processing %s: %v", inPath, err)
		return
	}
	defer outFile.Close()

	in, out := bufio.NewReader(inFile), bufio.NewWriter(outFile)
	defer out.Flush()

	convert(in, out, context.WithValue(context.Background(), "path", inPath))
}

func convert(in *bufio.Reader, out *bufio.Writer, ctx context.Context) {
	path := ctx.Value("path")

	var readErr, writeErr error
	for {
		var record binary.Record
		if readErr == nil {
			readErr = binary.ReadRecord(&record, in)
		} else {
			readErr = binary.TryReadRecord(&record, in)
		}
		if readErr == binary.EOF {
			infof("processing %s ... done", path)
			return
		}
		if readErr == nil {
			_, writeErr = out.Write(text.FormatRecord(&record, flagColoring))
		} else {
			if ioErr, ok := readErr.(*binary.IOError); ok {
				errorf("processing %s: %v", path, ioErr)
				return
			} else {
				warnf("processing %s: corrupted log, locating next log ...", path)
				log := "!!!!!!!! one or more corrupted logs !!!!!!!!"
				if flagColoring {
					log = fmt.Sprintf("%s%s%s", text.Magenta, log, text.Reset)
				}
				_, writeErr = out.WriteString(log + "\n")
			}
		}
		if writeErr != nil {
			errorf("processing %s: %v", path, writeErr)
			return
		}
	}
}
