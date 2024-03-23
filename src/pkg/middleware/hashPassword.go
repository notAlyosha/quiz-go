package middleware

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

// Generate bytes randomly and securely using the
// Cryptographically secure pseudorandom number generator (CSPRNG)
// in the crypto.rand package
func GenerateRandomSalt(saltSize int) ([]byte, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		return nil, err
	}

	return salt, nil
}

// Combine password and salt then hash them using the SHA-512
// hashing algorithm and then return the hashed password
// as a hex string
func HashPassword(password string, salt []byte) string {
	var passwordBytes = []byte(password)

	passwordBytes = append(passwordBytes, salt...)

	sha512Hasher := sha512.New()

	sha512Hasher.Write(passwordBytes)

	hashedPasswordBytes := sha512Hasher.Sum(nil)

	hashedPasswordHex := hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

// Check if two passwords match
func DoPasswordsMatch(hashedPassword, password string, salt []byte) bool {

	return hashedPassword == HashPassword(password, salt)
}
