package Files

import "testing"

func TestReadAllLines(t *testing.T) {
	lines, err := ReadAllLines("file_reader.go")
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range lines {
		t.Logf("%s\n", line)
	}
}
