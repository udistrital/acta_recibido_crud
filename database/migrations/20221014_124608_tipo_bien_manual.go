package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type TipoBienManual_20221014_124608 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TipoBienManual_20221014_124608{}
	m.Created = "20221014_124608"

	migration.Register("TipoBienManual_20221014_124608", m)
}

// Run the migrations
func (m *TipoBienManual_20221014_124608) Up() {
	file, err := os.ReadFile("../scripts/20221014_124608_tipo_bien_manual_up.sql")

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
func (m *TipoBienManual_20221014_124608) Down() {
	file, err := os.ReadFile("../scripts/20221014_124608_tipo_bien_manual_down.sql")

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
