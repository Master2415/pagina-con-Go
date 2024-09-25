package main

import (
    "html/template"
    
    "net/http"
    
    "os"
    "fmt"
)

type PageData struct {
    HostName   string
    RandomImgs []string
    Name       string
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Servidor corriendo en http://localhost:8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        fmt.Println("Error al iniciar el servidor:", err)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))

    // Obtener el nombre del host
    hostName, err := os.Hostname()
    if err != nil {
        hostName = "Desconocido"
    }

    // Obtener 4 imágenes aleatorias codificadas en base64
    randomImages := getRandomImagesBase64(4)

    // Datos para pasar a la plantilla
    data := PageData{
        HostName:   hostName,
        RandomImgs: randomImages,
        Name:"Douglas - Erick - jhojan",
    }

    tmpl.Execute(w, data)
}
