package main

////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2016, Gianluca Fiore
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
////////////////////////////////////////////////////////////////////////////////

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

var mgnt_type string
var watchdir string

// Parse flags
func init() {

	const (
		def_watchdir	= "/media/private/torrent/watch"
		def_type		= "others"
	)

	flag.StringVar(&mgnt_type, "type", def_type, "")
	flag.StringVar(&mgnt_type, "t", def_type, "")
	flag.StringVar(&watchdir, "watchdir", def_watchdir, "")
	flag.StringVar(&watchdir, "w", def_watchdir, "")

}

func main() {
	var torrentFilename string

	// initialize cli arguments
	flag.Parse()

	// at least the magnet link is needed, bail out if it was not given
	if flag.NArg() == 0 {
		fmt.Println("No magnet link given, exiting...")
		os.Exit(1)
	}

	// Compile a couple of regexp we need
	var validMagnet = regexp.MustCompile(`xt=urn:btih:([^&/]+)`)
	var displayName = regexp.MustCompile(`dn=([^&/]+)`)

	if validMagnet.MatchString(flag.Arg(0)) {
		if displayName.MatchString(flag.Arg(0)) {
			torrentFilename = displayName.FindString(flag.Arg(0))
			// split at '='
			fileName := regexp.MustCompile(`=`).Split(torrentFilename, -1)[1]
			if (len(fileName) == 0) {
				xt := validMagnet.FindString(flag.Arg(0))
				fileName = regexp.MustCompile(`:`).Split(xt, -1)[2]
			}
			fmt.Println(torrentFilename)
			fmt.Println(fileName)
		}
	} else {
		// not a valid magnet URI given
		fmt.Println("The magnet URI is not correct or unparseable")
		os.Exit(1)
	}
}
