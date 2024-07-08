package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("Rama")
	if result != "Hello Rama" {
		t.Fail()
	}
	fmt.Println("Done")
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("Rama")
	if result != "Hello Rama" {
		t.FailNow()
	}
	fmt.Println("Done")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Rama")
	if result != "Hello Rama" {
		t.Error("Harusnya Hello Rama")
	}
	fmt.Println("Done")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Rama")
	if result != "Hello Rama" {
		t.Fatal("Harusnya Hello Rama")
	}
	fmt.Println("Done")
}

// Assert test
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Rama")
	assert.Equal(t, "Hello Rama", result, "Result must be Hello Rama")
	// Akan mengembalikan fail, maka kode di bawahnya bisa dieksekusi meskipun error
	fmt.Println("Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Rama")
	require.Equal(t, "Hello Rama", result, "Result must be Hello Rama")
	// Akan mengembalikan failnow, maka kode di bawahnya tidak dieksekusi meskipun error
	fmt.Println("Done")
}

// Skip test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Tidak bisa berjalan di windows")
	}
	result := HelloWorld("Rama")
	require.Equal(t, "Hello Rama", result)
}

// Before after test
func TestMain(m *testing.M) {
	fmt.Println("Before unit test")
	m.Run() // Digunakan untuk menjalankan semua unit test
	fmt.Println("After unit test")
}

// Subtest
func TestSubTest(t *testing.T) {
	t.Run("Dimas", func(t *testing.T) {
		result := HelloWorld("Dimas")
		require.Equal(t, "Hello Dimas", result, "Result must be Hello Dimas")
	})
	t.Run("Rama", func(t *testing.T) { // nama subtest, func
		result := HelloWorld("Rama")
		require.Equal(t, "Hello Rama", result, "Result must be Hello Rama")
	})
}

// Table test
func TestTableTest(t *testing.T) {
	tests := []struct { // Membuat slice struct literal
		name     string
		request  string
		expected string
	}{
		{
			name:     "Rama",
			request:  "Rama",
			expected: "Hello Rama",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
