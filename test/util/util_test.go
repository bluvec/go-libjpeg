package util

import (
	"testing"
)

func TestOpenFile(t *testing.T) {
	for _, file := range SubsampledImages {
		OpenFile(file)
	}
}

func TestReadFile(t *testing.T) {
	for _, file := range SubsampledImages {
		ReadFile(file)
	}
}

func TestCreateFile(t *testing.T) {
	f := CreateFile("util_test")
	f.Write([]byte{'o', 'k'})
	f.Close()
}
