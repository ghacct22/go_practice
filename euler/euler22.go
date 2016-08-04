package euler

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Euler22() {
	data, err := ioutil.ReadFile("p022_names.txt")
	check(err)
	names := strings.Split(strings.Replace(string(data), "\"", "", -1), ",")
	sort.Strings(names)
	nameValueSum := 0
	for index, name := range names {
		nameValue := 0
		for _, char := range name {
			nameValue += int(char) - 64
		}
		// if name == "COLIN" {
		//     fmt.Printf("%s has an index of %d and a value of %d\n", name, index, nameValue)
		//     for _, char := range name {
		//         fmt.Println(int(char) - 64)
		//     }
		// }
		nameValueSum += nameValue * (index + 1)
	}

	fmt.Println(nameValueSum)
}
