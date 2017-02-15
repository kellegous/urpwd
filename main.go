package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// GetHome ...
func GetHome() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return u.HomeDir, nil
}

// Err ...
func Err(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

// Dir ...
func Dir() string {
	if flag.NArg() > 0 {
		return flag.Arg(0)
	}

	d, err := os.Getwd()
	if err != nil {
		Err(err.Error())
	}

	return d
}

// Split ...
func Split(home, path string, n int) []string {
	vals := make([]string, n+1)
	var file string

	for i := len(vals) - 1; i > 0; i-- {
		path = strings.TrimRight(path, "/")

		if path == "" {
			return vals[i:]
		}

		if path == home {
			vals[i] = "~"
			return vals[i:]
		}

		path, file = filepath.Split(path)
		vals[i] = file
	}

	if path != "/" {
		vals[0] = "..."
	}

	return vals
}

func main() {
	flagN := flag.Int("n", 2, "the number of path components to take from the end")
	flag.Parse()

	home, err := GetHome()
	if err != nil {
		Err(err.Error())
	}

	path := Dir()
	if path == "/" {
		fmt.Println("/")
	} else {
		fmt.Println(strings.Join(Split(home, path, *flagN), "/"))
	}
}
