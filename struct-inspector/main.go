package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Name    string `json:"name"`
	Age     int8   `json:"age"`
	Email   string `json:"email"`
	Address Address
	Company Company
	Manager *Maneger
}

type Maneger struct {
	Name string
	Exp  int
}

type Address struct {
	City       string `json:"city"`
	State      string `json:"state"`
	PinCode    int32  `json:"pincode"`
	Country    string `json:"country"`
	GeoDetails GeoLocation
}

type GeoLocation struct {
	Latitude    float64
	Longitude   float64
	Coordinates Coordinates
}

type Coordinates struct {
	X float64
	Y float64
	Z float64
}

type Company struct {
	Name       string
	Department Department
}

type Department struct {
	Name   string
	Floor  int
	Office Address
}

func printMetadata(field reflect.StructField, prefix string, isLast bool) {
	branch := "├──"
	if isLast {
		branch = "└──"
	}

	fmt.Printf("%s%s %s\n", prefix, branch, field.Name)

	detailPrefix := prefix
	if isLast {
		detailPrefix += "    "
	} else {
		detailPrefix += "│   "
	}

	fmt.Printf("%s├── Type        : %s\n", detailPrefix, field.Type)
	fmt.Printf("%s├── Kind        : %s\n", detailPrefix, field.Type.Kind())
	fmt.Printf("%s├── Package     : %s\n", detailPrefix, field.PkgPath)
	fmt.Printf("%s├── Exported    : %t\n", detailPrefix, field.IsExported())
	fmt.Printf("%s├── Anonymous   : %t\n", detailPrefix, field.Anonymous)
	fmt.Printf("%s├── Index       : %v\n", detailPrefix, field.Index)
	fmt.Printf("%s├── Offset      : %d bytes\n", detailPrefix, field.Offset)
	fmt.Printf("%s├── Align       : %d\n", detailPrefix, field.Type.Align())
	fmt.Printf("%s├── Field Align : %d\n", detailPrefix, field.Type.FieldAlign())
	fmt.Printf("%s├── Size        : %d bytes\n", detailPrefix, field.Type.Size())

	if field.Tag != "" {
		fmt.Printf("%s└── Tags        : %s\n", detailPrefix, field.Tag)
	} else {
		fmt.Printf("%s└── Tags        : No tags\n", detailPrefix)
	}
}

func getNestedStruct(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Struct {
		return t
	}

	return nil
}

func inspectStruct(t reflect.Type, depth int, visited map[reflect.Type]bool) {
	if t.Kind() != reflect.Struct {
		return
	}

	indent := strings.Repeat("    ", depth)

	fmt.Println()
	fmt.Printf("%s╔══════════════════════════════════════════════════╗\n", indent)
	fmt.Printf("%s║              STRUCT: %-25s ║\n", indent, t.Name())
	fmt.Printf("%s╚══════════════════════════════════════════════════╝\n", indent)

	fmt.Printf("%sFields: %d | Size: %d bytes | Align: %d\n",
		indent, t.NumField(), t.Size(), t.Align())

	if visited[t] {
		fmt.Printf("%s[Already visited: recursive reference]\n", indent)
		return
	}

	visited[t] = true
	defer delete(visited, t)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		isLast := i == t.NumField()-1

		printMetadata(field, indent, isLast)

		nestedType := getNestedStruct(field.Type)

		if nestedType != nil {
			fmt.Printf("\n%s    ▼ NESTED STRUCT\n", indent)

			if visited[nestedType] {
				fmt.Printf("%s    ↳ Recursive reference to %s\n",
					indent, nestedType.Name())
				continue
			}

			inspectStruct(nestedType, depth+1, visited)
		}
	}
}

func main() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════╗")
	fmt.Println("║                                                      ║")
	fmt.Println("║             GO REFLECTION INSPECTOR                  ║")
	fmt.Println("║                                                      ║")
	fmt.Println("╚══════════════════════════════════════════════════════╝")

	runtimeType := reflect.TypeFor[User]()

	inspectStruct(runtimeType, 0, make(map[reflect.Type]bool))

	fmt.Println("\n========================================================")
	fmt.Println("              INSPECTION COMPLETED")
	fmt.Println("========================================================")
}
