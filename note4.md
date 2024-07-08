Benchmark: Menghitung kecepatan performa kode aplikasi
    Melakukan iterasi kode berkali-kali sampai waktu tertentu
    Ada attribute N yang digunakan untuk iterasi
    nama function harus diawali Benchmar dengan paramter (b *testing.B)
    nama file diakhiri _test
    Untuk menjalankan benchmar dengan unit test: go test -v -bench=.
    Running seluruh benchmark: go test -v -run=TestTidakAda -bench=.
    Untuk running tertentu: go test -v -run=TidakAda -bench=BenchmarkHelloWorld
    Running benchmark di dalam package yang lain: go test -v -run=TidakAda -bench=. ./...
    Untuk running subbenchmark go test -v -run=TidakAda -bench=Benchmark/subbenchamrk