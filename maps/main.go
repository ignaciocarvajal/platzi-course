package main

import "fmt"

func main() {
	map1 := make(map[string]int)

	map1["a"] = 8

	fmt.Println(map1["a"])
}

// Apuntes

// lo mismo que el range
// for i := 0; i < len(listaTareas.tasks); i++ {
// 	fmt.Println("INDEX", i, "NOMBRE", listaTareas.tasks[i].nombre)
// }

// for i, task := range listaTareas.tasks {
// 	fmt.Println("INDEX", i, "NOMBRE", task.nombre)
// }
// Break
// for i := 0; i < 10; i++ {
// 	if i == 5 {
// 		break
// 	}
// 	fmt.Println(i)
// }
// Continue
// for i := 0; i < 10; i++ {
// 	if i == 5 {
// 		continue
// 	}
// 	fmt.Println(i)
// }
