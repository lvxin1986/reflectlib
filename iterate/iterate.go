package iterate

import (
	"github.com/lvxin1986/reflectlib/convert"
	"golang.org/x/exp/errors/fmt"
	"log"
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
//create by bjlvxin at 12:57 for project reflectlib

//Iterate all the inferface to print all the value of interface i
func  IterateValue(i interface{})(info string){
	v:=reflect.ValueOf(i)
	t:=reflect.TypeOf(i)
	k:=t.Kind()
	if k==reflect.Ptr {
		if vv,e:= convert.Prt2Struct(i);e!=nil {
			log.Println(e)
			return
		} else {
			return IterateValue(vv)
		}
	} else  if k == reflect.Struct{
		tinfo:=fmt.Sprintln(t.Name()," ",k, "{")
		for i:=0; i<t.NumField(); i++ {
			f:=t.Field(i)
			fv:=v.Field(i)
			if fv.CanInterface() {
				tinfo = fmt.Sprintln(tinfo, f.Name, " ", f.Type, ": ", IterateValue(fv.Interface()))
			}
		}
		tinfo=fmt.Sprintln(tinfo,"}")
		return tinfo
	} else {
		return fmt.Sprintln(v)
	}
}
