// Filewriter implements buffered I/O operations
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Application log file contains events if BACKEND is set "file"
var logfile *LogFile

// customTimeLayout (set to empty to unset)
var customTimeLayout = "15:04:05 - 02.01.2006"

// FileWriter type
type FileWriter struct {
	file       *os.File      // File descriptor.
	writer     *bufio.Writer // Buffered writer bufio instance.
	filename   string        // The path to the file.
	bufferSize int           // The buffer size for buffered I/O.
}

// LogFile type operates an append only file to store events
type LogFile struct {
	fileWriter *FileWriter // File writer used to operate the underlaying file
	TimeLayout string      // Sets custom time layout for TimeNow()
	enabled    bool        // Specifies if the logfile is enabled
}

// NewFileWriter creates a new instance of FileWriter with the given filename
// and buffer size.
//
// trunc variadic parameter when set to true enables trancation of the file
func NewFileWriter(filename string, bufferSize int, trunc ...bool) *FileWriter {
	var file *os.File
	var err error

	if len(trunc) > 0 && trunc[0] {
		file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	} else {
		file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	}

	if err != nil {
		logger.Panic(err)
	}

	return &FileWriter{
		filename:   filename,
		bufferSize: bufferSize,
		file:       file,
		writer:     bufio.NewWriterSize(file, bufferSize),
	}
}

// Syncs and close the FileWriter
func (fw *FileWriter) Sync() {
	err := fw.writer.Flush()
	if err != nil {
		logger.Error(err)
	}
	err = fw.file.Close()
	if err != nil {
		logger.Error(err)
	}
}

// Writes bytes array into the buffer writer
func (fw *FileWriter) Write(data []byte) error {
	_, err := fw.writer.Write(data)
	return err
}

// Writes are new line into the buffer writer
func (fw *FileWriter) WriteString(line string) error {
	_, err := fw.writer.WriteString(line)
	return err
}

// Creates a new LogFile instance to write events to a buffered log file
// Note that Sync() has to invoked before exist to preserve the data.
//
// By this we achieve basic optimization, since we couldn't use an efficient
// logger due to the task requirements.
func NewLogFile(filename string) *LogFile {
	bufferSize := 1024 // 1024 byte buffer
	lf := &LogFile{
		fileWriter: NewFileWriter(filename, bufferSize),
		TimeLayout: time.Layout,
		enabled:    true,
	}
	return lf
}

// Flush and Sync the LogFile
func (lf *LogFile) Sync() {
	lf.fileWriter.Sync()
}

// Write a line into the log
func (lf *LogFile) Writeln(a ...any) error {
	if !lf.enabled {
		return nil
	}
	return lf.fileWriter.WriteString(fmt.Sprintln(a...))
}

// Write a formatted line line into the log
func (lf *LogFile) Writef(format string, a ...any) error {
	if !lf.enabled {
		return nil
	}
	return lf.fileWriter.WriteString(fmt.Sprintf(format, a...))
}

// Write a formatted line line into the log
func (lf *LogFile) TimeNow() string {
	return time.Now().Format(lf.TimeLayout)
}

func initLogFile() {
	logfile = NewLogFile(appConfig.FilePath)
	if customTimeLayout != "" {
		logfile.TimeLayout = customTimeLayout
	}

	if strings.ToLower(appConfig.Backend) != "logfile" {
		logfile.enabled = false
		return
	}

	logger.Warnw("Initialized logfile backend", "filepath", appConfig.FilePath)
}
