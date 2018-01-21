package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"os"
)

func getClient(encryption string, servername string, host string, tlsconfig *tls.Config) *smtp.Client {
	if encryption == "ssl" {
		// conn := ssl(servername, tlsconfig)
		conn, err := tls.Dial("tcp", servername, tlsconfig)
		if err != nil {
			log.Panic(err)
		}
		c, err := smtp.NewClient(conn, host)
		if err != nil {
			log.Panic(err)
		}
		return c
	}
	c, err := smtp.Dial(servername)
	if err != nil {
		log.Panic(err)
	}
	return c
}

// SSL/TLS Email Example

func main() {

	fmt.Println("HOST:", os.Getenv("HOST"))
	fmt.Println("PORT:", os.Getenv("PORT"))
	fmt.Println("FROM:", os.Getenv("FROM"))
	fmt.Println("TO:", os.Getenv("TO"))
	fmt.Println("SUBJECT:", os.Getenv("SUBJECT"))
	fmt.Println("BODY:", os.Getenv("BODY"))
	fmt.Println("USERNAME:", os.Getenv("USERNAME"))
	fmt.Println("ENCRYPTION:", os.Getenv("ENCRYPTION"))
	from := mail.Address{"", os.Getenv("FROM")}
	to := mail.Address{"", os.Getenv("TO")}
	subj := os.Getenv("SUBJECT")
	body := os.Getenv("BODY")

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	host, _, _ := net.SplitHostPort(servername)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	var err error
	c := getClient(os.Getenv("ENCRYPTION"), servername, host, tlsconfig)

	if os.Getenv("ENCRYPTION") == "tls" {
		if err = c.StartTLS(tlsconfig); err != nil {
			log.Printf("Error performing StartTLS: %s\n", err)
			return
		}
	}

	// Auth
	if os.Getenv("AUTH") == "1" {
		fmt.Println("AUTH")
		auth := smtp.PlainAuth("", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), host)
		if err = c.Auth(auth); err != nil {
			log.Panic(err)
		}
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

}
