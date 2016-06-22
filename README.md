M2T
======

[![Build Status](https://travis-ci.org/Donearm/m2t.svg?branch=master)](https://travis-ci.org/Donearm/m2t)


M2T (aka Magnet To Torrent) is an utility to convert magnet URI into torrent files. It is meant to be used in combination with rTorrent or any other torrent downloader that support a "watch directory" (a directory that is periodically checked for new torrent files which are then automatically downloaded). rTorrent, µtorrent, Transmission do support this feature, among others (you know more? PR please!)

It is heavily inspired by the bash script on the [rTorrent Wiki](http://community.rutorrent.org/MagnetUri#Handling_.22magnet:.22_URIs_via_a_bash_script)

Install
=======

Pretty standard procedure

	go get github.com/Donearm/m2t
	go install github.com/Donearm/m2t


License
=======

MIT. See LICENSE
