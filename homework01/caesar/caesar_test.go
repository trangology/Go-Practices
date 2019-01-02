package caesar

import "testing"

func TestEncryptCaesar(t *testing.T) {
	result := EncryptCaesar("PYTHON")
	expectedResult := "SBWKRQ"

	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result)
	}

	result = EncryptCaesar("python")
	expectedResult = "sbwkrq"

	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result)
	}

	result = EncryptCaesar("Python3.6")
	expectedResult = "Sbwkrq3.6"

	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result)
	}

	result = EncryptCaesar("")
	expectedResult = ""

	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result)
	}
}

func TestDecryptCaesar(t *testing.T) {
	result := DecryptCaesar("SBWKRQ")
	expected_result := "PYTHON"

	if result != expected_result {
		t.Fatalf("Expected '%s' but got '%s'", expected_result, result)
	}

	result = DecryptCaesar("sbwkrq")
	expected_result = "python"

	if result != expected_result {
		t.Fatalf("Expected '%s' but got '%s'", expected_result, result)
	}

	result = DecryptCaesar("Sbwkrq3.6")
	expected_result = "Python3.6"

	if result != expected_result {
		t.Fatalf("Expected '%s' but got '%s'", expected_result, result)
	}

	result = DecryptCaesar("")
	expected_result = ""

	if result != expected_result {
		t.Fatalf("Expected '%s' but got '%s'", expected_result, result)
	}
}
