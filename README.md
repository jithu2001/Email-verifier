# Email Verification Tool

This Go program allows you to verify the validity of an email address and perform a series of checks to determine if the domain associated with the email has the following:

1. **MX Records**: Ensures the domain can receive emails.
2. **SPF (Sender Policy Framework)**: Validates if the domain is authorized to send emails.
3. **DMARC (Domain-based Message Authentication, Reporting, and Conformance)**: Specifies how emails from the domain should be handled to prevent spoofing.

---

## Features

- **Interactive User Input**: Prompts the user to input an email address for verification.
- **Domain Checks**: Verifies the presence of MX, SPF, and DMARC records for the email domain.
- **Validation Output**: Displays details such as email validity, domain, MX status, SPF status and record, and DMARC status and record.

---

## Prerequisites

- **Go**: Ensure Go is installed on your system. You can download it [here](https://golang.org/dl/).
- **Internet Access**: The program performs DNS lookups, requiring an active internet connection.

---

## How to Run

1. Clone this repository or copy the code into a file named `main.go`.
2. Open a terminal and navigate to the directory containing `main.go`.
3. Run the following command to build and execute the program:
   ```bash
   go run main.go
   ```
4. Enter email addresses when prompted, or type `exit` to quit the program.

---

## Sample Input and Output

### Example 1: Valid Email with Records
**Input**: `user@example.com`  
**Output**:
```
Enter the Mail ID (or type 'exit' to quit): user@example.com
Email :  user@example.com
Email validity :  true
Domain :  example.com
Does it have MX :  true
Does it have SPF :  true
SPF Record :  v=spf1 include:_spf.example.com ~all
Does it have DMARC :  true
DMARC Record :  v=DMARC1; p=none; rua=mailto:dmarc-reports@example.com
```

### Example 2: Invalid Email Format
**Input**: `invalid-email`  
**Output**:
```
Enter the Mail ID (or type 'exit' to quit): invalid-email
Email :  invalid-email
Email validity :  false
Domain :  
Does it have MX :  false
Does it have SPF :  false
SPF Record :  
Does it have DMARC :  false
DMARC Record :  
```

---

## Code Structure

- **Main Program**:
  - Prompts user input and verifies email format.
- **Email Parsing**:
  - Uses `net/mail` to validate the email address syntax.
- **Domain Checks**:
  - Uses `net.LookupMX` to check for MX records.
  - Uses `net.LookupTXT` to find SPF and DMARC records.

---

## Error Handling

The program logs DNS lookup errors but continues execution to ensure uninterrupted user interaction. Invalid email inputs are gracefully handled with appropriate messages.

---

## Dependencies

- **Standard Library**: The program relies solely on Go's standard library (`bufio`, `fmt`, `log`, `net`, `os`, `strings`).

---

## Future Enhancements

- Add support for batch email verification via file input.
- Improve error handling for specific DNS lookup failures.
- Display more detailed validation for SPF and DMARC records.

---

## License

This project is open-source and available under the MIT License.

---

## Author

Developed by [Your Name]. Feel free to contribute or reach out for improvements!
