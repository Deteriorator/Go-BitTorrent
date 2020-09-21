/*
 * @program: Go-BitTorrent
 * @author: Leon
 * @create: 2020-09-21 23:16
 */
package main

import (
	"log"
	"os"

	"github.com/Deteriorator/Go-BitTorrent/torrentfile"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
