package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type PermitirSoportesNulos_20201222_231049 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PermitirSoportesNulos_20201222_231049{}
	m.Created = "20201222_231049"

	migration.Register("PermitirSoportesNulos_20201222_231049", m)
}

// Run the migrations
func (m *PermitirSoportesNulos_20201222_231049) Up() {
	// to make schema update
	file, err := os.ReadFile("../scripts/20201222_231049_permitir_soportes_nulos_up.sql")

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
func (m *PermitirSoportesNulos_20201222_231049) Down() {
	// to reverse schema update	file, err := os.ReadFile("../scripts/20201222_231049_permitir_soportes_nulos_down.sql")
	file, err := os.ReadFile("../scripts/20201222_231049_permitir_soportes_nulos_down.sql")

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
