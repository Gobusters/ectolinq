package ectolinq

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConcurrentDictionary(t *testing.T) {
	t.Run("NewConcurrentDictionary", func(t *testing.T) {
		d := NewConcurrentDictionary[int]()
		require.NotNil(t, d, "NewConcurrentDictionary should not return nil")
	})

	t.Run("ToConcurrentDictionary", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		d := ToConcurrentDictionary(m)
		assert.Equal(t, 2, d.Count(), "Expected count 2")
	})

	t.Run("Get and Set", func(t *testing.T) {
		d := NewConcurrentDictionary[int]()
		d.Set("key", 42)
		value, ok := d.Get("key")
		assert.True(t, ok, "Get should return true for existing key")
		assert.Equal(t, 42, value, "Get should return correct value")
	})

	t.Run("Keys", func(t *testing.T) {
		d := ToConcurrentDictionary(map[string]int{"a": 1, "b": 2})
		keys := d.Keys()
		assert.Len(t, keys, 2, "Keys should return correct number of keys")
		assert.Contains(t, keys, "a", "Keys should contain 'a'")
		assert.Contains(t, keys, "b", "Keys should contain 'b'")
	})

	t.Run("ContainsKey and ContainsValue", func(t *testing.T) {
		d := ToConcurrentDictionary(map[string]int{"a": 1, "b": 2})
		assert.True(t, d.ContainsKey("a"), "ContainsKey should return true for existing key")
		assert.True(t, d.ContainsValue(2), "ContainsValue should return true for existing value")
	})

	t.Run("ContainsWhere", func(t *testing.T) {
		d := ToConcurrentDictionary(map[string]int{"a": 1, "b": 2})
		assert.True(t, d.ContainsWhere(func(k string, v int) bool { return k == "b" && v == 2 }), "ContainsWhere should return true for existing condition")
	})

	t.Run("Remove", func(t *testing.T) {
		d := ToConcurrentDictionary(map[string]int{"a": 1, "b": 2})
		d.Remove("a")
		assert.False(t, d.ContainsKey("a"), "Remove should remove the key")
	})

	t.Run("RemoveValue", func(t *testing.T) {
		d := ToConcurrentDictionary(map[string]int{"a": 1, "b": 2})
		d.RemoveValue(1)
		assert.False(t, d.ContainsValue(1), "RemoveValue should remove the value")
	})

	t.Run("RemoveWhere", func(t *testing.T) {
		d := ToConcurrentDictionary(map[string]int{"a": 1, "b": 2})
		d.RemoveWhere(func(k string, v int) bool { return v > 1 })
		assert.False(t, d.ContainsKey("b"), "RemoveWhere should remove 'b'")
	})

	t.Run("Clear and Count", func(t *testing.T) {
		d := ToConcurrentDictionary(map[string]int{"a": 1, "b": 2})
		d.Clear()
		assert.Equal(t, 0, d.Count(), "Clear should remove all elements")
	})

	t.Run("ToArray and ToList", func(t *testing.T) {
		d := ToConcurrentDictionary(map[string]int{"a": 1, "b": 2})
		arr := d.ToArray()
		list := d.ToList()
		assert.Len(t, arr, 2, "ToArray should return correct number of elements")
		assert.Len(t, list, 2, "ToList should return correct number of elements")
		assert.Contains(t, arr, 1, "ToArray should contain 1")
		assert.Contains(t, arr, 2, "ToArray should contain 2")
		assert.Contains(t, list, 1, "ToList should contain 1")
		assert.Contains(t, list, 2, "ToList should contain 2")
	})

	t.Run("ToMap", func(t *testing.T) {
		original := map[string]int{"a": 1, "b": 2}
		d := ToConcurrentDictionary(original)
		m := d.ToMap()
		assert.Equal(t, original, m, "ToMap should return the original map")
	})

	t.Run("Merge", func(t *testing.T) {
		d1 := ToConcurrentDictionary(map[string]int{"a": 1, "b": 2})
		d2 := ToConcurrentDictionary(map[string]int{"c": 3, "d": 4})
		d1.Merge(d2)
		assert.Equal(t, 4, d1.Count(), "Merge should result in correct number of elements")
		assert.True(t, d1.ContainsKey("c"), "Merge should add 'c'")
		assert.True(t, d1.ContainsKey("d"), "Merge should add 'd'")
	})

	t.Run("MergeMaps", func(t *testing.T) {
		d := ToConcurrentDictionary(map[string]int{"a": 1, "b": 2})
		d.MergeMaps(map[string]int{"c": 3, "d": 4})
		assert.Equal(t, 4, d.Count(), "MergeMaps should result in correct number of elements")
		assert.True(t, d.ContainsKey("c"), "MergeMaps should add 'c'")
		assert.True(t, d.ContainsKey("d"), "MergeMaps should add 'd'")
	})

	t.Run("Concurrent access", func(t *testing.T) {
		d := NewConcurrentDictionary[int]()
		var wg sync.WaitGroup
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				d.Set(string(rune(i)), i)
			}(i)
		}
		wg.Wait()
		assert.Equal(t, 1000, d.Count(), "Concurrent Set operations should result in 1000 elements")
	})
}
