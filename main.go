package main

import (
	"fmt"

	"github.com/dslipak/pdf"
	"github.com/go-pdf/fpdf"
)

func main() {
	content, err := readPdf2("./FunProjectsGo/p2.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	//
	// pdfController()
}

//func pdfController() *fpdf.Fpdf {
// pdf.AddPage()
// pdf.SetFont("Arial", "B", 16)
// pdf.Cell(40, 10, "Hello, world\n")
// pdf.SetLineWidth(2)
// pdf.Ln(-1)
// pdf.Cell(40, 10, "another line")
// pdf.OutputFileAndClose("hello.pdf")
// return pdf
//return nil
//}

func readPdf2(path string) (string, error) {

	pdfCtx := fpdf.New("P", "mm", "A4", "")
	r, err := pdf.Open(path)
	// remember close file
	if err != nil {
		return "", err
	}

	//totalPage := r.NumPage()

	page := r.Page(1)
	pdfCtx.AddPage()
	// fmt.Println(page.Content().Text)
	// j, _ := json.MarshalIndent(page.Content().Text, "", " ")
	// fmt.Println(string(j))
	currentLine := float64(0)
	complete_word := ""
	for _, v := range page.Content().Text {
		//fmt.Println(v.Font)
		// pdfCtx.SetFont("Arial", "", v.FontSize)
		// primero debo de armar el texto
		// pdfCtx.Cell(2, 0, v.S)
		// fmt.Println(v.S)

		if v.S != " " {
			complete_word += v.S
		} else {
			fmt.Print(complete_word, " ")
			complete_word = ""

			if currentLine != v.Y {
				fmt.Println("")
			}

			currentLine = v.Y

		}
		// pdfCtx.Cell(40, 10, "Hola")
	}
	// pdfCtx.SetFont("Arial", "", 16)

	// pdfCtx.Cell(40, 10, "AAAAAAAAAA")

	// pdfCtx.OutputFileAndClose("./test.pdf")
	// fmt.Println("END content")
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
