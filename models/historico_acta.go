package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type HistoricoActa struct {
	Id                int           `orm:"column(id);pk;auto"`
	ProveedorId       int           `orm:"column(proveedor_id);null"`
	UbicacionId       int           `orm:"column(ubicacion_id);null"`
	RevisorId         int           `orm:"column(revisor_id)"`
	PersonaAsignadaId int           `orm:"column(persona_asignada_id)"`
	Observaciones     string        `orm:"column(observaciones);null"`
	FechaVistoBueno   time.Time     `orm:"column(fecha_visto_bueno);type(date);null"`
	ActaRecibidoId    *ActaRecibido `orm:"column(acta_recibido_id);rel(fk)"`
	EstadoActaId      *EstadoActa   `orm:"column(estado_acta_id);rel(fk)"`
	Activo            bool          `orm:"column(activo)"`
	FechaCreacion     time.Time     `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion time.Time     `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *HistoricoActa) TableName() string {
	return "historico_acta"
}

func init() {
	orm.RegisterModel(new(HistoricoActa))
}

// AddHistoricoActa insert a new HistoricoActa into database and returns
// last inserted Id on success.
func AddHistoricoActa(m *HistoricoActa) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHistoricoActaById retrieves HistoricoActa by Id. Returns error if
// Id doesn't exist
func GetHistoricoActaById(id int) (v *HistoricoActa, err error) {
	o := orm.NewOrm()
	v = &HistoricoActa{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHistoricoActa retrieves all HistoricoActa matches certain condition. Returns empty list if
// no records exist
func GetAllHistoricoActa(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HistoricoActa)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.Contains(k, "__in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
		} else {
			qs = qs.Filter(k, v)
		}
	}

	count, err = qs.Count()
	if err != nil {
		return
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, 0, errors.New("Error: unused 'order' fields")
		}
	}

	var l []HistoricoActa
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, count, nil
	}
	return nil, 0, err
}

// UpdateHistoricoActa updates HistoricoActa by Id and returns error if
// the record to be updated doesn't exist
func UpdateHistoricoActaById(m *HistoricoActa) (err error) {
	o := orm.NewOrm()
	v := HistoricoActa{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHistoricoActa deletes HistoricoActa by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHistoricoActa(id int) (err error) {
	o := orm.NewOrm()
	v := HistoricoActa{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HistoricoActa{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
