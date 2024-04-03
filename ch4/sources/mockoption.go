package main

import "fmt"

type record struct {
	name    string
	gender  string
	age     uint16
	city    string
	country string
}

func enroll(args ...any) (*record, error) {
	if len(args) > 5 || len(args) < 3 {
		return nil, fmt.Errorf("the number of arguments passed is wrong")
	}

	r := &record{
		city:    "Shanghai",
		country: "China",
	}

	for i, v := range args {
		switch i {
		case 0:
			name, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("name is not passed as string")
			}
			r.name = name
		case 1:
			gender, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("gender is not passed as string")
			}
			r.gender = gender
		case 2:
			age, ok := v.(int)
			if !ok {
				return nil, fmt.Errorf("age is not passed as int")
			}
			r.age = uint16(age)
		case 3:
			city, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("city is not passed as string")
			}
			r.city = city
		case 4:
			country, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("contry is not passed as string")
			}
			r.country = country
		default:
			return nil, fmt.Errorf("unknown argument passed")
		}
	}

	return r, nil
}

func main() {
	r, _ := enroll("Emory.Du", "Male", 24)
	fmt.Printf("%+v\n", *r)

	r, _ = enroll("H", "Female", 24, "Hangzhou")
	fmt.Printf("%+v\n", *r)
}
