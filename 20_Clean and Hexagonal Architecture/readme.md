Summary: Clean and Hexagonal Architecture

- Arsitektur yang baik adalah separation of concern menggunakan layer untuk membangun aplikasi yang modular, skalabel, mudah dipelihara, dan mudah diuji
- Hexagonal Architecture, juga dikenal sebagai Ports and Adapters Architecture, adalah sebuah pola arsitektur perangkat lunak yang bertujuan untuk memisahkan logika bisnis dari detail implementasi luar, seperti database, UI, atau sistem lainnya
- Manfaat CA:
    1. Struktur standar memudahkan untuk menemukan jalan Anda di proyek.
    2. Pengembangan lebih cepat dalam jangka panjang.
    3. Mocking dependensi menjadi sepele dalam pengujian unit.
    4. Mudah beralih dari prototipe ke solusi yang tepat (misalnya, mengubah penyimpanan di memori menjadi database SQL).
- CA layer:
    - Entitiy Layer (opsional)
    - Use Case -> Domain Layer
    - Controller -> Presentation Layer
    - Drivers -> Data Layer
- DDD (Domain Driven Design) adalah pendekatan pengembangan perangkat lunak yang berfokus pada kolaborasi antara pengembang dan ahli domain untuk membangun model yang mencerminkan kebutuhan bisnis dan logika inti aplikasi
- Context adalah suatu package yang membawa deadline, cancellation signal, atau value lain pada request API