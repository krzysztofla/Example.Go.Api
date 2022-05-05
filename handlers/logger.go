package handler

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "Basket-Api", log.LstdFlags)
