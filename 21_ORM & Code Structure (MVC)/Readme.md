# (21) ORM and Code Structure
- ORM (Object Relational Mapping) adalah teknik pemrograman untuk mengubah data antara sistem tipe yang tidak kompatibel menggunakan bahasa pemrograman berorientasi objek.
- GORM merupakan ORM library untuk Golang.
- Keuntungan dari ORM: Kueri yang lebih sedikit berulang, Ambil data secara otomatis ke objek siap pakai, Cara sederhana jika Anda ingin menyaring data sebelum menyimpannya di basis data, dan Beberapa memiliki kueri cache fitur.
- Kerugian dari ORM: Tambahkan lapisan dalam kode dan biaya proses overhead, Memuat data hubungan yang tidak diperlukan, Permintaan mentah yang kompleks dapat lama ditulis dengan ORM (> 10 tabel joins), Fungsi SQL khusus yang terkait dengan satu vendor mungkin tidak didukung atau tidak ada fungsi khusus ( MySQL : FORCE INDEX).
- DB Migration merupakan Cara memperbarui versi basis data untuk menyesuaikan perubahan versi aplikasi. Perubahan dapat berupa upgrade ke versi terbaru atau rollback ke versi sebelumnya.
- Kelebihan DB Migration: Kesederhanaan pembaruan basis data, Kesederhanaan rollback basis data, Lacak perubahan pada struktur basis data, Riwayat struktur basis data ditulis pada kode, Selalu kompatibel dengan perubahan versi aplikasi.
- Install GORM: go get -u gorm.io/gorm
- MySQL Driver: go get -u gorm.io/driver/mysql