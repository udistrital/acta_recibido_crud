package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EliminarLongMax_20220315_133140 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EliminarLongMax_20220315_133140{}
	m.Created = "20220315_133140"

	migration.Register("EliminarLongMax_20220315_133140", m)
}

// Run the migrations
func (m *EliminarLongMax_20220315_133140) Up() {
	file, err := os.ReadFile("../scripts/20220315_133140_eliminar_long_max_up.sql")

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
func (m *EliminarLongMax_20220315_133140) Down() {
	file, err := os.ReadFile("../scripts/20220315_133140_eliminar_long_max_down.sql")

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
