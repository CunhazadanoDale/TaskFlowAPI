package services


func hashPassword(password string) string {
	// Implement password hashing logic here (e.g., using bcrypt)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// Handle error appropriately
		return ""
	}
	// For demonstration purposes, we'll just return the password as-is.
	return string(passwordHash)
}