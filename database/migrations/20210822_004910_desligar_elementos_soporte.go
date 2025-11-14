package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type DesligarElementosSoporte_20210822_004910 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &DesligarElementosSoporte_20210822_004910{}
	m.Created = "20210822_004910"

	migration.Register("DesligarElementosSoporte_20210822_004910", m)
}

// Run the migrations
func (m *DesligarElementosSoporte_20210822_004910) Up() {
	file, err := os.ReadFile("../scripts/20210822_004910_desligar_elementos_soporte_up.sql")

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
func (m *DesligarElementosSoporte_20210822_004910) Down() {
	file, err := os.ReadFile("../scripts/20210822_004910_desligar_elementos_soporte_down.sql")

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
