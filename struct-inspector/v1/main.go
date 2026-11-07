package v1

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Name  string `json:"name" validate:"required,max=2,min=2"`
	Email string `json:"email" validate:"required,email" custom:"unique"`
	Age   int    `json:"age" validate:"required,min=18"`
}

func yield(field reflect.StructField, fieldValue reflect.Value) {
	fieldTag := field.Tag
	kvMpp := make(map[string][]string)

	tagString := string(fieldTag)

	if tagString != "" {
		for _, tag := range strings.Split(tagString, " ") {
			if tag == "" {
				continue
			}

			parts := strings.SplitN(tag, ":", 2)
			if len(parts) != 2 {
				continue
			}

			key := parts[0]
			value := strings.Trim(parts[1], `"`)

			kvMpp[key] = strings.Split(value, ",")
		}
	}

	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════╗")
	fmt.Println("║                 STRUCT FIELD                     ║")
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

	fmt.Println("\n  STRUCT TAGS")
	fmt.Println("  ─────────────────────────────────────────────────")

	if len(kvMpp) == 0 {
		fmt.Println("  No tags defined")
	}

	for key, values := range kvMpp {
		fmt.Printf("  %-12s : %s\n", key, strings.Join(values, ", "))
	}

	fmt.Println("────────────────────────────────────────────────────")

	fmt.Print("  Runtime Value : ")

	if fieldValue.IsValid() && fieldValue.CanInterface() {
		fmt.Printf("%v\n", fieldValue.Interface())
	} else {
		fmt.Println("<not available - type metadata only>")
	}

	fmt.Println("────────────────────────────────────────────────────")
}

func test(v reflect.Type, instance reflect.Value) {
	fmt.Println("\nInspecting:", v)

	for i := 0; i < v.NumField(); i++ {
		var fieldValue reflect.Value

		if instance.IsValid() {
			fieldValue = instance.Field(i)
		}

		yield(v.Field(i), fieldValue)
	}
}

func main() {
	u := User{
		Name:  "Sahil",
		Age:   16,
		Email: "test@gmail.com",
	}

	// 1. Compile-time type specification. No instance required.
	fmt.Println("\n========== TYPE FOR ==========")
	test(
		reflect.TypeFor[User](),
		reflect.Value{},
	)

	// 2. Runtime type inspection with an actual instance.
	fmt.Println("\n========== TYPE OF ==========")
	test(
		reflect.TypeOf(u),
		reflect.ValueOf(u),
	)
}
