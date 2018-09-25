package models

type UsuarioRelacionado struct {
	UsuarioPerteneceId *Usuario `orm:"column(usuario_pertenece_id);rel(fk)"`
	UsuarioId          *Usuario `orm:"column(usuario_id);rel(fk)"`
}
