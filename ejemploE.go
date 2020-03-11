package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "bufio"
)

type usuario struct {
    Cali, nombre string
}

func main() {
    menu := `¿Qué quieres hacer?
[1] -- Capturar
[2] -- Mostrar
[3] -- Salir
`
    var eleccion int
    
    for eleccion != 3 {
        fmt.Print(menu)
        fmt.Scanln(&eleccion)
        scanner := bufio.NewScanner(os.Stdin)
        switch eleccion {
        case 1:
            var Usuarios []usuario
            for i:= 0; i<5; i++ {
            var c usuario
            fmt.Println("Ingresa el nombre:")
            if scanner.Scan() {
                c.nombre = scanner.Text()
            }
            fmt.Println("Ingresa las Calificaciones: ")
            if scanner.Scan() {
                c.Cali = scanner.Text()
            }
        
            Usuarios = append(Usuarios,c)
        }
            
        for i:= 0; i<len(Usuarios); i++ {
            InsertarDatos(Usuarios[i].nombre, Usuarios[i].Cali)
        }
        
        case 2:
            
            data, err := ioutil.ReadFile("text.txt")
            if err != nil {
                fmt.Println(err)
            }

    
            fmt.Println("El archivo contiene \n " + string(data))

            
        
            
            
        }
    }
}
func InsertarDatos (nombre string, Cali string) {
    f, err := os.OpenFile("text.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) 
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    if _, err := f.WriteString(nombre+" "+Cali+"\n"); err != nil {
        fmt.Println(err)
    }

    }

