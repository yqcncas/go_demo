package gormtest

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestUUid(t *testing.T) {
	u2 := uuid.NewV4()
	fmt.Println(u2)
	fmt.Printf("UUIDv4: %s\n", u2)
}
