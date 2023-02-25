# (7) Time Complexity & Space Complexity
- Time complexity merupakan waktu yang dibutuhkan untuk menjalankan suatu program, sedangkan space complexity merupakan memori yang dibutuhkan untuk menjalankan suatu program.
- Ada beberapa jenis time complexity, diantaranya :
    1. Constant time atau O(1) jika tidak terdapat loop sama sekali dalam sebuah program.
    2. Linear time atau O(n) jika terdapat 1 loop yang dimana loop tersebut akan dijalankan hingga n kali atau batasan dari loop nya yaitu sampai n.
    3. Linear time yang kedua atau O(n+m) jika terdapat 2 loop tetapi tidak nested yang dimana loop tersebut akan dijalankan hingga n kali dan yang satunya hingga m kali.
    4. Logarithmic time atau O(log n) jika terdapat 1 loop yang dimana nilai n akan di div atau dibagi dengan 2 pada setiap looping nya sehingga loop tersebut dijalankan sebanyak 2 log n kali. Contoh n = 8, maka loop dijalankan sebanyak 3 kali.
    5. Quadratic time atau O(n^2) jika terdapat 2 loop yang bersarang, maka looping akan dijalankan sebanyak n*n atau n^2.
    6. Faktorial time atau O(n!)
    7. Exponential time atau O(2^n)
- Yang terakhir untuk nilai space complexity dipengaruhi oleh banyaknya variable.