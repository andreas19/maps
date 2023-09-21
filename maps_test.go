package maps

import (
	"reflect"
	"sort"
	"testing"
)

func TestClone(t *testing.T) {
	var tests = []map[string]int{
		nil,
		{},
		{"a": 1},
		{"a": 1, "b": 2},
	}
	for i, test := range tests {
		got := Clone(test)
		if !reflect.DeepEqual(got, test) {
			t.Errorf("%d: got %v, want %v", i, got, test)
		}
	}
}

func TestUpdate(t *testing.T) {
	var tests = []struct{ m1, m2, want map[string]int }{
		{map[string]int{}, map[string]int{}, map[string]int{}},
		{map[string]int{}, nil, map[string]int{}},
		{map[string]int{}, map[string]int{"a": 1}, map[string]int{"a": 1}},
		{map[string]int{"a": 1}, map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2}},
	}
	for i, test := range tests {
		Update(test.m1, test.m2)
		if !reflect.DeepEqual(test.m1, test.want) {
			t.Errorf("%d: got %v, want %v", i, test.m1, test.want)
		}
	}
}

func TestUpdateNil(t *testing.T) {
	defer func() { _ = recover() }()
	var m1 map[string]int
	m2 := map[string]int{}
	Update(m1, m2)
	t.Errorf("did not panic")
}

func TestMerge(t *testing.T) {
	var tests = []struct{ m1, m2, want map[string]int }{
		{map[string]int{}, map[string]int{}, map[string]int{}},
		{map[string]int{}, nil, map[string]int{}},
		{map[string]int{}, map[string]int{"a": 1}, map[string]int{"a": 1}},
		{map[string]int{"a": 1}, map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2}},
	}
	for i, test := range tests {
		m := Merge(test.m1, test.m2)
		if !reflect.DeepEqual(m, test.want) {
			t.Errorf("%d: got %v, want %v", i, m, test.want)
		}
	}
}

func TestMergeNil(t *testing.T) {
	defer func() { _ = recover() }()
	var m1 map[string]int
	m2 := map[string]int{}
	Merge(m1, m2)
	t.Errorf("did not panic")
}

func TestClear(t *testing.T) {
	var tests = []map[string]int{
		{},
		{"a": 1},
		{"a": 1, "b": 2},
	}
	var m map[string]int
	Clear(m)
	if m != nil {
		t.Errorf("got %v, want nil map", m)
	}
	want := map[string]int{}
	for i, test := range tests {
		Clear(test)
		if !reflect.DeepEqual(test, want) {
			t.Errorf("%d: got %v, want empty map", i, test)
		}
	}
}

func TestContains(t *testing.T) {
	var tests = []struct {
		m    map[string]int
		want bool
	}{
		{nil, false},
		{map[string]int{}, false},
		{map[string]int{"a": 1}, true},
		{map[string]int{"b": 2}, false},
	}
	for i, test := range tests {
		if got := Contains(test.m, "a"); got != test.want {
			t.Errorf("%d: got %v, want %v", i, got, test.want)
		}
	}
}

func TestGet(t *testing.T) {
	key := "a"
	dflt := -1
	var tests = []struct {
		m    map[string]int
		want int
	}{
		{nil, dflt},
		{map[string]int{}, dflt},
		{map[string]int{key: 1}, 1},
		{map[string]int{"b": 2}, dflt},
	}
	for i, test := range tests {
		if got := Get(test.m, key, dflt); got != test.want {
			t.Errorf("%d: got %d, want %d", i, got, test.want)
		}
	}
}

func TestKeys(t *testing.T) {
	var tests = []struct {
		m    map[string]int
		want []string
	}{
		{nil, nil},
		{map[string]int{}, []string{}},
		{map[string]int{"a": 1}, []string{"a"}},
		{map[string]int{"a": 1, "b": 2}, []string{"a", "b"}},
	}
	for i, test := range tests {
		got := Keys(test.m)
		sort.Slice(got, func(i, j int) bool { return got[i] < got[j] })
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %v, want %v", i, got, test.want)
		}
	}
}

func TestValues(t *testing.T) {
	var tests = []struct {
		m    map[string]int
		want []int
	}{
		{nil, nil},
		{map[string]int{}, []int{}},
		{map[string]int{"a": 1}, []int{1}},
		{map[string]int{"a": 1, "b": 2}, []int{1, 2}},
	}
	for i, test := range tests {
		got := Values(test.m)
		sort.Slice(got, func(i, j int) bool { return got[i] < got[j] })
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %v, want %v", i, got, test.want)
		}
	}
}

func TestEqual(t *testing.T) {
	var tests = []struct {
		m1, m2 map[string]int
		want   bool
	}{
		{nil, nil, true},
		{map[string]int{}, nil, false},
		{map[string]int{"a": 1}, map[string]int{}, false},
		{map[string]int{"a": 1}, map[string]int{"a": 1}, true},
		{map[string]int{"a": 1}, map[string]int{"a": 2}, false},
		{map[string]int{"a": 1}, map[string]int{"a": 1, "b": 2}, false},
		{map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2}, true},
		{map[string]int{"a": 1, "b": 2}, map[string]int{"b": 2, "a": 1}, true},
	}
	for i, test := range tests {
		got1 := Equal(test.m1, test.m2)
		got2 := Equal(test.m2, test.m1)
		if got1 != test.want || got2 != test.want {
			t.Errorf("%d: got %v and %v, want %v", i, got1, got2, test.want)
		}
	}
}

