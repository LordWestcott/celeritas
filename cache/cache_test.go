package cache

import "testing"

func TestRedisCache_Has(t *testing.T) {
	err := trc.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := trc.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("Foo IS Found in Cache and it shouldn't be there!")
	}

	err = trc.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	inCache, err = trc.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("Foo NOT Found in Cache but it should be there!")
	}
}

func TestRedisCache_Get(t *testing.T) {
	err := trc.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	x, err := trc.Get("foo")
	if err != nil {
		t.Error(err)
	}

	if x != "bar" {
		t.Error("Foo not found in cache")
	}
}

func TestRedisCache_Forget(t *testing.T) {
	err := trc.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	err = trc.Forget("alpha")
	if err != nil {
		t.Error(err)
	}

	inCache, err := trc.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("Alpha IS Found in Cache and it shouldn't be there!")
	}
}

func TestRedisCache_Empty(t *testing.T) {
	err := trc.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	err = trc.Empty()
	if err != nil {
		t.Error(err)
	}

	inCache, err := trc.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("Alpha IS Found in Cache and it shouldn't be there!")
	}
}

func TestRedisCache_EmptyByMatch(t *testing.T) {
	err := trc.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	err = trc.Set("alpha2", "foo")
	if err != nil {
		t.Error(err)
	}

	err = trc.Set("beta", "funk")
	if err != nil {
		t.Error(err)
	}

	err = trc.EmptyByMatch("alpha")
	if err != nil {
		t.Error(err)
	}

	inCache, err := trc.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("Alpha IS Found in Cache and it shouldn't be there!")
	}

	inCache, err = trc.Has("alpha2")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("Alpha2 IS Found in Cache and it shouldn't be there!")
	}

	inCache, err = trc.Has("beta")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("Beta NOT Found in Cache and it should be there!")
	}
}

func TestEncodeDecode(t *testing.T) {
	entry := Entry{}
	entry["foo"] = "bar"

	bytes, err := encode(entry)
	if err != nil {
		t.Error(err)
	}

	decoded, err := decode(string(bytes))
	if err != nil {
		t.Error(err)
	}

	if decoded["foo"] != "bar" {
		t.Error("Foo not found in decoded map")
	}
}
