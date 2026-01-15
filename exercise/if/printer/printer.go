package printer

import "fmt"

type FormatType string

const (
	Line   FormatType = "line"
	Format FormatType = "format"
)

func Print(x string, format FormatType) {
	switch format {
	case Line:
		fmt.Println(x)
	case Format:
		fmt.Printf("format time: %s\n", x)
	default:
		fmt.Println(x)
	}
}
