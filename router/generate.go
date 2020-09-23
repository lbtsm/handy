package router

import (
	"reflect"
	"runtime"
	"strings"
)

const (
	upAOfASCII          = 65
	upZOfASCII          = 90
	num0OfASCII         = 48
	num9OfASCII         = 57
	caseIntervalOfASCII = 32
)

var (
	routerVersion = "v1.0"
	projectName   = ""
)

func SetProjectName(name string) {
	projectName = name
}

func GenerateUrlPathByFunc(f interface{}, prefix ...string) string {
	if projectName == "" {
		panic("projectName is empty by generate url path")
	}
	if f == nil {
		panic("receive param is nil")
	}
	fk := reflect.TypeOf(f).Kind()
	if fk != reflect.Func {
		panic("receive first param is must be func")
	}
	fullPath := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	// split by .
	funcPaths := strings.Split(fullPath, ".")
	funcName := funcPaths[len(funcPaths)-1]
	// generate url path
	urlPath := strings.Builder{}
	for i := 0; i < len(funcName); i++ {
		w := funcName[i]
		var isNum, isUp bool
		if w >= upAOfASCII && w <= upZOfASCII {
			isUp = true
		}
		if w >= num0OfASCII && w <= num9OfASCII {
			isNum = true
		}
		if (isNum || isUp) && urlPath.Len() != 0 {
			urlPath.WriteByte('-')
		}

		if isUp {
			w += caseIntervalOfASCII
		}

		urlPath.WriteByte(w)
	}

	paths := make([]string, 0, 3)
	paths = append(paths, projectName)
	if len(prefix) > 0 {
		paths = append(paths, prefix[0])
	} else {
		paths = append(paths, routerVersion)
	}
	paths = append(paths, urlPath.String())
	return strings.Join(paths, "/")
}
