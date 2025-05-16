package scheduling

import (
	"testing"
	"time"
)

var s *Scheduling
var name = "John Doe"
var email = "johndoe@mail.com"
var phone = "99999999"
var cpf = "33333333333"
var dtime = time.Now().Add(time.Hour)
var status = Scheduled

func setupScheduling() {
	s = New()
	s.Name = name
	s.Email = email
	s.Phone = phone
	s.CPF = cpf
	s.Datetime = dtime
	s.Status = status
}

func TestCreate(t *testing.T) {
	setupScheduling()

	if s.Name != name {
		t.Fatalf("invalid name: expected: %v, got: %v", name, s.Name)
	}
	if s.Email != email {
		t.Fatalf("invalid email: expected: %v, got: %v", email, s.Email)
	}
	if err := s.Validate(); err != nil {
		t.Fatalf("method Validate should've not returned error: %v", err)
	}

}

func TestValidateName(t *testing.T) {
	setupScheduling()

	s.Name = ""

	if err := s.Validate(); err == nil {
		t.Fatal("method Validate should've returned error for name")
	}
}

func TestValidateEmail(t *testing.T) {
	setupScheduling()

	s.Email = "wrong-email-format"

	if err := s.Validate(); err == nil {
		t.Fatal("method Validate should've returned error for email")
	}
}

func TestValidatePhone(t *testing.T) {
	setupScheduling()

	s.Phone = "123"

	if err := s.Validate(); err == nil {
		t.Fatal("method Validate should've returned error for phone")
	}
}

func TestValidatePhoneAndEmail(t *testing.T) {
	setupScheduling()

	s.Phone = ""
	s.Email = ""

	if err := s.Validate(); err == nil {
		t.Fatal("method Validate should've returned error for phone and email")
	}
}

func TestValidateCPF(t *testing.T) {
	setupScheduling()

	s.CPF = "123"

	if err := s.Validate(); err == nil {
		t.Fatal("method Validate should've returned error for cpf")
	}
}

func TestValidateDatetime(t *testing.T) {
	setupScheduling()

	s.Datetime = time.Now().Add(-1 * time.Minute)

	if err := s.Validate(); err == nil {
		t.Fatal("method Validate should've returned error for datetime")
	}
}

func TestValidateStatus(t *testing.T) {
	setupScheduling()

	s.Status = "asdf"

	if err := s.Validate(); err == nil {
		t.Fatal("method Validate should've returned error for status")
	}
}
