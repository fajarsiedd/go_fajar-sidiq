Summary: ORM & Code Structure (MVC)

- ORM (Object Relational Mapping) adalah sebuah teknik pemrograman untuk mengkonversikan data antara incompatible type systems menggunakan bahasa object-oriented programming
- GORM adalah package atau ORM library yang diguanakan untuk bahasa pemrograma Go
- Keuntungan menggunakan ORM:
    1. Mengurangi query berulang
    2. Secara otomatis mengambil data ke objek yang siap digunakan
    3. Cara yang simpel jika ingin screening data sebelum menyimpan ke database
    4. Beberapa ORM memiliki fitur caching
- Kelemahan ORM:
    1. Menambah layer baru pada code dan menghitung biaya proses overhead
    2. Memuat relasi data yang tidak perlu
    3. Untuk query yang kompleks akan sangat panjang dibuat jika dengan ORM (>10 tabel join)
- Databse Migration adalah cara untuk mengupdate versi database untuk mengikuti perubahan versi aplikasi
- DB relations di GORM: Belongs To, Has One, Has Many, Many to Many
- Selebihnya untuk cara menginstall GORM dan cara menggunakannya bisa dilihat di dokumentasi resmi GORM
- MVC adalah singkatan dari Model, View dan Controller. MVC adalah cara populer dalam mengorganisir sebuah code
- MVC digunakan untuk membuat aplikasi modular, mengimplementasikan concerns terpisah dan mengurangi conflict pada versioning
- Struktur folder sederhana yang bisa digunakan untuk menerapkan konsep MVC pada project Go:
    - configs
    - controllers
    - models
    - routes