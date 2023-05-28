package location

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Location struct {
	ID        uuid.UUID `json:"id"`
	X         float32   `json:"x"`
	Y         float32   `json:"y"`
	CreatedAt time.Time `json:"created_at"`
}

func NewLocation() *Location {
	return &Location{
		ID:        uuid.New(),
		X:         generateRandomFloat(),
		Y:         generateRandomFloat(),
		CreatedAt: time.Now(),
	}
}

func (l *Location) Marshal() ([]byte, error) {
	j, err := json.MarshalIndent(l, "\t", "")
	if err != nil {
		return j, nil
	}

	return j, nil
}

func (l *Location) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, l); err != nil {
		return err
	}

	return nil
}

func generateRandomFloat() float32 {
	rand.NewSource(time.Now().UnixNano()) // Seed the random number generator with the current time

	min := 0.0
	max := 640.0

	// Generate a random float between min and max
	randomFloat := min + rand.Float64()*(max-min)

	// Round the random float to 2 decimal places
	roundedFloat := float32(int(randomFloat*100)) / 100

	return roundedFloat
}
