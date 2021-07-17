package persistance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersonTable_Delete(t *testing.T) {
	table := Connect()
	person := table.Get("quinty6==")
	assert.NotNil(t, person)
	table.Delete("quinty6==")
	person = table.Get("quinty6==")
	assert.Nil(t, person)
}
