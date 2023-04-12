# (24) Clean and Hexagonal Architecture
- Arsitektur yang baik adalah pemisahan perhatian menggunakan lapisan untuk membangun aplikasi yang modular, dapat diskalakan, dapat dipelihara, dan dapat diuji.
- Kendala yang biasa ditemukan sebelum men-design clean architecture, diantaranya pada: independent of frameworks, testable, independent of UI, independent of Database, dan independent of any external.
- Keuntungan dari clean architecture: struktur standar, sehingga mudah menemukan jalan Anda dalam proyek, perkembangan yang lebih cepat dalam jangka panjang, ketergantungan mocking menjadi sepele dalam pengujian unit,beralih dengan mudah dari prototipe ke solusi yang tepat (mis., mengubah penyimpanan dalam memori ke database SQL).
- Terdapat 3 CA Layer, diantaranya: Use case, Controller, dan drivers.
- Use case berisi logic bisnis.
- Controller menampilkan data ke layar dan mengelola interaksi dengan user.
- Drivers memanage data dari sebuah aplikasi.
- DDD (Domain Driven Design) merupakan salah satu teknik design perangkat lunak.
