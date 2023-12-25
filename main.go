package main

import (
	"github.com/joho/godotenv"
	"github.com/forumGamers/Octo-Cat/errors"
)

func main() {
	errors.PanicIfError(godotenv.Load())
}
