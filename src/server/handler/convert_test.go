package handler

import (
"testing"
)

func TestGenerateID(t *testing.T) {
id1 := generateID()
id2 := generateID()

if id1 == "" {
t.Error("ID should not be empty")
}

if id1 == id2 {
t.Error("IDs should be unique")
}
}
