package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "Error opeing File")
	}

	if err := os.Remove(pidFile); err != nil {
		fmt.Println("Error closing file")
	}
	strPid := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPid)
	if err != nil {
		return errors.Wrap(err, "Invalid Pid format")
	}
	fmt.Printf("Killing Process : %d \n", pid)
	return nil
}

func main() {

	if err := killServer("pidFile.txt"); err != nil {
		fmt.Println("Error Killig server : ", err.Error())
	}
}
