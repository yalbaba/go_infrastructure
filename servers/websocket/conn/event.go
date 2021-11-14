package conn

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	jsoniter "github.com/json-iterator/go"
	"liveearth/infrastructure/pkg/errno"
	"reflect"
)

type Event struct {
	EventName string `json:"event_name"`
	Body      string `json:"body"`
}

func WriteRequest(data []byte) (Event, error) {
	e := Event{}
	err := jsoniter.Unmarshal(data, &e)
	return e, err
}

type Response struct {
	ErrCode   int         `json:"errcode"`
	ErrMsg    string      `json:"errmsg"`
	EventName string      `json:"event_name"`
	Data      interface{} `json:"data"`
}

// https://github.com/asaskevich/govalidator
// 字段必填   valid:"required"
// 邮箱   	valid:"email"
// 范围  	valid:"range(min|max)"
// byte长度  valid:"length(min|max)"
// rune长度  valid:"runelength(min|max)"
// string长度 valid:"stringlength(min|max)"
// in  valid:"in(string1|string2|...|stringN)"
func (e *Event) Bind(obj interface{}) error {
	err := jsoniter.UnmarshalFromString(e.Body, obj)
	if err != nil {
		return errno.New(errno.ErrBind, fmt.Errorf("参数格式错误 err: %v,body: %s", err, e.Body))
	}

	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil
	}
	if _, err := govalidator.ValidateStruct(obj); err != nil {
		return errno.New(errno.ErrParam, fmt.Errorf("输入参数有误 %v", err))
	}

	return nil
}

func WriteResponse(event string, result interface{}) *Response {

	r := []byte("{}")
	var code int
	message := "OK"

	if v, ok := result.(error); ok {
		code, message = errno.DecodeErr(v)
		return &Response{
			ErrCode:   code,
			ErrMsg:    message,
			EventName: event,
			Data:      r,
		}
	}

	if result != nil {
		r, _ = jsoniter.Marshal(result)
	}

	return &Response{
		ErrCode:   code,
		ErrMsg:    message,
		EventName: event,
		Data:      r,
	}
}
