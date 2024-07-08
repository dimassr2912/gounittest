testing.M
    Nama function harus TestMain dengan paramter testing M
    Hanya dieksekusi 1x tiap package 
    Bisa dimanfaatkan sebagai before after test
    Before after test dieksekusi sebelum unit test 
    Biasanya di awal untuk memanggil koneksi ke database

Subtest: Mendukung membuat function unit test di dalam function unit test
    Untuk menjalankan subtest saja: go test -v -run=namafunction/subtest
    Semua subtest di semua function: go test -v -run=/rama (Mnejalankan semua yang ada nama ramanya)

Table test: Data berupa slice yang berisi ekspektasi hasil dari unit test