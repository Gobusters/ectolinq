package ectolinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Name    string
	Age     int
	Nested  nestedStruct
	PtrNest *nestedStruct
}

type nestedStruct struct {
	Value string
}

func TestEquals(t *testing.T) {
	t.Run("Equal values", func(t *testing.T) {
		assert.True(t, Equals(1, 1), "Expected 1 to equal 1")
	})

	t.Run("Unequal values", func(t *testing.T) {
		assert.False(t, Equals(1, 2), "Expected 1 to not equal 2")
	})

	t.Run("Complex types", func(t *testing.T) {
		a := testStruct{Name: "Test", Age: 30}
		b := testStruct{Name: "Test", Age: 30}
		assert.True(t, Equals(a, b), "Expected structs to be equal")
	})

	t.Run("Unequal complex types", func(t *testing.T) {
		a := testStruct{Name: "Test", Age: 30}
		b := testStruct{Name: "Test", Age: 31}
		assert.False(t, Equals(a, b), "Expected structs to be unequal")
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		assert.True(t, IsEmpty(""), "Expected empty string to be empty")
	})

	t.Run("Non-empty string", func(t *testing.T) {
		assert.False(t, IsEmpty("test"), "Expected non-empty string to not be empty")
	})

	t.Run("Zero int", func(t *testing.T) {
		assert.True(t, IsEmpty(0), "Expected 0 to be empty")
	})

	t.Run("Non-zero int", func(t *testing.T) {
		assert.False(t, IsEmpty(1), "Expected 1 to not be empty")
	})

	t.Run("Zero float", func(t *testing.T) {
		assert.True(t, IsEmpty(0.0), "Expected 0.0 to be empty")
	})

	t.Run("Non-zero float", func(t *testing.T) {
		assert.False(t, IsEmpty(1.0), "Expected 1.0 to not be empty")
	})

	t.Run("Nil pointer", func(t *testing.T) {
		assert.True(t, IsEmpty[*testStruct](nil), "Expected nil pointer to be empty")
	})

	t.Run("Non-nil pointer", func(t *testing.T) {
		assert.False(t, IsEmpty(&testStruct{}), "Expected non-nil pointer to not be empty")
	})

	t.Run("Empty slice", func(t *testing.T) {
		assert.True(t, IsEmpty([]int{}), "Expected empty slice to be empty")
	})

	t.Run("Non-empty slice", func(t *testing.T) {
		assert.False(t, IsEmpty([]int{1, 2, 3}), "Expected non-empty slice to not be empty")
	})

	t.Run("Empty map", func(t *testing.T) {
		assert.True(t, IsEmpty(map[string]int{}), "Expected empty map to be empty")
	})

	t.Run("Non-empty map", func(t *testing.T) {
		assert.False(t, IsEmpty(map[string]int{"a": 1}), "Expected non-empty map to not be empty")
	})

	t.Run("Nil channel", func(t *testing.T) {
		var ch chan int
		assert.True(t, IsEmpty(ch), "Expected nil channel to be empty")
	})

	t.Run("Non-nil channel with zero capacity", func(t *testing.T) {
		ch := make(chan int)
		assert.True(t, IsEmpty(ch), "Expected zero-capacity channel to be empty")
	})

	t.Run("Non-nil channel with non-zero capacity", func(t *testing.T) {
		ch := make(chan int, 1)
		assert.False(t, IsEmpty(ch), "Expected non-zero capacity channel to not be empty")
	})

	t.Run("Nil interface", func(t *testing.T) {
		var i interface{}
		assert.True(t, IsEmpty(i), "Expected nil interface to be empty")
	})

	t.Run("Empty interface", func(t *testing.T) {
		var i interface{} = nil
		assert.True(t, IsEmpty(i), "Expected empty interface to be empty")
	})

	t.Run("Non-nil interface", func(t *testing.T) {
		var i interface{} = 1
		assert.False(t, IsEmpty(i), "Expected non-nil interface to not be empty")
	})

	t.Run("Uninitialized value", func(t *testing.T) {
		var v int
		assert.True(t, IsEmpty(v), "Expected uninitialized value to be empty")
	})

	t.Run("Nil function", func(t *testing.T) {
		var f func()
		assert.True(t, IsEmpty(f), "Expected nil function to be empty")
	})

	t.Run("Non-nil function", func(t *testing.T) {
		f := func() {}
		assert.False(t, IsEmpty(f), "Expected non-nil function to not be empty")
	})
}

