package main

import (
	ascii "ascii-art-web/pkg/ascii-art"
	"os"
	"strings"
	"testing"
)

func TestCheckString(t *testing.T) {
	a, err := ascii.CheckString("h")
	if a != true || err != nil {
		t.Error("Error: there is problem in CheckString function!\n")
	}
}

func TestGetBannerFile(t *testing.T) {
	file, _ := os.ReadFile("banners/standard.txt")
	text := string(file)
	check, err := ascii.GetBannerFile("banners/standard.txt")
	if check != text || err != nil {
		t.Error("Error: there is problem in GetBannerFile function!\n")
	}
}

func TestGetAscii(t *testing.T) {
	banner, _ := ascii.GetBannerFile("banners/standard.txt")
	lines := strings.Split(banner, "\n")
	check, err := ascii.GetAscii('h', 1, lines)
	if check != " _      " || err != nil {
		t.Error("Error: there is problem in GetAscii function!\n")
	}
}

func TestAsciiArt(t *testing.T) {
	bytes := []byte("hello")
	banner := "standard"
	expected :=
		` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`
	check, err := ascii.AsciiArt(bytes, banner)
	if check != expected || err != nil {
		t.Error("Expected does not match the actual output needed\n")
	}

}
