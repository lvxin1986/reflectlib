package convert

import (
	"reflect"
)
import "errors"

// Copyright 2019 The ChuBao Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.
//create by bjlvxin at 17:33 for project reflectlib


//StructCopy copy the exported value of a struct to a likely struct , with reflect.src and dst must be ptr
func StructCopy(src, dst interface{}) error {
	if !IsPtr(src)||!IsPtr(dst) {
		return  errors.New("type error: source and target must be a pointer!")

	}
	srcKeys := make(map[string]bool)
	srcV:=reflect.ValueOf(src).Elem()
	for i := 0; i < srcV.NumField(); i++ {
		srcKeys[srcV.Type().Field(i).Name] = true
	}
	dstV:=reflect.ValueOf(dst).Elem()
	for i := 0; i < dstV.NumField(); i++ {
		fName := dstV.Type().Field(i).Name
		if _, ok := srcKeys[fName]; ok {
			fv := srcV.FieldByName(fName)
			if fv.CanInterface() && dstV.Field(i).CanSet(){
				dstV.Field(i).Set(fv)
			}
		}
	}

	return nil
}

//copy the map to the struct
func Map2Struct(data map[string]interface{}, inStructPtr interface{}) ( e error){
	if !IsPtr(inStructPtr) {
		return errors.New("type error: target must be a pointer!")

	}
	rType := reflect.TypeOf(inStructPtr).Elem()
	rVal := reflect.ValueOf(inStructPtr).Elem()
	// 遍历结构体
	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)
		f := rVal.Field(i)
		if !f.CanSet(){
			continue
		}
		// 得到tag中的字段名
		key := t.Tag.Get("key")
		if v, ok := data[key]; ok {
			// 检查是否需要类型转换
			dataType := reflect.TypeOf(v)
			filedType := f.Type()
			if filedType == dataType {
				f.Set(reflect.ValueOf(v))
			} else {
				if dataType.ConvertibleTo(filedType) {
					// 转换类型
					f.Set(reflect.ValueOf(v).Convert(filedType))
				} else {
					panic(t.Name + " type mismatch")
				}
			}
		} else {
			continue
		}
	}
	return nil
}