package persistance

import "math/rand"

var chars = "acdefghijkmnpqrtuvwxyz234789ABCDEFGHJKLMNPQRSTUVWXYZ"

type Person struct {
	Id 			string 	`json:"id"`
	Name 		string 	`json:"name"`
	Age 		int8	`json:"age"`
	Favorite 	string 	`json:"favorite"`
}

type PersonTable struct {
	Persons	map[string]Person
}

func Connect() *PersonTable {
	table := PersonTable{make(map[string]Person, 2)}
	person1 := Person{Id: "quinty6==", Name: "Douglas", Age: 19, Favorite: "Fresh air"}
	person2 := Person{Id: "dublious=", Name: "Manerva", Age: 24, Favorite: "Comfortable shoes"}
	table.Update(person1)
	table.Update(person2)
	return &table
}

func (table *PersonTable) Update(p Person) Person {
	if p.Id == "" {
		newId := ""
		for i := 0; i < 8; i++ {
			newId += string(chars[rand.Intn(len(chars))])
		}
		p.Id = newId
	}
	table.Persons[p.Id] = p
	return p
}

func (table PersonTable) Get(id string) Person {
	if person, ok := table.Persons[id]; ok {
		return person
	}
	return Person{}
}

func (table PersonTable) GetAll() []Person {
	persons := make([]Person, 0, len(table.Persons))
	for _, person := range table.Persons {
		persons = append(persons, person)
	}
	return persons
}

func (table *PersonTable) Delete(id string) {
	if _, ok := table.Persons[id]; ok {
		delete(table.Persons, id)
	}
}
