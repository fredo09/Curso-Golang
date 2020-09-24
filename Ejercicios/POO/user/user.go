package user

import "time"

// Modelado de un objeto de tipo Usuario
type Usuario struct {
	id        int
	nombre    string
	edad      int
	fechaAlta time.Time
	estatus   bool
}

// Meotodo para dar de alta a usuarios
func (this *usuario) altaUsuario(id int, nombre string, edad int, fechaAlta time.Time, estatus bool) {
	this.id = id
	this.nombre = nombre
	this.fechaAlta = fechaAlta
	this.estatus = estatus
}
