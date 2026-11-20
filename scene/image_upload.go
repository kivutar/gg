package scene

import (
	"image"
	"unsafe"

	"github.com/gogpu/gg"
)

// uploadBuf is the GPU upload for this scene image.
// A replay must reuse it. ImageBufFromImage mints a new GenerationID, and the
// GPU texture cache keys on that id, so a fresh buffer uploads the picture again.
type imageUpload struct {
	buf *gg.ImageBuf
	ptr unsafe.Pointer
	n   int
}

func (img *Image) uploadBuf() *gg.ImageBuf {
	if img == nil || img.Width <= 0 || img.Height <= 0 || len(img.Data) < img.Width*img.Height*4 {
		return nil
	}
	ptr := unsafe.Pointer(&img.Data[0])
	if img.upload.buf != nil && img.upload.ptr == ptr && img.upload.n == len(img.Data) {
		return img.upload.buf
	}
	rgba := &image.RGBA{
		Pix:    img.Data,
		Stride: img.Width * 4,
		Rect:   image.Rect(0, 0, img.Width, img.Height),
	}
	img.upload.buf = gg.ImageBufFromImage(rgba)
	img.upload.ptr = ptr
	img.upload.n = len(img.Data)
	return img.upload.buf
}
