package repository

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	m "github.com/krzysztofla/Example.Go.Api/models"
)

var itemsList = []*m.Item{
	&m.Item{
		UUID:        "ffe5c744-3060-4413-9512-b2da5afea26e",
		Name:        "Item 01",
		Price:       1.11,
		Description: "1",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	},
	&m.Item{
		UUID:        "e54bcbc9-5d9b-4ef2-8b29-b954a5765bb3",
		Name:        "Item 02",
		Price:       2.22,
		Description: "2",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	},
	&m.Item{
		UUID:        "c8fcc2e2-4d68-48b1-8c01-30b26f013482",
		Name:        "Item 03",
		Price:       3.33,
		Description: "3",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	},
}

func GetAllItems() []*m.Item {
	return itemsList
}

func GetItemById(id string) *m.Item {
	for _, item := range itemsList {
		if item.UUID == id {
			return item
		}
	}
	return nil
}

func DeleteItemById(id string) bool {
	for index, item := range itemsList {
		if item.UUID == id {
			itemsList = append(itemsList[:index], itemsList[+1:]...)
			return true
		}
	}
	return false
}

func CreateItem(r *http.Request) {
	var item m.Item
	er := json.NewDecoder(r.Body).Decode(&item)
	if er == nil {
		item.UUID = uuid.New().String()
		itemsList = append(itemsList, &item)

	}
}
