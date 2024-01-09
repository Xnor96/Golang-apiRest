package models

type Users struct {
    ID              int     `json:"id"`
    NumeroEmpleado  string  `json:"numeroEmpleado"`
    Nombre          string  `json:"nombre"`
    Apellido        string  `json:"apellido"`
    Email           string  `json:"email"`
    Estatus         string  `json:"estatus"`
    DescripcionPuesto string `json:"descripcionPuesto"`
    Ciudad          string  `json:"ciudad"`
    NumeroCelular   string  `json:"numeroCelular"`
    Sexo            string  `json:"sexo"`
    Centro          string  `json:"centro"`
    NumeroCentro    string  `json:"numeroCentro"`
    FechaNacimiento string  `json:"fechaNacimiento"`
    FechaAlta       string  `json:"fechaAlta"`
}