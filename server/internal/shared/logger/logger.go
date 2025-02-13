package logger

import (
	"io"
	"os"
)

type LoggerWriter struct {
	destinationFile string

	file        *os.File
	multiWriter io.Writer
}

func NewLoggerWriter(destinationFile string) (LoggerWriter, error) {
	file, err := os.OpenFile(destinationFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return LoggerWriter{}, err
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	return LoggerWriter{destinationFile: destinationFile, file: file, multiWriter: multiWriter}, nil
}

func (l LoggerWriter) Write(p []byte) (n int, err error) {
	return l.multiWriter.Write(p)
}

func (l LoggerWriter) Close() error {
	return l.file.Close()
}
