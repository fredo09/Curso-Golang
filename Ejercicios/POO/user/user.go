package user

import "time"

// Modelado de un objeto de tipo Usuario
type usuario struct {
	id        int
	nombre    string
	edad      int
	fechaAlta time.Time
	estatus   bool
}

// Meotodo para dar de alta a usuarios
func (t *usuario) altaUsuario(id int, nombre string, edad int, fechaAlta time.Time, estatus bool) {
	t.id = id
	t.nombre = nombre
	t.fechaAlta = fechaAlta
	t.estatus = estatus
}
