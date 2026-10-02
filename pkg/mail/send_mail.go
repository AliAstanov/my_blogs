package mail

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendMail(to []string, otp string) error {
	// Jo'natuvchi email manzili (Gmail)
	fromGmail := "alibekastan1998@gmail.com"
	// Gmail uchun App Password (2FA yoqilgan bo'lsa oddiy parol ishlamaydi)
	password := "hbce batn wzbs fwvn"

	// SMTP server manzili va porti
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Emailning subject qismi (majburiy bo'lib, SMTP standarti uchun kerak)
	subject := "Subject: Your OTP Code\n"

	// Jo'natuvchi (From) headeri, SMTP server uchun aniq manzil ko'rsatiladi
	from := "From: " + fromGmail + "\n"

	// Qabul qiluvchi (To) headeri, bu yerda faqat birinchi email olinmoqda
	toHeader := "To: " + to[0] + "\n"

	// MIME va Content-Type headerlari — email matni UTF-8 formatda ekanligini bildiradi
	contentType := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"

	// Email matni (body) — OTP kodi shu yerda jo'natiladi
	body := "Your OTP code is: " + otp + "\n"

	// Barcha headerlar va body ni bitta xabarga birlashtiramiz
	// Headerlar va body orasida bo'sh qator (\n\n) muhim
	message := []byte(subject + from + toHeader + contentType + body)

	// SMTP serverga autentifikatsiya uchun login ma'lumotlarini tayyorlaymiz
	auth := smtp.PlainAuth("", fromGmail, password, smtpHost)

	// SMTP serverga emailni yuboramiz
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromGmail, to, message)
	if err != nil {
		// Agar xato bo'lsa, logga yozamiz va xatoni qaytaramiz
		log.Println("error sending mail:", err)
		return err
	}

	// Muvaffaqiyatli yuborilganini konsolga chiqaramiz
	fmt.Println("Mail sent successfully")
	return nil
}





// func SendMail(to []string, otp string) error {

// 	var (
// 		fromGmail = "alibekastan1998@gmail.com"
// 		password  = "hbce batn wzbs fwvn "	
// 	)

// 	smsHost := "smtp.gmail.com"
// 	port := "587"

// 	message := []byte("Your otp code is\n" + otp)

// 	byteM := []byte(message)

// 	auth := smtp.PlainAuth("", fromGmail, password, smsHost)

// 	err := smtp.SendMail(smsHost+":"+port, auth, fromGmail, to, byteM)

// 	if err != nil {
// 		log.Println("error", err)
// 		return err
// 	}

// 	fmt.Println("sms successfully")

// 	return nil
// }
