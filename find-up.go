package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var appname = os.Args[0]

func main() {
	if len(os.Args) != 2 {
		if strings.HasPrefix(appname, "/tmp") {
			appname = "go run find-up.go"
		}
		fmt.Fprintf(os.Stderr, "Usage: %s <name>", appname)
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Looks for the named directory entry, climbing up the directory tree until")
		fmt.Fprintln(os.Stderr, "either finding it, reaching your home dir, or reaching the root of the")
		fmt.Fprintln(os.Stderr, "filesystem.")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintf(os.Stderr, `e.g., "%s .git" would find the git directory "/foo/bar/.git" if you were `, appname)
		fmt.Fprintf(os.Stderr, `in the subdir "/foo/bar/baz/quux".`)
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr)
		os.Exit(255)
	}

	var needle = os.Args[1]
	var wd, err = os.Getwd()
	if err != nil {
		log.Printf("%s: error getting current dir: %s", appname, err)
		os.Exit(1)
	}

	var home string
	home, err = os.UserHomeDir()
	if err != nil {
		log.Printf("%s: error getting home dir: %s", appname, err)
		os.Exit(1)
	}

	var dir string
	dir, err = find(home, wd, needle)
	if err != nil {
		os.Exit(0)
	}
	fmt.Println(dir)
}

func find(home, dir, needle string) (string, error) {
	home = filepath.Clean(home)
	var fullpath = filepath.Join(dir, needle)
	var _, err = os.Stat(fullpath)
	if err != nil {
		if os.IsNotExist(err) {
			var newdir, _ = filepath.Split(dir)
			newdir = filepath.Clean(newdir)
			if newdir == dir {
				return "", os.ErrNotExist
			}
			return find(home, newdir, needle)
		}
		log.Printf("%s: error statting %q: %s", appname, fullpath, err)
		return "", err
	}

	return fullpath, nil
}
