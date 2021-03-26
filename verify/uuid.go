package verify

import "github.com/google/uuid"

// UUIDIsValid uuid valid
func UUIDIsValid(uid string) bool {
	_, err := uuid.Parse(uid)
	if err != nil {
		return false
	}
	return true
}
