package personal

type Personal struct {
	ID_DNI_C          string `form:"ID_DNI_C" json:"ID_DNI_C"`
	CIP_C             *string
	AP_PATERNO_V      string `form:"AP_PATERNO" json:"AP_PATERNO" comment:"ABRAHAM"`
	AP_MATERNO_V      *string
	NOMBRE_V          string
	PASSWORD_V        string `json:"-"`
	SECRET_PASS_V     *string
	ESTADO_B          *bool
	FECHA_NAC_D       *string
	FOTO_V            *string
	DOMIC_I           *string
	CELULAR_V         *string
	FECHA_CREATE_D    *string `form:"FECHA_CREATE_D" json:"FECHA_CREATE_D,omitempty" comment:""`
	GRADO_ID_I        uint
	ESPECIALIDAD_ID_I uint
}
