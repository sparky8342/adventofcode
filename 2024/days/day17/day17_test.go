package day17

import "testing"

func Test1(t *testing.T) {
	data := []string{
		"Register A: 0",
		"Register B: 0",
		"Register C: 9",
		"",
		"Program: 2,6",
	}

	computer := parse_data(data)
	_ = computer.run_program()

	got := computer.b
	want := 1

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"Register A: 10",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 5,0,5,1,5,4",
	}

	computer := parse_data(data)

	got := computer.run_program()
	want := "0,1,2"

	if want != got {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"Register A: 2024",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,1,5,4,3,0",
	}

	computer := parse_data(data)

	got := computer.run_program()
	want := "4,2,5,6,7,7,7,7,3,1,0"

	if want != got {
		t.Errorf("got %s, wanted %s", got, want)
	}

	got_a := computer.a
	want_a := 0

	if want_a != got_a {
		t.Errorf("got %d, wanted %d", got_a, want_a)
	}

}

func Test4(t *testing.T) {
	data := []string{
		"Register A: 0",
		"Register B: 29",
		"Register C: 0",
		"",
		"Program: 1,7",
	}

	computer := parse_data(data)
	_ = computer.run_program()

	got := computer.b
	want := 26

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}

}

func Test5(t *testing.T) {
	data := []string{
		"Register A: 0",
		"Register B: 2024",
		"Register C: 43690",
		"",
		"Program: 4,0",
	}

	computer := parse_data(data)
	_ = computer.run_program()

	got := computer.b
	want := 44354

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}

}

func Test6(t *testing.T) {
	data := []string{
		"Register A: 729",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,1,5,4,3,0",
	}

	computer := parse_data(data)

	got := computer.run_program()
	want := "4,6,3,5,6,3,5,2,1,0"

	if want != got {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
