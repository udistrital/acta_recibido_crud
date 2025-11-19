package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddCampoPersonaAsignadaEnTablaActaRecibido_20201229_093708 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddCampoPersonaAsignadaEnTablaActaRecibido_20201229_093708{}
	m.Created = "20201229_093708"

	migration.Register("AddCampoPersonaAsignadaEnTablaActaRecibido_20201229_093708", m)
}

// Run the migrations
func (m *AddCampoPersonaAsignadaEnTablaActaRecibido_20201229_093708) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := os.ReadFile("../scripts/20201229_093708_add_campo_persona_asignada_en_tabla_acta_recibido_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *AddCampoPersonaAsignadaEnTablaActaRecibido_20201229_093708) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := os.ReadFile("../scripts/20201229_093708_add_campo_persona_asignada_en_tabla_acta_recibido_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
