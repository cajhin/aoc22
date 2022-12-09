package main

// start with go test -bench=. -benchmem -memprofile mprofile.out -cpuprofile cprofile.out

import (
	"os"
	"testing"
)

// jjj 1042, 2980
const day06SizePart1 = 4
const day06SizePart2 = 100 //14

var day06Tests = []struct {
	stream string
	part1  int
	part2  int
}{
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 19},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 23},
	{"nppdvjthqldpwncqszvftbrmjlhg", 6, 23},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 29},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 26},
}

func TestDay06ExamplesPart1(t *testing.T) {
	for _, tt := range day06Tests {
		id := tt.stream
		t.Run(id, func(t *testing.T) {
			want := tt.part1
			got := FindFirstMarker(tt.stream, day06SizePart1)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay06Part1(t *testing.T) {
	const want = 1042
	buf, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	got := FindFirstMarker(string(buf), day06SizePart1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay06Part2(t *testing.T) {
	const want = 2980
	buf, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	got := FindFirstMarker(string(buf), day06SizePart2)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay06Part1(b *testing.B) {
	buf, err := os.ReadFile("input")
	if err != nil {
		b.Fatal(err)
	}
	s := string(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindFirstMarker(s, day06SizePart1)
	}
}
func BenchmarkDay06Part2(b *testing.B) {
	buf, err := os.ReadFile("input")
	if err != nil {
		b.Fatal(err)
	}
	s := string(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindFirstMarker(s, day06SizePart2)
	}
}

func Benchmark_jot1(b *testing.B) {
	buf, err := os.ReadFile("input")
	if err != nil {
		b.Fatal(err)
	}
	s := string(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jot_day06OnesCount(s, day06SizePart1)
	}
}
func Benchmark_jot2(b *testing.B) {
	buf, err := os.ReadFile("input")
	if err != nil {
		b.Fatal(err)
	}
	s := string(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jot_day06OnesCount(s, day06SizePart2)
	}
}
