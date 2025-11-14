package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarEspacioFisico_20210317_195101 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarEspacioFisico_20210317_195101{}
	m.Created = "20210317_195101"

	migration.Register("AgregarEspacioFisico_20210317_195101", m)
}

// Run the migrations
func (m *AgregarEspacioFisico_20210317_195101) Up() {
	// to make schema update
	file, err := os.ReadFile("../scripts/20210317_195101_agregar_espacio_fisico_up.sql")

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
func (m *AgregarEspacioFisico_20210317_195101) Down() {
	// to make schema update
	file, err := os.ReadFile("../scripts/20210317_195101_agregar_espacio_fisico_down.sql")

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
