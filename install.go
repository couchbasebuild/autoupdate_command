package main

import (
	"io"
	"os"
)

func doInstall(name string) (err error) {
	// open files r and w
	inputPath, err := os.Executable()
	if err != nil {
		return
	}

	r, err := os.Open(inputPath)
	if err != nil {
		return
	}
	defer r.Close()

	w, err := os.Create(name)
	if err != nil {
		return
	}
	defer w.Close()

	// do the actual work
	_, err = io.Copy(w, r)
	if err != nil {
		return
	}

	err = w.Sync()
	if err != nil {
		return
	}

	return
}
