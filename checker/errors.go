// Copyright Cristian Echeverría Rabí

// Utilidades para verificar parámetros o argumentos dentro de rangos
package checker

//----------------------------------------------------------------------------------------

// NewCheckError crea CheckError con argumentos s: mensaje n: cuenta
func NewCheckError(s string, n int) *CheckError {
	return &CheckError{s, n}
}

// CheckError clase para identificar Errores en verificación de valores
type CheckError struct {
	s string // Descripción del error o errores
	n int    // Cantidad de errores
}

func (e *CheckError) Error() string {
	return e.s
}

func (e *CheckError) Count() int {
	return e.n
}
