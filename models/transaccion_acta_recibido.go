package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type TransaccionActaRecibido struct {
	ActaRecibido            *ActaRecibido
	UltimoEstado			*HistoricoActa
	SoportesActa			*[]TransaccionSoporteActa
}

// GetTransaccionActaRecibido Transacción para registrar toda la información del Acta de Recibido
func GetTransaccionActaRecibido(id int) (v []interface{}, err error) {
	o := orm.NewOrm()
	var Acta ActaRecibido

	if _, err := o.QueryTable(new(ActaRecibido)).RelatedSel().Filter("Id",id).Filter("Activo",true).All(&Acta); err == nil{

		fmt.Println("Acta :",Acta)
		var UltimoEstado HistoricoActa

		if _, err := o.QueryTable(new(HistoricoActa)).RelatedSel().Filter("ActaRecibidoId__Id",id).Filter("Activo",true).All(&UltimoEstado); err == nil{
			fmt.Println("Historico :" ,UltimoEstado)
			var Soportes []SoporteActa

			if _, err := o.QueryTable(new(SoporteActa)).RelatedSel().Filter("ActaRecibidoId__Id",id).Filter("Activo",true).All(&Soportes); err == nil{
				fmt.Println("Soporte :" ,Soportes)
				var w []interface{}

				for _, Soporte_ := range Soportes{

					var Elementos []Elemento

					if _, err := o.QueryTable(new(Elemento)).RelatedSel().Filter("SoporteActaId__Id",Soporte_.Id).Filter("Activo",true).All(&Elementos); err == nil{
						fmt.Println(Elementos)
						var q []interface{}

						for _, Elemento_ := range Elementos{

							q = append(q,map[string]interface{}{
								"Id": Elemento_.Id,
								"Nombre": Elemento_.Nombre,
								"Cantidad": Elemento_.Cantidad,
								"Marca": Elemento_.Marca,
								"Serie": Elemento_.Serie,
								"UnidadMedida": Elemento_.UnidadMedida,
								"ValorUnitario": Elemento_.ValorUnitario,
								"Subtotal": Elemento_.Subtotal,
								"Descuento": Elemento_.Descuento,
								"ValorTotal": Elemento_.ValorTotal,
								"PorcentajeIvaId": Elemento_.PorcentajeIvaId,
								"ValorIva": Elemento_.ValorIva,
								"ValorFinal": Elemento_.ValorFinal,
								"SubgrupoCatalogoId": Elemento_.SubgrupoCatalogoId,
								"Verificado": Elemento_.Verificado,
								"TipoBienId": Elemento_.TipoBienId,
								"EstadoElementoId": Elemento_.EstadoElementoId,
								"SoporteActaId": Elemento_.SoporteActaId,
								"Activo": Elemento_.Activo,
								"FechaCreacion": Elemento_.FechaCreacion,
								"FechaModificacion": Elemento_.FechaModificacion,
							})
						}
						w = append(w,map[string]interface{}{
							"SoporteActa": Soporte_,
							"Elementos": &q,
						})
					}
				}
				v = append(v,map[string]interface{}{
					"ActaRecibido": Acta,
					"UltimoEstado": UltimoEstado,
					"SoportesActa": &w,
				})
			}
		}
		return v, nil
	}
	return nil, err
}

// AddTransaccionActaRecibido Transacción para registrar toda la información del Acta de Recibido
func AddTransaccionActaRecibido(m *TransaccionActaRecibido) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	if idActa, errTr := o.Insert(m.ActaRecibido); errTr == nil {
		fmt.Println("idActa: ", idActa)
		fmt.Println("m: ", m.UltimoEstado)
		m.UltimoEstado.ActaRecibidoId.Id = int(idActa)
		m.UltimoEstado.Activo = true
		fmt.Println("m: ", m.UltimoEstado)
		if _, errTr := o.Insert(m.UltimoEstado); errTr == nil {

			for _, v := range *m.SoportesActa {
				v.SoporteActa.ActaRecibidoId.Id = int(idActa)

				if IdSoporte, errTr := o.Insert(v.SoporteActa); errTr == nil {

					for _, w := range *v.Elementos {

						w.SoporteActaId.Id = int(IdSoporte)

						if _, errTr = o.Insert(&w); errTr != nil {
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}

					}

				} else {
					err = errTr
					fmt.Println(err)
					_ = o.Rollback()
				}
			}

		}else {
			err = errTr
			fmt.Println(err)
			_ = o.Rollback()
		}
		_ = o.Commit()

	} else {
		err = errTr
		fmt.Println(errTr)
		_ = o.Rollback()
	}
	return
}

