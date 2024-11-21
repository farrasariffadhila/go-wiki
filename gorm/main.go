package main

import (
    "fmt"
    "gorm/config"
    "gorm/models"
)

func main() {
    // Koneksi ke database
    db, err := config.ConnectDB()
    if err != nil {
        panic("Gagal terhubung ke database")
    }

    fmt.Println("Koneksi berhasil!")

    // Migrasi model ke database
    db.AutoMigrate(&models.User{})
}
