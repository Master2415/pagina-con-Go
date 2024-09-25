package main

import (
    "encoding/base64"
    "io/ioutil"
    "math/rand"
    "os"
    "time"
)

func getRandomImagesBase64(n int) []string {
    var base64Images []string
    images := getImagesFromDir("imagenes")

    rand.Seed(time.Now().UnixNano())

    // Seleccionar `n` imágenes aleatorias
    for i := 0; i < n && i < len(images); i++ {
        randIdx := rand.Intn(len(images))
        imgBase64 := encodeImageToBase64(images[randIdx])
        base64Images = append(base64Images, imgBase64)
    }

    return base64Images
}

// Lee imágenes desde un directorio específico
func getImagesFromDir(dir string) []string {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        panic(err)
    }

    var imageFiles []string
    for _, file := range files {
        if !file.IsDir() && (file.Name()[len(file.Name())-4:] == ".jpg" || file.Name()[len(file.Name())-5:] == ".jpeg" || file.Name()[len(file.Name())-4:] == ".png") {
            imageFiles = append(imageFiles, dir+"/"+file.Name())
        }
    }
    return imageFiles
}

// Codifica una imagen a base64
func encodeImageToBase64(imagePath string) string {
    file, err := os.Open(imagePath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        panic(err)
    }

    return base64.StdEncoding.EncodeToString(fileBytes)
}
