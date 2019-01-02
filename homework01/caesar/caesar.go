package main

func EncryptCaesar(plaintext string) string {
	var ciphertext string

	for i := 0; i < len(plaintext); i++ {
		char := plaintext[i]
		if char <= 'Z' && char >= 'A' || char >= 'a' && char <= 'z' {
			char += 3
			if char > 'Z' && char < 'a' || char > 'z' {
				char -= 26
			}
		}
		ciphertext += string(char)
	}
	return ciphertext
}

func DecryptCaesar(ciphertext string) string {
	var plaintext string

	for i := 0; i < len(ciphertext); i++ {
		char := ciphertext[i]
		if char <= 'Z' && char >= 'A' || char >= 'a' && char <= 'z' {
			char -= 3
			if char > 'Z' && char < 'a' || char < 'A' {
				char += 26
			}
		}
		plaintext += string(char)
	}
	return plaintext
}
