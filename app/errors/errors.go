package errors

import (
	"errors"
	"fmt"

	errs "github.com/pkg/errors"
)

type Code string
type Err string

const (
	ErrHoge     Code = "err Hoge"       // sample error
	ErrUnknown  Code = "unknown error"  // unknown error
	ErrDataBase Code = "database error" // database error
	ErrNewUUID  Code = "uuid error"     // uuid error
)

type Error struct {
	Code Code
	Err  error
}

// カスタムエラーのErrorメソッド
func (e *Error) Error() string {
	return fmt.Sprintf("Code: %s, Msg: %s", e.Code, e.Err)
}

// カスタムエラーの生成
func SetError(c Code, err string) error {
	return &Error{
		Code: c,
		Err:  errs.New(err), // github.com/pkg/errorsでラップする
	}
}

// スタックトレースを返す
func StackTrace(err error) string {
	var e *Error
	if errors.As(err, &e) {
		return fmt.Sprintf("%+v\n", e.Err)
	}

	return ""
}

// Codes はエラーコードを返す
// カスタムエラーとして定義されていない場合は、Unknownを返す
func Codes(err error) Code {
	var e *Error
	if errors.As(err, &e) {
		switch e := err.(type) {
		case *Error:
			return e.Code
		}
	}
	// カスタムエラーとして定義されていない場合はUnknownを返す
	return ErrUnknown
}
