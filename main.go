package main

import (
	"fmt"
	"latihan-gin/database"
	"latihan-gin/router"
	"log"
	"os"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Gagal terhubung ke database: ", err)
	}
	defer db.Close()
	fmt.Println("✅ Terhubung ke database PostgreSQL")

	r := router.SetupRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	fmt.Printf("🚀 Server Gin berjalan di http://localhost:%s\n", port)
	if err := r.Run(addr); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
