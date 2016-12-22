// Copyright Cristian Echeverría Rabí

// Utilidades para verificar parámetros o argumentos dentro de rangos
package checker

//----------------------------------------------------------------------------------------

// NewCheckError crea CheckError con argumentos s: mensaje n: cuenta
func NewCheckError(errmsg string, msg string, n int) *CheckError {
	return &CheckError{errmsg, msg, n}
}

// CheckError clase para identificar Errores en verificación de valores
type CheckError struct {
	errmsg string // Descripción del error o errores
	msg    string // Mensaje(s) simplificado(s)
	n      int    // Cantidad de errores
}

func (e *CheckError) Error() string {
	return e.errmsg
}

func (e *CheckError) Msg() string {
	return e.msg
}

func (e *CheckError) Count() int {
	return e.n
}
