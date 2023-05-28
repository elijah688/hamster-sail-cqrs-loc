package location

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
)

func TestLocation_Marshal(t *testing.T) {
	loc := NewLocation()

	data, err := loc.Marshal()
	if err != nil {
		t.Errorf("Marshal returned an error: %v", err)
	}

	// Verify that the marshaled data can be successfully unmarshaled back into a Location struct
	var unmarshaledLoc Location
	err = json.Unmarshal(data, &unmarshaledLoc)
	if err != nil {
		t.Errorf("Unmarshal returned an error: %v", err)
	}

	// Verify that the unmarshaled location matches the original location
	if unmarshaledLoc.ID != loc.ID {
		t.Errorf("Unmarshaled ID does not match the original ID")
	}
	if unmarshaledLoc.X != loc.X {
		t.Errorf("Unmarshaled X does not match the original X")
	}
	if unmarshaledLoc.Y != loc.Y {
		t.Errorf("Unmarshaled Y does not match the original Y")
	}
}

func TestLocation_Unmarshal(t *testing.T) {
	// Define a sample JSON data representing a Location struct
	jsonData := []byte(`{
		"id": "123e4567-e89b-12d3-a456-426655440000",
		"x": 12.34,
		"y": 56.78
	}`)

	loc := &Location{}
	err := loc.Unmarshal(jsonData)
	if err != nil {
		t.Errorf("Unmarshal returned an error: %v", err)
	}

	// Verify that the unmarshaled location matches the expected values
	expectedID, _ := uuid.Parse("123e4567-e89b-12d3-a456-426655440000")
	expectedX := float32(12.34)
	expectedY := float32(56.78)

	if loc.ID != expectedID {
		t.Errorf("Unmarshaled ID does not match the expected ID")
	}
	if loc.X != expectedX {
		t.Errorf("Unmarshaled X does not match the expected X")
	}
	if loc.Y != expectedY {
		t.Errorf("Unmarshaled Y does not match the expected Y")
	}
}

func TestNewLocation(t *testing.T) {
	loc := NewLocation()

	// Verify that the generated location has a valid UUID
	if _, err := uuid.Parse(loc.ID.String()); err != nil {
		t.Errorf("NewLocation generated an invalid UUID: %v", err)
	}

	// Verify that the generated location has valid X and Y values within the specified range
	if loc.X < 0.0 || loc.X > 640.0 {
		t.Errorf("NewLocation generated an invalid X value: %f", loc.X)
	}
	if loc.Y < 0.0 || loc.Y > 640.0 {
		t.Errorf("NewLocation generated an invalid Y value: %f", loc.Y)
	}
}
