package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestBadgerCache_Has(t *testing.T) {
	err := tbc.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := tbc.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache when it should not be")
	}

	_ = tbc.Set("foo", "bar")

	inCache, err = tbc.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found in cache when it should be")
	}

	err = tbc.Forget("foo")
}

func TestBadgerCache_Get(t *testing.T) {
	err := tbc.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	x, err := tbc.Get("foo")
	if err != nil {
		t.Error(err)
	}

	if x != "bar" {
		t.Error("foo not found in cache when it should be")
	}

	err = tbc.Forget("foo")
}

func TestBadgerCache_Forget(t *testing.T) {
	err := tbc.Set("foo", "foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := tbc.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found in cache when it should be")
	}

	err = tbc.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err = tbc.Has("foo")
	if err != nil {
		t.Error(err)
	}
	if inCache {
		t.Error("foo found in cache when it should not be")
	}
}

func TestBadgerCache_Empty(t *testing.T) {
	err := tbc.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	err = tbc.Empty()
	if err != nil {
		t.Error(err)
	}

	inCache, err := tbc.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("alpha found in cache when it should not be")
	}
}

func TestBadgerCache_EmptyByMatch(t *testing.T) {
	err := tbc.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	err = tbc.Set("alpha2", "beta2")
	if err != nil {
		t.Error(err)
	}

	err = tbc.Set("beta", "beta3")
	if err != nil {
		t.Error(err)
	}

	err = tbc.EmptyByMatch("a")
	if err != nil {
		t.Error(err)
	}

	inCache, err := tbc.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("alpha found in cache when it should not be")
	}

	inCache, err = tbc.Has("alpha2")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("alpha2 found in cache when it should not be")
	}

	inCache, err = tbc.Has("beta")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("beta not found in cache when it should be")
	}
}

func TestBadgerCache_Set_WithTTL(t *testing.T) {
	err := tbc.Set("foo", "bar", 1)
	if err != nil {
		t.Error(err)
	}

	inCache, err := tbc.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found in cache when it should be")
	}

	err = tbc.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	//wait for TTL to expire
	time.Sleep(1100 * time.Millisecond) //This will slow tests.

	inCache, err = tbc.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache when it should not be")
	}
}

func TestBadgerCache_Empty_MoreThanMaxDeleteCapacity(t *testing.T) {
	for i := 0; i < 100005; i++ {
		err := tbc.Set(fmt.Sprintf("foo-%d", i), "bar")
		if err != nil {
			t.Error(err)
		}
	}

	err := tbc.Empty()
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 100005; i++ {
		inCache, err := tbc.Has(fmt.Sprintf("foo-%d", i))
		if err != nil {
			t.Error(err)
		}

		if inCache {
			t.Error("foo" + fmt.Sprintf("%d", i) + " found in cache when it should not be")
		}
	}
}
