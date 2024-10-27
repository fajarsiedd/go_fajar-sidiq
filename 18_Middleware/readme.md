Summary: Middleware

- Middleware adalah suatu entitas yang menyisipkan dirinya ke dalam proses permintaan dan respons server
- Middleware memungkinkan kita untuk mengimplementasikan berbagai fungsionalitas di sekitar permintaan HTTP yang masuk ke server dan respons yang keluar
- Contoh third-party middleware di Go
    1. Negroni
    2. Echo
    3. Interpose
    4. Alice
- Echo Pre() dieksekusi sebelum router memproses request
- Echo Use() dieksekusi setelah router memproses request dan memiliki ful akses pada echo.Context API
- CORS adalah singkatan dari Cross Origin Resource Sharing
- CORS memungkinkan untuk berbagi resource dengan origins atau domain yang berbeda
- Rate Limiter berfungsi untuk membatasi jumlah request yang masuk ke server
- Rate Limiter berguna untuk mencegah serangan seperti DDos atau Bruce Force Attack
- Logger Middleware digunakan untuk mencatat informasi dari HTTP request
- Basic Authentication adalah sebuah teknik autentikasi http request, method ini membutuhkan username dan password untuk dimasukkan ke request header
- JWT adalah singkatan dari JSON Web Token adalah standar terbuka (RFC 7519) yang digunakan untuk berbagi informasi yang terverifikasi dan aman antara dua pihak
- Auth Middleware digunakan untuk memvalidasi otentikasi user