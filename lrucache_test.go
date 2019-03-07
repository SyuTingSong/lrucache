package lrucache

import "testing"

func TestLRUCache_Get(t *testing.T) {
	lc := New(3)

	lc.Put("abc", "cba")
	lc.Put("def", 123)
	if lc.Len() == 2 {
		t.Log("Put 2 elements")
	} else {
		t.Error("Cache length error")
	}

	v, _ := lc.Get("abc")
	if v == "cba" {
		t.Log("abc -> cba")
	} else {
		t.Error("cache value mismatch abc -> cba")
	}
	v, _ = lc.Get("def")
	if v == 123 {
		t.Log("def -> 123")
	} else {
		t.Error("cache value mismatch def -> 123")
	}
}

func TestOverflow(t *testing.T) {
	lc := New(3)

	lc.Put("abc", "cba")
	lc.Put("def", 123)
	lc.Put("ghi", "welcome")
	lc.Put("jkl", "lp")

	t.Log("Put 4 elements")

	if lc.Len() == 3 {
		t.Log("Cache length ", lc.Len())
	} else {
		t.Error("Cache overflow")
	}

	_, ok := lc.Get("abc")
	if ok {
		t.Error("abc hasn't been kick out")
	} else {
		t.Log("abc has been kick out")
	}
}

func TestOverwrite(t *testing.T) {
	lc := New(3)

	lc.Put("abc", "cba")
	lc.Put("abc", "wel")
	v, _ := lc.Get("abc")
	if v == "wel" {
		t.Log("abc -> wel")
	} else {
		t.Error("overwrite existing key failed")
	}

	t.Log("Cache length ", lc.Len())
}

func TestLRUCache_Remove(t *testing.T) {
	lc := New(3)

	lc.Put("abc", 123)
	lc.Put("def", 123)
	lc.Del("abc")

	t.Log("Cache length ", lc.Len())
	_, ok := lc.Get("abc")
	if ok {
		t.Error("Del abc failed")
	} else {
		t.Log("abc not found")
	}
}

func BenchmarkLRUCache_Put(b *testing.B) {
	lc := New(3000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lc.Put(i, "abc")
	}
}

func BenchmarkLRUCache_Get(b *testing.B) {
	lc := New(3000)

	for i := 0; i < 1000; i++ {
		lc.Put(i, "abc")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lc.Get(i)
	}
}
