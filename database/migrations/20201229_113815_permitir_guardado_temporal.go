package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type PermitirGuardadoTemporal_20201229_113815 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PermitirGuardadoTemporal_20201229_113815{}
	m.Created = "20201229_113815"

	migration.Register("PermitirGuardadoTemporal_20201229_113815", m)
}

// Run the migrations
func (m *PermitirGuardadoTemporal_20201229_113815) Up() {
	// to make schema update
	file, err := os.ReadFile("../scripts/20201229_113815_permitir_guardado_temporal_up.sql")

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
func (m *PermitirGuardadoTemporal_20201229_113815) Down() {
	// to reverse schema update
	file, err := os.ReadFile("../scripts/20201229_113815_permitir_guardado_temporal_down.sql")

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
