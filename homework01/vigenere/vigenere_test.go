package vigenere

import "testing"

func TestEncryptVigenere(t *testing.T) {
	result := EncryptVigenere("PYTHON", "A")
	expected_result := "PYTHON"

	if result != expected_result {
		t.Fatalf("Expected '%s' but got '%s'", expected_result, result)
	}

	result = EncryptVigenere("python", "a")
	expected_result = "python"

	if result != expected_result {
		t.Fatalf("Expected '%s' but got '%s'", expected_result, result)
	}

	result = EncryptVigenere("python3.6", "a")
	expected_result = "python3.6"

	if result != expected_result {
		t.Fatalf("Expected '%s' but got '%s'", expected_result, result)
	}

	result = EncryptVigenere("ATTACKATDAWN", "LEMON")
	expected_result = "LXFOPVEFRNHR"

	if result != expected_result {
		t.Fatalf("Expected '%s' but got '%s'", expected_result, result)
	}
}

func TestDecryptVigenere(t *testing.T) {
	result := DecryptVigenere("PYTHON", "A")
	expected_result := "PYTHON"

	if result != expected_result {
		t.Fatalf("Expected '%s' but got '%s'", expected_result, result)
	}

	result = DecryptVigenere("python", "a")
	expected_result = "python"

	if result != expected_result {
		t.Fatalf("Expected '%s' but got '%s'", expected_result, result)
	}

	result = DecryptVigenere("python3.6", "a")
	expected_result = "python3.6"

	if result != expected_result {
		t.Fatalf("Expected '%s' but got  '%s'", expected_result, result)
	}

	result = DecryptVigenere("LXFOPVEFRNHR", "LEMON")
	expected_result = "ATTACKATDAWN"

	if result != expected_result {
		t.Fatalf("Expected '%s' but got '%s'", expected_result, result)
	}
}
