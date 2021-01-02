package sql

import (
	"fmt"
	"testing"
)

func TestNewStorageConnectionOk(t *testing.T) {
	db, err := NewStorage()
	if err != nil {
		fmt.Println("error")
	}
	db.db.Ping()
}