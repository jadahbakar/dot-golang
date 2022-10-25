package logger

import (
	"fmt"
	"os"
)

type FiberLog struct {
	file *os.File
	err  *error
}

type FiberLogger interface {
	Write(path string) (*os.File, error)
}

// jika ingin advance lagi bisa menggunakan interface,
func NewFiberLogger(path string) (*os.File, error) {
	err := createFolderLogIfNotExist(path)
	if err != nil {
		return nil, err
	}

	return os.OpenFile(fmt.Sprintf("%s/%s", path, "fiber.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

func createFolderLogIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
