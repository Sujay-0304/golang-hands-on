package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := `Contact us at support@example.comfgh or sales@example.com.
             Visit https://golang.org or http://example.com for more info.
             Call us at +1-800-555-1234 or (555) 123-4567 `


	emailPattern := `[\w\.-]+@[\w\.-]+\.\w+`
	// phonePattern := `\+?\d{1,2}?[-.\s]?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`
	phonePattern := `\+?\d{1}?[-\s]?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`
	// phonePattern := `\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`

	urlPattern := `https?://[a-zA-Z0-9./-]+`


	emailRegex := regexp.MustCompile(emailPattern)
	phoneRegex := regexp.MustCompile(phonePattern)
	urlRegex := regexp.MustCompile(urlPattern)

	
	emails := emailRegex.FindAllString(text, -1)
	phones := phoneRegex.FindAllString(text, -1)
	urls := urlRegex.FindAllString(text, -1)


	fmt.Println("Emails:", emails)
	fmt.Println("Phone Numbers:", phones)
	fmt.Println("URLs:", urls)
}
