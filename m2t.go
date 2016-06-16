package main

////////////////////////////////////////////////////////////////////////////////
//
// Copyright (c) 2016, Gianluca Fiore
//
////////////////////////////////////////////////////////////////////////////////

import (
	"flag"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
)

var mgnt_type string
var watchdir string

// Parse flags
func init() {

	const (
		def_watchdir = "/media/private/torrent/watch"
		def_type     = "others"
	)

	flag.StringVar(&mgnt_type, "type", def_type, "")
	flag.StringVar(&mgnt_type, "t", def_type, "")
	flag.StringVar(&watchdir, "watchdir", def_watchdir, "")
	flag.StringVar(&watchdir, "w", def_watchdir, "")

}

func main() {
	var torrentFilename, fileName string

	// initialize cli arguments
	flag.Parse()

	// at least the magnet link is needed, bail out if it was not given
	if flag.NArg() == 0 {
		fmt.Println("No magnet link given, exiting...")
		os.Exit(1)
	}

	// Compose the watch dir's path for the torrent file
	outputDir := path.Join(watchdir, mgnt_type)
	// Reassign first flag to a variable
	magnetLink := flag.Arg(0)

	// Compile a couple of regexp we need
	var validMagnet = regexp.MustCompile(`xt=urn:btih:([^&/]+)`)
	var displayName = regexp.MustCompile(`dn=([^&/]+)`)

	if validMagnet.MatchString(magnetLink) {
		if displayName.MatchString(magnetLink) {
			torrentFilename = displayName.FindString(magnetLink)
			// split at '='
			fileName = regexp.MustCompile(`=`).Split(torrentFilename, -1)[1]
			if len(fileName) == 0 {
				xt := validMagnet.FindString(magnetLink)
				fileName = regexp.MustCompile(`:`).Split(xt, -1)[2]
			}
			// Add torrent extension
			fileName = fileName + ".torrent"
		}
	} else {
		// not a valid magnet URI given
		fmt.Println("The magnet URI is not correct or unparseable")
		os.Exit(1)
	}

	// Create torrent file and output magnet link string to it
	f, err := os.Create(path.Join(outputDir, fileName))
	if err != nil {
		panic(err)
	}
	_, w_err := f.WriteString("d10:magnet-uri" + strconv.Itoa(len(magnetLink)) + ":" + magnetLink + "e")
	if w_err != nil {
		panic(err)
	}

	defer f.Close()

	os.Exit(0)
}
