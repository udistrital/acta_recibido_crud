package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearEsquemaActaRecibido_20190723_112831 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearEsquemaActaRecibido_20190723_112831{}
	m.Created = "20190723_112831"

	migration.Register("CrearEsquemaActaRecibido_20190723_112831", m)
}

// Run the migrations
func (m *CrearEsquemaActaRecibido_20190723_112831) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := os.ReadFile("../scripts/20190723_112831_crear_esquema_acta_recibido_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *CrearEsquemaActaRecibido_20190723_112831) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := os.ReadFile("../scripts/20190723_112831_crear_esquema_acta_recibido_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
