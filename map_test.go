package ectolinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	t.Run("non-empty map", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		keys := Keys(m)
		assert.ElementsMatch(t, []string{"a", "b", "c"}, keys)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		keys := Keys(m)
		assert.Empty(t, keys)
	})
}

func TestValues(t *testing.T) {
	t.Run("non-empty map", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		values := Values(m)
		assert.ElementsMatch(t, []int{1, 2, 3}, values)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		values := Values(m)
		assert.Empty(t, values)
	})
}

func TestEntries(t *testing.T) {
	t.Run("non-empty map", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		entries := Entries(m)
		assert.Len(t, entries, 2)
		assert.Contains(t, entries, MapEntry[string, int]{Key: "a", Value: 1})
		assert.Contains(t, entries, MapEntry[string, int]{Key: "b", Value: 2})
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		entries := Entries(m)
		assert.Empty(t, entries)
	})
}

func TestMapFromEntries(t *testing.T) {
	t.Run("non-empty entries", func(t *testing.T) {
		entries := []MapEntry[string, int]{
			{Key: "a", Value: 1},
			{Key: "b", Value: 2},
		}
		m := MapFromEntries(entries)
		assert.Equal(t, map[string]int{"a": 1, "b": 2}, m)
	})

	t.Run("empty entries", func(t *testing.T) {
		entries := []MapEntry[string, int]{}
		m := MapFromEntries(entries)
		assert.Empty(t, m)
	})
}

func TestFilterMap(t *testing.T) {
	t.Run("filter even values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		filtered := FilterMap(m, func(key string, value int) bool {
			return value%2 == 0
		})
		assert.Equal(t, map[string]int{"b": 2, "d": 4}, filtered)
	})

	t.Run("filter all", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		filtered := FilterMap(m, func(key string, value int) bool {
			return false
		})
		assert.Empty(t, filtered)
	})
}

func TestMapValues(t *testing.T) {
	t.Run("double values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		doubled := MapValues(m, func(v int) int {
			return v * 2
		})
		assert.Equal(t, map[string]int{"a": 2, "b": 4, "c": 6}, doubled)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		result := MapValues(m, func(v int) int {
			return v
		})
		assert.Empty(t, result)
	})
}

func TestMapKeys(t *testing.T) {
	t.Run("uppercase keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		upper := MapKeys(m, func(k string) string {
			return string(k[0] - 32)
		})
		assert.Equal(t, map[string]int{"A": 1, "B": 2}, upper)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		result := MapKeys(m, func(k string) string {
			return k
		})
		assert.Empty(t, result)
	})
}

func TestMapEntries(t *testing.T) {
	t.Run("modify keys and values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		result := MapEntries(m, func(key string, value int) (string, int) {
			return string(key[0] - 32), value * 2
		})
		assert.Equal(t, map[string]int{"A": 2, "B": 4}, result)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		result := MapEntries(m, func(key string, value int) (string, int) {
			return key, value
		})
		assert.Empty(t, result)
	})
}

func TestFindValue(t *testing.T) {
	t.Run("find existing value", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		value, found := FindValue(m, func(key string, value int) bool {
			return value > 2
		})
		assert.True(t, found)
		assert.Equal(t, 3, value)
	})

	t.Run("value not found", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		value, found := FindValue(m, func(key string, value int) bool {
			return value > 5
		})
		assert.False(t, found)
		assert.Equal(t, 0, value)
	})
}

func TestFindKey(t *testing.T) {
	t.Run("find existing key", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		key, found := FindKey(m, func(key string, value int) bool {
			return key > "b"
		})
		assert.True(t, found)
		assert.Equal(t, "c", key)
	})

	t.Run("key not found", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		key, found := FindKey(m, func(key string, value int) bool {
			return key > "c"
		})
		assert.False(t, found)
		assert.Equal(t, "", key)
	})
}

