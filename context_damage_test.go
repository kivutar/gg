package gg

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

func TestFillRectCPUDevicePixelRoundTrip(t *testing.T) {
	for _, scale := range []float64{1, 1.25, 4.0 / 3, 1.5, 2.3} {
		t.Run(fmt.Sprint(scale), func(t *testing.T) {
			dc := NewContext(180, 180, WithDeviceScale(scale))
			t.Cleanup(func() { _ = dc.Close() })
			// A dirty rectangle aligned to physical pixels, expressed in
			// logical coordinates. Floating-point division/multiplication
			// must not cause an extra row or column to be cleared.
			dc.FillRectCPU(149/scale, 97/scale, 20/scale, 20/scale, RGBA{R: 1, A: 1})
			img := dc.Image().(*image.RGBA)
			bounds := image.Rect(149, 97, 169, 117)
			for y := range img.Bounds().Dy() {
				for x := range img.Bounds().Dx() {
					var want color.RGBA
					if image.Pt(x, y).In(bounds) {
						want = color.RGBA{R: 255, A: 255}
					}
					if got := img.RGBAAt(x, y); got != want {
						t.Fatalf("pixel (%d,%d) = %v, want %v", x, y, got, want)
					}
				}
			}
		})
	}
}
