package exemplo3

import (
	"os"
	"testing"
)

func setupFile(t *testing.T) (string, func()) {
	t.Helper()
	file, err := os.CreateTemp("/", "my_test_file")
	if err != nil {
		t.Fatalf("could't create temp file: %q", err)
	}
	return file.Name(), func() {
		file.Close()
		os.Remove(file.Name())
	}
}

func TestSaveToFile(t *testing.T) {
	file, cleanUp := setupFile(t)
	defer cleanUp()

	person := Person{
		Name:  "Eduardo",
		Email: "eduardohitek@gmail.com",
	}

	err := saveToFile(file, person)
	if err != nil {
		t.Errorf("expected err to be nil, got %q", err)
	}

}
