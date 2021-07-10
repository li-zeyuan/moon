package meta

import (
	"fmt"
	"testing"

	"github.com/fatih/structs"
	"github.com/stretchr/testify/assert"
)

func TestMeta_Init(t *testing.T) {
	userModel := struct {
		Name string `gorm:"name"`
		Age  string `gorm:"age"`
	}{}

	str := GetFieldName(&userModel.Name)
	assert.Equal(t, str, "name")
}

func GetFieldName(i interface{}) string {

	return ""
}

func TestStruct(t *testing.T) {
	// can stop processing them via "omitempty" tag option.
	type Server struct {
		Name     string `structs:",omitempty"`
		ID       int32  `structs:"server_id,omitempty"`
		Location string
	}

	// Only add location
	s := &Server{
		Location: "Tokyo",
	}

	m := structs.Map(s)

	// map contains only the Location field
	fmt.Printf("%v\n", m)
	// Output:
	// map[Location:Tokyo]
}
