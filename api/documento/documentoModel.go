package documento

type Personal struct {
	ID_DNI     string `json:"ID_DNI_C "`
	AP_PATERNO string `json:"AP_PATERNO_V"`
	AP_MATERNO string `json:"AP_MATERNO_V"`
	NOMBRE     string `json:"NOMBRE_V"`
	FECHA_NAC  string `json:"FECHA_NAC_D"`
	GRADO      int    `json:"GRADO_I"`
}

type Producto struct {
	ID_PRODUCTO_I int    `json:"ID_PRODUCTO_I"`
	DESCRIPCION_V string `json:"DESCRIPCION_V"`
	PRECIO_I      int    `json:"PRECIO_I"`
}
