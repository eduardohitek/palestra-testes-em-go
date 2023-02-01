//go:build integration

package exemplo7

import "testing"

func TestGetStarWarsCharacterInfo(t *testing.T) {

	char, err := getStarWarsCharacterInfo(1)
	if err != nil {
		t.Errorf("Expected to be nil, got %q", err)
	}
	if char.Name != "Luke Skywalker" {
		t.Errorf("Expected to be Luke, got %q", char.Name)
	}
}
