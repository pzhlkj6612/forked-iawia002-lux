package utils

import (
	"os"
	"testing"
)

func TestCheckFFmpeg(t *testing.T) {
	originalPath := os.Getenv("PATH")

	t.Run("ffmpeg not found", func(t *testing.T) {
		// Set PATH to empty so ffmpeg won't be found
		t.Setenv("PATH", "")
		err := CheckFFmpeg()
		if err == nil {
			t.Error("CheckFFmpeg() should return error when ffmpeg is not in PATH")
		}
	})

	t.Run("ffmpeg found", func(t *testing.T) {
		t.Setenv("PATH", originalPath)
		err := CheckFFmpeg()
		// We don't fail if ffmpeg is not installed in the test environment,
		// but if it is installed, the check should pass
		if err != nil {
			t.Skipf("ffmpeg not available in test environment: %v", err)
		}
	})
}
