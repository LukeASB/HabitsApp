package validation

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	maxNameLength  = 50
	maxEmailLength = 320
	minPassword    = 8
	maxPassowrd    = 72 // Bcrypt has a limit of 72 bytes
)

func IsValidName(name string) bool {
	success := true
	if len(name) == 0 || len(name) > maxNameLength {
		return !success
	}

	validName := regexp.MustCompile(`[a-zA-Z\s'-]`)

	if !validName.MatchString(name) {
		return !success
	}

	return success
}

// https://emaillistvalidation.com/blog/demystifying-email-validation-understanding-the-maximum-length-of-email-addresses/
func IsValidEmail(email string) bool {
	success := true
	if len(email) == 0 || len(email) > maxEmailLength {
		return !success
	}

	validEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !validEmail.MatchString(email) {
		return !success
	}

	return success
}

func IsValidatePassword(plainTextPassword string) bool {
	success := true

	if len(plainTextPassword) == 0 || len(plainTextPassword) < minPassword || len(plainTextPassword) > maxPassowrd {
		return !success
	}

	validPassword := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]{8,20}$`)

	if !validPassword.MatchString(plainTextPassword) {
		return !success
	}

	return success

}

func HashPassword(plainTextPassword string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("userValidation.HashedPassword - failed to hash password")
	}

	return hashedPassword, nil
}

// Compares the userAuth (plain text password) with userData (hashed password)
func VerifyUserPassword(plainTextPassword, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword)); err != nil {
		return false
	}

	return true
}
