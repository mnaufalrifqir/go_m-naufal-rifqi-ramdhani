# (9) String - Advance Function - Pointer - Method - Struct and Interface
- String
    1. Len pada string berfungsi untuk melihat panjang dari string tersebut / banyaknya huruf dari string tersebut (termasuk space, dll).
    2. String dapat kita bandingkan menggunakan "==" untuk mengecek apakah nilai nya sama atau tidak.
    3. Pada packages strings ada yang namanya strings.Contains(a, b) untuk mengecek apakah string b terdapat pada string a atau tidak.
    4. Pada packages strings juga ada yang namanya strings.Replace(string, "-", " ", 2) untuk mengubah nilai pada string, jika cara pemanggilannya seperti diatas berarti ia akan merubah "-" sebanyak 2x menjadi " " pada string string.
- Function
    1. Kita dapat menggunakan (number ...int) pada parameter sebuah function agar inputannya berubah menjadi sebuah slice.
    2. Kita dapat menggunakan keyword defer pada function agar function tersebut dieksekusi di akhir program.
- Pointer
    1. Pointer merupakan sebuah variable yang dapat menyimpan alamat dari sebuah data.
    2. Cara untuk membuat variable pointer, contohnya : var p *int = &a.
    3. Kita dapat menggunakan "*" pada pemanggilan nama variable pointer untuk melihat isi datanya.
    4. Kita dapat menggunakan "&" pada pemanggilan nama variable string/int/float atau yang lainnya untuk melihat alamatnya.
    5. Pointer juga dapat digunakan untuk parameter sebuah fungsi untuk mengubah nilai berdasarkan reference nya.
- Struct merupakan tipe data bentukan yang memiliki properti satu atau lebih.
- Method
    1. Metode adalah fungsi yang melekat pada suatu tipe (bisa struct atau tipe data lainnya)
    2. Cara mendeklarasikan sebuah method, func (receiver StructType) MethodName(parameterList) (returnTypes){}
- Interface
    1. Interface adalah kumpulan tanda tangan metode yang dapat diimplementasikan objek. Karenanya Inteface mendefinisikan perilaku objek.
    2. Cara mendeklarasikan sebuah interface, type namaInterface interface{}, didalamnya berisi namaMethod dengan returnType nya.
    3. Interface dapat digunakan sebagai variable dinamic yang nantinya isi datanya bisa berubah-ubah tipe datanya.
