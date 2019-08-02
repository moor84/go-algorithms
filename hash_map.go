package algorithms

import "hash/fnv"

// HashMapNode hash map node
type HashMapNode struct {
	key   string
	value string
	next  *HashMapNode
}

// HashMap implementation
type HashMap struct {
	data [16]*HashMapNode
	size int
}

func (hm *HashMap) getIndex(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	hashCode := h.Sum32()
	return (hashCode & 0x7fffffff) % 16
}

// Put value into the hash map
func (hm *HashMap) Put(key string, value string) {
	idx := hm.getIndex(key)
	entry := hm.data[idx]
	node := &HashMapNode{key, value, nil}

	if entry == nil {
		hm.data[idx] = node
		hm.size++
		return
	}

	if entry.key == key {
		entry.value = value
		return
	}

	for entry.next != nil {
		entry = entry.next
		if entry.key == key {
			entry.value = value
			return
		}
	}
	entry.next = node
	hm.size++
}

// Remove key/value from the map
func (hm *HashMap) Remove(key string) {
	idx := hm.getIndex(key)
	entry := hm.data[idx]
	if entry == nil {
		return
	}
	if entry.key == key {
		if entry.next != nil {
			hm.data[idx] = entry.next
		} else {
			hm.data[idx] = nil
		}
		hm.size--
		return
	}
	for entry.next != nil {
		if entry.next.key == key {
			entry.next = entry.next.next
			hm.size--
			return
		}
		entry = entry.next
	}
}

// Get value from the hash map
func (hm HashMap) Get(key string) string {
	idx := hm.getIndex(key)
	entry := hm.data[idx]
	if entry == nil {
		return ""
	}
	for entry != nil {
		if entry.key == key {
			return entry.value
		}
		entry = entry.next
	}
	return ""
}

// Size of the hash map
func (hm HashMap) Size() int {
	return hm.size
}

func (hm HashMap) traverse() []*HashMapNode {
	var entries []*HashMapNode
	for _, entry := range hm.data {
		for entry != nil {
			entries = append(entries, entry)
			entry = entry.next
		}
	}
	return entries
}

// Keys - returns all keys
func (hm HashMap) Keys() []string {
	var keys []string
	for _, entry := range hm.traverse() {
		keys = append(keys, entry.key)
	}
	return keys
}

// Values - returns all keys
func (hm HashMap) Values() []string {
	var values []string
	for _, entry := range hm.traverse() {
		values = append(values, entry.value)
	}
	return values
}
