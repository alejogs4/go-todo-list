package uuid

import "github.com/google/uuid"

type UUIDGenerator struct{}

func (u UUIDGenerator) Generate() string {
	return uuid.New().String()
}
