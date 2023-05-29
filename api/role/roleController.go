package role

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rolando-d3v/server-examen-app-go/db"
)

var DBconnect = db.ConexionBD()

func GetAllRole(c *fiber.Ctx) error {
	ro := Role{}

	rol := `ID_ROLE_I, DESC_CORTA_V, DESC_LARGA_V, ESTADO_B`
	queryGetAll := `SELECT ` + rol + ` FROM role;`

	rows, err := DBconnect.Query(queryGetAll)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msj":   "Personal no encontrado",
			"error": err.Error(),
		})
	}

	varRole := []Role{}

	for rows.Next() {
		err := rows.Scan(
			&ro.ID_ROLE_I,
			&ro.DESC_CORTA_V,
			&ro.DESC_LARGA_V,
			&ro.ESTADO_B,
		)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"msj":   "Personal no encontrado",
				"error": err.Error(),
			})
		}

		varRole = append(varRole, ro)

	}

	return c.Status(400).JSON(fiber.Map{
		"msj": "successfully", "role": varRole,
	})

}
