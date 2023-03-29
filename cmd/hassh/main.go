package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/edutko/hassh-go/internal/hassh"
	"github.com/edutko/hassh-go/x/crypto/ssh"
)

func init() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [<options>] <host>[:<port>]\n", os.Args[0])
		flag.PrintDefaults()
	}
}

var verbose = flag.Bool("v", false, "print kex/cipher/mac details")

func main() {
	flag.Parse()
	addr := flag.Arg(0)

	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		if ae, ok := err.(*net.AddrError); ok && ae.Err == "missing port in address" {
			host = addr
			port = "22"
		} else {
			log.Fatalln(err)
		}
	}

	info, err := hassh.Connect(host, port)
	if err != nil {
		log.Fatalln(err)
	}
	printBanner(info.Banner)
	printHassh(info.Kexinit)
}

func printBanner(b string) {
	if *verbose {
		fmt.Println("Banner:")
		fmt.Println(b)
	}
}

func printHassh(info ssh.KexinitInfo) {
	if *verbose {
		fmt.Println("KEX:")
		fmt.Println("  " + strings.Join(info.KexAlgos, "\n  "))
		fmt.Println("Ciphers:")
		fmt.Println("  " + strings.Join(info.PeerCiphers, "\n  "))
		fmt.Println("MACs:")
		fmt.Println("  " + strings.Join(info.PeerMACs, "\n  "))
		fmt.Println("Compression:")
		fmt.Println("  " + strings.Join(info.PeerCompression, "\n  "))

		fmt.Print("\nhassh:\n  ")
	}
	fmt.Println(hex.EncodeToString(hassh.Digest(info)))
}
