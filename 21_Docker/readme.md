Summary: Docker

- Docker adalah platform open-source yang digunakan untuk mengembangkan, mengirim, dan menjalankan aplikasi dalam wadah atau container
- Container adalah sebuah unit yang membungkus kode beserta dependensi sehingga aplikasi dapat dijalankan dengna cepat dan andal di berbagai environment
- Perbedaan container dan virtual machine adalah vm digunakan untuk menjalankan komputer dalam bentuk virtual sedangkan container digunakan untuk membungkus sebuah aplikasi. Container adalah abstraksi di layer aplikasi sedangkan vm adalah abstraksi pada layer hardware fisik
- Docker Volume adalah sebuah mekanisme penyimpanan data yang dapat digunakan agar dapat tetap ada (persistent)
- Mengapa menggunakan Docker Volume?
    1. Backup dan migrasi lebih muda
    2. Dapat dibagi lebih aman dengan container lain
    3. Menambah fungsionalitas lain
- Docker Network adalah sebuah jaringan yang dapat digunakan antar container untuk berkomunikasi satu sama lain
- Container Orchestration adalah sebuah mekanisme untuk mengelola container. Container Orchestration cocok digunakan untuk mengelola sistem yang dibangun dengan berbagai container seperti microservices
- Docker compose adalah sebuah container orchestration yang sudah disediakan oleh docker
- Docker compose digunakan untuk menjalankan berbagai container dengan membuat file konfigurasi dalam format YAML