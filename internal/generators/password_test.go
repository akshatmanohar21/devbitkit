package generators

import (
	"strings"
	"testing"
)

func TestGeneratePassword_Length(t *testing.T) {
	password, err := GeneratePassword(16, false, false, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(password) != 16 {
		t.Errorf("expected password length 16, got %d\n", len(password))
	}
}

func TestGeneratePassword_Randomness(t *testing.T) {
	first, err := GeneratePassword(16, false, false, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	second, err := GeneratePassword(16, false, false, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if first == second {
		t.Errorf("expected two calls to produce different passwords, got identical: %s", first)
	}
}

func TestGeneratePassword_NoSymbols(t *testing.T) {
	symbolCharset := "!@#$%^&*()-_=+"

	for i := 0; i<50; i++ {
		password, err := GeneratePassword(20, false, false, true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, char := range password {
			if strings.ContainsRune(symbolCharset, char) {
				t.Errorf("password contains excluded symbol character %q: %s", char, password)
			}
		}
	}
}

func TestGeneratePassword_NoLetters(t *testing.T) {
	letterCharset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	for i := 0; i<50; i++ {
		password, err := GeneratePassword(20, true, false, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, char := range password {
			if strings.ContainsRune(letterCharset, char) {
				t.Errorf("password contains excluded letter character %q: %s", char, password)
			}
		}
	}
}

func TestGeneratePassword_NoNumbers(t *testing.T) {
	numberCharset := "0123456789"

	for i := 0; i<50; i++ {
		password, err := GeneratePassword(20, false, true, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, char := range password {
			if strings.ContainsRune(numberCharset, char) {
				t.Errorf("password contains excluded numeric character %q: %s", char, password)
			}
		}
	}
}

func TestGeneratePassword_InvalidLength(t *testing.T) {
	_, err := GeneratePassword(0, false, false, false)
	if err == nil {
		t.Errorf("expected an error for length 0, got nil")
	}
}

func TestGeneratePassword_AllExclude(t *testing.T) {
	_, err := GeneratePassword(20, true, true, true)
	if err == nil {
		t.Errorf("expected an error for length 0, got nil")
	}
}