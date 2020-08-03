package render

import (
	"net/http"
	"reflect"
	"unsafe"
)

// Render インターフェースはJSON、XML、YAMLなどで実装します。
type Render interface {
	// Render 与えられたインターフェースオブジェクトをマーシャルし、カスタムContentTypeでデータを書き込みます
	Render(http.ResponseWriter) error
	// WriteContentType レスポンスにContentTypeを書き込みます
	WriteContentType(w http.ResponseWriter)
}

var (
	_ Render = JSON{}
	_ Render = JSONAscii{}
	_ Render = YAML{}
)

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

// stringToBytes はメモリー割り当てを行わずに文字列型からバイト型配列へ変換します
func stringToBytes(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}

// bytesToString はメモリー割り当てを行わずにバイト型配列から文字列型へ変換します
func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
