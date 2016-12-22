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
func New(s string) *Checker {
	return &Checker{title: s}
}

//----------------------------------------------------------------------------------------

type Checker struct {
	title   string   // Error title
	msgs    []string // Error msg without title and final resume
	i_name  string   // Item name
	i_value float64  // Item value
}

func (vc *Checker) add(sim string, limit float64) {
	vc.msgs = append(vc.msgs, fmt.Sprintf("%s (%v) required value %s %v", vc.i_name, vc.i_value, sim, limit))
}

// Incorpora un mensaje de error (sin verificar) y aumenta la cuenta
func (vc *Checker) Append(msg string) {
	vc.msgs = append(vc.msgs, msg)
}

// Incorpora mensajes de otro Checker
func (vc *Checker) AppendError(err error) {
	switch err.(type) {
	case *CheckError:
		ckerr, _ := err.(*CheckError)
		vc.msgs = append(vc.msgs, ckerr.Error())
		return
	default:
		vc.msgs = append(vc.msgs, err.Error())
		return
	}
}

func (vc *Checker) Reset(s string) {
	vc.title = s
	vc.msgs = []string{}
	vc.i_name = ""
	vc.i_value = 0.0
}

//func (vc *Checker) Msg() string {
//	return vc.msg
//}

//func (vc *Checker) Count() int {
//	return vc.n
//}

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
	n := len(vc.msgs)
	if n == 0 {
		return nil
	}
	title := fmt.Sprintf("%s [%d errors]: ", vc.title, n)

	msg := title + strings.Join(vc.msgs, "; ")
	return &CheckError{msg, "", len(vc.msgs)}
}
