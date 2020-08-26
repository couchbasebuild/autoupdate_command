package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/akamensky/argparse"
)

func otherMain() (err error) {
	return
}

func autoupdMain() (err error) {
	parser := argparse.NewParser("", "Wraps a command-line tool with auto-update behaviour")
	s := parser.String("", "name", &argparse.Options{Required: true, Help: "Name of tool"})
	err = parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return nil
	}
	fmt.Println(*s)
	err = doInstall(*s)
	return
}

func main() {
	// We only parse our own arguments when invoked as "autoupdate_command".
	// Otherwise we'll pass through all arguments to the underlying command.
	// Note: Don't use os.Executable() here, as that will resolve symlinks
	// to the underlying binary, which is exactly what we DON'T want!
	mycmd := strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0]))
	fmt.Println(mycmd)

	var err error
	if mycmd == "autoupdate_command" {
		err = autoupdMain()
	} else {
		err = otherMain()
	}

	if err != nil {
		panic(err)
	}
}
