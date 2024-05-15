package server

import (
	"fmt"
	"os"
)

// - store
//	- meta.json
// 	- tables
// 		- table1.json
// 		- table2.json
// 	- data
// 		- table1_data.json
// 		- table2_data.json

func CheckDirs() {
	fmt.Println("Checking directories...")
	if _, err := os.Stat("store"); os.IsNotExist(err) {
		fmt.Println("Creating store directory...")
		os.Mkdir("store", 0755)
	}

	if _, err := os.Stat("store/tables"); os.IsNotExist(err) {
		fmt.Println("Creating store/tables directory...")
		os.Mkdir("store/tables", 0755)
	}

	if _, err := os.Stat("store/data"); os.IsNotExist(err) {
		fmt.Println("Creating store/data directory...")
		os.Mkdir("store/data", 0755)
	}

	if _, err := os.Stat("store/meta.json"); os.IsNotExist(err) {
		fmt.Println("Creating store/meta.json file...")
		os.Create("store/meta.json")
	}

}
