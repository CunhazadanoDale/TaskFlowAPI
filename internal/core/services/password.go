package services

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	// Implement password hashing logic here (e.g., using bcrypt)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return string(err.Error())
	}
	// For demonstration purposes, we'll just return the password as-is.
	return string(passwordHash)
}


func CheckPasswordHash(password, hash string) bool {
 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	
	return err == nil
}