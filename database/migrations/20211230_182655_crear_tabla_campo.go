package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaCampo_20211230_182655 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaCampo_20211230_182655{}
	m.Created = "20211230_182655"

	migration.Register("CrearTablaCampo_20211230_182655", m)
}

// Run the migrations
func (m *CrearTablaCampo_20211230_182655) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := os.ReadFile("../scripts/20211228_104027_crear_tabla_campo_up.sql")

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
func (m *CrearTablaCampo_20211230_182655) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := os.ReadFile("../scripts/20211228_104027_crear_tabla_campo_down.sql")

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
