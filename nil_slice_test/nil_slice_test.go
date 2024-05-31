package main

import "testing"

type Person struct {
	name string
}

func BenchmarkNonEmptySlice1000(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(1000)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 100)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNonEmptySlice500(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(500)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 100)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNonEmptySlice250(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(250)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 100)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNonEmptySlice100(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(100)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 100)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNonEmptySlice47(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(47)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 100)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNonEmptySlice10(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(10)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 100)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNonEmptySlice1(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(1)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 100)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNilSlice1000(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(1000)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var names []string
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNilSlice500(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(500)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var names []string
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNilSlice250(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(250)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var names []string
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNilSlice100(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(100)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var names []string
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNilSlice47(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(47)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var names []string
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNilSlice10(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(10)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var names []string
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNilSlice1(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(1)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var names []string
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func genPeople(numPeople int) []Person {
	people := make([]Person, numPeople, numPeople)
	for x := 0; x < numPeople; x++ {
		// everyone is named Ted, deal with it.
		people[x] = Person{name: "Ted"}
	}
	return people
}

func BenchmarkNonEmptyAdjustedSlice1000(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(1000)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 1000)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNonEmptyAdjustedSlice500(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(500)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 1000)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNonEmptyAdjustedSlice250(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(250)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 1000)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNonEmptyAdjustedSlice100(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(100)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 1000)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNonEmptyAdjustedSlice47(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(47)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 1000)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNonEmptyAdjustedSlice10(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(10)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 1000)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}

func BenchmarkNonEmptyAdjustedSlice1(b *testing.B) {
	// sshhh pretend we don't know how many their
	// are.
	people := genPeople(1)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		names := make([]string, 0, 1000)
		for _, person := range people {
			names = append(names, person.name)
		}
	}
}
