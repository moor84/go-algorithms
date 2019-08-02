package algorithms

import "testing"
import "reflect"

func TestHashMap_Put(t *testing.T) {
	hm := HashMap{}

	// Put
	hm.Put("one", "1")
	hm.Put("two", "2")
	hm.Put("dhdjkfdhsfjkdsfjkdshfjkdshfdsk", "value")
	hm.Put("Dfdjkfdhsfjkdsfjkdshfjkdshfdsk", "value2")
	hm.Put("one", "1")       // set again
	hm.Put("three", "here")  // collision
	hm.Put("three", "here2") // set again, collision

	// Get
	if hm.Get("one") != "1" {
		t.Error("one is not found")
	}
	if hm.Get("two") != "2" {
		t.Error("two is not found")
	}
	value := hm.Get("something")
	if value != "" {
		t.Errorf("something is not empty: %s\n", value)
	}
	value = hm.Get("Dfdjkfdhsfjkdsfjkdshfjkdshfdsk")
	if value != "value2" {
		t.Errorf("Dfdjkfdhsfjkdsfjkdshfjkdshfdsk is not found: %s\n", value)
	}
	value = hm.Get("three")
	if value != "here2" {
		t.Errorf("three is not found: %s\n", value)
	}

	// Size
	if hm.Size() != 5 {
		t.Errorf("Size is wrong: %d", hm.Size())
	}

	// Keys
	keys := hm.Keys()
	if !reflect.DeepEqual(keys, []string{
		"dhdjkfdhsfjkdsfjkdshfjkdshfdsk", "Dfdjkfdhsfjkdsfjkdshfjkdshfdsk", "three", "two", "one",
	}) {
		t.Error(keys)
	}

	// Values
	values := hm.Values()
	if !reflect.DeepEqual(values, []string{
		"value", "value2", "here2", "2", "1",
	}) {
		t.Error(values)
	}

	// Remove
	hm.Remove("one")
	if hm.Size() != 4 {
		t.Errorf("Size after removing is wrong: %d", hm.Size())
	}
	value = hm.Get("one")
	if value != "" {
		t.Errorf("one is not empty: %s\n", value)
	}
	hm.Remove("two")
	hm.Remove("something") // not there
	hm.Remove("dhdjkfdhsfjkdsfjkdshfjkdshfdsk")
	hm.Remove("Dfdjkfdhsfjkdsfjkdshfjkdshfdsk")
	value = hm.Get("Dfdjkfdhsfjkdsfjkdshfjkdshfdsk")
	if value != "" {
		t.Errorf("Dfdjkfdhsfjkdsfjkdshfjkdshfdsk is not empty: %s\n", value)
	}
	value = hm.Get("three")
	if value != "here2" {
		t.Errorf("three is not found: %s\n", value)
	}
	hm.Remove("three")
	value = hm.Get("three")
	if value != "" {
		t.Errorf("three is not empty: %s\n", value)
	}
	if hm.Size() != 0 {
		t.Errorf("Size after all keys is wrong: %d", hm.Size())
	}
	hm.Remove("something") // not there
	if hm.Size() != 0 {
		t.Errorf("Size after all keys is wrong: %d", hm.Size())
	}
}
