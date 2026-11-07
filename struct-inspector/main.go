package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name    string `json:"name"`
	Age     int8   `json:"age"`
	Email   string `json:"email"`
	Address Address
}

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	PinCode int32  `json:"pincode"`
	Country string `json:"country"`
}

func printMetadata(field reflect.StructField, structName string) {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════╗")
	fmt.Printf(" ║                %s-STRUCT FIELD                   ║\n", structName)
	fmt.Println("╚══════════════════════════════════════════════════╝")

	fmt.Printf("  Name          : %s\n", field.Name)
	fmt.Printf("  Type          : %s\n", field.Type)
	fmt.Printf("  Kind          : %s\n", field.Type.Kind())
	fmt.Printf("  Package       : %s\n", field.PkgPath)
	fmt.Printf("  Anonymous     : %t\n", field.Anonymous)
	fmt.Printf("  Exported      : %t\n", field.IsExported())
	fmt.Printf("  Index         : %v\n", field.Index)
	fmt.Printf("  Offset        : %d bytes\n", field.Offset)
	fmt.Printf("  Align         : %d\n", field.Type.Align())
	fmt.Printf("  Field Align   : %d\n", field.Type.FieldAlign())
	fmt.Printf("  Size          : %d bytes\n", field.Type.Size())
}

func runtimeMetadata(t reflect.Type) {
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			if field.Type.Kind() == reflect.Struct {
				runtimeMetadata(field.Type)
			} else {
				printMetadata(field, t.Name())
			}
		}
	}
}

func main() {
	u := User{
		Name:  "Sahil",
		Age:   18,
		Email: "test@gmail.com",
		Address: Address{
			City:    "Kalewal Beet",
			State:   "Punjab",
			PinCode: 144523,
			Country: "India",
		},
	}

	runtimeType := reflect.TypeOf(u)
	runtimeMetadata(runtimeType)
}
