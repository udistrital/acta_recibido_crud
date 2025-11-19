package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type PermiteGuardadoTemporal_20201229_125940 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PermiteGuardadoTemporal_20201229_125940{}
	m.Created = "20201229_125940"

	migration.Register("PermiteGuardadoTemporal_20201229_125940", m)
}

// Run the migrations
func (m *PermiteGuardadoTemporal_20201229_125940) Up() {
	// to make schema update
	file, err := os.ReadFile("../scripts/20201229_125940_permite_guardado_temporal_up.sql")

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
func (m *PermiteGuardadoTemporal_20201229_125940) Down() {
	// to make schema update
	file, err := os.ReadFile("../scripts/20201229_125940_permite_guardado_temporal_down.sql")

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
