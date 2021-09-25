package task9

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//Wget runs wget util
func Wget(args []string) error {
	// fmt.Println("args:", args)
	fileName := args[0]
	urlPath := args[1]
	response, err := http.Get(urlPath)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	fmt.Println("Response status: ", response.Status)

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	size, err := io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Saved in %s with size %d\n", fileName, size)
	return nil
}
