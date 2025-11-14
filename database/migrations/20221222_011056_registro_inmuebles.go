package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type RegistroInmuebles_20221222_011056 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RegistroInmuebles_20221222_011056{}
	m.Created = "20221222_011056"

	migration.Register("RegistroInmuebles_20221222_011056", m)
}

// Run the migrations
func (m *RegistroInmuebles_20221222_011056) Up() {
	file, err := os.ReadFile("../scripts/20221222_011056_registro_inmuebles_up.sql")

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
func (m *RegistroInmuebles_20221222_011056) Down() {
	file, err := os.ReadFile("../scripts/20221222_011056_registro_inmuebles_down.sql")

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
