package convert

import (
	"github.com/pkg/errors"
	"reflect"
)

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


//将指针类型的变量转换成实际对象类型的变量
func Prt2Struct(i interface{}) (ri interface{},err error){
	if IsPtr(i) {
		return reflect.ValueOf(i).Elem().Interface(),nil
	} else {
		return nil, errors.New("the input value must be ptr!")
	}
}

//whether the i is a pointer
func IsPtr(i interface{})(is bool){
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		return true
	} else {
		return false
	}
}
