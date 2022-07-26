package shorter

import "crypto/sha256"

// GetShort - пока заглушка
func GetShort(original string) string {
	hash := sha256.Sum256([]byte(original))
	return string(hash[:10])
}
