package user

import "github.com/rolando-d3v/server-examen-app-go/api/role"

type User struct {
	ID_USUARIO_I   uint64
	ID_DNI_C       string
	PASSWORD_V     string `json:"-"`
	ESTADO_B       bool
	SECRET_PASS_V  string
	FECHA_CREATE_D string
	ROL_ID_I       int
	ROLE           role.Role
}
