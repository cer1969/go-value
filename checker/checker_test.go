// Copyright Cristian Echeverría Rabí

package checker

import (
	"fmt"
	"testing"
)

func TestNewAndReset(t *testing.T) {
	vc := New("Prueba")
	if vc.title != "Prueba" {
		t.Errorf("Wrong Title: %v", vc.title)
	}
	if len(vc.msgs) != 0 {
		t.Errorf("Wrong Count: %d", len(vc.msgs))
	}

	vc.Ck("Número", 30.0).Lt(30.0)
	err := vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
	if len(vc.msgs) != 1 {
		t.Errorf("Wrong Count: %d", len(vc.msgs))
	}

	//vc.Reset("")
	//if vc.msg != "" {
	//	t.Errorf("Wrong Msg: %v", vc.msg)
	//}
	//if len(vc.msgs) != 0 {
	//	t.Errorf("Wrong Count: %d", len(vc.msgs))
	//}
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
}

func ExampleChecker() {
	vc := New("Prueba")
	vc.Ck("Número", 30).Gt(40.0).Gt(100.0)
	vc.Ck("Sol", 0.0).Gt(0.0).Lt(1.0)
	vc.Ck("Sol", 0.5).Gt(0.0).Lt(1.0)
	err, _ := vc.Error().(*CheckError) // Verifica y convierte vc.Error() en *CheckError
	fmt.Printf(err.Error())
	fmt.Printf("\n%d", err.Count())
	// Output:
	// Prueba
	//   Número required value > 40.000000 [30.000000 received]
	//   Número required value > 100.000000 [30.000000 received]
	//   Sol required value > 0.000000 [0.000000 received]
	// Total errors: 3
	// 3
}

func ExampleIn() {
	vc := New("Prueba")
	vc.Ck("Número", 3.0).In(1.0, 2.0, 4.0)
	fmt.Printf("%v", vc.Error())
	// Output:
	// Prueba
	//   Número required value in [1 2 4] [3.000000 received]
	// Total errors: 1
}

func ExampleAppend() {
	vc := New("Prueba")
	vc.Ck("Número", 3.0).In(1.0, 2.0, 4.0)
	vc.Append("Este es un error libre")
	fmt.Printf("%v", vc.Error())
	// Output:
	// Prueba
	//   Número required value in [1 2 4] [3.000000 received]
	//   Este es un error libre
	// Total errors: 2
}

func ExampleAppendChecker() {
	vc1 := New("Prueba")
	vc1.Ck("Número", 3.0).In(1.0, 2.0, 4.0)
	vc1.Append("Este es un error libre")
	//err1, _ := vc1.Error().(*CheckError)

	vc2 := New("Super prueba")
	vc2.Ck("Sol", 0).Gt(0).Lt(1)
	vc2.AppendError(vc1.Error())

	fmt.Printf("%v", vc2.Error())
	// Output:
	// Super prueba
	//   Sol required value > 0.000000 [0.000000 received]
	//   Número required value in [1 2 4] [3.000000 received]
	//   Este es un error libre
	// Total errors: 3
}
