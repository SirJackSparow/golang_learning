package helper

import "testing"

//harus awalan nama dikasih "Test" dan parameter harus seperti contoh dibawah
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("ABC")

	if result != "Hello AC" {
		t.Fail()
	}

	print("print tetap exsekusi jika fail")
}

func TestHelloWord2(t *testing.T) {
	result := HelloWorld("s")
	if result != "Hello ABC" {
		t.FailNow()
	}

	print("Print ini tidak di eksekusi jika fail")
}

//-----------------------------------------------------------------------------------------------------

func TestHelloWorld3(t *testing.T) {
	result := HelloWorld("ABC")

	if result != "Hello AC" {
		t.Error("Error :( ada alasan")
	}

	print("print tetap exsekusi jika fail")
}

func TestHelloWord4(t *testing.T) {
	result := HelloWorld("s")
	if result != "Hello ABC" {
		t.Fatal("Fatal Error")
	}

	print("Print ini tidak di eksekusi jika fail")
}
