package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EliminarTablaTipoBien_20210831_031754 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EliminarTablaTipoBien_20210831_031754{}
	m.Created = "20210831_031754"

	migration.Register("EliminarTablaTipoBien_20210831_031754", m)
}

// Run the migrations
func (m *EliminarTablaTipoBien_20210831_031754) Up() {
	file, err := os.ReadFile("../scripts/20210831_031754_eliminar_tabla_tipo_bien_up.sql")

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
func (m *EliminarTablaTipoBien_20210831_031754) Down() {
	file, err := os.ReadFile("../scripts/20210831_031754_eliminar_tabla_tipo_bien_down.sql")

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
