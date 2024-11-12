package main

import (
	"context"
	"fmt"
	"os"
	"image/png"
	"io"
	"log"
	"strings"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
)

// App struct
type App struct {
	ctx context.Context
	rootPath string
	printerPath string
}

type Response struct {
	Success bool `json:"success"`
	Error string `json:"error"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.rootPath = os.Getenv("PRINT_FILE_PATH")
	a.printerPath = os.Getenv("PRINTER_PATH")
}

func (a *App) PrintFile(filepath string, quantity int) Response {
	for i := 0; i < quantity; i++ {
		db_file, err := os.Open(fmt.Sprintf("%v\\%v", a.rootPath, filepath))
		if err != nil {
			return Response{Success: false, Error: err.Error()}
		}
		defer db_file.Close()
		success, err := Copy(db_file, a.printerPath)
		if err != nil {
			log.Printf("could not print file %v: %v", filepath, err)
			return Response{Success: false, Error: err.Error()}

		}
		if !success {
			return Response{Success: false, Error: "something went wrong"}
		}
	}
	return Response{Success: true, Error: ""}
}

func (a *App) Generate(filepath string) Response {
	if filepath == "" {
		log.Printf("cannot generate barcode from empty filepath")
		return Response{Success: false, Error: "could not generate barcode from empty filepath"}
	}

	writer := oned.NewCode128Writer()

	img, err := writer.Encode(filepath, gozxing.BarcodeFormat_CODE_128, 250, 50, nil)
	if err != nil {
		log.Printf("cannot encode filepath as barcode: %v", err)
		return Response{Success: false, Error: err.Error()}

	}

	filename := strings.TrimSuffix(filepath, ".prn")


	file, err := os.Create(fmt.Sprintf("%v\\%v_barcode.png", a.rootPath, filename))
	if err != nil {
		log.Printf("cannot create barcode png image: %v", err)
		return Response{Success: false, Error: err.Error()}

	}

	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Printf("cannot encode barcode to png file: %v", err)
		return Response{Success: false, Error: err.Error()}

	}

	return Response{Success: false, Error: ""}

}

func Copy(source *os.File, dest string) (bool, error) {

	fd2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("could not open destination file: %v", err)
		return false, err
	}


	_, err = io.Copy(fd2, source)
	fd2.Close()
	if err != nil {
		nested_err := os.Remove(fd2.Name())
		if nested_err != nil {
			log.Printf("could not delete destination file: %v", err)
			err = fmt.Errorf("%w\n%w", err, nested_err)
		}
		return false, err
	}
	return true, nil
}
