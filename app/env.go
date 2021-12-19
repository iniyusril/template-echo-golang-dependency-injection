package app

import "github.com/joho/godotenv"

func InitializedEnvirontment() {
	godotenv.Load(".env")
}
