Summary: Unit Testing

- Software Testing adalah sebuah proses menganalisis suatu item perangkat lunak untuk mendeteksi perbedaan antara kondisi yang ada dan kondisi yang diperlukan (yaitu, cacat) serta untuk mengevaluasi fitur dari item perangkat lunak tersebut
- Tujuan testing:
    1. Mencegah regresi: Pastikan fitur yang sudah jadi tidak rusak saat ada perubahan
    2. Meningkatkan kepercayaan saat mengubah kode: Yakin bahwa perubahan kode tidak menimbulkan masalah baru
    3. Perbaiki desain kode: Identifikasi bagian kode yang kurang baik dan ubah
    4. Dokumentasi kode: Pengujian bisa jadi seperti contoh cara kerja kode
    5. Permudah kerja sama tim: Pengujian yang baik membantu anggota tim baru memahami proyek
- Level testing:
    1. Unit: Menguji bagian terkecil dari aplikasi, seperti sebuah fungsi. Tujuannya memastikan setiap bagian kecil bekerja dengan benar
    2. Integration: Menguji bagaimana beberapa bagian kecil yang sudah diuji tadi bekerja sama.
    3. UI (end to end): Menguji keseluruhan aplikasi dari sudut pandang pengguna. Apakah mudah digunakan? Apakah tampilannya menarik?
- Framework yang digunakan di Go adalah testify
- Untuk mocking package yang digunakan adalah mockery
- Hal yang tidak boleh di test:
    1. 3rd party modules
    2. Database
    3. 3rd party api atau sistem eksternal
    4. Sebuah class object yang sudah ditest di tempat lain