package transcribe

import (
	"testing"
)

func TestBool(t *testing.T) {
	for _, b := range []bool{true, false} {
		v := Transcribe(b).(bool)
		if v != b {
			t.Errorf("got %t, expect %t", v, b)
		}
	}
}

func TestSlice(t *testing.T) {
	v := [][]int{{1, 2, 3}, {4, 5, 6}}
	g := Transcribe(v).([][]int)

	for i, row := range v {
		for j, e := range row {
			if g[i][j] != e {
				t.Errorf("copy slice faield: g[%d][%d] != %d", i, j, e)
			}
		}
	}

	v[1][1] = 1000
	if g[1][1] == 1000 {
		t.Error("failed to deep copy slice")
	}
}

func TestArray(t *testing.T) {
	v := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	g := Transcribe(v).([2][3]int)

	for i, row := range v {
		for j, e := range row {
			if g[i][j] != e {
				t.Errorf("faield to copy array: g[%d][%d] != %d", i, j, e)
			}
		}
	}

	v[1][1] = 1000
	if g[1][1] == 1000 {
		t.Error("failed to deep copy array")
	}
}

func TestStruct(t *testing.T) {
	type dummy struct {
		foo    int
		Slice  []int
		Slice2 [][]int
	}
	v := dummy{
		foo:    2,
		Slice:  []int{1, 2, 3},
		Slice2: [][]int{{1, 2, 3}, {4, 5, 6}},
	}

	g := Transcribe(v).(dummy)
	if g.foo != 2 {
		t.Error("failed to copy private field ")
	}
	v.Slice[1] = 1000
	if g.Slice[1] == 1000 {
		t.Error("failed to deep copy slice in struct")
	}
}

func TestMap(t *testing.T) {
	v := make(map[string][]int)
	v["Slice1"] = []int{1, 2, 3}
	v["Slice2"] = []int{4, 5, 6}

	g := Transcribe(v).(map[string][]int)
	for key, slice := range v {
		for i, val := range slice {
			if g[key][i] != val {
				t.Error("failed to copy slice in map")
			}
		}
	}
}
