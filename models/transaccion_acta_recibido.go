package models

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type TransaccionActaRecibido struct {
	ActaRecibido *ActaRecibido
	UltimoEstado *HistoricoActa
	SoportesActa []*SoporteActa
	Elementos    []*Elemento
}

func GetTransaccionActaRecibido(id int, getElementos bool) (v TransaccionActaRecibido, err error) {
	o := orm.NewOrm()

	var acta ActaRecibido
	err = o.QueryTable(new(ActaRecibido)).RelatedSel().Filter("Id", id).One(&acta)
	if err != nil {
		return
	}

	v.ActaRecibido = &acta
	var ultimoEstado HistoricoActa
	_, err = o.QueryTable(new(HistoricoActa)).RelatedSel().Filter("ActaRecibidoId__Id", id).Filter("Activo", true).All(&ultimoEstado)
	if err != nil {
		return
	}

	v.UltimoEstado = &ultimoEstado
	var soportes []*SoporteActa

	_, err = o.QueryTable(new(SoporteActa)).RelatedSel().Filter("ActaRecibidoId__Id", id).Filter("Activo", true).All(&soportes)
	if err != nil {
		return
	}

	v.SoportesActa = soportes
	if getElementos {
		var elementos []*Elemento
		_, err = o.QueryTable(new(Elemento)).RelatedSel().Filter("ActaRecibidoId__Id", id).Filter("Activo", true).All(&elementos)
		if err != nil {
			return
		}

		v.Elementos = elementos
	}

	return
}

func AddTransaccionActaRecibido(m *TransaccionActaRecibido) (err error) {

	o := orm.NewOrm()

	defer func() {
		r := recover()
		if r != nil {
			err = r.(error)
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()

	err = o.Begin()
	if err != nil {
		return
	}

	if m.ActaRecibido == nil || m.ActaRecibido.UnidadEjecutoraId <= 0 ||
		m.ActaRecibido.TipoActaId == nil || m.ActaRecibido.TipoActaId.Id <= 0 ||
		m.UltimoEstado == nil || m.UltimoEstado.EstadoActaId.Id <= 0 {
		return
	}

	now := time.Now()

	m.ActaRecibido.FechaCreacion = now
	m.ActaRecibido.FechaModificacion = now
	idActa, err := o.Insert(m.ActaRecibido)
	if err != nil {
		return err
	}

	m.UltimoEstado.ActaRecibidoId = &ActaRecibido{Id: int(idActa)}
	m.UltimoEstado.Activo = true
	m.UltimoEstado.FechaCreacion = now
	m.UltimoEstado.FechaModificacion = now
	_, err = o.Insert(m.UltimoEstado)
	if err != nil {
		return
	}

	for _, v := range m.SoportesActa {
		v.ActaRecibidoId = &ActaRecibido{Id: int(idActa)}
		v.FechaCreacion = now
		v.FechaModificacion = now
		_, err = o.Insert(v)
		if err != nil {
			return
		}
	}

	for _, w := range m.Elementos {
		w.ActaRecibidoId = &ActaRecibido{Id: int(idActa)}
		w.FechaCreacion = now
		w.FechaModificacion = now
		_, err = o.Insert(w)
		if err != nil {
			return
		}
	}

	return
}

func UpdateTransaccionActaRecibido(m *TransaccionActaRecibido) (err error) {
	o := orm.NewOrm()

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()

	err = o.Begin()
	if err != nil {
		return
	}

	v := ActaRecibido{Id: m.ActaRecibido.Id}
	err = o.Read(&v)
	if err != nil {
		return
	}

	_, err = o.Update(m.ActaRecibido, "UnidadEjecutoraId")
	if err != nil {
		return
	}

	var Historico_ HistoricoActa
	err = o.QueryTable(new(HistoricoActa)).RelatedSel().Filter("Activo", true).Filter("ActaRecibidoId__Id", m.ActaRecibido.Id).One(&Historico_)
	if err != nil {
		return
	}

	Historico_.Activo = false
	_, err = o.Update(&Historico_, "Activo")
	if err != nil {
		return
	}

	now := time.Now()

	m.UltimoEstado.ActaRecibidoId = &ActaRecibido{Id: m.ActaRecibido.Id}
	m.UltimoEstado.Activo = true
	m.UltimoEstado.FechaCreacion = now
	m.UltimoEstado.FechaModificacion = now
	_, err = o.Insert(m.UltimoEstado)
	if err != nil {
		return
	}

	var listSoportes orm.ParamsList
	_, err = o.QueryTable(new(SoporteActa)).RelatedSel().Filter("ActaRecibidoId__Id", m.ActaRecibido.Id).Filter("Activo", true).ValuesFlat(&listSoportes, "Id")
	if err != nil {
		return
	}

	for _, soporte := range m.SoportesActa {
		if soporte.Id > 0 {
			if i := findIdInArray(listSoportes, soporte.Id); i > -1 {
				listSoportes = append(listSoportes[:i], listSoportes[i+1:]...)
				existing := SoporteActa{Id: soporte.Id}
				err = o.Read(&existing)
				if err != nil {
					return
				}
				soporte.FechaCreacion = existing.FechaCreacion
				soporte.FechaModificacion = now
				_, err = o.Update(soporte)
				if err != nil {
					return
				}
			}
		} else {
			soporte.ActaRecibidoId = &ActaRecibido{Id: m.ActaRecibido.Id}
			soporte.FechaCreacion = now
			soporte.FechaModificacion = now
			_, err = o.Insert(soporte)
			if err != nil {
				return
			}
		}
	}

	for _, id := range listSoportes {
		id_, _ := id.(int64)
		soporteInactivo := SoporteActa{Id: int(id_), Activo: false}
		_, err = o.Update(&soporteInactivo, "Activo")
		if err != nil {
			return
		}
	}

	var listElementos orm.ParamsList
	_, err = o.QueryTable(new(Elemento)).RelatedSel().Filter("ActaRecibidoId__Id", m.ActaRecibido.Id).Filter("Activo", true).ValuesFlat(&listElementos, "Id")
	if err != nil {
		return
	}

	for _, elemento := range m.Elementos {
		if elemento.Id > 0 {
			if i := findIdInArray(listElementos, elemento.Id); i > -1 {
				listElementos = append(listElementos[:i], listElementos[i+1:]...)
				existing := Elemento{Id: elemento.Id}
				err = o.Read(&existing)
				if err != nil {
					return
				}
				elemento.FechaCreacion = existing.FechaCreacion
				elemento.FechaModificacion = now
				_, err = o.Update(elemento)
				if err != nil {
					return
				}
			}
		} else {
			elemento.ActaRecibidoId = &ActaRecibido{Id: m.ActaRecibido.Id}
			elemento.FechaCreacion = now
			elemento.FechaModificacion = now
			_, err = o.Insert(elemento)
			if err != nil {
				return
			}
		}
	}

	for _, id := range listElementos {
		id_, _ := id.(int64)
		elementoInactivo := Elemento{Id: int(id_), Activo: false}
		_, err = o.Update(&elementoInactivo, "Activo")
		if err != nil {
			return
		}
	}

	return
}

func findIdInArray(idsList orm.ParamsList, id int) (i int) {
	for i, id_ := range idsList {
		if id_ == int64(id) {
			return i
		}
	}
	return -1
}
