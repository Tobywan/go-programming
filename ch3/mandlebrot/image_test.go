package mandlebrot

import (
	"reflect"
	"testing"
)

func TestNewArgand(tst *testing.T) {
	type input struct {
		z complex128
		h float64
		w float64
	}
	type test struct {
		in   input
		want Argand
	}

	tests := []test{
		{
			in:   input{0 + 0i, 2, 2},
			want: Argand{-1, 1, -1, 1, 0},
		},
		{
			in:   input{1 + 1i, 0.25, 0},
			want: Argand{1, 1, 0.875, 1.125, 1 + 1i},
		},
	}

	for _, t := range tests {
		got := NewArgand(t.in.z, t.in.h, t.in.w)
		if !reflect.DeepEqual(*got, t.want) {
			tst.Errorf("newArgand(%v)=%v, want %v", t.in, *got, t.want)
		}
	}
}

func TestMapComplexCenter(tst *testing.T) {
	c := Canvas{1000, 1000, nil}
	a := NewArgand(0+0i, 1, 1)

	type canvasPoint struct {
		x, y int
	}

	tests := []struct {
		in   canvasPoint
		want complex128
	}{
		{
			in:   canvasPoint{500, 500},
			want: 0 + 0i,
		},
		{
			in:   canvasPoint{0, 0},
			want: -0.5 + 0.5i,
		},
		{
			in:   canvasPoint{1000, 1000},
			want: 0.5 - 0.5i,
		},
		{
			in:   canvasPoint{0, 1000},
			want: -0.5 - 0.5i,
		},
		{
			in:   canvasPoint{500, 0},
			want: 0 + 0.5i,
		},
	}

	for _, t := range tests {
		got := mapComplex(t.in.x, t.in.y, c, *a)
		if !reflect.DeepEqual(t.want, got) {
			tst.Errorf("mapComplex(%v)=%g, want %g", t.in, got, t.want)
		}
	}

}
func TestMapComplexOffset(tst *testing.T) {
	c := Canvas{1000, 1000, nil}
	a := NewArgand(-1-1i, 2, 2)

	type canvasPoint struct {
		x, y int
	}

	tests := []struct {
		in   canvasPoint
		want complex128
	}{
		{
			in:   canvasPoint{500, 500},
			want: -1 - 1i,
		},
		{
			in:   canvasPoint{0, 0},
			want: -2 + 0i,
		},
		{
			in:   canvasPoint{1000, 1000},
			want: 0 - 2i,
		},
		{
			in:   canvasPoint{0, 1000},
			want: -2 - 2i,
		},
		{
			in:   canvasPoint{500, 0},
			want: -1 + 0i,
		},
	}

	for _, t := range tests {
		got := mapComplex(t.in.x, t.in.y, c, *a)
		if !reflect.DeepEqual(t.want, got) {
			tst.Errorf("mapComplex(%v)=%g, want %g", t.in, got, t.want)
		}
	}

}