func TestItems(t *testing.T) {
	var tests = []struct {
		m    map[string]int
		want []Item[string, int]
	}{
		{nil, nil},
		{map[string]int{}, []Item[string, int]{}},
		{map[string]int{"a": 1}, []Item[string, int]{{"a", 1}}},
		{map[string]int{"a": 1, "b": 2}, []Item[string, int]{{"a", 1}, {"b", 2}}},
	}
	for i, test := range tests {
		got := Items(test.m)
		sort.Slice(got, func(i, j int) bool { return got[i].Key < got[j].Key })
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %v, want %v", i, got, test.want)
		}
	}
}

func TestFromItems(t *testing.T) {
	var tests = []struct {
		items []Item[string, int]
		want  map[string]int
	}{
		{nil, nil},
		{[]Item[string, int]{}, map[string]int{}},
		{[]Item[string, int]{{"a", 1}}, map[string]int{"a": 1}},
		{[]Item[string, int]{{"a", 1}, {"b", 2}}, map[string]int{"a": 1, "b": 2}},
	}
	for i, test := range tests {
		if got := FromItems(test.items); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %v, want %v", i, got, test.want)
		}
	}
}

func TestFromSlices(t *testing.T) {
	var tests = []struct {
		keys   []string
		values []int
		want   map[string]int
	}{
		{nil, nil, nil},
		{nil, []int{}, nil},
		{[]string{}, nil, map[string]int{}},
		{[]string{"a", "b"}, nil, map[string]int{"a": 0, "b": 0}},
		{[]string{"a", "b"}, []int{1}, map[string]int{"a": 1, "b": 0}},
		{[]string{"a", "b"}, []int{1, 2}, map[string]int{"a": 1, "b": 2}},
		{[]string{"a", "b"}, []int{1, 2, 3}, map[string]int{"a": 1, "b": 2}},
	}
	for i, test := range tests {
		if got := FromSlices(test.keys, test.values); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %v, want %v", i, got, test.want)
		}
	}
}

func TestFromFuncs(t *testing.T) {
	first := true
	keyFn := func() string {
		if first {
			first = false
			return "a"
		}
		return "b"
	}
	valFn := func() int { return 1 }
	var tests = []struct {
		size int
		valF func() int
		want map[string]int
	}{
		{0, valFn, map[string]int{}},
		{1, valFn, map[string]int{"a": 1}},
		{2, valFn, map[string]int{"a": 1, "b": 1}},
		{0, nil, map[string]int{}},
		{1, nil, map[string]int{"a": 0}},
		{2, nil, map[string]int{"a": 0, "b": 0}},
	}
	for i, test := range tests {
		first = true
		if got := FromFuncs(test.size, keyFn, test.valF); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestFromFuncsNegSize(t *testing.T) {
	defer func() { _ = recover() }()
	FromFuncs(-1, func() string { return "a" }, func() int { return 1 })
	t.Error("did not panic")
}

func TestFromFuncsNilKeys(t *testing.T) {
	defer func() { _ = recover() }()
	FromFuncs[string](1, nil, func() int { return 1 })
	t.Error("did not panic")
}

func TestKeysForValue(t *testing.T) {
	value := 1
	var tests = []struct {
		m    map[string]int
		want []string
	}{
		{nil, nil},
		{map[string]int{}, []string{}},
		{map[string]int{"b": 2}, []string{}},
		{map[string]int{"a": 1, "b": 2}, []string{"a"}},
		{map[string]int{"a": 1, "b": 2, "c": 1}, []string{"a", "c"}},
	}
	for i, test := range tests {
		got := KeysForValue(test.m, value)
		sort.Slice(got, func(i, j int) bool { return got[i] < got[j] })
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: got %v, want %v", i, got, test.want)
		}
	}
}

func TestDelete(t *testing.T) {
	var tests = []struct {
		m        map[string]int
		fn       func(k string, v int) bool
		want_int int
		want_map map[string]int
	}{
		{
			nil,
			func(k string, v int) bool { return true },
			0,
			nil,
		},
		{
			map[string]int{},
			func(k string, v int) bool { return true },
			0,
			map[string]int{},
		},
		{
			map[string]int{"a": 1, "b": 2, "c": 1},
			func(k string, v int) bool { return true },
			3,
			map[string]int{},
		},
		{
			map[string]int{"a": 1, "b": 2, "c": 1},
			func(k string, v int) bool { return false },
			0,
			map[string]int{"a": 1, "b": 2, "c": 1},
		},
		{
			map[string]int{"a": 1, "b": 2, "c": 1},
			func(k string, v int) bool { return v < 2 },
			2,
			map[string]int{"b": 2},
		},
		{
			map[string]int{"a": 1, "b": 2, "c": 1},
			func(k string, v int) bool { return k == "b" },
			1,
			map[string]int{"a": 1, "c": 1},
		},
	}
	for i, test := range tests {
		if got := Delete(test.m, test.fn); got != test.want_int || !reflect.DeepEqual(test.m, test.want_map) {
			t.Errorf("%d: got %d and %#v, want %d and %#v", i, got, test.m, test.want_int, test.want_map)
		}
	}
}