// UpdateTransaccionActaRecibido updates ActaRecibido by Id and returns error if
// the record to be updated doesn't exist
func UpdateTransaccionActaRecibido(m *TransaccionActaRecibido) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	v := ActaRecibido{Id: m.ActaRecibido.Id}
	// ascertain id exists in the database
	fmt.Println(v)
	if errTr := o.Read(&v); errTr == nil {

		if _, errTr = o.Update(m.ActaRecibido,"UbicacionId","FechaVistoBueno","RevisorId","Observaciones","FechaModificacion","PersonaAsignada"); errTr == nil {
			fmt.Println("Acta: ",m.ActaRecibido)

			var Historico_ HistoricoActa

			if errTr := o.QueryTable(new(HistoricoActa)).RelatedSel().Filter("Activo", true).Filter("ActaRecibidoId__Id", m.ActaRecibido.Id).One(&Historico_); errTr == nil {

				fmt.Println("Historico :", Historico_)

				Historico_.Activo = false
				if _, errTr = o.Update(&Historico_,"Activo"); errTr == nil{

					m.UltimoEstado.Activo = true

					if _, errTr = o.Insert(m.UltimoEstado); err != nil{
						err = errTr
						fmt.Println(err)
						_ = o.Rollback()
						return
					}
				} else {
					err = errTr
					fmt.Println(err)
					_ = o.Rollback()
					return
				}
			} else {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}

			var Soportes []SoporteActa
			var q []int

			if _, errTr = o.QueryTable(new(SoporteActa)).RelatedSel().Filter("ActaRecibidoId__Id",m.ActaRecibido.Id).All(&Soportes); err == nil{
				for _, Soporte := range Soportes{
					q = append(q,Soporte.Id)
					fmt.Println(q)
				}
				fmt.Println(q)

			} else {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}

				for _, w := range *m.SoportesActa {
					var metadato SoporteActa
					if errTr = o.QueryTable(new(SoporteActa)).RelatedSel().Filter("Id",w.SoporteActa.Id).Filter("ActaRecibidoId__Id",m.ActaRecibido.Id).One(&metadato); err == nil{

						if (metadato.Id != 0) {

							metadato.Consecutivo = w.SoporteActa.Consecutivo
							metadato.ProveedorId = w.SoporteActa.ProveedorId
							metadato.FechaSoporte = w.SoporteActa.FechaSoporte
							metadato.FechaModificacion = w.SoporteActa.FechaModificacion

							if _, errTr = o.Update(&metadato); errTr == nil {

								var Elementos []Elemento
								var g []int

								if _, errTr = o.QueryTable(new(Elemento)).RelatedSel().Filter("SoporteActaId__Id",metadato.Id).All(&Elementos); err == nil{
									for _, Elemento := range Elementos{
										g = append(g,Elemento.Id)
										fmt.Println(g)
									}
									fmt.Println(g)
								} else {
									err = errTr
									fmt.Println(err)
									_ = o.Rollback()
									return
								}


								for _, u := range *w.Elementos {

									var metadato_elemento Elemento

									if errTr = o.QueryTable(new(Elemento)).RelatedSel().Filter("Id",u.Id).Filter("SoporteActaId__Id",metadato.Id).One(&metadato_elemento); err == nil{

										if (metadato_elemento.Id != 0){
											metadato_elemento.Nombre = u.Nombre
											metadato_elemento.Cantidad = u.Cantidad
											metadato_elemento.Marca = u.Marca
											metadato_elemento.Serie = u.Serie
											metadato_elemento.UnidadMedida = u.UnidadMedida
											metadato_elemento.ValorUnitario = u.ValorUnitario
											metadato_elemento.Subtotal = u.Subtotal
											metadato_elemento.Descuento = u.Descuento
											metadato_elemento.ValorTotal = u.ValorTotal
											metadato_elemento.PorcentajeIvaId = u.PorcentajeIvaId
											metadato_elemento.ValorIva = u.ValorIva
											metadato_elemento.ValorFinal = u.ValorFinal
											metadato_elemento.SubgrupoCatalogoId = u.SubgrupoCatalogoId
											metadato_elemento.Verificado = u.Verificado
											metadato_elemento.FechaModificacion = u.FechaModificacion
											if _, errTr = o.Update(&metadato_elemento); errTr == nil {

												g = RemoveSoporte(g,metadato_elemento.Id)

											} else {
												err = errTr
												fmt.Println(err)
												_ = o.Rollback()
												return
											}
										} else {
											u.SoporteActaId.Id = metadato.Id
											if _, errTr = o.Insert(&u); errTr != nil {
												err = errTr
												fmt.Println(err)
												_ = o.Rollback()
												return
											}
										}
									} else {
										err = errTr
										fmt.Println(err)
										_ = o.Rollback()
										return
									}
								}

								q = RemoveSoporte(q, metadato.Id)

								if (g != nil){

									var _Elemento Elemento
									for _, Elemento := range g {
										_Elemento.Id = Elemento
										_Elemento.Activo = false
										if _, errTr = o.Update(&_Elemento,"Activo"); err != nil{
											err = errTr
											fmt.Println(err)
											_ = o.Rollback()
											return
										}
									}
									fmt.Println(g)
								}

							} else {
								err = errTr
								fmt.Println(err)
								_ = o.Rollback()
								return
							}
						} else {
							w.SoporteActa.ActaRecibidoId.Id = m.ActaRecibido.Id
							if _, errTr = o.Insert(w.SoporteActa); errTr == nil {
								for _, u := range *w.Elementos {

									u.SoporteActaId.Id = w.SoporteActa.Id
									if _, errTr = o.Insert(&u); errTr != nil {
										err = errTr
										fmt.Println(err)
										_ = o.Rollback()
										return
									}
								}
							} else {
								err = errTr
								fmt.Println(err)
								_ = o.Rollback()
								return
							}
						}
					} else {
						err = errTr
						fmt.Println(err)
						_ = o.Rollback()
						return
					}		
				}
			
				if (q != nil){
					fmt.Println(q)
					var _Soporte SoporteActa
					for _, Soporte := range q {
						_Soporte.Id = Soporte
						_Soporte.Activo = false
						if _, errTr = o.Update(&_Soporte,"Activo"); err != nil{
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}

					}

				}
			
			_ = o.Commit()
		}	else {
			err = errTr
			fmt.Println(err)
			_ = o.Rollback()
			return
		}
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}
	return
}

func RemoveSoporte(m []int, i int )(v []int){

	var Soportes []int

	for _, Soporte := range m {
		if (Soporte != i){
			Soportes = append(Soportes,Soporte)
			fmt.Print(Soporte)
		}
	} 
	
	return Soportes
}

