package utils

import (
	"crypto/rand"
	"fmt"
)

func NewUUID() string {
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		return ""
	}

	// Versão 4 (aleatória)
	u[6] = (u[6] & 0x0f) | 0x40

	// Variante (RFC 4122)
	u[8] = (u[8] & 0x3f) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		u[0:4],
		u[4:6],
		u[6:8],
		u[8:10],
		u[10:16],
	)
}
