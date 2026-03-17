package main

import (
	"fmt"
	"strings"
	"unicode"
)

var s1 = "    Hello, World!               "
var s2 = "!!!!!Hello World!!"

// Hàm trim()
func trim() {
	// Loại bỏ khoảng trắng 2 đầu
	fmt.Println("TrimSpace:", strings.TrimSpace(s1))

	// Loại bỏ khoảng trắng bên trái
	fmt.Println("TrimLeft:", strings.TrimLeft(s1, " "))

	// Loại bỏ khoảng trắng bên phải
	fmt.Println("TrimRight:", strings.TrimRight(s1, " "))

	// Loại bỏ ký tự được chỉ định ở 2 đầu
	fmt.Println("Trim '!':", strings.Trim(s2, "!"))
}

func contains() {
	// Kiểm tra tồn tại
	fmt.Println("Contains 'World':", strings.Contains(s1, "World"))
	fmt.Println("Contains 'My':", strings.Contains(s1, "My"))

	// ContainsAny(s, chars): -> True nếu s chứa ít nhất một ký tự trong chars
	fmt.Println("ContainsAny 'Hello':", strings.ContainsAny(s1, "Hello"))
	fmt.Println("Contains 'My':", strings.Contains("We are champions", "you"))

	// Kiểm tra một chuỗi có chứa một ký tự (rune - unicode) cụ thể hay không
	fmt.Println("", strings.ContainsRune("Hello, 世界", '世'))

	// ContainsFunc(s string, f func(rune) bool): -> Kiểm tra ký tự unicode có thỏa mãn hàm nào đó
}

func count() {
	// Count(s, substr): -> Đếm số lần xuất hiện của substr trong s
	fmt.Println("Count 'l':", strings.Count(s1, "l"))
	fmt.Println("Count 'o':", strings.Count(s1, "el"))
}

func cut() {
	// Tách chuỗi tại lần xuất hiện đầu tiên của một delimiter
	s := "key=value"

	// Tìm thấy separator
	before, after, found := strings.Cut(s, "=")
	fmt.Println(before, after, found)

	// Khi không tìm thấy separator
	before1, after1, found1 := strings.Cut(s, ":")
	fmt.Println(before1, after1, found1)

	// CutPrefix(s, prefix): Cắt bỏ một prefix khỏi chuỗi nếu prefix đó tồn tại
	url := "https://www.google.com"
	after2, found2 := strings.CutPrefix(url, "https://")
	fmt.Println(after2, found2)

	// CutSuffix(s, suffix): Cắt bỏ một suffix khỏi chuỗi nếu suffix đó tồn tại
	filename := "file.txt"
	before2, found2 := strings.CutSuffix(filename, ".txt")
	fmt.Println(before2, found2)

	// EqualFold(s, t): So sánh hai chuỗi mà không phân biệt chữ hoa thường
	fmt.Println("GO=go:", strings.EqualFold("Go", "go"))
}

func splitAndFields() {
	s := "go   is   very   fast"
	split := strings.Split(s, " ") // = ["go", "", "", "is", "", "", "fast"]
	fields := strings.Fields(s)    // = ["go", "is", "very", "fast"]
	fmt.Println(split)
	fmt.Println(fields) // len(fields) -> Có thể hiểu là đếm số từ trong câu

	// FieldsFunc() -> Tách chuỗi dựa vào hàm tự định nghĩa
	s1 := "go,java;python|rust"
	result := strings.FieldsFunc(s1, func(r rune) bool {
		return r == ',' || r == ';' || r == '|'
	})
	fmt.Println("FieldsFunc:", result)

	// Tách theo mọi ký tự không phải chữ
	s2 := "go123java456python"
	words := strings.FieldsFunc(s2, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	fmt.Println("FieldsFunc:", words)

	// Tách theo số
	s3 := "abc1def2ghi3"
	result1 := strings.FieldsFunc(s3, func(r rune) bool {
		return r >= '0' && r <= '9'
	})
	fmt.Println("FieldsFunc:", result1)

	// Cắt chuỗi nhưng vẫn giữa delimiter ở cuối
	s4 := "a,b,c,d"
	fmt.Println(strings.SplitAfter(s4, ",")) // = ["a,", "b,", "c,", "d"]

	// SplitAfterN(s, ",", 3): Tương tự như SplitAfter() trả về đúng n phần tử.
}

func prefixAndSuffix() {
	s := "golang"
	fmt.Println(strings.HasPrefix(s, "go"))
	fmt.Println(strings.HasSuffix(s, "lang"))
}

func index() {
	// Index(s, substr): Trả về vị trí đầu tiên của substr trong s, nếu không tìm thấy trả về -1
	s := "Hello, World!"
	fmt.Println("Index 'l':", strings.Index(s, "l"))

	// IndexAny(s, chars): Tìm ký tự đầu tiên trong s xuất hiện trong chuỗi chars.
	fmt.Println("IndexAny:", strings.IndexAny("golang", "aeiou"))

	// IndexByte(s, c byte): Tìm kiếm c trong s.
	fmt.Println("IndexByte:", strings.IndexByte("golang", 'l'))

	// IndexFunc(): Ký tự đầu tiên thỏa mãn hàm
	// Ký tự viết hoa đầu tiên
	fmt.Println("IndexFunc:", strings.IndexFunc("helloWorld", unicode.IsUpper))

	// Tương tự có: LastIndex(), LastIndexAny(), LastIndexByte(), LastIndexFunc()
}

func join() {
	// Nối chuỗi bằng ký tự
	s := []string{"a", "b", "c"}
	join := strings.Join(s, "-")
	fmt.Println(join)
}

func stringBuilder() {
	// Strings.Builder
	var b strings.Builder
	words := []string{"go", "is", "very", "fast"}
	for _, w := range words {
		b.WriteString(w) //	Ngoài ra còn có: WriteByte(), WriteRune()
		b.WriteString(" ")
	}
	fmt.Println(b.String())
}

func stringsMap() {
	// Xử lý từng ký tự trong chuỗi theo hàm được truyền vào
	result := strings.Map(unicode.ToUpper, s2)
	fmt.Println(result)
}

func repeat() {
	s := "var"
	// Lặp chuỗi n lần
	fmt.Println(strings.Repeat(s, 3))
}

func replace() {
	s := "go go go"
	// Thay thế 1 lần
	fmt.Println("Replace:", strings.Replace(s, "go", "G_O", 1))
	// Thay thế tất cả
	fmt.Println("Replace(-1)", strings.Replace(s, "go", "G_O", -1))
	// ReplaceAll()
	date := "2026-03-16"
	fmt.Println("ReplaceAll:", strings.ReplaceAll(date, "-", "/"))
}

func toValidUTF8() {
	// Chuyển đổi chuỗi sang UTF-8 hợp lệ, thay thế các byte không hợp lệ bằng ký tự được chỉ định
	s := "Hello\xffWorld"
	result := strings.ToValidUTF8(s, "?")
	fmt.Println(result)
}

func x() {
	r := strings.NewReplacer(
		"Hello", "Hi",
		"Go", "Golang",
	)
	input := "Hello Go! Hello everyone!"
	output := r.Replace(input)
	fmt.Println(output) // Hi Golang! Hi everyone!
}

func main() {
	// trim()
	// contains()
	// count()
	// cut()
	// splitAndFields()
	// prefixAndSuffix()
	// index()
	// join()
	// stringBuilder()
	// stringsMap()
	// repeat()
	// replace()
	toValidUTF8()
}
