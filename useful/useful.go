package useful

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func swap() {
	a, b, c := 1, 2, 3
	fmt.Println(a, b, c)
	a, b, c = c, a, b
	fmt.Println(a, b, c)
}
func selectColumn() {
	filepath := []string{"./test.txt", "./test1.txt"}
	column := "1"

	col, err := strconv.Atoi(column)
	if err != nil {
		fmt.Println("column is not integer")
		return
	}

	for _, filename := range filepath {
		fmt.Println("filename : ", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error to open file %s", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)

		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("eror to read file %s", err)
			}

			fmt.Print("line : ", line)
			data := strings.Fields(line) // unicode.IsSpace() 함수에 정의된 whitespace 기준으로 자른다.
			if len(data) >= col {
				fmt.Println(data[col-1])
			}
		}
	}
}
func Useful() {

	//swap()
	selectColumn()

}
