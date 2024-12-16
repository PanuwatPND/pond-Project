package main

import (
	"app/cmd"              // ใช้งานคำสั่ง Cobra
	"github.com/joho/godotenv" // โหลดไฟล์ .env
	"log"
)

func main() {
	// โหลดค่าจากไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// เรียกใช้คำสั่ง Cobra
	cmd.Execute()
}