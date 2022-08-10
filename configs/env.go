package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ENVCLOUNDNAME() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR env file")
	}
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func ENVCLOUNDAPIKEY() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR env file")
	}
	return os.Getenv("CLOUDINARY_API_KEY")
}

func ENVCLOUNDAPISECRET() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR env file")
	}
	return os.Getenv("CLOUDINARY_API_SECRET")
}

func ENVCLOUNDUPLOADFOLDER() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR env file")
	}
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}
