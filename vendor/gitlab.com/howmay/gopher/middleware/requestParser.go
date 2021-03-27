package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"

	"gitlab.com/howmay/gopher/errors"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// APIKey is a helper to validate request
type APIKey struct {
	header      interface{}
	params      interface{}
	query       interface{}
	body        interface{}
	wantRawBody bool
}

// NewAPIKey new a apiKey object
func NewAPIKey() *APIKey {
	return &APIKey{}
}

// WantRawBody 之後需要取得 raw body
func (a *APIKey) WantRawBody() {
	a.wantRawBody = true
}

// BindHeader 檢查 header
func (a *APIKey) BindHeader(t interface{}) error {
	if t != nil {
		v := reflect.Indirect(reflect.ValueOf(t))
		count := v.NumField()
		for i := 0; i < count; i++ {
			f := v.Field(i)
			switch f.Kind() {
			case reflect.String:
			case reflect.Int:
			case reflect.Float32, reflect.Float64:
			default:
				return errors.NewWithMessage(errors.ErrInvalidInput, "type error")
			}
		}
		a.header = t
	}
	return nil
}

// BindQuery 檢查 query
func (a *APIKey) BindQuery(t interface{}) error {
	if t != nil {
		a.query = t
	}
	return nil
}

// BindParmas 檢查 params
func (a *APIKey) BindParmas(t interface{}) error {
	if t != nil {
		v := reflect.Indirect(reflect.ValueOf(t))
		count := v.NumField()
		for i := 0; i < count; i++ {
			f := v.Field(i)
			switch f.Kind() {
			case reflect.String:
			case reflect.Int:
			case reflect.Float32, reflect.Float64:
			default:
				return errors.NewWithMessage(errors.ErrInvalidInput, "type error")
			}
		}
		a.params = t
	}
	return nil
}

// BindBody 檢查 body
func (a *APIKey) BindBody(t interface{}) {
	if t != nil {
		a.body = t
	}
}

// CheckBind 檢查參數
func (a *APIKey) CheckBind() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			out := ""
			if a.header != nil {
				_a := newStruct(a.header)
				err = checkAndBuildSturct("header", c.Request().Header.Get, _a, c)
				if err != nil {

					errors.HTTPErrorHandlerForEcho(errors.NewWithMessage(errors.ErrInternalError, err.Error()), c)
					return
				}
				c.Set("header", _a)
				b, _ := json.Marshal(_a)
				if out != "" {
					out += "\t"
				}
				out += fmt.Sprintf("Header= %+v", string(b))
			}
			if a.params != nil {
				_a := newStruct(a.params)
				err = checkAndBuildSturct("parameter", c.Param, _a, c)
				if err != nil {
					errors.HTTPErrorHandlerForEcho(errors.NewWithMessage(errors.ErrInvalidAmount, err.Error()), c)
					return
				}
				c.Set("params", _a)
				b, _ := json.Marshal(_a)
				if out != "" {
					out += "\t"
				}
				out += fmt.Sprintf("Params= %+v", string(b))
			}

			if a.query != nil {
				_a := newStruct(a.query)

				err = checkAndBuildSturct("query", c.Param, _a, c)
				if err != nil {
					errors.HTTPErrorHandlerForEcho(errors.NewWithMessage(errors.ErrInvalidInput, err.Error()), c)
					return
				}
				c.Set("query", _a)
				b, _ := json.Marshal(_a)
				if out != "" {
					out += "\t"
				}
				out += fmt.Sprintf("Query= %+v", string(b))
			}

			if a.body != nil {
				if a.wantRawBody {
					rawBytes, _ := ioutil.ReadAll(c.Request().Body)
					c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(rawBytes))
					c.Set("rawBody", string(rawBytes))
				}
				body := newStruct(a.body)
				err = c.Bind(body)
				if err != nil {
					errors.HTTPErrorHandlerForEcho(errors.NewWithMessage(errors.ErrInvalidInput, err.Error()), c)
					return
				}
				c.Set("body", body)
				b, _ := json.Marshal(body)
				if out != "" {
					out += "\t"
				}
				out += fmt.Sprintf("Body= %+v", string(b))
			}
			if out != "" {
				//			logger := GetLogger(c)
				ctx := c.Request().Context()
				logger := log.Ctx(ctx)
				logger.Info().Msg(out)
			}

			return next(c)
		}
	}
}

func checkAndBuildSturct(part string, getValue func(string) string, _a interface{}, c echo.Context) error {

	v := reflect.Indirect(reflect.ValueOf(_a))
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		n := v.Type().Field(i)
		field := ""
		if n.Tag.Get("json") != "" {
			field = n.Tag.Get("json")
		} else {
			field = n.Name
		}

		r := getValue(field)
		if r == "" && n.Tag.Get("binding") == "required" {
			return errors.NewWithMessage(errors.ErrInvalidInput, part+" '"+field+"' is missing.")
		} else if r != "" {
			if reflect.TypeOf(r).Kind() != f.Kind() {
				if reflect.TypeOf(r).ConvertibleTo(f.Type()) {
					v := reflect.ValueOf(r).Convert(f.Type())
					f.Set(v)
				} else {
					switch f.Kind() {
					case reflect.Int:
						i, err := strconv.ParseInt(r, 10, 0)
						if err != nil {
							return errors.NewWithMessage(errors.ErrInvalidInput, part+" '"+field+"' invalid value: "+err.Error())
						}
						f.Set(reflect.ValueOf(int(i)))
					case reflect.Float32, reflect.Float64:
						i, err := strconv.ParseFloat(r, 0)
						if err != nil {
							return errors.NewWithMessage(errors.ErrInvalidInput, part+" '"+field+"' invalid value: "+err.Error())

						}
						f.Set(reflect.ValueOf(float32(i)))
					}
				}
			} else {
				f.Set(reflect.ValueOf(r))
			}
		}
	}
	return nil
}

func newStruct(i interface{}) (s interface{}) {
	return reflect.New(
		reflect.Indirect(
			reflect.ValueOf(i),
		).Type(),
	).Interface()
}

// BodyUnmarshal unmarshal body
func BodyUnmarshal(c echo.Context, t interface{}) error {
	s := c.Get("body")
	return unmarshal(s, t)
}

// HeaderUnmarshal unmarshal header
func HeaderUnmarshal(c echo.Context, t interface{}) error {
	s := c.Get("header")
	return unmarshal(s, t)
}

// ParamsUnmarshal unmarshal params
func ParamsUnmarshal(c echo.Context, t interface{}) error {
	s := c.Get("params")
	return unmarshal(s, t)
}

// QueryUnmarshal unmarshal query
func QueryUnmarshal(c echo.Context, t interface{}) error {
	s := c.Get("query")
	return unmarshal(s, t)
}

func unmarshal(s, d interface{}) error {
	_s := reflect.Indirect(reflect.ValueOf(s))
	if _s.Type() == reflect.Indirect(reflect.ValueOf(d)).Type() {
		reflect.Indirect(reflect.ValueOf(d)).Set(_s)
	} else {
		d = nil
		return errors.NewWithMessage(errors.ErrInvalidInput, "can not unmarshal")
	}
	return nil
}