func TestIfEmpty(t *testing.T) {
	t.Run("Empty value", func(t *testing.T) {
		result := IfEmpty("", "default")
		assert.Equal(t, "default", result, "Expected default value")
	})

	t.Run("Non-empty value", func(t *testing.T) {
		result := IfEmpty("test", "default")
		assert.Equal(t, "test", result, "Expected original value")
	})

	t.Run("Zero int", func(t *testing.T) {
		result := IfEmpty(0, 1)
		assert.Equal(t, 1, result, "Expected default value")
	})

	t.Run("Non-zero int", func(t *testing.T) {
		result := IfEmpty(1, 0)
		assert.Equal(t, 1, result, "Expected original value")
	})

	t.Run("Zero float", func(t *testing.T) {
		result := IfEmpty(0.0, 1.0)
		assert.Equal(t, 1.0, result, "Expected default value")
	})

	t.Run("Non-zero float", func(t *testing.T) {
		result := IfEmpty(1.0, 0.0)
		assert.Equal(t, 1.0, result, "Expected original value")
	})
}

func TestGet(t *testing.T) {
	s := testStruct{
		Name: "Test",
		Age:  30,
		Nested: nestedStruct{
			Value: "NestedValue",
		},
		PtrNest: &nestedStruct{
			Value: "PtrNestedValue",
		},
	}

	t.Run("Simple field", func(t *testing.T) {
		result, err := Get(s, "Name")
		require.NoError(t, err)
		assert.Equal(t, "Test", result)
	})

	t.Run("Nested field", func(t *testing.T) {
		result, err := Get(s, "Nested.Value")
		require.NoError(t, err)
		assert.Equal(t, "NestedValue", result)
	})

	t.Run("Pointer nested field", func(t *testing.T) {
		result, err := Get(s, "PtrNest.Value")
		require.NoError(t, err)
		assert.Equal(t, "PtrNestedValue", result)
	})

	t.Run("Non-existent field", func(t *testing.T) {
		_, err := Get(s, "NonExistent")
		assert.Error(t, err)
	})
}

func TestSet(t *testing.T) {
	t.Run("Simple field", func(t *testing.T) {
		s := &testStruct{}
		err := Set(s, "Name", "NewName")
		require.NoError(t, err)
		assert.Equal(t, "NewName", s.Name)
	})

	t.Run("Nested field", func(t *testing.T) {
		s := &testStruct{Nested: nestedStruct{}}
		err := Set(s, "Nested.Value", "NewValue")
		require.NoError(t, err)
		assert.Equal(t, "NewValue", s.Nested.Value)
	})

	t.Run("Non-existent field", func(t *testing.T) {
		s := &testStruct{}
		err := Set(s, "NonExistent", "Value")
		assert.Error(t, err)
	})
}

func TestHasField(t *testing.T) {
	s := testStruct{}

	t.Run("Existing field", func(t *testing.T) {
		assert.True(t, HasField(s, "Name"), "Expected 'Name' field to exist")
	})

	t.Run("Non-existent field", func(t *testing.T) {
		assert.False(t, HasField(s, "NonExistent"), "Expected 'NonExistent' field to not exist")
	})
}

func TestGetFieldNames(t *testing.T) {
	s := testStruct{}

	t.Run("Get all field names", func(t *testing.T) {
		names := GetFieldNames(s)
		expected := []string{"Name", "Age", "Nested", "PtrNest"}
		assert.ElementsMatch(t, expected, names)
	})
}

func TestToMap(t *testing.T) {
	s := testStruct{
		Name: "Test",
		Age:  30,
		Nested: nestedStruct{
			Value: "NestedValue",
		},
	}

	t.Run("Convert struct to map", func(t *testing.T) {
		m, err := ToMap(s)
		require.NoError(t, err)
		assert.Equal(t, "Test", m["Name"])
		assert.Equal(t, 30, m["Age"])
	})
}

func TestFromMap(t *testing.T) {
	t.Run("Create struct from map", func(t *testing.T) {
		m := map[string]interface{}{
			"Name": "Test",
			"Age":  30,
		}
		var s testStruct
		err := FromMap(m, &s)
		require.NoError(t, err)
		assert.Equal(t, "Test", s.Name)
		assert.Equal(t, 30, s.Age)
	})
}
