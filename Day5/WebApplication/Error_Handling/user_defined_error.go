package main

import "fmt"
import "errors"

func main() {

	id, err := GetData()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

}

func GetData() (int, error) {

	var id = 0

	if id == 0 {
		return 0, errors.New("Value shoudl be greater than 0")
	} else {
		return 100, nil
	}
}
