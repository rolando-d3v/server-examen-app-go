package user

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/rolando-d3v/server-examen-app-go/db"
)

var connect = db.ConexionBD()

// ? CREATE ONE USUARIO
// ********************************/
func CreateUser(c *fiber.Ctx) error {

	var us = User{
		PASSWORD_V: "secreto peru",
	}

	if err := c.BodyParser(&us); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj": "No hay campos disponibles", "error": err.Error(),
		})
	}

	queryCreate := `INSERT INTO usuario (PASSWORD_V, ESTADO_B, SECRET_PASS_V, ROL_ID_I, ID_DNI_C)  VALUES(?, ?, ?, ?, ?);`

	create, err := connect.Query(queryCreate, us.PASSWORD_V, us.ESTADO_B, us.SECRET_PASS_V, us.ROL_ID_I, us.ID_DNI_C)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj":   "problema en la base de datos",
			"error": err.Error(),
		})
	}

	log.Println(create)

	return c.Status(201).JSON(fiber.Map{
		"msj": "usuario creado", "usuario": us,
	})

}

// ? GET ALL USUARIOS
// ? **************************************************************************************************************
func GetAllUser(c *fiber.Ctx) error {

	us := User{}

	vUser := "u.ID_USUARIO_I, u.ID_DNI_C, u.ROL_ID_I, u.PASSWORD_V, u.SECRET_PASS_V, u.ESTADO_B, u.FECHA_CREATE_D,"
	vRole := "r.ID_ROLE, r.DESC_CORTA_V, r.DESC_LARGA_V, r.ESTADO_B"

	getQuery := `SELECT ` + vUser + vRole + ` FROM usuario AS u INNER JOIN role AS r ON u.ROL_ID_I = r.ID_ROLE ORDER BY FECHA_CREATE_D DESC;`

	getAll, err := connect.Query(getQuery)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj": "problema en la base de datos", "error": err.Error(),
		})
	}

	defer getAll.Close()

	varUser := []User{}

	for getAll.Next() {
		err := getAll.Scan(
			&us.ID_USUARIO_I,
			&us.ID_DNI_C,
			&us.ROL_ID_I,
			&us.PASSWORD_V,
			&us.SECRET_PASS_V,
			&us.ESTADO_B,
			&us.FECHA_CREATE_D,
			&us.ROLE.ID_ROLE,
			&us.ROLE.DESC_CORTA_V,
			&us.ROLE.DESC_LARGA_V,
			&us.ROLE.ESTADO_B,
		)

		if err != nil {
			log.Fatal(err.Error())
		}

		varUser = append(varUser, us)

	}

	return c.Status(201).JSON(fiber.Map{
		"msj": "Successfully", "usuario": varUser,
	})

}

// ? UPDATED ONE USUARIO
// ? **************************************************************************************************************
func UpdateUser(c *fiber.Ctx) error {

	us := User{}

	if err := c.BodyParser(&us); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj": "Error de Parseo de datos", "Error": err.Error(),
		})
	}

	var idx = c.Params("id")

	putQuery := `UPDATE usuario SET PASSWORD_V = ?, SECRET_PASS_V = ?, ROL_ID_I = ?, ESTADO_B = ? WHERE ID_DNI_C = ?;`

	update, err := connect.Query(putQuery, us.PASSWORD_V, us.SECRET_PASS_V, us.ROL_ID_I, us.ESTADO_B, idx)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj": "Error de base de datos", "Error": err.Error(),
		})
	}

	log.Println(update.NextResultSet())

	return c.Status(200).JSON(fiber.Map{
		"msj": "Dato updated", "personal": us,
	})
}

// ? DELETED ONE USUARIO
// ? **************************************************************************************************************
func DeleteUser(c *fiber.Ctx) error {
	us := User{}

	var idx = c.Params("id")

	// Parse body into struct
	if err := c.BodyParser(&us); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj": "Error de parseo de datos", "Error": err.Error(),
		})
	}

	// Delete Employee from database
	delete, err := connect.Query("DELETE FROM usuario WHERE ID_DNI_C = ?", idx)
	if err != nil {
		return err
	}

	// Print result
	log.Println(delete)

	return c.Status(200).JSON(fiber.Map{
		"msj": "Dato Eliminado", "usuario": idx,
	})
}
