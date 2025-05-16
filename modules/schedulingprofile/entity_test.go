package schedulingprofile

import (
	"testing"
	"time"
)

var sp = New()

var name = "Name"
var sh = &SchedulingHour{
	StartAt: time.Now(),
	EndAt:   time.Now().Add(10 * time.Minute),
}

func setupSchedulingProfile() {
	sp.Name = name
	sp.Monday = sh
	sp.Tuesday = sh
	sp.Wednesday = sh
	sp.Thursday = sh
	sp.Friday = sh
	sp.Saturday = sh
	sp.Sunday = sh

}
func TestCreate(t *testing.T) {
	sp = New()

	setupSchedulingProfile()

	if sp.Name != name {
		t.Fatalf("invalid name: expected: %v, got %v", name, sp.Name)
	}

	if sp.Monday != sh {
		t.Fatalf("invalid scheduling hour: expected %v, got %v", sh, sp.Monday)
	}
	if sp.Tuesday != sh {
		t.Fatalf("invalid scheduling hour: expected %v, got %v", sh, sp.Tuesday)
	}
	if sp.Wednesday != sh {
		t.Fatalf("invalid scheduling hour: expected %v, got %v", sh, sp.Wednesday)
	}
	if sp.Thursday != sh {
		t.Fatalf("invalid scheduling hour: expected %v, got %v", sh, sp.Thursday)
	}
	if sp.Friday != sh {
		t.Fatalf("invalid scheduling hour: expected %v, got %v", sh, sp.Friday)
	}
	if sp.Saturday != sh {
		t.Fatalf("invalid scheduling hour: expected %v, got %v", sh, sp.Saturday)
	}
	if sp.Sunday != sh {
		t.Fatalf("invalid scheduling hour: expected %v, got %v", sh, sp.Sunday)
	}

	if err := sp.Validate(); err != nil {
		t.Fatalf("method Validate should've not returned error: %v", err)
	}
}

func TestValidateName(t *testing.T) {
	setupSchedulingProfile()

	sp.Name = ""

	if err := sp.Validate(); err == nil {
		t.Fatal("method Validate should've returned an error for name")
	}
}

func TestValidateSchedulingHour(t *testing.T) {
	setupSchedulingProfile()

	sp.Monday = &SchedulingHour{
		StartAt: time.Now(),
		EndAt:   time.Now().Add(-1 * time.Second),
	}

	if err := sp.Validate(); err == nil {
		t.Fatal("method Validate should've returned an error for monday")
	}
}
