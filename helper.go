package reflecthelper

import (
    "fmt"
    "reflect"
)

func Join(destination interface{}, source interface{}) interface{} {
    if source == destination {
        return destination 
    }
    td := reflect.TypeOf(destination)
    ts := reflect.TypeOf(source)

    if td != ts || td.Kind() != reflect.Ptr {
        panic("Can't join different types OR non pointers")
    }

    tdValue := reflect.ValueOf(destination)
    tsValue := reflect.ValueOf(source)


    for i := 0; i < td.Elem().NumField(); i++ {
        fSource := tsValue.Elem().Field(i)
        fDest := tdValue.Elem().Field(i)

        if fDest.CanSet(){
            switch fSource.Kind() {
                case reflect.Int:
                    if fDest.Int() == 0 {
                        fDest.SetInt(fSource.Int())
                    }
                case reflect.Bool: 
                    if fDest.Bool() == false {
                        fDest.SetBool(fSource.Bool())    
                    }
                case reflect.String: 
                    if fDest.String() == "" && fSource.String() != "" {
                        fDest.SetString(fSource.String())
                    }
                case reflect.Slice:
                    fDest.Set(reflect.AppendSlice(fDest, fSource))
                case reflect.Map:
                    if fDest.IsNil(){
                        fDest.Set(reflect.MakeMap(fDest.Type()))
                    }
                    for _, key := range fSource.MapKeys() {
                        fDest.SetMapIndex(key, fSource.MapIndex(key))
                    }
                default:
                    fmt.Println(fSource.Kind())
            }
        } else {
            fmt.Println("Can't set", tdValue.Field(i))
        }
    }

    return destination
}