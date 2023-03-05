# (11) Problem Solving Paradigm - Brute Force, Greedy, and Dynamic Programming
- Brute Force
    1. Algoritma ini bukan merupakan prioritas utama untuk digunakan.
    2. Algoritma ini mudah untuk dirancang namun membutuhkan waktu yang lama, karena dia akan mencari seluruh kemungkinan yang mungkin terjadi.
- Divide and Conquer
    1. Cara kerja algoritma ini membagi menjadi bagian yang lebih kecil lalu menyelesaikan setiap bagian tersebut.
    2. Terdapat 3 tahapan dari algoritma ini, diantaranya : divide, conquer, dan combine.
    3. Pada tahap divide kita membagi masalah menjadi bagian yang lebih kecil.
    4. Kemudian, pada tahap conquer kita menyelesaikan masalahnya.
    5. Terakhir pada tahap combine kita menyatukan kembali bagian kecil tadi menjadi solusi untuk masalah yang besar.
    6. Contoh kasus nya yaitu binary search dan power(perpangkatan).
- Greedy
    1. Algoritma ini akan menggunakan pilihan lokal optimal yang harapannya itu juga merupakan solusi untuk global optimal.
    2. Untuk beberapa kasus algoritma ini dapat digunakan dan berjalan dengan cepat.
    3. Contoh kasus nya adalah coin change, huffman coding, activity selection, dijkstra algorithm, dll.
- Dynamic Programming
    1. Cara kerja dari algoritma ini yaitu dengan cara memecahnya menjadi submasalah yang lebih sederhana dan memanfaatkan fakta bahwa solusi optimal untuk keseluruhan masalah bergantung pada solusi optimal untuk submasalahnya.
    2. Karakteristik dari dynamic programming : Overlapping subproblems dan optimal substructures property.
    3. Method yang ada pada dynamic programming : Top down with memoization dengan menyimpan nilai data dari subproblem yang sudah dicari, dan yang kedua ada bottom up with tabulation.
    4. Contoh kasus nya adalah fibonacci, dll.