func TestReduceEntries(t *testing.T) {
	t.Run("sum of values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		sum := ReduceEntries(m, 0, func(acc int, key string, value int) int {
			return acc + value
		})
		assert.Equal(t, 6, sum)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		result := ReduceEntries(m, 10, func(acc int, key string, value int) int {
			return acc + value
		})
		assert.Equal(t, 10, result)
	})
}

func TestForEachEntry(t *testing.T) {
	t.Run("modify map", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		result := make(map[string]int)
		ForEachEntry(m, func(key string, value int) {
			result[key] = value * 2
		})
		assert.Equal(t, map[string]int{"a": 2, "b": 4, "c": 6}, result)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		count := 0
		ForEachEntry(m, func(key string, value int) {
			count++
		})
		assert.Equal(t, 0, count)
	})
}

func TestMerge(t *testing.T) {
	t.Run("merge two maps", func(t *testing.T) {
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{"b": 3, "c": 4}
		merged := Merge(m1, m2)
		assert.Equal(t, map[string]int{"a": 1, "b": 3, "c": 4}, merged)
	})

	t.Run("merge multiple maps", func(t *testing.T) {
		m1 := map[string]int{"a": 1}
		m2 := map[string]int{"b": 2}
		m3 := map[string]int{"c": 3}
		merged := Merge(m1, m2, m3)
		assert.Equal(t, map[string]int{"a": 1, "b": 2, "c": 3}, merged)
	})

	t.Run("merge with empty map", func(t *testing.T) {
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{}
		merged := Merge(m1, m2)
		assert.Equal(t, map[string]int{"a": 1, "b": 2}, merged)
	})
}

func TestCountEntries(t *testing.T) {
	t.Run("count even values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		count := CountEntries(m, func(key string, value int) bool {
			return value%2 == 0
		})
		assert.Equal(t, 2, count)
	})

	t.Run("count all", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		count := CountEntries(m, func(key string, value int) bool {
			return true
		})
		assert.Equal(t, 2, count)
	})

	t.Run("count none", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		count := CountEntries(m, func(key string, value int) bool {
			return false
		})
		assert.Equal(t, 0, count)
	})
}

func TestGroupBy(t *testing.T) {
	t.Run("group by value parity", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		grouped := GroupBy(m, func(k string, v int) string {
			if v%2 == 0 {
				return "even"
			}
			return "odd"
		})
		assert.Len(t, grouped, 2)
		assert.Len(t, grouped["even"], 2)
		assert.Len(t, grouped["odd"], 2)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		grouped := GroupBy(m, func(k string, v int) string {
			return "group"
		})
		assert.Empty(t, grouped)
	})
}

func TestInvert(t *testing.T) {
	t.Run("invert map", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		inverted := Invert(m)
		assert.Equal(t, map[int]string{1: "a", 2: "b", 3: "c"}, inverted)
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		inverted := Invert(m)
		assert.Empty(t, inverted)
	})
}

func TestPick(t *testing.T) {
	t.Run("pick existing keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		picked := Pick(m, "a", "c")
		assert.Equal(t, map[string]int{"a": 1, "c": 3}, picked)
	})

	t.Run("pick non-existing keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		picked := Pick(m, "c", "d")
		assert.Empty(t, picked)
	})
}

func TestOmit(t *testing.T) {
	t.Run("omit existing keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		omitted := Omit(m, "a", "c")
		assert.Equal(t, map[string]int{"b": 2, "d": 4}, omitted)
	})

	t.Run("omit non-existing keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		omitted := Omit(m, "c", "d")
		assert.Equal(t, map[string]int{"a": 1, "b": 2}, omitted)
	})
}

func TestContainsValue(t *testing.T) {
	t.Run("contains element", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		assert.True(t, ContainsValue(slice, 3))
	})

	t.Run("does not contain element", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		assert.False(t, ContainsValue(slice, 6))
	})

	t.Run("empty slice", func(t *testing.T) {
		var slice []int
		assert.False(t, ContainsValue(slice, 1))
	})
}
