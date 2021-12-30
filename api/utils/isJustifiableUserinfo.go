package utils

import (
	"unicode"
	"unicode/utf8"
)

// IsJustifiableUserinfo(private)
// 機能 : ID，パスワードの正当性チェック
// 引数 :
//  * id   : 検証したいID
//  * pass : 検証したいパスワード
// 戻り値 : 正しいID，パスワードならばtrue
func IsJustifiableUserinfo(id, pass string) bool {
	// ASCIIかどうかのチェック
	if !(utf8.ValidString(id) && utf8.RuneCountInString(id) == len(id)) ||
		!(utf8.ValidString(pass) && utf8.RuneCountInString(pass) == len(pass)) {
		return false
	}

	// idに空白文字もしくは制御文字があるかどうかのチェック
	for _, c := range id {
		if c == ' ' {
			return false
		}
		if unicode.IsControl(c) {
			return false
		}
	}

	// パスワードに空白文字もしくは制御文字があるかどうかのチェック
	for _, c := range pass {
		if c == ' ' {
			return false
		}
		if unicode.IsControl(c) {
			return false
		}
	}

	return true
}