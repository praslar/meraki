package main

import (
	"unicode"
)

func main() {
	// a-z ; B xóa; C bật/tắt caplocks
	// input: string

	// loop qua từng ký tự của input và add vào output
	// nếu gặp ký tự B
	// 1. không add vào output
	// 2. xóa kí tự cuối cùng của output nếu len của out > 0

	// nếu gặp ký tự C
	// 1. đổi trạng thái của capslock (false -> true; true -> false)
	// 2. Không add vào out

	// Có một biến capslock := false
	// trước khi add bất kỳ ký tự nào vào output thì phải kiểm tra caplock đang = true/false
	// nếu capslock = true thì viết hoa (in hoa kí tự xong rồi add kí tự vào out)
	// nếu capslock = false thì không viết hoa (add kí tự vào out)
	//
	//out := Solution("aaBB")
	//fmt.Println(out)

	//kyTuXoa := []rune("B")

}

func ConvertToRune(in string) rune {
	return []rune(in)[0]
}

func Solution(input string) string {
	Kytuxoa := ConvertToRune("B")
	KyTuInHoa := ConvertToRune("C")
	var output []rune
	capslock := false
	for _, value := range input {
		if value == Kytuxoa {
			if len(output) > 0 {
				output = output[:len(output)-1]
			}
		}
		if value == KyTuInHoa {
			capslock = !capslock
		}
		if value != KyTuInHoa && value != Kytuxoa {
			if capslock == true {
				value = unicode.ToUpper(value)
			}
			output = append(output, value)
		}
	}
	return string(output)
}
