package main

import (
	"encoding/json"
	"fmt"
	"github.com/devfeel/mapper"
	"github.com/fatih/color"
	"time"
)

type (
	User struct {
		Name string
		Age  int
		Id   string `mapper:"_id"`
		AA   string `json:"Score"`
		Time time.Time
	}

	Student struct {
		Name  string
		Age   int
		Id    string `mapper:"_id"`
		Score string
	}

	Teacher struct {
		Name  string
		Age   int
		Id    string `mapper:"_id"`
		Level string
	}
)

func init() {
	mapper.Register(&User{})
	mapper.Register(&Student{})
}

func main() {
	user := &User{}
	userMap := &User{}
	teacher := &Teacher{}
	student := &Student{Name: "test", Age: 10, Id: "testId", Score: "100"}
	valMap := make(map[string]interface{})
	valMap["Name"] = "map"
	valMap["Age"] = 10
	valMap["_id"] = "x1asd"
	valMap["Score"] = 100
	valMap["Time"] = time.Now()

	// mapper.Mapper(student, user)
	// mapper.AutoMapper(student, teacher)
	// mapper.MapperMap(valMap, userMap)
	JSONCopy(user, student)

	fmt.Println("student:", student)
	fmt.Println("user:", user)
	fmt.Println("teacher", teacher)
	fmt.Println("userMap:", userMap)
}

// JSONCopy 用b的所有字段覆盖a的
// a应该为结构体指针
func JSONCopy(a interface{}, b interface{}) {
	val, e := json.Marshal(b)
	if e != nil {
		color.Red(e.Error())
	} else {
		json.Unmarshal(val, a)
	}
}
