// Copyright Cristian Echeverría Rabí

package checker

import (
	"fmt"
	"strings"
)

//----------------------------------------------------------------------------------------

// New retorna Checker, objeto para verificar conjunto de valores
// s: Inicializa Error message que pasará a Checker.
//    Identifica al grupo de verificaciones, que puede ser un método o función
func New(title string) *Checker {
	return &Checker{msg: fmt.Sprintf("%s: ", title), ok: true}
}

//----------------------------------------------------------------------------------------

type Checker struct {
	msg     string  // Error title
	ok      bool    // True if no errors
	i_name  string  // Item name
	i_value float64 // Item value
}

func (vc *Checker) Reset(title string) {
	vc.msg = fmt.Sprintf("%s: ", title)
	vc.ok = true
	vc.i_name = ""
	vc.i_value = 0.0
}

func (vc *Checker) Msg() string {
	return vc.msg
}

func (vc *Checker) Ok() bool {
	return vc.ok
}

func (vc *Checker) add(sim string, limit float64) {
	vc.msg += fmt.Sprintf("%s (%v) required value %s %v, ", vc.i_name, vc.i_value, sim, limit)
	vc.ok = false
}

// Incorpora un mensaje de error (sin verificar) y aumenta la cuenta
func (vc *Checker) Append(msg string) {
	vc.msg += fmt.Sprintf("%s, ", msg)
	vc.ok = false
}

// Incorpora mensajes de otro error
func (vc *Checker) AppendError(err error) {
	vc.Append(err.Error())
}

// Incorpora mensajes de error como sub proceso
func (vc *Checker) AppendSub(err error) {
	lines := strings.Split(err.Error(), "\n")
	for i := range lines {
		lines[i] = fmt.Sprintf("  %s", lines[i])
	}
	vc.msg = vc.msg + "\n" + strings.Join(lines, "\n")
	vc.ok = false
}

func (vc *Checker) Ck(name string, val float64) *Checker {
	vc.i_name = name
	vc.i_value = val
	return vc
}

func (vc *Checker) Lt(limit float64) *Checker {
	if vc.i_value >= limit {
		vc.add("<", limit)
	}
	return vc
}

func (vc *Checker) Le(limit float64) *Checker {
	if vc.i_value > limit {
		vc.add("<=", limit)
	}
	return vc
}

func (vc *Checker) Gt(limit float64) *Checker {
	if vc.i_value <= limit {
		vc.add(">", limit)
	}
	return vc
}

func (vc *Checker) Ge(limit float64) *Checker {
	if vc.i_value < limit {
		vc.add(">=", limit)
	}
	return vc
}

func (vc *Checker) In(values ...float64) *Checker {
	for _, x := range values {
		if vc.i_value == x {
			return vc
		}
	}
	vc.Append(fmt.Sprintf("%s (%v) required value in %v", vc.i_name, vc.i_value, values))
	return vc
}

func (vc *Checker) Error() error {
	if vc.ok {
		return nil
	}
	return &CheckError{vc.msg}
}
