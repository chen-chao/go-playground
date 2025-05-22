## Transcribe

Transcribe is to get a deep copy of any variable in golang.

Usage:

``` go
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
```

Issues and improvments are welcome.
