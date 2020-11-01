package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

type taskList struct {
	tasks []*task
}

func (tl *taskList) agregarALista(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) EliminarDeLista(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl *taskList) mostrarLista() {
	for _, task := range tl.tasks {
		fmt.Println("NOMBRE", task.nombre)
		fmt.Println("DESCRIPCION", task.descripcion)
	}
}

func (tl *taskList) mostrarListaCompletdos() {
	for _, task := range tl.tasks {
		if task.completado {
			fmt.Println("NOMBRE", task.nombre)
			fmt.Println("DESCRIPCION", task.descripcion)
		}

	}
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go esta semana",
	}

	t2 := &task{
		nombre:      "Completar mi curso de C#",
		descripcion: "Completar mi curso de C# esta semana",
	}

	t3 := &task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python esta semana",
	}

	listaTareas := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	listaTareas.agregarALista(t3)
	listaTareas.mostrarLista()
	listaTareas.tasks[0].marcarCompleta()
	fmt.Println("-----------------------------------------------------")
	listaTareas.mostrarListaCompletdos()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Nestor"] = listaTareas

	t4 := &task{
		nombre:      "Completar mi curso de PHP",
		descripcion: "Completar mi curso de PHP esta semana",
	}

	t5 := &task{
		nombre:      "Completar mi curso de Javascript",
		descripcion: "Completar mi curso de Javascript esta semana",
	}

	listaTareas2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Pedro"] = listaTareas2

	fmt.Println("Tareas Nestor")
	mapaTareas["Nestor"].mostrarLista()
	fmt.Println("Tareas Pedro")
	mapaTareas["Pedro"].mostrarLista()

}
