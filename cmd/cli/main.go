package main

import (
	"encoding/base64"
	"flag"
	"github.com/mwildt/golang-acme-extract/pkg/acme"
	"log"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {

	var src string
	var target string

	var crt_file_extention string
	var key_file_extention string

	var domain string

	flag.StringVar(&domain, "d", "*", "Specifiy Domain. Default is *")
	flag.StringVar(&domain, "domain", "*", "Specifiy Domain. Default is *")

	// flags declaration using flag package
	flag.StringVar(&src, "source", "acme.json", "Specify Source file. Default is acme.json")
	flag.StringVar(&src, "s", "acme.json", "Specify Source file. Default is acme.json (short)")

	flag.StringVar(&target, "target", "target", "Specify pass. Default is password")
	flag.StringVar(&target, "t", "target", "Specify pass. Default is password (short)")

	flag.StringVar(&crt_file_extention, "crt-ext", ".crt", "Certificate file extension")
	flag.StringVar(&key_file_extention, "key-ext", ".key", "Key file extension")

	flag.Parse()

	container, err := acme.ReadAcmeFile(src)
	checkError(err)

	checkError(os.MkdirAll(target, os.ModePerm))

	for _, cert := range container.Acmetls.Certificates {
		if domain == "*" || domain == cert.Domain.Main {
			crt_fileame := cert.Domain.Main + crt_file_extention
			crt_path := filepath.Join(target, crt_fileame)
			crt_decoded, err := base64.StdEncoding.DecodeString(cert.Certificate)
			checkError(err)
			os.WriteFile(crt_path, crt_decoded, 0644)

			key_fileame := cert.Domain.Main + key_file_extention
			key_path := filepath.Join(target, key_fileame)
			key_decoded, err := base64.StdEncoding.DecodeString(cert.Key)
			checkError(err)
			os.WriteFile(key_path, key_decoded, 0644)
		}
	}

}
