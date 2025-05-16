package base

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreate(t *testing.T) {
	b := New()

	if err := uuid.Validate(b.ID.String()); err != nil {
		t.Fatalf("invalid uuid: %v", b.ID.String())
	}

	if b.CreatedAt.IsZero() {
		t.Fatalf("invalid created_at date: %v", b.CreatedAt)
	}

	if b.UpdatedAt.IsZero() {
		t.Fatalf("invalid updated_at date: %v", b.UpdatedAt)
	}

	if err := b.Validate(); err != nil {
		t.Fatalf("method Validate should've not returned error: %v", err)
	}
}

func TestError(t *testing.T) {
	b := &BaseEntity{}

	err := b.Validate()

	if err == nil {
		t.Fatal("method Validate should've returned error")
	}
}
