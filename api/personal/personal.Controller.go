package personal

import (
	"log"
	// "strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rolando-d3v/server-examen-app-go/db"
)

var DBconnect = db.ConexionBD()

// var per Personal

// ?  CREATE ONE PERSONAL
// ? **************************************************************************************************************
func CreatePersonal(c *fiber.Ctx) error {

	var per Personal

	//Parsea los datos de los campos
	if err := c.BodyParser(&per); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj": "No hay campos disponibles", "error": err.Error(),
		})
	}

	//Valida si hay los campos disponibles
	if DNI := c.FormValue("DNI"); len(DNI) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"msj": "El campo DNI está vacío",
		})
	}

	if grado := c.FormValue("GRADO_ID_I"); len(grado) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"msj": "El campo GRADO está vacío",
		})
	}

	if especialidad := c.FormValue("ESPECIALIDAD_ID_I"); len(especialidad) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"msj": "El campo ESPECIALIADAD está vacío",
		})
	}

	perso := "ID_DNI_C, CIP_C, AP_PATERNO_V, AP_PATERNO_V, AP_MATERNO_V, NOMBRE_V, DOMIC_I, CELULAR_V, FOTO_V, FECHA_NAC_D, GRADO_ID_I, ESPECIALIDAD_ID_I"

	queryCreate := `INSERT INTO personal ( ` + perso + ` ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	create, err := DBconnect.Exec(queryCreate,
		per.ID_DNI_C,
		per.CIP_C,
		per.AP_PATERNO_V,
		per.AP_MATERNO_V,
		per.NOMBRE_V,
		per.DOMIC_I,
		per.CELULAR_V,
		per.FOTO_V,
		per.FECHA_NAC_D,
		per.GRADO_ID_I,
		per.ESPECIALIDAD_ID_I)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj":   "problema en la base de datos",
			"error": err.Error(),
		})
	}

	id, err := create.LastInsertId()

	if err != nil {
		log.Fatal(err.Error())
	}

	if id != 0 {
		return c.Status(201).JSON(fiber.Map{
			"msj": "Error al crear",
		})
	} else {
		return c.Status(201).JSON(fiber.Map{
			"msj": "producto creado", "producto": per, "pro": id,
		})
	}

	// per.ESPECIALIDAD_ID_I = uint(id)

}

// ?  GET ALL PERSONAL
// ? **************************************************************************************************************
func GetAllPersonal(c *fiber.Ctx) error {

	per := Personal{}

	// Paginación de resultados
	nroPage := c.QueryInt("page")

	// //  convierte de string a Int y Si hay un error, el valor de `i` será cero
	// i, err := strconv.Atoi(nroPage)
	// if err != nil {
	// 	return c.JSON(fiber.Map{
	// 		"msj": `Enviar Nro de pagina => /personal/?page=1`,"Error": err.Error(),
	// 	})
	// }

	page := nroPage
	pageSize := 10
	offset := (page - 1) * pageSize

	perso := "ID_DNI_C, CIP_C, AP_PATERNO_V, AP_MATERNO_V, NOMBRE_V, DOMIC_I, CELULAR_V, FOTO_V, FECHA_NAC_D, FECHA_CREATE_D, FECHA_UPDATE_D, GRADO_ID_I, ESPECIALIDAD_ID_I"
	queryGetAll := `SELECT ` + perso + ` FROM personal ORDER BY FECHA_CREATE_D DESC LIMIT ? OFFSET ?;`
	rows, err := DBconnect.Query(queryGetAll, pageSize, offset)
	if err != nil {
		log.Fatal(err)
	}

	// Crear un slice para almacenar los productos
	varPersonal := []Personal{}

	for rows.Next() {
		err := rows.Scan(
			&per.ID_DNI_C,
			&per.CIP_C,
			&per.AP_PATERNO_V,
			&per.AP_MATERNO_V,
			&per.NOMBRE_V,
			&per.DOMIC_I,
			&per.CELULAR_V,
			&per.FOTO_V,
			&per.FECHA_NAC_D,
			&per.FECHA_CREATE_D,
			&per.FECHA_UPDATE_D,
			&per.GRADO_ID_I,
			&per.ESPECIALIDAD_ID_I,
		)
		varPersonal = append(varPersonal, per)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	defer rows.Close()

	return c.JSON(fiber.Map{
		"msj": `succesfully ok`, "personal": varPersonal, "page": nroPage,
	})

}

// ?  GET ONE PERSONAL
// ? **************************************************************************************************************
func GetIdPersonal(c *fiber.Ctx) error {

	idx := c.Params("id")
	var per Personal

	idPer := "ID_DNI_C, CIP_C, AP_PATERNO_V, AP_MATERNO_V, NOMBRE_V, DOMIC_I, CELULAR_V, FOTO_V, FECHA_NAC_D, GRADO_ID_I, ESPECIALIDAD_ID_I"
	getId := `SELECT ` + idPer + ` FROM personal WHERE ID_DNI_C = ?;`

	rowGetId := DBconnect.QueryRow(getId, idx)

	err := rowGetId.Scan(
		&per.ID_DNI_C,
		&per.CIP_C,
		&per.AP_PATERNO_V,
		&per.AP_MATERNO_V,
		&per.NOMBRE_V,
		&per.DOMIC_I,
		&per.CELULAR_V,
		&per.FOTO_V,
		&per.FECHA_NAC_D,
		&per.GRADO_ID_I,
		&per.ESPECIALIDAD_ID_I,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msj":   "Personal no encontrado",
			"error": err.Error(),
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"msj":      "successfully",
			"personal": per,
		})

	}

}

// ? UPDATED ONE PERSONAL
// ? **************************************************************************************************************
func UpdatePersonal(c *fiber.Ctx) error {

	var per Personal

	var idx = c.Params("id")

	//PONER HORA
	currentTime := time.Now()
	dateTime := currentTime.Format("2006-01-02 15:04:05")

	//Parsea los datos de los campos
	if err := c.BodyParser(&per); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj": "No hay campos disponibles", "error": err.Error(),
		})
	}

	update := `UPDATE personal SET NOMBRE_V = ?, FOTO_V = ?, FECHA_DELETE_D = ? WHERE ID_DNI_C = ?;`

	result, err := DBconnect.Exec(update, &per.NOMBRE_V, &per.FOTO_V, dateTime, idx)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj": "Error de base de datos", "Error": err.Error(),
		})
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj": "Error de lectura de datos", "Error": err.Error(),
		})
	}

	if rows == 0 {
		return c.Status(200).JSON(fiber.Map{
			"msj": "No se pudo actualizar", "Error": err.Error(),
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"msj": "Dato updated", "personal": per,
		})
	}

}
