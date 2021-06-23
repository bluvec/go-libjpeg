package util

import (
	"testing"

	"github.com/bluvec/go-libjpeg/test/util"
)

func TestOpenFile(t *testing.T) {
	for _, file := range util.SubsampledImages {
		util.OpenFile(file)
	}
}

func TestReadFile(t *testing.T) {
	for _, file := range util.SubsampledImages {
		util.ReadFile(file)
	}
}

func TestCreateFile(t *testing.T) {
	f := util.CreateFile("util_test")
	f.Write([]byte{'o', 'k'})
	f.Close()
}
