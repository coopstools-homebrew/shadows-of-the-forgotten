package persistance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersonTable_Delete(t *testing.T) {
	table := Connect()
	person := table.Get("quinty6==")
	assert.Equal(t, "quinty6==", person.Id)
	table.Delete("quinty6==")
	person = table.Get("quinty6==")
	assert.Equal(t, "", person.Id)
}

func TestPersonTable_Update_new(t *testing.T) {
	table := Connect()
	allPersons := table.GetAll()
	assert.Equal(t, 2, len(allPersons))
	newPerson := table.Update(Person{Name: "Jacqueline", Favorite: "Brightly colored fish"})
	assert.Equal(t, 8, len(newPerson.Id))

	newPersonRetrieved := table.Get(newPerson.Id)
	assert.Equal(t, "Jacqueline", newPersonRetrieved.Name)
}
