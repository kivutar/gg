package scene

import "testing"

func TestSceneImageUploadIDStaysStable(t *testing.T) {
	img := NewImage(2, 2)
	img.Data = make([]byte, 16)
	img.Data[0] = 255
	first := img.uploadBuf()
	second := img.uploadBuf()
	if first == nil || second == nil {
		t.Fatal("missing upload buffer")
	}
	if first.GenerationID() != second.GenerationID() {
		t.Fatalf("replay id %d then %d", first.GenerationID(), second.GenerationID())
	}
	img.Data = make([]byte, 16)
	third := img.uploadBuf()
	if third.GenerationID() == first.GenerationID() {
		t.Fatal("replaced pixels kept the old upload id")
	}
}
