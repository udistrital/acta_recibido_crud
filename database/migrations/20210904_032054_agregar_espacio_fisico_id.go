package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarEspacioFisicoId_20210904_032054 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarEspacioFisicoId_20210904_032054{}
	m.Created = "20210904_032054"

	migration.Register("AgregarEspacioFisicoId_20210904_032054", m)
}

// Run the migrations
func (m *AgregarEspacioFisicoId_20210904_032054) Up() {
	file, err := os.ReadFile("../scripts/20210904_032054_agregar_espacio_fisico_id_up.sql")

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
func (m *AgregarEspacioFisicoId_20210904_032054) Down() {
	file, err := os.ReadFile("../scripts/20210904_032054_agregar_espacio_fisico_id_down.sql")

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
