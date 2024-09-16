package models

//Example for strusctur of table
type Transparasi struct{
	Id int64 `gorm:"primary_key"`
	NamaInstitusi string `json:"nama_institusi" gorm:"varchar(300)"`
	JenisAnggaran int16 `json:"jenis_anggaran" gorm:"int(10)"`
	JumlahAnggaran string `json:"jumlah_anggaran" gorm:"varchar(300)"`
	Kategori int16 `json:"kategori" gorm:"int(10)"`
	Uraian string  `json:"uraian" gorm:"text"`
	Alamat string  `json:"alamat" gorm:"text"`
	
}