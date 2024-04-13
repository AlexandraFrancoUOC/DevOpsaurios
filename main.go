package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Escribir el contenido HTML con la etiqueta de imagen
	htmlContent := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>¡Hola!</title>
		</head>
		<body>
			<p>Soy alumno de la UOC</p>
			<img src="/prueba.png" alt="Imagen UOC">
		</body>
		</html>
	`

	// Escribir el contenido HTML en la respuesta
	w.Write([]byte(htmlContent))
}

// imageHandler servirá la imagen cuando se acceda a la ruta /prueba.png
func imageHandler(w http.ResponseWriter, r *http.Request) {
	imagePath := "D:/UOC/DevOps/prueba.png" //Ruta de la Imagen
	http.ServeFile(w, r, imagePath)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/prueba.png", imageHandler)

	// Iniciar el servidor
	log.Println("Servidor corriendo en http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
