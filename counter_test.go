package counter

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		args []string
		want map[string]int
	}{
		{[]string{}, map[string]int{}},
		{[]string{"a"}, map[string]int{"a": 1}},
		{[]string{"a", "b", "a"}, map[string]int{"a": 2, "b": 1}},
	}
	for _, test := range tests {
		c := New(test.args...)
		if !reflect.DeepEqual(c.data, test.want) {
			t.Errorf("got %v, want %v", c.data, test.want)
		}
	}
}

func TestSub(t *testing.T) {
	c := New("a")
	if cnt := c.Sub("a"); cnt != 0 {
		t.Errorf("got %d, want 0", cnt)
	}
	if cnt := c.Sub("b"); cnt != -1 {
		t.Errorf("got %d, want 0", cnt)
	}
	m := map[string]int{"a": 0, "b": -1}
	if !reflect.DeepEqual(c.data, m) {
		t.Errorf("got %v, want %v", c.data, m)
	}
}

func TestRemove(t *testing.T) {
	c := New("a")
	if ok := c.Remove("a"); !ok {
		t.Errorf("got %v, want true", ok)
	}
	if ok := c.Remove("b"); ok {
		t.Errorf("got %v, want false", ok)
	}
	if !reflect.DeepEqual(c.data, map[string]int{}) {
		t.Errorf("got %v, want empty map", c.data)
	}
}

func TestGet(t *testing.T) {
	c := New("a", "a")
	if cnt := c.Get("a"); cnt != 2 {
		t.Errorf("got %d, want 2", cnt)
	}
	if cnt := c.Get("b"); cnt != 0 {
		t.Errorf("got %d, want 0", cnt)
	}
}

func TestContains(t *testing.T) {
	c := New("a")
	if ok := c.Contains("a"); !ok {
		t.Errorf("got %v, want true", ok)
	}
	if ok := c.Contains("b"); ok {
		t.Errorf("got %v,, want false", ok)
	}
}

func TestLen(t *testing.T) {
	c := New("a", "a", "b")
	if n := c.Len(); n != 2 {
		t.Errorf("got %d, want 2", n)
	}
}

func TestTotal(t *testing.T) {
	c := New("a", "a", "b")
	if n := c.Total(); n != 3 {
		t.Errorf("got %d, want 3", n)
	}
}

func TestMostCommon(t *testing.T) {
	var tests = []struct {
		args []string
		want []ItemCount[string]
	}{
		{[]string{}, []ItemCount[string]{}},
		{[]string{"a", "b", "a", "b", "c", "a"},
			[]ItemCount[string]{
				{Item: "a", Count: 3},
				{Item: "b", Count: 2},
				{Item: "c", Count: 1},
			},
		},
	}
	for _, test := range tests {
		c := New(test.args...)
		for i := 0; i < 5; i++ {
			got := c.MostCommon(uint(i))
			want := test.want
			if i != 0 && i < len(want) {
				want = want[:i]
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	}
}

func TestItems(t *testing.T) {
	var tests = []struct {
		args []string
		want []string
	}{
		{[]string{}, []string{}},
		{[]string{"a", "b", "a", "b", "c", "a"}, []string{"a", "a", "a", "b", "b", "c"}},
	}
	for _, test := range tests {
		c := New(test.args...)
		got := c.Items()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}

func TestMap(t *testing.T) {
	c := New("a", "a", "b")
	want := map[string]int{"a": 2, "b": 1}
	if m := c.Map(); !reflect.DeepEqual(m, want) {
		t.Errorf("want %v, got %v", m, want)
	}
}

func TestClone(t *testing.T) {
	c := New("a", "a", "b")
	c2 := c.Clone()
	if !reflect.DeepEqual(c.data, c2.data) {
		t.Errorf("want %v, got %v", c.data, c2.data)
	}
}

func TestString(t *testing.T) {
	c := New("a", "a", "b")
	want := `Counter{Items: 2, Total: 3}`
	if s := c.String(); s != want {
		t.Errorf("got %q, want %q", s, want)
	}
}
