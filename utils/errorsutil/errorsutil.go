package errorsutil

import "fmt"

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
