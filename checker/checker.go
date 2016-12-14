// Copyright Cristian Echeverría Rabí

package checker

import (
	"fmt"
)

//----------------------------------------------------------------------------------------

// New retorna Checker, objeto para verificar conjunto de valores
// s: Inicializa Error message que pasará a Checker.
//    Identifica al grupo de verificaciones, que puede ser un método o función
func New(s string) *Checker {
	return &Checker{msg: s}
}

//----------------------------------------------------------------------------------------

type Checker struct {
	msg     string  // Error msg
	n       int     // Error count
	i_name  string  // Item name
	i_value float64 // Item value
}

func (vc *Checker) add(sim string, limit float64) {
	vc.n += 1
	vc.msg += fmt.Sprintf("\n  %s required value %s %f [%f received]", vc.i_name, sim, limit, vc.i_value)
}

// Incorpora un mensaje de error (sin verificar) y aumenta la cuenta
func (vc *Checker) Append(msg string) {
	vc.n += 1
	vc.msg += fmt.Sprintf("\n  %s", msg)
}

func (vc *Checker) Reset(s string) {
	vc.msg = s
	vc.n = 0.0
	vc.i_name = ""
	vc.i_value = 0.0
}

func (vc *Checker) Msg() string {
	return vc.msg
}

func (vc *Checker) Count() int {
	return vc.n
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
	vc.n += 1
	vc.msg += fmt.Sprintf("\n  %s required value in %v [%f received]", vc.i_name, values, vc.i_value)
	return vc
}

func (vc *Checker) Error() error {
	if vc.n == 0 {
		return nil
	}
	msg := vc.msg + fmt.Sprintf("\nTotal errors: %d", vc.n)
	return &CheckError{msg, vc.n}
}
