package ui

import (
	"embed"
	"io/fs"
	"log"
)

// Embed the build directory from the frontend.
//
//go:embed build/*
var BuildFs embed.FS

// Get the subtree of the embedded files with `build` directory as a root.
func BuildHTTPFS() fs.FS {
	build, err := fs.Sub(BuildFs, "build")
	if err != nil {
		log.Fatal(err)
	}
	return build
}
