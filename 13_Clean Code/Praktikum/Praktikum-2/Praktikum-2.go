package main

//Untuk nama awal Struct gunakan huruf besar, untuk menandakan bahwa itu nama struct bukan variable
type Kendaraan struct {
	//Untuk setiap pergantian kata pada nama variable diawali dengan huruf kapital
	totalRoda       int
	kecepatanPerJam int
}

//Untuk nama awal Struct gunakan huruf besar, untuk menandakan bahwa itu nama struct bukan variable
type Mobil struct {
	Kendaraan
}

//Gunakan nama variable "mobil" agar mudah dimengerti jangan gunakan "m" saja
func (mobil *Mobil) berjalan() {
	mobil.tambahKecepatan(10)
}

/*
Untuk setiap pergantian kata pada nama fungsi diawali dengan huruf kapital
Gunakan nama variable "mobil" agar mudah dimengerti jangan gunakan "m" saja
Untuk setiap pergantian kata pada nama variable diawali dengan huruf kapital
*/
func (mobil *Mobil) tambahKecepatan(kecepatanBaru int) {
	mobil.kecepatanPerJam = mobil.kecepatanPerJam + kecepatanBaru
}

func main() {
	//Untuk setiap pergantian kata pada nama variable diawali dengan huruf kapital
	mobilCepat := Mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	//Untuk setiap pergantian kata pada nama variable diawali dengan huruf kapital
	mobilLamban := Mobil{}
	mobilLamban.berjalan()
}

// Berapa banyak kekurangan dalam penulisan kode tersebut? 11