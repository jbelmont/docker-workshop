package model

import (
	"fmt"
	"testing"
)

func TestInitDB(t *testing.T) {
	db := InitDB()
	fmt.Println(db)
}
