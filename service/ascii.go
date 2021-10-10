package server

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

type Ascii struct {
	Input  string
	Output string
	Fs     string
}

const (
	fileLen = 855
)

func isValid(s string) bool {
	for _, s := range s {
		if ((s < 32 && s != 10) || s > 126) && s != '\r' {
			return false
		}
	}
	return true
}

func GetArt(input, fs string) (string, error) {
	arr := []string{}
	if !isValid(input) {
		log.Println("non valid symbols")
		return "", errors.New("non valid symbols")
	}
	readFile, err := os.Open("fonts/" + fs + ".txt")
	defer readFile.Close()
	if err != nil {
		return "", err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}
	if len(arr) != fileLen {
		log.Println("Font file is corrupted")
		return "", errors.New("Font file is corrupted")
	}
	argsArr := strings.Split(input, "\r\n")
	ans := ""
	for _, ch := range argsArr {
		if ch == "" {
			ans += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				n := (j-32)*9 + 1
				ans += arr[int(n)+i]
			}
			ans += "\n"
		}
		ans += "\n"
	}
	return ans, nil
}
