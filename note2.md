Jangan pakai panic jika ingin menggagalkan unit test
Untuk menggagalkan unit test: 
    t.Fail(): Dianggap gagal, tetapi tetap melanjutkan eksekusi unit test
    t.FailNow(): Diangga gagal dan tidak melanjutkan eksekusi unit test
    t.Error(): Bisa ditambahkan argument errornya, kemudian akan memanggil Fail() sehingga berperilaku seperti Fail
    t.Fatal(): Bisa ditambahkan argument errornya, kemudian akan memanggil Fail() sehingga berperilaku seperti FailNow()

Disarankan menggunakan assertion untuk pengecekan, assertions semacam library
    Jangan menggunakan if else karena ribet
    assertion tidak tersedia di golang, harus menggunakan library: Testify
    ada 2 package di testify: 
        assert dan require
        assert: Memanggil fail()
        require: Memanggil failnow()

Skip test: Untuk membatalkan unit test  
    Contohnya: Jika berjalan di mac akan di skip
