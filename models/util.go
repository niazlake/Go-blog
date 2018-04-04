package models

import "fmt"
import "crypto/rand"

func GenerateId() string {
	b := make([]byte, 16)

	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
