package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		val  int
		want int
	}{
		{
			name: "value present in middle",
			arr:  []int{1, 4, 5, 7, 9},
			val:  5,
			want: 2,
		},
		{
			name: "value present at start",
			arr:  []int{1, 4, 5, 7, 9},
			val:  1,
			want: 0,
		},
		{
			name: "value present at end",
			arr:  []int{1, 4, 5, 7, 9},
			val:  9,
			want: 4,
		},
		{
			name: "value not present",
			arr:  []int{1, 4, 5, 7, 9},
			val:  11,
			want: -1,
		},
		{
			name: "empty slice",
			arr:  []int{},
			val:  5,
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.arr, tt.val)
			if got != tt.want {
				t.Fatalf("BinarySearch(%v,%v) = %d; want %d", tt.arr, tt.val, got, tt.want)
			}
		})
	}

}

func TestBinarySearchStr(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		val  string
		want int
	}{
		{
			name: "value present in middle",
			arr:  []string{"a", "b", "c", "d", "e"},
			val:  "c",
			want: 2,
		},
		{
			name: "value present at start",
			arr:  []string{"a", "b", "c", "d", "e"},
			val:  "a",
			want: 0,
		},
		{
			name: "value present at end",
			arr:  []string{"a", "b", "c", "d", "e"},
			val:  "e",
			want: 4,
		},
		{
			name: "value not present",
			arr:  []string{"a", "b", "c", "d", "e"},
			val:  "x",
			want: -1,
		},
		{
			name: "empty slice",
			arr:  []string{},
			val:  "e",
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.arr, tt.val)
			if got != tt.want {
				t.Fatalf("BinarySearch(%v,%v) = %d; want %d", tt.arr, tt.val, got, tt.want)
			}
		})
	}

}
