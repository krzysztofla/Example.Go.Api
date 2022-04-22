package models

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	UUID        string  `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

func NewItem(name string, price float32) (*Item, error) {
	return &Item{
		UUID:      uuid.NewString(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	}, nil
}

func GetAllItems() []*Item {
	return itemsList
}

var itemsList = []*Item{
	&Item{
		UUID:        uuid.NewString(),
		Name:        "Item 01",
		Price:       1.11,
		Description: "1",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	},
	&Item{
		UUID:        uuid.NewString(),
		Name:        "Item 02",
		Price:       2.22,
		Description: "2",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	},
	&Item{
		UUID:        uuid.NewString(),
		Name:        "Item 03",
		Price:       3.33,
		Description: "3",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	},
}
