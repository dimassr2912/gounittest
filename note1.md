Software testing: Memastikan minim bug dan kesalahan
Testing (dari paling bawah)     
    Unit testing 
        Test komponen paling kecil
    Service testing / integration testing 
    UI testing / End to end test
        Harus buat bot
Makin bawah makin mudah dan murah implementasi sehingga porsinya banyak
Makin ke atas makin mahal resource dan implementasi
Unit test: Simple, murah, bisa bikin banyak
Semakin bawah semakin cepat eksekusi

- Jika ada high level architecture application
Ada front end terhubung 2 backend, masing2 backend terhubung db yang berbeda, dan db salah satu terhubung dengan payment gateway

UI testing 
    Di testing dengan bot (seakan-akan user biasa)
    Harus running semua (dari frontend hingga db)
    + makin real karena seperti behavior dari user aslinya
Service testing 
    Lebih fokus ke 1 aplikasi
        Contoh ke salah satu backend 
        - Perlu run database dan other system (payment gateway)
Unit test
    testing komponen terkecil (fokus test function dan method)
    Tidak perlu run database

Aturan pembuatan unit test: nama file diakhiri _test
	Contoh: hello_world_test.go
Nama function untuk unit test harus diawali Test
	Contoh: TestHelloWorld
Parameter harus t *testing.T dan tidak boleh ada return value
Untuk menjalankan unit test
	go test: test seluruh unit test di package
	go test -v: test seluruh unit test di package dengan informasi lebih detail
	go test -v -run NamaFile: test unit test salah satu function
	go test ./... : Menjalankan seluruh unit test dari modulenya

Menggagalkan test 
    Fail(): Menggagalkan tetapi tetap melanjutkan eksekusi code program (bukan unit test) selanjutnya
    FailNow(): Menggagalkan unit test saat itu juga serta eksekusi program selanjutnya, tetapi unit test selanjutnya tetap dilanjutkan
    Error(): melakukan log(print) lalu memanggil function Fail
    Fatal(): Melakukan log(print) lalu memanggil function FailNow

Assertion
Melakukan menggunakan if else akan menyebalkan ketika banyak return 
    Contoh: Struct dengan banyak attribute di cek
Menggunakan module testify: go get github.com/stretchr/testify
ada 2 package: 
    assert.Equal(t, expected, result, "ERROR")
        Jika false maka akan memanggil Fail()
    require.Equal(t, expected, result, "ERROR")
        Jika false maka akan memanggil FailNow()

Skip test 
Jika ingin skip test di os yang berbeda 
function Skip()
Statement di bawah Skip() tidak dieksekusi
GOOS: Go operating System

Before after 
Before biasanya untuk koneksi ke db
TestMain dieksekusi sekali tiap package
m.Run: Untuk test semua unit test

Subtest
Functin unit test di dalam function unit test 
Untuk run sub test: go test -v -run NamaTest/NamaSubTest

Table Test 
Ketika unit test isi sama tetapi hanya ekpektasi berbeda lebih baik menggunakan perulangan



TIDAK PAHAM
1. Third party service 
2. Mocking service
3. parameter t *testing.T, t *testing.M, t *testing.B
4. testing.T = struct? kenapa bukan interface
5. Before after unit test untuk apa?
