# M2T

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme) [![Build Status](https://travis-ci.org/Donearm/m2t.svg?branch=master)](https://travis-ci.org/Donearm/m2t)

M2T (aka Magnet To Torrent) is an utility to convert magnet URI into torrent files. It is meant to be used in combination with rTorrent or any other torrent downloader that support a "watch directory" (a directory that is periodically checked for new torrent files which are then automatically downloaded). rTorrent, µtorrent, Transmission do support this feature, among others (you know more? PR please!)

It is heavily inspired by the bash script on the [rTorrent Wiki](http://community.rutorrent.org/MagnetUri#Handling_.22magnet:.22_URIs_via_a_bash_script)

## Install

Pretty standard procedure

	go get github.com/Donearm/m2t
	go install github.com/Donearm/m2t

## Usage

M2T accepts only 2 parameters, `-t` and `w`. The former is the type of file to download (preset are `music`, `books`, `vids` and `others`). This is mostly useful to scripting then with your torrent client saving directories, notification messages etc. according to different types of files. So using M2T like

	m2t -t vids $long_magnet_link

will save the resulting torrent files inside the watch directory, in a directory named "vids". That's for movies, tv series, video lessons and any files containing audio+video. It's up to you to classify or not the torrent files according to these types or not. If not inserted, it defaults to `others`.

With `w` instead you can choose a different watch directory. Yours will be most definitely different than mine so for frequent use, directly change the constant `defWatchdir` in the source and build M2T again to save the data.

## License

MIT © Gianluca Fiore. See LICENSE

[![ko-fi](https://www.ko-fi.com/img/donate_sm.png)](https://ko-fi.com/W7W7KA0Z)
