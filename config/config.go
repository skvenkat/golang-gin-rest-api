package config

import (
	"log"
)

type GoAppTools struct {
	ErrorLogger log.Logger
	InfoLogger  log.Logger
	Validate    *validate.Validate
}