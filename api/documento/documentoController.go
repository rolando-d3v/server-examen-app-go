package documento

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rolando-d3v/server-examen-app-go/db"
)

var DBconnect = db.ConexionBD()

// !  CREATE ONE REGISTRO
// ********************************/

func CreatePersonal(c *fiber.Ctx) error {

	var pro Producto

	// primera forma
	if err := c.BodyParser(&pro); err != nil {
		return c.JSON(fiber.Map{
			"msj": "No hay campos disponibles", "error": err.Error(),
		})

	}

	//  segunda forma
	// err := c.BodyParser(&pro)
	// if err != nil {
	// 	return c.JSON(err.Error())
	// }


	//se puede cambiar los datos del body por ejemplo descripcion si quiero poner un ID
	//pro.DESCRIPCION_V = "peru gana 2024"


	//Valida si hay los campos disponibles
	if descripcion := c.FormValue("DESCRIPCION_V"); len(descripcion) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msj": "El campo descripcion está vacío",
		})
	}

	if precio := c.FormValue("PRECIO_I"); len(precio) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msj": "El campo precio está vacío",
		})
	}


	queryCreate := `INSERT INTO producto (DESCRIPCION_V, PRECIO_I) VALUES (?, ?)`

	result, err := DBconnect.Exec(queryCreate, pro.DESCRIPCION_V, pro.PRECIO_I)
	if err != nil {
		log.Fatal(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}

	pro.ID_PRODUCTO_I = int(id)

	return c.Status(201).JSON(fiber.Map{
		"msj": "producto creado", "producto": pro,
	})

}

// !  GET ALL REGISTRO
// ********************************/
func GetAllPersonal(c *fiber.Ctx) error {

	per := Personal{}

	koler := c.BaseURL()
	log.Println(koler)

	queryGetAll := `SELECT ID_DNI_C, AP_PATERNO_V,AP_MATERNO_V, NOMBRE_V, FECHA_NAC_D, GRADO_I FROM personal`
	rows, err := DBconnect.Query(queryGetAll)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj": "error base de datos", "error": err.Error(),
		})
		// log.Fatal(err.Error()) //para la app
		
	}

	// Crear un slice para almacenar los productos
	varPersonal := []Personal{}

	// Iterar a través de cada fila y agregar un producto al slice
	for rows.Next() {
		err := rows.Scan(&per.ID_DNI, &per.AP_PATERNO, &per.AP_MATERNO, &per.NOMBRE, &per.FECHA_NAC, &per.GRADO)
		varPersonal = append(varPersonal, per)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	defer rows.Close()

	return c.JSON(fiber.Map{
		"msj": "succesfully ok", "personal": varPersonal,
	})
}

// !  GET ONE REGISTRO
// ********************************/

func GetIdPersonal(c *fiber.Ctx) error {

	idx := c.Params("id")
	var per Personal
	// per := Personal{}

	queryGetID := `SELECT ID_DNI_C, AP_PATERNO_V,AP_MATERNO_V, NOMBRE_V, FECHA_NAC_D, GRADO_I FROM personal WHERE ID_DNI_C = ?`

	rowGetId := DBconnect.QueryRow(queryGetID, idx)

	err := rowGetId.Scan(&per.ID_DNI, &per.AP_PATERNO, &per.AP_MATERNO, &per.NOMBRE, &per.FECHA_NAC, &per.GRADO)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msj":   "personal no encontrado encontrado",
			"error": err.Error(),
		})
	} else {
		c.Status(fiber.StatusOK)

		return c.JSON(fiber.Map{
			"msj": "succesfully", "personal": per,
		})
	}

}

// !  DELETE ONE REGISTRO
// ********************************/

func DeleteIdPersonal(c *fiber.Ctx) error {

	idx := c.Params("id")

	var pro Producto

	queryGetID := `SELECT ID_PRODUCTO_I, DESCRIPCION_V, PRECIO_I FROM producto WHERE ID_PRODUCTO_I = ?`

	rowGetId := DBconnect.QueryRow(queryGetID, idx)

	err := rowGetId.Scan(&pro.ID_PRODUCTO_I, &pro.DESCRIPCION_V, &pro.PRECIO_I)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msj":   "prodcuto no encontrado encontrado",
			"error": err.Error(),
		})

	} else {

		delet, err := DBconnect.Prepare("DELETE FROM producto WHERE ID_PRODUCTO_I = ?;")
		if err != nil {
			log.Println(err.Error())
		}

		_, err = delet.Exec(idx)

		if err != nil {
			log.Println(err.Error())
		}

		return c.JSON(fiber.Map{
			"msn":           "producto deleted",
			"producto_id":   idx,
			"producto_name": pro.DESCRIPCION_V,
		})

	}

}
