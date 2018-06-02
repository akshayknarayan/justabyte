package justabyte

import (
	"io"
	"testing"
)

func TestRead(t *testing.T) {
	// small read = 1 MB
	// chunks = 1 KB
	buf := make([]byte, 1024)
	tot := 0
	b := New(1)
	for {
		n, err := b.Read(buf)
		if n != 1024 && n != 0 {
			t.Fatalf("Wrong read size: %d", n)
		}

		if err != nil && err == io.EOF {
			if tot != 1024 { // 1 MB / 1 KB = 1024 chunks
				t.Fatalf("wrong total size: %d", tot)
			}

			return
		} else if err != nil {
			panic(err)
		}

		tot++
	}
}

func benchReadSize(b *testing.B, totSizeMB int, chunkSize int) {
	buf := make([]byte, chunkSize)
	for i := 0; i < b.N; i++ {
		tot := 0
		b.SetBytes(int64(totSizeMB) * 1024 * 1024)
		r := New(uint32(totSizeMB))
		for {
			n, err := r.Read(buf)
			if n != 1024 && n != 0 {
				b.Fatalf("Wrong read size: %d", n)
			}

			if err != nil && err == io.EOF {
				if tot != totSizeMB*1024*1024/chunkSize {
					b.Fatalf("Wrong total size: %d", tot)
				}
				break
			} else if err != nil {
				panic(err)
			}

			tot++
		}
	}
}

func BenchmarkSmallRead(b *testing.B) {
	benchReadSize(b, 1, 1024)
}

func BenchmarkMediumRead(b *testing.B) {
	benchReadSize(b, 128, 1024)
}

func BenchmarkLongRead(b *testing.B) {
	benchReadSize(b, 1024, 1024)
}
