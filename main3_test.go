package main

import (
	helper "ascii-art-web/pkg/helper"
	"testing"
)

func TestCheckBannerName(t *testing.T) {
	if helper.CheckBannerName("standard") != true {
		t.Error("Error: there is problem in CheckBannerName function!\n")
	}
}

func TestCheckFileExtension(t *testing.T) {
	if helper.CheckFileExtension("doc") != true {
		t.Error("Error: there is problem in CheckFileExtension function!\n")
	}
}

func TestCheckFileName(t *testing.T) {
	if helper.CheckFileName("file.txt") != true {
		t.Error("Error: there is problem in CheckFileName function!\n")
	}
}

func TestConvertHexToInt(t *testing.T) {
	if helper.ConvertHexToInt("#7f7f7f") != 127 {
		t.Error("Error: there is problem in ConvertHexToInt function!\n")
	}
}

func TestCreatefile(t *testing.T) {
	if helper.Createfile("txt", "hello") != nil {
		t.Error("Error: there is problem in Createfile function!\n")
	}
}
