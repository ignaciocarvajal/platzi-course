package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()
	// create channel que tranmitira  un string
	channel := make(chan string)
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0

	for {
		if i > 10 {
			break
		}
		for _, servidor := range servidores {
			// se crea un go rutina
			// se pasa el canal creado como parametro a la go rutina
			go revisarServidor(servidor, channel)
		}
		time.Sleep(1 * time.Second)
		// se tiene que esperar explicitamiente a la go rutina por eso esta enmvuelto con un for infinito
		// se imprime el mensaje(string) transmitido desde revisa servidor
		fmt.Println(<-channel)
		i++
	}

	tiempoPaso := time.Since(inicio)

	fmt.Printf("tiempo de ejecucuion %s\n", tiempoPaso)

}

func revisarServidor(servidor string, channel chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		// se transmite el mensaje(sring) al canal
		channel <- servidor + "no esta disponible =("
	} else {
		channel <- servidor + "esta funcionando normalmente"
	}

}
