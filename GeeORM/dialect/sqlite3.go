package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type sqlite3 struct{}

var _ Dialect = (*sqlite3)(nil)

func init(){
	RegisterDialect("sqlite3", &sqlite3{})
}

func(s *sqlite3) DataTypeOf(typ reflect.Value) string{
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		// use reflect.Value.Interface() to get the object
		// and use object.(Type) to judge the type of this object
		if _, ok := typ.Interface().(time.Time); ok{
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Type().Kind()))
}

func (s *sqlite3) TableExistSQL(tableName string) (string, []interface{}){
	args := []interface{}{tableName}
	return "select name from sqlite_master where type = 'table' and name = ?", args
}

