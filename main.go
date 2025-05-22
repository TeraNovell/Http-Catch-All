package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func main() {
	exPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	folderPath := path.Join(filepath.Dir(exPath), "files")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, os.ModePerm)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s -> %s\n", r.Method, r.RequestURI)

		if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
			r.ParseMultipartForm(0)

			for key, val := range r.MultipartForm.Value {
				fmt.Printf("%s = %s\n", key, val[0])
			}

			for key, val := range r.MultipartForm.File {
				file, err := val[0].Open()
				if err != nil {
					fmt.Println(err)
					return
				}
				defer file.Close()

				id, err := uuid.NewUUID()
				if err != nil {
					fmt.Println(err)
					return
				}

				filePath := path.Join(folderPath, id.String()+filepath.Ext(val[0].Filename))

				dest, err := os.Create(filePath)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer dest.Close()

				_, err = io.Copy(dest, file)
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Printf("%s = Saved to: %s\n", key, filePath)
			}
		} else {
			bytes, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
			}

			if len(bytes) == 0 {
				return
			}

			fmt.Println(string(bytes))
		}

		w.WriteHeader(http.StatusOK)
	})

	port := 8080

	log.Printf("Server listens on port %s\n", strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(port)), nil))
}
