package personal

type Personal struct {
	ID_DNI_C          string `form:"DNI" json:"DNI"`
	CIP_C             *string
	AP_PATERNO_V      *string `form:"AP_PATERNO" json:"AP_PATERNO" comment:"ABRAHAM"`
	AP_MATERNO_V      *string
	NOMBRE_V          string
	FECHA_NAC_D       *string
	FOTO_V            *string
	DOMIC_I           *string
	CELULAR_V         *string
	FECHA_CREATE_D    *string `form:"FECHA_CREATE_D" json:"FECHA_CREATE_D,omitempty" comment:""`
	FECHA_UPDATE_D    *string `form:"FECHA_UPDATE_D" json:"FECHA_UPDATE_D,omitempty" comment:""`
	GRADO_ID_I        uint
	ESPECIALIDAD_ID_I uint
}
