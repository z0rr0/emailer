// Copyright (c) 2015, Alexander Zaytsev. All rights reserved.
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

package main

import (
    "flag"
    "fmt"
    "net/smtp"
    "os"
    "strings"
)

const (
    defaultSmtpHost string = "localhost"
    defaultSmtpPort int = 25
    defaultEmailSubject string = "NoSubject"
)

func send(from, subject, msg, username, password, host string, to []string, port int) error {
    const mime string = "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n";
    header := fmt.Sprintf("From: %v\nSubject: %v\n", from, subject)
    content := []byte(header + mime + msg)
    addr := fmt.Sprintf("%v:%v", host, port)
    auth := smtp.PlainAuth("", username, password, host)
    return smtp.SendMail(addr, auth, username, to, content)
}

func main() {
    host := flag.String("host", defaultSmtpHost, "smtp host")
    port := flag.Int("port", defaultSmtpPort, "smtp port")
    username := flag.String("username", "", "smtp username")
    password := flag.String("password", "", "smtp password")
    from := flag.String("from", "", "sender email")
    to := flag.String("to", "", "email recipients")
    subject := flag.String("subject", defaultEmailSubject, "email subject")
    flag.Parse()
    msg := strings.Join(flag.Args(), " ")
    if *to == "" {
        fmt.Println("not found recipients")
        os.Exit(1)
    }
    recipients := strings.Split(*to, ",")
    if err := send(*from, *subject, msg, *username, *password, *host, recipients, *port); err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
}