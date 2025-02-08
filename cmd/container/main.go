package main

import (
	"encoding/base64"
	"github.com/mwildt/golang-acme-extract/pkg/acme"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if acmeFilePath := os.Getenv("ACME_FILE"); acmeFilePath == "" {
		log.Fatal("no acme path defined")
	} else if target := os.Getenv("CERTIFICATE_DIR"); target == "" {
		log.Fatal("no target dir defined")
	} else if err := os.MkdirAll("outPath", 0755); err != nil {
		log.Fatalf("unable to init target dir %v", err)
	} else if container, err := acme.ReadAcmeFile(acmeFilePath); err != nil {
		log.Fatalf("unable to read acme file  %v", err)
	} else {
		for _, cert := range container.Acmetls.Certificates {
			crt_path := filepath.Join(target, cert.Domain.Main+".crt")
			key_path := filepath.Join(target, cert.Domain.Main+".key")
			if crt_decoded, err := base64.StdEncoding.DecodeString(cert.Certificate); err != nil {
				log.Printf("error decoding cert %v", err)
			} else if err := os.WriteFile(crt_path, crt_decoded, 0644); err != nil {
				log.Printf("error writig cert %v", err)
			} else if key_decoded, err := base64.StdEncoding.DecodeString(cert.Key); err != nil {
				log.Printf("error decoding cert %v", err)
			} else if err := os.WriteFile(key_path, key_decoded, 0644); err != nil {
				log.Printf("error writig key %v", err)
			} else {
				log.Printf("certificate saved (%s)", cert.Domain.Main)
			}
		}
	}
}
