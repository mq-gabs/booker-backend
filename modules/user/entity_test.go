package user

import "testing"

var u *User
var name = "John Doe"
var email = "johndoe@mail.com"
var pass = "john123doe"

func setupUser() {
	u = New()
	u.Name = name
	u.Email = email
	u.Password = pass
}

func TestCreate(t *testing.T) {
	setupUser()

	if u.Name != name {
		t.Fatalf("invalid name: expected: %v, got: %v", name, u.Name)
	}
	if u.Email != email {
		t.Fatalf("invalid email: expected: %v, got: %v", email, u.Email)
	}
	if u.Password != pass {
		t.Fatalf("invalid password: expected: %v, got: %v", pass, u.Password)
	}
	if err := u.Validate(); err != nil {
		t.Fatalf("method Validate should've not returned error: %v", err)
	}

}

func TestValidateName(t *testing.T) {
	setupUser()

	u.Name = ""

	if err := u.Validate(); err == nil {
		t.Fatal("method Validate should've returned error for name")
	}
}

func TestValidateEmail(t *testing.T) {
	setupUser()

	u.Email = "wrong-email-format"

	if err := u.Validate(); err == nil {
		t.Fatal("method Validate should've returned error for email")
	}
}

func TestValidatePassword(t *testing.T) {
	setupUser()

	u.Password = ""

	if err := u.Validate(); err == nil {
		t.Fatal("method Validate should've returned error for password")
	}
}
