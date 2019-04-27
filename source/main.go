package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
)

const (
	EXIT_CODE_POSITIVE = 0
	EXIT_CODE_NEGATIVE = 1
	EXIT_CODE_ERROR = 2
)

func positive () {
	fmt.Printf("TRUE: Running in a container.\r\n")
	os.Exit(EXIT_CODE_POSITIVE)
}

func negative () {
	fmt.Printf("FALSE: Not running in a container.\r\n")
	os.Exit(EXIT_CODE_NEGATIVE)
}

func checkUnix () {

	data, err := ioutil.ReadFile("/proc/1/cgroup")
	if err != nil {
		if os.IsNotExist(err) {
			negative()
		} else {
			log.Fatal(err)
			os.Exit(EXIT_CODE_ERROR)
		}
	}

	matched1, err := regexp.Match(`.*/docker/.*`, data)
	if err != nil {
		log.Fatal(err)
		os.Exit(EXIT_CODE_ERROR)
	}
	if matched1 {
		positive()
	}

	matched2, err := regexp.Match(`.*/lxc/.*`, data)
	if err != nil {
		log.Fatal(err)
		os.Exit(EXIT_CODE_ERROR)
	}
	if matched2 {
		positive()
	}

	negative()
}

func main () {
	switch kernel := runtime.GOOS; kernel {
	// case "aix":
	// 	checkUnix()
	// case "android":
	// 	checkUnix()
	case "darwin":
		checkUnix()
	// case "dragonfly":
	// 	checkUnix()
	// case "freebsd":
	// 	checkUnix()
	// case "js":
	case "linux":
		checkUnix()
	// case "nacl":
	// case "netbsd":
	// 	checkUnix()
	// case "openbsd":
	// 	checkUnix()
	// case "plan9":
	// case "solaris":
	// 	checkUnix()
	// case "windows":
	default:
		fmt.Printf("Unsupported operating system: %s.\r\n", kernel)
		os.Exit(EXIT_CODE_ERROR)
	}
}