package src

import "fmt"

func Maps() {
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
	fmt.Println(sammy)
	fmt.Println(sammy["color"])
}
