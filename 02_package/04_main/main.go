package main

import (
	"fmt"

	bear "github.com/GoesToTwentyOne/Learnig_from_code_Golang_Training-Todd_Sir/02_package/01_bear"
	dog "github.com/GoesToTwentyOne/Learnig_from_code_Golang_Training-Todd_Sir/02_package/02_dog"
	reverseone "github.com/GoesToTwentyOne/Learnig_from_code_Golang_Training-Todd_Sir/02_package/03_reverse_func"
)

func main() {
	fmt.Println(reverseone.Reverse("Nihad"))
	fmt.Println(bear.BearName)
	fmt.Println(dog.DogName)
}
