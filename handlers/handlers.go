package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

type EmailData struct {
	EmailTo     string `json:"EmailTo"`
	EmailFrom   string `json:"EmailFrom"`
	EmailSub    string `json:"EmailSub"`
	EmailBody   string `json:"EmailBody"`
	AppPassword string `json:"AppPassword"`
	Host        string `json:"Host"`
	Port        string `json:"Port"`
}

var EmailDatas []EmailData

func SendMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emaildata EmailData
	_ = json.NewDecoder(r.Body).Decode(&emaildata)
	EmailDatas = append(EmailDatas, emaildata)
	json.NewEncoder(w).Encode(emaildata)
	EmailDatas = append(EmailDatas, emaildata)

	from := emaildata.EmailFrom

	// Array of recipients address
	to := []string{emaildata.EmailTo}

	// Create a message and convert it into bytes
	msg := []byte("To: " + emaildata.EmailTo + "\r\n" +
		"From: " + emaildata.EmailFrom + "\r\n" +
		"Subject: " + emaildata.EmailSub + "\r\n" +
		"\r\n" +
		emaildata.EmailBody + "\r\n")

	// Call the sendEmail function
	status := sendEmail(emaildata, from, to, msg)

	// check if email sent successfully or not
	if status {
		fmt.Printf("Email sent successfully.\n")
	} else {
		fmt.Printf("Email sent failed.\n")
	}
}

func sendEmail(emaildata EmailData, from string, to []string, msg []byte) bool {
	// Set up authentication information.
	auth := smtp.PlainAuth("", from, emaildata.AppPassword, emaildata.Host)

	// format smtp address
	smtpAddress := fmt.Sprintf("%s:%v", emaildata.Host, emaildata.Port)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(smtpAddress, auth, from, to, msg)

	if err != nil {
		log.Fatal(err)
		return false
	}

	// return true on success
	return true
}


func GetFormat(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello from my SMTP client\n")
	fmt.Fprintf(w, "To use it make a post request to /sendmail\n")
	fmt.Fprintf(w, `Example Object: {"EmailTo": "Reciever Email Address", "EmailFrom": "Your Email Address", "EmailSub":"Email Subject","EmailBody":"Email Body","AppPassword":"Your Email password.For gmail generate and provide app password","Host":"Your SMTP server host (for gmail its 'smtp.gmail.com' )","Port":"SMTP server port (for gmail its 587)"}`)
}