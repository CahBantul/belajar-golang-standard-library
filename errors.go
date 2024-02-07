package main

import "errors"

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Nozami"{
		return NotFoundError
	}

	//sukses
	return nil
}

func main(){
	err := GetById("Nozami")

	if err != nil {
		 if errors.Is(err, ValidationError){
		 	println("Terjadi kesalahan validasi")
		 }else if errors.Is(err, NotFoundError){
		 	println("Data tidak ditemukan")
		 }else{
			println("uknown error")
		 }
	}
}