package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/mail"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter the Mail ID (or type 'exit' to quit): ") // Prompt before user input
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		verifyEmail(input)
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}

}

func verifyEmail(email string) {
	parsedEmail, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Println("Email : ", email)
		fmt.Println("Email validity : ", false)
		fmt.Println("Domain : ", "")
		fmt.Println("Does it have MX : ", false)
		fmt.Println("Does it have SPF : ", false)
		fmt.Println("SPF Record : ")
		fmt.Println("Does it have DMARC : ", false)
		fmt.Println("DMARC Record : ")
		return
	}

	domain := strings.Split(parsedEmail.Address, "@")[1]

	checkDomain(email, domain)
}

func checkDomain(email, domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecods, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecods) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Println("Email : ", email)
	fmt.Println("Email validity : ", true)
	fmt.Println("Domain : ", domain)
	fmt.Println("Does it have MX : ", hasMX)
	fmt.Println("Does it have SPF : ", hasSPF)
	fmt.Println("SPF Record : ", spfRecord)
	fmt.Println("Does it have DMARC : ", hasDMARC)
	fmt.Println("DMARC Record : ", dmarcRecord)

}
