package session

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alexedwards/scs/v2"
)

func TestSession_InitSession(t *testing.T) {
	s := &Session{
		CookieLifetime: "100",
		CookiePersist:  "true",
		CookieName:     "celeritas",
		CookieDomain:   "localhost",
		SessionType:    "cookie",
	}

	var sm *scs.SessionManager

	ses := s.InitSession()

	var sessKind reflect.Kind
	var sessType reflect.Type

	rv := reflect.ValueOf(ses)

	for rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
		fmt.Println("For loop: ", rv.Kind(), rv.Type(), rv)
		sessKind = rv.Kind()
		sessType = rv.Type()

		rv = rv.Elem()
	}

	//Check its valid
	if !rv.IsValid() {
		t.Error("invalid type of kind; kind: ", rv.Kind(), " type: ", rv.Type())
	}

	if sessKind != reflect.ValueOf(sm).Kind() {
		t.Error("Wrong kind returned, testing cookie session. \nExpected: ", reflect.ValueOf(sm).Kind(), "\nGot: ", sessKind)
	}

	if sessType != reflect.ValueOf(sm).Type() {
		t.Error("Wrong type returned, testing cookie session. \nExpected: ", reflect.ValueOf(sm).Type(), "\nGot: ", sessType)
	}
}
