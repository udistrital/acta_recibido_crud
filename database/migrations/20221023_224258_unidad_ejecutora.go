package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UnidadEjecutora_20221023_224258 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UnidadEjecutora_20221023_224258{}
	m.Created = "20221023_224258"

	migration.Register("UnidadEjecutora_20221023_224258", m)
}

// Run the migrations
func (m *UnidadEjecutora_20221023_224258) Up() {
	file, err := os.ReadFile("../scripts/20221023_224258_unidad_ejecutora_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}

}

// Reverse the migrations
func (m *UnidadEjecutora_20221023_224258) Down() {
	file, err := os.ReadFile("../scripts/20221023_224258_unidad_ejecutora_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
