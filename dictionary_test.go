package ectolinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDictionary(t *testing.T) {
	d := NewDictionary[int]()
	require.NotNil(t, d, "NewDictionary should not return nil")
	assert.Empty(t, d.values, "New dictionary should be empty")
}

func TestNewDictionaryWithCapacity(t *testing.T) {
	d := NewDictionaryWithCapacity[string](10)
	require.NotNil(t, d, "NewDictionaryWithCapacity should not return nil")
	// Note: We can't directly test the capacity of the underlying map
}

func TestToDictionary(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	d := ToDictionary(m)
	assert.Equal(t, m, d.values, "ToDictionary should create correct dictionary")
}

func TestDictionaryGetSet(t *testing.T) {
	d := NewDictionary[int]()
	d.Set("key", 42)

	value, ok := d.Get("key")
	assert.True(t, ok, "Get should return true for existing key")
	assert.Equal(t, 42, value, "Get should return correct value")

	_, ok = d.Get("nonexistent")
	assert.False(t, ok, "Get should return false for nonexistent key")
}

func TestDictionaryKeys(t *testing.T) {
	d := ToDictionary(map[string]int{"one": 1, "two": 2})
	keys := d.Keys()
	assert.Len(t, keys, 2, "Keys should return correct number of keys")
	assert.Contains(t, keys, "one", "Keys should contain 'one'")
	assert.Contains(t, keys, "two", "Keys should contain 'two'")
}

func TestDictionaryContainsKey(t *testing.T) {
	d := ToDictionary(map[string]int{"key": 42})
	assert.True(t, d.ContainsKey("key"), "ContainsKey should return true for existing key")
	assert.False(t, d.ContainsKey("nonexistent"), "ContainsKey should return false for nonexistent key")
}

func TestDictionaryContainsValue(t *testing.T) {
	d := ToDictionary(map[string]int{"key": 42})
	assert.True(t, d.ContainsValue(42), "ContainsValue should return true for existing value")
	assert.False(t, d.ContainsValue(0), "ContainsValue should return false for nonexistent value")
}

func TestDictionaryContainsWhere(t *testing.T) {
	d := ToDictionary(map[string]int{"one": 1, "two": 2})
	assert.True(t, d.ContainsWhere(func(k string, v int) bool { return k == "one" && v == 1 }), "ContainsWhere should return true for existing condition")
	assert.False(t, d.ContainsWhere(func(k string, v int) bool { return v > 10 }), "ContainsWhere should return false for nonexistent condition")
}

func TestDictionaryRemove(t *testing.T) {
	d := ToDictionary(map[string]int{"key": 42})
	d.Remove("key")
	assert.False(t, d.ContainsKey("key"), "Remove should remove the key")
}

func TestDictionaryRemoveValue(t *testing.T) {
	d := ToDictionary(map[string]int{"one": 1, "two": 2})
	d.RemoveValue(1)
	assert.False(t, d.ContainsValue(1), "RemoveValue should remove the value")
	assert.False(t, d.ContainsKey("one"), "RemoveValue should remove the key associated with the value")
}

func TestDictionaryRemoveWhere(t *testing.T) {
	d := ToDictionary(map[string]int{"one": 1, "two": 2, "three": 3})
	d.RemoveWhere(func(k string, v int) bool { return v%2 == 1 })
	assert.False(t, d.ContainsKey("one"), "RemoveWhere should remove 'one'")
	assert.False(t, d.ContainsKey("three"), "RemoveWhere should remove 'three'")
	assert.True(t, d.ContainsKey("two"), "RemoveWhere should not remove 'two'")
}

func TestDictionaryMerge(t *testing.T) {
	d1 := ToDictionary(map[string]int{"one": 1, "two": 2})
	d2 := ToDictionary(map[string]int{"three": 3, "four": 4})
	d1.Merge(*d2)
	assert.Equal(t, 4, d1.Count(), "Merge should result in correct number of elements")
	assert.True(t, d1.ContainsKey("three"), "Merge should add 'three'")
	assert.True(t, d1.ContainsKey("four"), "Merge should add 'four'")
}

func TestDictionaryMergeMaps(t *testing.T) {
	d := ToDictionary(map[string]int{"one": 1, "two": 2})
	m := map[string]int{"three": 3, "four": 4}
	d.MergeMaps(m)
	assert.Equal(t, 4, d.Count(), "MergeMaps should result in correct number of elements")
	assert.True(t, d.ContainsKey("three"), "MergeMaps should add 'three'")
	assert.True(t, d.ContainsKey("four"), "MergeMaps should add 'four'")
}

func TestDictionaryClear(t *testing.T) {
	d := ToDictionary(map[string]int{"one": 1, "two": 2})
	d.Clear()
	assert.Equal(t, 0, d.Count(), "Clear should remove all elements")
}

func TestDictionaryCount(t *testing.T) {
	d := ToDictionary(map[string]int{"one": 1, "two": 2})
	assert.Equal(t, 2, d.Count(), "Count should return correct number of elements")
}

func TestDictionaryToArray(t *testing.T) {
	d := ToDictionary(map[string]int{"one": 1, "two": 2})
	arr := d.ToArray()
	assert.Len(t, arr, 2, "ToArray should return correct number of elements")
	assert.Contains(t, arr, 1, "ToArray should contain 1")
	assert.Contains(t, arr, 2, "ToArray should contain 2")
}

func TestDictionaryToList(t *testing.T) {
	d := ToDictionary(map[string]int{"one": 1, "two": 2})
	list := d.ToList()
	assert.Len(t, list, 2, "ToList should return correct number of elements")
	assert.Contains(t, list, 1, "ToList should contain 1")
	assert.Contains(t, list, 2, "ToList should contain 2")
}

func TestDictionaryToMap(t *testing.T) {
	original := map[string]int{"one": 1, "two": 2}
	d := ToDictionary(original)
	m := d.ToMap()
	assert.Equal(t, original, m, "ToMap should return the original map")
}

func TestDictionaryValues(t *testing.T) {
	d := ToDictionary(map[string]int{"one": 1, "two": 2})
	values := d.Values()
	assert.Len(t, values, 2, "Values should return correct number of elements")
	assert.Contains(t, values, 1, "Values should contain 1")
	assert.Contains(t, values, 2, "Values should contain 2")
}
