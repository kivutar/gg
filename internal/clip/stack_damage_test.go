package clip

import "testing"

func TestRasterizedMaskFractionalOriginMatchesCoverage(t *testing.T) {
	for _, clipped := range []bool{false, true} {
		stack := NewClipStack(NewRect(0, 0, 100, 100))
		stack.PushRRect(NewRect(10.75, 10.75, 40, 40), 4)
		if clipped {
			stack.PushRect(NewRect(0, 0, 31, 31))
		}
		mask := stack.RasterizedMask()
		for y := range 100 {
			for x := range 100 {
				var got byte
				mx, my := x-mask.OriginX, y-mask.OriginY
				if mx >= 0 && my >= 0 && mx < mask.Width && my < mask.Height {
					got = mask.Pixels[my*mask.Width+mx]
				}
				want := stack.Coverage(float64(x)+.5, float64(y)+.5)
				if got != want {
					t.Fatalf("clipped=%v: mask at (%d,%d) = %d, coverage = %d", clipped, x, y, got, want)
				}
			}
		}
	}
}
