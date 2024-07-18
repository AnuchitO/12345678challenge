package main

import (
	"testing"
)

func Test12345678Challenge(t *testing.T) {
	t.Run("topHalf", func(t *testing.T) {
		words := [8]string{"1", "2", "3", "4", "5", "6", "7", "8"}
		want := ` 1 2 3 4 5 6 7 8
  2 3 4 5 6 7 8
   3 4 5 6 7 8
    4 5 6 7 8
     5 6 7 8
      6 7 8
       7 8
        8
`

		lyric := topHalf(words)

		if lyric != want {
			t.Errorf("lyric %q, want %q", lyric, want)
		}
	})

	t.Run("bottomHalf", func(t *testing.T) {
		words := [8]string{"1", "2", "3", "4", "5", "6", "7", "8"}
		want := `        8
       7 8
      6 7 8
     5 6 7 8
    4 5 6 7 8
   3 4 5 6 7 8
  2 3 4 5 6 7 8
 1 2 3 4 5 6 7 8
`

		lyric := bottomHalf(words)

		if lyric != want {
			t.Errorf("lyric %q, want %q", lyric, want)
		}
	})

	t.Run("challenge", func(t *testing.T) {
		words := [8]string{"1", "2", "3", "4", "5", "6", "7", "8"}
		want := ` 1 2 3 4 5 6 7 8
  2 3 4 5 6 7 8
   3 4 5 6 7 8
    4 5 6 7 8
     5 6 7 8
      6 7 8
       7 8
        8
        8
       7 8
      6 7 8
     5 6 7 8
    4 5 6 7 8
   3 4 5 6 7 8
  2 3 4 5 6 7 8
 1 2 3 4 5 6 7 8
`

		lyric := challenge(words)

		if lyric != want {
			t.Errorf("lyric %q, want %q", lyric, want)
		}
	})

	t.Run("mor lum", func(t *testing.T) {
		words := [8]string{"หมอ", "ลำ", "นี่", "มา", "ตั้ง", "แต่", "เวียง", "จันทน์"}
		want := ` หมอ ลำ นี่ มา ตั้ง แต่ เวียง จันทน์
  ลำ นี่ มา ตั้ง แต่ เวียง จันทน์
   นี่ มา ตั้ง แต่ เวียง จันทน์
    มา ตั้ง แต่ เวียง จันทน์
     ตั้ง แต่ เวียง จันทน์
      แต่ เวียง จันทน์
       เวียง จันทน์
        จันทน์
`

		lyric := topHalf(words)

		if lyric != want {
			t.Errorf("lyric %q, want %q", lyric, want)
		}
	})
}
