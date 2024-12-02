package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//BENCMARK TEST
func BenchmarkHelloWorldAdiva(b *testing.B){
	for i := 0; i < b.N; i++{
		HelloWorld("Adiva")
	}
}

func BenchmarkHelloWorldEko(b *testing.B){
	for i := 0; i < b.N; i++{
		HelloWorld("HAJI BOLOT NANDA")
	}
}

// TABLE TEST
// Konsep ini digunakan apabila dari test yang kita buat hanya mengubah isi data saja, namun test yang digunakan sama
// Didalam table test kita menggunakan SUBTEST
// Kita bisa membuat slice struct untuk mengisi data dari setiap variabel nya

func TestTableTest(t *testing.T)  {
	// Mendekalarasikan slice nya dulu
	value := []struct {
		name string
		result string
		expected string
	}{
		{
			name : "HelloWorld(Adiva)",
			result : "Adiva",
			expected : "Hello There, Adiva",
		},
		{
			name : "HelloWorld(Eko)",
			result : "Eko",
			expected : "Hello There, Eko",
		},
		{
			name : "HelloWorld(Dustin)",
			result : "Dustin",
			expected : "Hello There, Dustin",
		},
	}

	// Lalu mengambil data dari setiap array yang dibuat
	for _, test := range value{
		// Setelah setiap data diambil, maka buat SUBTEST nya
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.result)
			assert.Equal(t, test.expected, result, "Result not match")
			fmt.Printf("SUB TEST 1 = %s \n", test.result)
		})
	}
}


// SUB TEST
// fitur unit test untuk function di dalam function
// diawali dengan t.Run(nama test, function test nya )

func TestSubTest(t *testing.T){
	// 1
	t.Run("Adiva", func(t *testing.T) {
		result := HelloWorld("Adiva")
		assert.Equal(t, "Hello There, Adiva", result, "Result not match")
		fmt.Println("SUB TEST 1 = Adiva")
	})
	// 2
	t.Run("Darwin", func(t *testing.T) {
		result := HelloWorld("Darwin")
		require.Equal(t, "Hello, Darwin", result, "Result not match")
		fmt.Println("SUB TEST 2 = Darwin")
	})
	// 3
	t.Run("Skip Test", func(t *testing.T) {
		if runtime.GOOS == "windows"{
			t.Skip("Only running on Mac")
		}
		result := HelloWorld("Windows")
		assert.Equal(t, "Hello There, Windows", result, "Result not match")
		fmt.Println("SUB TEST 3 = Windows")
	})
}


// Before n After Unit Test
// Hanya berlaku di satu package, jadi untuk di package lain harus dibuat lagi TEST MAIN
// nama functionya HARUS 'TESTMAIN' dengan parameter 'testing.M'
func TestMain(m *testing.M){
	// before test
	fmt.Println("BEFORE UNIT TEST")

	// untuk menjalankan semua testing nya
	m.Run()

	// after test
	fmt.Println("AFTER UNIT TEST")
}


// Membatalkan Test atau Skip Test
func TestSkip(t *testing.T){
	name := "Eko"
	
	if name == "Darwin"{
		t.Skip("Your Name " + name + " is BANNED!!") // Test akan dilewatkan
	}

	// if runtime.GOOS == "windows"{
	// 	t.Skip("Only running on Mac") // Test akan dilewatkan karna berjalan di system operasi windows
	// }
	
	result := HelloWorld(name)
	require.Equal(t, "Hello There, Eko", result, "Result Not Same")
}

// Menggunakan Assertion
func TestHelloAssert(t *testing.T)  {
	result := HelloWorld("Assertion")

	assert.Equal(t, "Hello There, Assertion", result, "Result must be 'Hello There, Assertion'") // Yang dikembalikan adalah Fail(), jadi program akan tetap berjalan
	fmt.Println("TestHelloAssert Done")
}

// Menggunakan Require
func TestHelloRequire(t *testing.T)  {
	result := HelloWorld("Assertion")

	require.Equal(t, "Hello There, Assertion", result, "Result must be 'Hello There, Assertion'") // Yang dikembalikan adalah FailNow(), jadi program tidak akan dilanjutkan
	fmt.Println("TestHelloRequire Done") // Tidak akan dieksekusi
}

// Cara manual tidak menggunakan Library
func TestHelloWorldAdiva(t *testing.T) {
	result := HelloWorld("Adiva")

	if result != "Hello There, Adiva"{
		// Jika error
		// t.Fail() -> masih berlanjut 
		t.Error("Harusnya muncul \"Hello There, Adiva\"") // -> masih berlanjut; lebih di rekomendasikan
	}

	fmt.Println("TestHelloWorldAdiva with assert Done")
}

func TestHelloWorldDustin(t *testing.T) {
	result := HelloWorld("Dustin")

	if result != "Hello There, Dustin"{	
		// Jika error
		// t.FailNow() -> langsung stop
		t.Fatal("Harusnya muncul \"Hello There, Dustin\"") //-> langsung stop; lebih di rekomendasikan
	}

	fmt.Println("TestHelloWorldDustin Done")
}