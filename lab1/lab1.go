// package main

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// type EncryptDecrypt struct {
// 	plainText string
// 	shift int
// 	encryptedText string
// 	decryptedText string
// }
// // Function to encrypt the text using Caesar Cipher
// func encrypt(text string, shift int) string {
// 	shift = shift % 26 // Normalize shift to be within the range of 0-25
// 	runes := []rune(text)
// 	for i, c := range runes {
// 		if c >= 'a' && c <= 'z' {
// 			runes[i] = 'a' + (c-'a'+rune(shift))%26
// 		} else if c >= 'A' && c <= 'Z' {
// 			runes[i] = 'A' + (c-'A'+rune(shift))%26
// 		}
// 	}
// 	return string(runes)
// }

// // Function to decrypt the text using Caesar Cipher
// func decrypt(text string, shift int) string {
// 	return encrypt(text, -shift)
// }

// func main() {
// 	frontPageHandler := func(w http.ResponseWriter, r *http.Request) {
// 		tmpl := template.Must(template.ParseFiles("index.html"))
// 		tmpl.Execute(w, nil)
// 	}
// 	encryptHandler := func(w http.ResponseWriter, r *http.Request) {
// 		plainText := r.PostFormValue("plainText")
// 		templ = template.Must(template.ParseFiles("index.html"))

// 	}
// 	http.HandleFunc("/", frontPageHandler)
// 	http.HandleFunc("/encrypt", encryptHandler)

//		log.Fatal(http.ListenAndServe(":80", nil))
//	}
package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type EncryptDecrypt struct {
	PlainText     string
	Shift         int
	EncryptedText string
	DecryptedText string
}

// Function to encrypt the text using Caesar Cipher
func encrypt(text string, shift int) string {
	shift = shift % 26 // Normalize shift to be within the range of 0-25
	runes := []rune(text)
	for i, c := range runes {
		if c >= 'a' && c <= 'z' {
			runes[i] = 'a' + (c-'a'+rune(shift))%26
		} else if c >= 'A' && c <= 'Z' {
			runes[i] = 'A' + (c-'A'+rune(shift))%26
		}
	}
	return string(runes)
}

// Function to decrypt the text using Caesar Cipher
func decrypt(text string, shift int) string {
	return encrypt(text, -shift)
}

func main() {
	frontPageHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	encryptHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			plainText := r.PostFormValue("plainText")
			shiftStr := r.PostFormValue("shift")
			shift, err := strconv.Atoi(shiftStr)
			if err != nil {
				http.Error(w, "Invalid shift value", http.StatusBadRequest)
				return
			}
			encryptedText := encrypt(plainText, shift)
			data := EncryptDecrypt{
				PlainText:     plainText,
				Shift:         shift,
				EncryptedText: encryptedText,
			}
			tmpl := template.Must(template.ParseFiles("index.html"))
			tmpl.Execute(w, data)
		}
	}

	decryptHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			encryptedText := r.PostFormValue("encryptedText")
			shiftStr := r.PostFormValue("shift")
			shift, err := strconv.Atoi(shiftStr)
			if err != nil {
				http.Error(w, "Invalid shift value", http.StatusBadRequest)
				return
			}
			decryptedText := decrypt(encryptedText, shift)
			data := EncryptDecrypt{
				EncryptedText: encryptedText,
				Shift:         shift,
				DecryptedText: decryptedText,
			}
			tmpl := template.Must(template.ParseFiles("index.html"))
			tmpl.Execute(w, data)
		}
	}

	http.HandleFunc("/", frontPageHandler)
	http.HandleFunc("/encrypt", encryptHandler)
	http.HandleFunc("/decrypt", decryptHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
