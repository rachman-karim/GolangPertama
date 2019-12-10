package main

import (
	raka "fmt"
	"os"
)

type person struct {
	name      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	var argsRaw = os.Args
	//raka.Printf("-> %#v\n", argsRaw)

	P := person{"Rachman Karim", "jalan jalan", "prigrammer", "Cari sesuatu yang baru"}
	P2 := person{"Riska Aprian", "jalan pondok gede", "prigrammer", "Cari sesuatu yang baru"}
	//raka.Printf("hasil", argsRaw)
	//configPath := os.Args

	if argsRaw[1] == "1" {
		//configpath = os.Args[1]
		raka.Println("Nama saya adalah", P.name)
		raka.Println("Alamat saya di", P.alamat)
		raka.Println("Pekerjaan saya adalah", P.pekerjaan)
		raka.Println("Alasan memilih kelas golang", P.pekerjaan)
	} else if argsRaw[1] == "2" {
		//configpath = os.Args[2]
		raka.Println("Nama saya adalah", P2.name)
		raka.Println("Alamat saya di", P2.alamat)
		raka.Println("Pekerjaan saya adalah", P2.pekerjaan)
		raka.Println("Alasan memilih kelas golang", P2.pekerjaan)
	}

}
