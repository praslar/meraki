package main

import "fmt"

func main() {

	type HocSinh struct {
		Ten        string
		Tuoi       int
		DangConHoc bool
	}

	type SinhVien struct {
		Ten          string
		Tuoi         int
		SinhVienNam  int
		TruongDaiHoc string
	}

	type Ly struct {
		CoQuai       bool
		Color        string
		DungTich     int
		ChieuCao     int
		ChatLieu     string
		Tuoi         int
		SinhVienNam  int
		TruongDaiHoc string
		Ten          string
		Tuoi         int
		DangConHoc   bool
	}

	hocSinh1 := HocSinh{
		Ten:        "Vu",
		Tuoi:       27,
		DangConHoc: true,
	}

	thinh := SinhVien{
		Ten:          "Thinh",
		Tuoi:         25,
		SinhVienNam:  7,
		TruongDaiHoc: "Dai hoc da lai",
	}

	// object

	// 1000 cai ly

	for i := 0; i < 1000; i++ {
		object := Ly{
			CoQuai:       true,
			Color:        "Trang",
			DungTich:     10,
			ChieuCao:     200,
			ChatLieu:     "Su",
			SinhVienNam:  1,
			TruongDaiHoc: "Dai hoc",
			Ten:          "Vu",
			Tuoi:         0,
			DangConHoc:   false,
		}
		if i%2 == 0 {
			object.ChieuCao = 120
			object.CoQuai = false
			object.Color = "Do"
			object.ChatLieu = "Nhua"
		}

		fmt.Println(object)
	}
	fmt.Println(hocSinh1, thinh)
}
