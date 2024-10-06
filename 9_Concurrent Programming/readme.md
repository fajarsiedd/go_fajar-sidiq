Summary: Concurrent Programming

- Sequential program: ketika tugas baru akan dijalankan, tugas sebelumnya harus selesai terlebih dahulu
- Pararrel program: mengeksekusi lebih dari satu tugas secara serentak
- Concurrent program: tugas dieksekusi secara independen dan mungkin akan memberi ilusi seperti dieksekusi secara serentak
- Fitur Go:
    1. Concurrent execution (Goroutines)
    2. Synchornization and messaging (channels)
    3. Multi-way concurrent control (select)
- Goroutine adalah metod atau fungsi yang akan dijalankan secara concurrent dengan fungsi atau metod lain
- Channel digunakan sebagai komunikasi antar goroutine
- Race condition terjadi ketika ada dua thread yang mengakses memori secara bersamaan, yang salah satunya adalah menulis
- Untuk mengatasi race condition ada beberapa cara:
    1. Menggunakan Waitgroups
    2. Menggunakan Channel blocking
    3. Menggunakan Mutex