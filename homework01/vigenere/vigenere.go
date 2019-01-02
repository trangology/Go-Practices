package vigenere

func EncryptVigenere(plaintext string, keyword string) string {
	var ciphertext string

	times := len(plaintext) / len(keyword)
	d := len(plaintext) % len(keyword) // d = div
	flag := keyword

	for i := 1; i < times; i++ {
		keyword += flag
	}
	if d > 0 {
		keyword += flag[0:d]
	}

	for i := 0; i < len(plaintext); i++ {
		char := plaintext[i]
		if char >= 'A' && char <= 'Z' {
			newChar := plaintext[i] + keyword[i] - 65
			if newChar > 'Z' {
				newChar -= 26
			}
			ciphertext += string(newChar)
		} else if char >= 'a' && char <= 'z' {
			newChar := plaintext[i] + keyword[i] - 97
			if newChar > 'z' {
				newChar -= 26
			}
			ciphertext += string(newChar)
		} else {
			ciphertext += string(char)
		}
	}

	return ciphertext
}

func DecryptVigenere(ciphertext string, keyword string) string {
	var plaintext string

	times := len(ciphertext) / len(keyword)
	d := len(ciphertext) % len(keyword) // d = div
	flag := keyword

	for i := 1; i < times; i++ {
		keyword += flag
	}

	if d > 0 {
		keyword += flag[0:d]
	}

	for i := 0; i < len(ciphertext); i++ {
		char := ciphertext[i]
		if char >= 'A' && char <= 'Z' {
			newChar := char - (keyword[i] - 65)
			if newChar < 'A' {
				newChar += 26
			}
			plaintext += string(newChar)
		} else if char >= 'a' && char <= 'z' {
			newChar := char - (keyword[i] - 97)
			if newChar < 'a' {
				newChar += 26
			}
			plaintext += string(newChar)
		} else {
			plaintext += string(char)
		}
	}

	return plaintext
}
