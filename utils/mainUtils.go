package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getLogFileName() string {
	return fmt.Sprintf("log%s.txt", time.Now().Format("2006-01-02_15-04-05"))
}

func readChoice(ch string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(ch))
}

func CreateLogFiles() (*os.File, error) {
	dir := "logs"

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("The error occured while creating logs dir: %s", err.Error())
		}
	}

	fileName := getLogFileName()

	file, err := os.Create("logs/" + fileName)
	if err != nil {
		return nil, fmt.Errorf("The error occured while creating file %s: %s", fileName, err.Error())
	}

	return file, nil
}

func ReadInt(reader *bufio.Reader) (int, error) {
	n, err := reader.ReadString('\n')
	if err != nil {
		return 0, errors.New("Enter error: " + err.Error())
	}

	value, err := readChoice(n)
	if err != nil {
		return 0, errors.New("Convert error: " + err.Error())
	}

	return value, nil
}

func ReadString(reader *bufio.Reader) (string, error) {
	n, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("Enter error: " + err.Error())
	}

	return strings.TrimSpace(n), nil
}
