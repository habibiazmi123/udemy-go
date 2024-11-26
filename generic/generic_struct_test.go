package generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
	Third  T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return first
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Muhamad",
		Second: "Habibi",
		Third:  "Azmi",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Muhamad",
		Second: "Habibi",
		Third:  "Azmi",
	}

	assert.Equal(t, "Budi", data.ChangeFirst("Budi"))
	assert.Equal(t, "Hello Eko", data.SayHello("Eko"))
}
