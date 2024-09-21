Summary: String - Advance Function - Pointer - Method - Struct and Interface

- Beberapa fungsi yang digunakan untuk string manipulation diantaranya:
    1. len
    2. compare
    3. contains
    4. substring
    5. Replace
    6. insert
- Variadic function digunakan agar kita bisa melewati proses pembuatan temporary slice ketika akan di-passing ke dalam function
- Anonymous function adalah sebuah function yang tidak memiliki nama, digunakan ketika membuat inline function
- Closure merupakan tipe dari anonymous function yang menunjuk variabel-variabel yang dideklarasikan di luar function itu sendiri
- Defer function merupakan sebuah function yang akan di eksekusi ketika parent functionnya mengembalikan nilai
- Pointer adalah variabel yang menyimpan alamat memori dari variabel lain. Penggunan pointer biasanya menggunakan simbol (&) untuk mengembalikan alamat variabel atau mengakses alamat sebuah variabel pointer, selain itu juga ada simbol (*) yang digunakan untuk mendeklarasikan variabel pointer dan mengakses nilai yang disimpan di alamat memori variabel tersebut
- Struct merupakan sebuah tipe data yang dapat berisi beberapa tipe data yang berbeda atau method
- Method adalah fungsi yang melekat pada sebuah tipe data tertentu bisa struct atau tipe data yang lain
- Interface adalah sekumpulan method abstrak yang dapat diimplementasi oleh tipe data, interface mendefinisikan perilaku dari sebuah objek
- Error handling, di bahasa Go tidak ada try catch akan tetapi menggunakan Panic dan Recover
- Panic terjadi ketika Go runtime mendeteksi ada sebuah kesalahan biasanya program akan terhenti
- Recover menjadi solusi ketika terjadi panic