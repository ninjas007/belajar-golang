package main

import "testing"

func Caculate(a, b int) int {
	return a + b
}

func TestCalculate(t *testing.T) {
	a, b := 10, 20
	must := 30
	result := Caculate(a, b)

	t.Logf("Hasil Calculate adalah: %v", result)

	if result != must {
		t.Errorf("Salah! hasil yang diharapkan : %v dan yang didapatkan : %v", must, result)
	}
}