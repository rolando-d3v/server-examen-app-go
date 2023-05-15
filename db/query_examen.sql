CREATE TABLE
    personal (
        `ID_DNI_C` char(8) NOT NULL,
        `CIP_C` char(9) DEFAULT NULL,
        `AP_PATERNO_V` varchar(50) DEFAULT NULL,
        `AP_MATERNO_V` varchar(50) DEFAULT NULL,
        `NOMBRE_V` varchar(191) DEFAULT NULL,
        `FECHA_NAC_D` DATE DEFAULT NULL,
        `FOTO_V` varchar(191) DEFAULT NULL,
        `DOMIC_I` varchar(191) DEFAULT NULL,
        `CELULAR_V` varchar(191) DEFAULT NULL,
        `FECHA_CREATE_D` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `GRADO_ID_I` int NOT NULL,
        `ESPECIALIDAD_ID_I` int NOT NULL,
        PRIMARY KEY (`ID_DNI_C`),
        CONSTRAINT `FK_especialidad_id` FOREIGN KEY (`ESPECIALIDAD_ID_I`) REFERENCES `especialidad_arma` (`ID_ESP_I`) ON DELETE RESTRICT ON UPDATE CASCADE,
        CONSTRAINT `FK_grado_id` FOREIGN KEY (`GRADO_ID_I`) REFERENCES `grado` (`ID_GRADO_I`) ON DELETE RESTRICT ON UPDATE CASCADE
    );

    
CREATE TABLE
    usuario (
        ID_USUARIO_I INT NOT NULL AUTO_INCREMENT,
        PASSWORD_V varchar(191) NOT NULL,
        ESTADO_B tinyint(1) DEFAULT "1",
        SECRET_PASS_V varchar(191) NOT NULL,
        FECHA_CREATE_D datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
        ROL_ID_I int DEFAULT NULL,
        ID_DNI_C char(8) NOT NULL UNIQUE,
        PRIMARY KEY (ID_USUARIO_I),
        CONSTRAINT FK_personal_id FOREIGN KEY (ID_DNI_C) REFERENCES personal(ID_DNI_C) ON DELETE CASCADE ON UPDATE CASCADE,
        CONSTRAINT FK_role_id FOREIGN KEY (ROL_ID_I) REFERENCES role(ID_ROLE) ON DELETE CASCADE ON UPDATE CASCADE
    );