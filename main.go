package main

import (
	"fmt"

	"github.com/dslipak/pdf"
)

func main() {
	content, err := readPdf2("./FunProjectsGo/p2.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

func readPdf2(path string) (string, error) {
	r, err := pdf.Open(path)
	// remember close file
	if err != nil {
		return "", err
	}

	//totalPage := r.NumPage()

	p := r.Page(1)
	fmt.Println(p.Resources().String())
	// fmt.Println(p.Content().Text)
	//j, _ := json.MarshalIndent(p, "", " ")
	//fmt.Println(string(j))

	// for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
	// 	p := r.Page(pageIndex)
	// 	if p.V.IsNull() {
	// 		continue
	// 	}
	// 	var lastTextStyle pdf.Text
	// 	texts := p.Content().Text
	// 	for _, text := range texts {
	// 		lastTextStyle.S = lastTextStyle.S + text.S
	// 		// fmt.Printf("Font: %s, Font-size: %f, x: %f, y: %f, content: %s \n", lastTextStyle.Font, lastTextStyle.FontSize, lastTextStyle.X, lastTextStyle.Y, lastTextStyle.S)
	// 		fmt.Print(text.S)
	// 		lastTextStyle = text
	// 	}
	// }
	return "", nil
}
