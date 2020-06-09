package main

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
//create by bjlvxin at 17:31 for project reflectlib

import (
	"github.com/lvxin1986/reflectlib/convert"
	"github.com/lvxin1986/reflectlib/iterate"
	"log"
)

var data = map[string]interface{}{
	"ID":    1001,
	"name":  "apple",
	"price": 16,
	"aa":"aaaa",
}

type Fruit struct {
	ID    int     `key:"id"`
	Name  string  `key:"name"`
	Price float64 `key:"price"`
	aa string
}

type At struct {
	Nm string
	Pr float64
}
type Fruit1 struct {
	ID    int     `key:"id"`
	Name  string  `key:"name"`
	Price float64 `key:"price"`
	Tt  *string
	Aa *At
}

type Fruit2 struct {
	ID    int     `key:"id"`
	Name  string  `key:"name"`
	price float64 `key:"price"`
	Tt *string
	Aa *At
}

func newFruit(data map[string]interface{}) *Fruit {
	s := Fruit{
		ID:    data["id"].(int),
		Name:  data["name"].(string),
		Price: data["price"].(float64),
	}
	return &s
}

type Persion struct {
	Name string
	age int
	Sex  *string
}


func main() {
	//usage
	fruit := Fruit{}

	convert.Map2Struct(data, &fruit)
	log.Println("fruit:", fruit)
	log.Println("-----------------------------")
	fruit2 := Fruit2{}
	tt:="test"
	at:=At{"aa1",11.2}
	fruit1 := Fruit1{1,"aaa",12.1,&tt,&at}
	convert.StructCopy(&fruit1,&fruit2)
	log.Print("aaaaa:",fruit2)
	log.Print("xxxx:",*fruit1.Aa)
	log.Print("xxxx:",*fruit2.Aa)
	log.Println("-------------------")
	s:="man"
	p:=Persion{"lvxin",28,&s}
	log.Println(iterate.IterateValue(&p))
	log.Println(iterate.IterateValue(&fruit2))

}