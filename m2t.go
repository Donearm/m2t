package main

////////////////////////////////////////////////////////////////////////////////
//
// Copyright (c) 2016-2019, Gianluca Fiore
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

var mgntType string
var watchdir string

var usageMessage = `
m2t [-t|-type <type>] [-w|watchdir <directory>] <magnetlink>

M2T (Magnet To Torrent) will convert magnet links to torrent files. It works by using the "watch directory" feature of many torrent clients

Arguments:
	-t|-type <type>
		The type of the file contained in the magnet link. Defaults to "others"

	-w|-watchdir <directory>
		The directory where the torrent client watches for new torrents.
		It must exist. Defaults to '/media/private/torrent/watch
	`

// Parse flags
func init() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usageMessage)
	}

	const (
		defWatchdir = "/media/private/torrent/watch"
		defType     = "others"
	)

	flag.StringVar(&mgntType, "type", defType, "")
	flag.StringVar(&mgntType, "t", defType, "")
	flag.StringVar(&watchdir, "watchdir", defWatchdir, "")
	flag.StringVar(&watchdir, "w", defWatchdir, "")

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
	outputDir := path.Join(watchdir, mgntType)
	// Reassign first flag to a variable
	magnetLink := flag.Arg(0)

	// Compile a couple of regexp we need
	var validMagnet = regexp.MustCompile(`xt=urn(\.[0-9]*)?:btih:([^&/]+)`)
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
	_, wErr := f.WriteString("d10:magnet-uri" + strconv.Itoa(len(magnetLink)) + ":" + magnetLink + "e")
	if wErr != nil {
		panic(err)
	}

	defer f.Close()

	os.Exit(0)
}
