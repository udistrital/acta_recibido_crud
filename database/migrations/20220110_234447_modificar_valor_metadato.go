package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ModificarValorMetadato_20220110_234447 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModificarValorMetadato_20220110_234447{}
	m.Created = "20220110_234447"

	migration.Register("ModificarValorMetadato_20220110_234447", m)
}

// Run the migrations
func (m *ModificarValorMetadato_20220110_234447) Up() {

	file, err := os.ReadFile("../scripts/20220110_234447_modificar_valor_metadato_up.sql")

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
func (m *ModificarValorMetadato_20220110_234447) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := os.ReadFile("../scripts/20220110_234447_modificar_valor_metadato_down.sql")

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
