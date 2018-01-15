package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/rusq/siebns"
)

const version = "1.0.0"

func main() {
	fmt.Printf("Siebnsfix %s - fix checksum in Siebel Gateway file\n", version)
	if len(os.Args) < 2 {
		fmt.Printf("\nUsage: %s <siebns.dat>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	//var ns NSFile = NSFile{}
	//var ns NSFile
	//ns := &NSFile{}
	var ns = siebns.NSFile{}

	err := ns.Load(os.Args[1])
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer ns.Close()

	if !ns.CorrectionNeeded {
		log.Printf("File %s:  OK:  No correction needed.\n", ns.Name)
		return
	}

	log.Printf("File %s:  Correction needed.\n", ns.Name)
	wrote, err := ns.FixSize()
	if err != nil {
		log.Fatalf("Error writing to file:  %s\n", err)
	}
	log.Printf("File %s:  OK: Updated %d bytes.\n", ns.Name, wrote)

}
