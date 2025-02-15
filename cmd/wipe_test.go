package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRunFiles tests the run_files function.
func TestRunFiles(t *testing.T) {
	// Temporäres Verzeichnis erstellen
	tempDir, err := os.MkdirTemp("", "shred-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Temporäre Dateien erstellen
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")
	err = os.WriteFile(file1, []byte("test data"), 0644)
	assert.NoError(t, err)
	err = os.WriteFile(file2, []byte("test data"), 0644)
	assert.NoError(t, err)

	// Teste run_files mit Dateien
	files, err := run_files([]string{file1, file2})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(files))
	assert.Contains(t, files, file1)
	assert.Contains(t, files, file2)

	files, err = run_files([]string{tempDir})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(files)) // Zwei Dateien im Verzeichnis
}

// TestRunMethod tests the runmethod function.
func TestRunMethod(t *testing.T) {
	// Temporäres Verzeichnis erstellen
	tempDir, err := os.MkdirTemp("", "shred-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Temporäre Dateien erstellen
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")
	err = os.WriteFile(file1, []byte("test data"), 0644)
	assert.NoError(t, err)
	err = os.WriteFile(file2, []byte("test data"), 0644)
	assert.NoError(t, err)

	// Teste runmethod mit der "fast"-Methode
	err = runmethod(1, []string{file1, file2})
	assert.NoError(t, err)

	// Überprüfe, ob die Dateien gelöscht wurden
	_, err = os.Stat(file1)
	assert.True(t, os.IsNotExist(err), "File %s should not exist after deletion", file1)
	_, err = os.Stat(file2)
	assert.True(t, os.IsNotExist(err), "File %s should not exist after deletion", file2)
}

// TestRunMethodWithKeep tests the runmethod function with the keep flag.
func TestRunMethodWithKeep(t *testing.T) {
	// Temporäres Verzeichnis erstellen
	tempDir, err := os.MkdirTemp("", "shred-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Temporäre Dateien erstellen
	file1 := filepath.Join(tempDir, "file1.txt")
	err = os.WriteFile(file1, []byte("test data"), 0644)
	assert.NoError(t, err)

	// Setze das keep-Flag
	keep = true

	// Teste runmethod mit der "fast"-Methode
	runmethod(1, []string{file1})

	// Überprüfe, ob die Datei noch existiert
	_, err = os.Stat(file1)
	assert.NoError(t, err)
}
