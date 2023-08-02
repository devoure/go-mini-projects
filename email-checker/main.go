package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func emailChecker(domain string) {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error : %v \n", err)
	}

	if len(mxRecord) > 0 {
		hasMx = true
	}

	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Println("Domain: ", domain)
	fmt.Println("hasMX: ", hasMx)
	fmt.Println("hasSPF: ", hasSPF)
	fmt.Println("spfRecord: ", spfRecord)
	fmt.Println("hasDMARC: ", hasDMARC)
	fmt.Println("dmarcRecord: ", dmarcRecord)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		emailChecker(scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal("Error: could not read from input %v \n", err)
	}
}
