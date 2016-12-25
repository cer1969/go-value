// Copyright Cristian Echeverría Rabí

// Utilidades para verificar parámetros o argumentos dentro de rangos
package checker

//----------------------------------------------------------------------------------------

// CheckError clase para identificar Errores en verificación de valores
type CheckError struct {
	msg string // Descripción del error o errores
}

func (e *CheckError) Error() string {
	return e.msg
}
