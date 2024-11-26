package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m MyManager) GetName() string {
	return m.Name
}

func (m MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresidents struct {
	Name string
}

func (m *MyVicePresidents) GetName() string {
	return m.Name
}

func (m *MyVicePresidents) GetVicePresidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Azmi", GetName[Manager](&MyManager{Name: "Azmi"}))
	assert.Equal(t, "Azmi", GetName[VicePresident](&MyVicePresidents{Name: "Azmi"}))
}
