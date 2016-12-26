// Copyright Cristian Echeverría Rabí

package checker

import (
	"fmt"
	//"strings"
	"testing"
)

func TestNewAndReset(t *testing.T) {
	vc := New("Prueba")
	if vc.msg != "Prueba: " {
		t.Errorf("Wrong msg: %v", vc.msg)
	}
	if !vc.ok {
		t.Errorf("Wrong ok: %v", vc.ok)
	}

	vc.Ck("Número", 30.0).Lt(30.0)
	err := vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
	if vc.ok {
		t.Errorf("Wrong ok: %v", vc.ok)
	}

	vc.Reset("")
	if vc.msg != ": " {
		t.Errorf("Wrong msg: %v", vc.msg)
	}
	if !vc.ok {
		t.Errorf("Wrong ok: %v", vc.ok)
	}
}

func TestLt(t *testing.T) {
	vc := New("Prueba Lt 1")
	vc.Ck("Número", 30.0).Lt(31.0)
	err := vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}

	vc.Reset("Prueba Lt 2")
	vc.Ck("Número", 30.0).Lt(30.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}

	vc.Reset("Prueba Lt 3")
	vc.Ck("Número", 30.0).Lt(29.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
}

func TestLe(t *testing.T) {
	vc := New("Prueba Le 1")
	vc.Ck("Número", 30.0).Le(31.0)
	err := vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}

	vc.Reset("Prueba Le 2")
	vc.Ck("Número", 30.0).Le(30.0)
	err = vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}

	vc.Reset("Prueba Le 3")
	vc.Ck("Número", 30.0).Le(29.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
}

func TestGt(t *testing.T) {
	vc := New("Prueba Gt 1")
	vc.Ck("Número", 30.0).Gt(29.0)
	err := vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}

	vc.Reset("Prueba Gt 2")
	vc.Ck("Número", 30.0).Gt(30.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}

	vc.Reset("Prueba Gt 3")
	vc.Ck("Número", 30.0).Gt(31.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
}

func TestGe(t *testing.T) {
	vc := New("Prueba Ge 1")
	vc.Ck("Número", 30.0).Ge(29.0)
	err := vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}

	vc.Reset("Prueba Ge 2")
	vc.Ck("Número", 30.0).Ge(30.0)
	err = vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}

	vc.Reset("Prueba Ge 3")
	vc.Ck("Número", 30.0).Ge(31.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
}

func TestMixed(t *testing.T) {
	vc := New("Prueba Ge Le 1")
	vc.Ck("Número", 35.0).Ge(30.0).Le(40.0)
	err := vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}

	vc.Reset("Prueba Ge Le 2")
	vc.Ck("Número", 30.0).Ge(30.0).Le(30.0)
	err = vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}

	vc.Reset("Prueba Ge Le 3")
	vc.Ck("Número", 30.0).Ge(31.0).Le(29.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}

	vc.Reset("Prueba Gt Lt 1")
	vc.Ck("Número", 35.0).Gt(30.0).Lt(40.0)
	err = vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}

	vc.Reset("Prueba Gt Lt 2")
	vc.Ck("Número", 30.0).Gt(30.0).Lt(30.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}

	vc.Reset("Prueba Gt Lt 3")
	vc.Ck("Número", 30.0).Gt(31.0).Lt(29.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
}

func TestIn(t *testing.T) {
	vc := New("Prueba In 1")
	vc.Ck("Número", 4).In(1, 4, 5)
	err := vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}

	vc.Reset("Prueba In 2")
	vc.Ck("Número", 3.0).In(1.0, 2.0, 4.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
}

func TestAppend(t *testing.T) {
	vc := New("Prueba Append")
	vc.Ck("Número", 30.0).Gt(29.0)
	err := vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}

	vc.Append("Este es un error libre")
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}

	vc.Reset("Prueba Append empty string")
	vc.Append("")
	err = vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
}

func TestAppendSub(t *testing.T) {
	vc := New("Prueba Append Sub")

	vcs := New("Sub Error")
	vcs.Ck("Número", 30.0).Gt(29.0)

	vc.AppendSub(vcs.Error())
	err := vc.Error()
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
}

func ExampleChecker() {
	vc := New("Prueba")
	vc.Ck("Número", 30).Gt(40.0).Gt(100.0)
	vc.Ck("Sol", 0.0).Gt(0.0).Lt(1.0)
	vc.Ck("Sol", 0.5).Gt(0.0).Lt(1.0)
	err := vc.Error()
	fmt.Printf(err.Error())
	// Output:
	// Prueba: Número (30) required value > 40, Número (30) required value > 100, Sol (0) required value > 0,
}

func ExampleIn() {
	vc := New("Prueba")
	vc.Ck("Número", 3.0).In(1.0, 2.0, 4.0)
	fmt.Printf("%v", vc.Error())
	// Output:
	// Prueba: Número (3) required value in [1 2 4],
}

func ExampleAppend() {
	vc := New("Prueba")
	vc.Ck("Número", 3.0).In(1.0, 2.0, 4.0)
	vc.Append("Este es un error libre")
	fmt.Printf("%v", vc.Error())
	// Output:
	// Prueba: Número (3) required value in [1 2 4], Este es un error libre,
}

func ExampleAppendError() {
	vc1 := New("Prueba")
	vc1.Ck("Número", 3.0).In(1.0, 2.0, 4.0)
	vc1.Append("Este es un error libre")

	vc2 := New("Super prueba")
	vc2.Ck("Sol", 0).Gt(0).Lt(1)
	vc2.AppendError(vc1.Error())

	fmt.Printf("%v", vc2.Error())
	// Output:
	// Super prueba: Sol (0) required value > 0, Prueba: Número (3) required value in [1 2 4], Este es un error libre, ,
}

func ExampleAppendSub() {
	vc1 := New("Prueba")
	vc1.Ck("Número", 3.0).In(1.0, 2.0, 4.0)
	vc1.Append("Este es un error libre")

	vc2 := New("Super prueba")
	vc2.Ck("Sol", 0).Gt(0).Lt(1)
	vc2.AppendSub(vc1.Error())

	fmt.Printf("%q", vc2.Msg())
	// Output:
	// "Super prueba: Sol (0) required value > 0, \n  Prueba: Número (3) required value in [1 2 4], Este es un error libre, "
}
