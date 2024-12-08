package postgres

import (
	"encoding/json"
	"github.com/rogue0026/door-locker/internal/models"
	"io"
	"os"
	"testing"
)

func TestDecodeImages(t *testing.T) {
	f, err := os.Open("test_data.txt")
	if err != nil {
		t.Error(err.Error())
	}
	request, err := io.ReadAll(f)
	testLock := models.Lock{}
	if err = json.Unmarshal(request, &testLock); err != nil {
		t.Error(err.Error())
	}
	filenames, err := decodeImages(testLock.Title, testLock.Images)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(filenames)
}
