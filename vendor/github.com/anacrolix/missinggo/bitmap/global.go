package bitmap

import "github.com/RoaringBitmap/roaring"

func Sub(left, right Bitmap) Bitmap {
	return Bitmap{
		RB: roaring.AndNot(left.lazyRB(), right.lazyRB()),
	}
}

func Flip(bm Bitmap, start, end int) Bitmap {
	return Bitmap{
		RB: roaring.FlipInt(bm.lazyRB(), start, end),
	}
}
