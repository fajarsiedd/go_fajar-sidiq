Summary: Intro Echo Golang

- Beberapa third party library yang bisa digunakan di bahasa pemrograman Golang seperti:
    1. echo
    2. Go Kit
    3. Logrus
    4. gorm.io
    5. cobra
- Echo adalah framework web yang ringan dan cepat untuk bahasa pemrograman Go (Golang). Echo dirancang untuk membangun aplikasi web dan RESTful API dengan mudah dan efisien
- Keuntungan menggunakan echo
    1. Optimized Router
    2. Data Rendering
    3. Data Binding
    4. Middleware
    5. Scalable
- Echo merupakan framework yang minimalis namun ekstensible. Dapat diintegrasikan dengan GORM untuk ORM dan Go JWT untuk autentikasi
- Untuk menginstansiasi echo dapat menggunakan -> e := echo.New()
- Untuk mendaftarkan route dapat menggunakan -> e.GET('/myroute', myController)
- Untuk mengambil url params dapat menggunakan -> echo.Context.Param("myUrlParam")
- Untuk mengambil query params dapat menggunakan -> echo.Context.QueryParams("match")
- Untuk mengambil request body dapat menggunakan -> echo.Context.Binding(&MyStruct{})