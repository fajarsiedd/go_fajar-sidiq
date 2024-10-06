package main

import "fmt"

// Format penamaan objek sebaiknya menggunakan Pascal Case. Objek 'user' sebaiknya diubah menjadi 'User'
type user struct {
	Email    string
	Password string
}

/*
	Format penamaan objek sebaiknya menggunakan Pascal Case. Objek 'userRepo' sebaiknya diubah menjadi 'UserRepo' atau 'UserRepository'.
	Penamaan variabel untuk tipe data slice of user tidak mewakili objek. Sebaiknya 'DB' diubah menjadi 'Users'.
*/
type userRepo struct {
	DB []user
}

func (r userRepo) Register(u user) {
	if u.Email == "" || u.Password == "" {
		fmt.Println("register failed")
	}

	r.DB = append(r.DB, u)
}

func (r userRepo) Login(u user) string {
	if u.Email == "" || u.Password == "" {
		fmt.Println("login failed")
	}

	/*
		Penamaan variabel ketika melakukan iterasi pada sebuah slices of object sebaiknya ditulis dengan jelas dan mewakili
		objek tunggal dari slices tersebut. Sebaiknya 'us' diubah menjadi 'user'
	*/
	for _, us := range r.DB {
		if us.Email == u.Email && us.Password == u.Password {
			return "auth token"
		}
	}

	/*
		Mengembalikan nilai kosong setelah melakukan pengecekan data user akan membingungkan,
		sebaiknya diberikan default message seperti 'username/password tidak valid'
	*/
	return ""
}
