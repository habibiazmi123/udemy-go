package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Azmi")
	}
}

func BenchmarkHelloWorldCumi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Cumi")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Before")

	m.Run()

	fmt.Println("After")
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "Cumi",
			request: "Cumi",
			expected: "Hello Cumi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)	
		})
	}
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac")
	}

	result := HelloWorld("Azmi")

	assert.Equal(t, "Hello Azmi", result, "Result must be 'Hello Azmi'")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Azmi")

	assert.Equal(t, "Hello Azmi", result, "Result must be 'Hello Azmi'")
}