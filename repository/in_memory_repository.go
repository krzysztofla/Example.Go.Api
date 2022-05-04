package repository

import (
	"time"

	"github.com/google/uuid"

	m "github.com/krzysztofla/Example.Go.Api/models"
)

var itemsList = []*m.Item{
	&m.Item{
		UUID:        uuid.NewString(),
		Name:        "Item 01",
		Price:       1.11,
		Description: "1",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	},
	&m.Item{
		UUID:        uuid.NewString(),
		Name:        "Item 02",
		Price:       2.22,
		Description: "2",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	},
	&m.Item{
		UUID:        uuid.NewString(),
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
