package main

import (
        "bytes"
        "log"
        "net/smtp"
)

func main() {
        // Connect to the remote SMTP server.
        c, err := smtp.Dial("127.0.0.1:2525")
        if err != nil {
                log.Fatal(err)
        }
        // Set the sender and recipient.
        c.Mail("sender@example.org")
        c.Rcpt("recipient@example.net")
        // Send the email body.
        wc, err := c.Data()
        if err != nil {
                log.Fatal(err)
        }
        defer wc.Close()
        buf := bytes.NewBufferString("This is the email body.")
        if _, err = buf.WriteTo(wc); err != nil {
                log.Fatal(err)
        }
}