package main

import (
	"flag"
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
	var port int
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.Parse()

	exPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	folderPath := path.Join(filepath.Dir(exPath), "files")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, os.ModePerm)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s -> %s\n", r.Method, r.RequestURI)

		if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
			r.ParseMultipartForm(0)

			for field, val := range r.MultipartForm.Value {
				fmt.Printf("%s = %s\n", field, val[0])
			}

			for field, headers := range r.MultipartForm.File {
				for _, header := range headers {
					file, err := header.Open()
					if err != nil {
						fmt.Printf("Failed to open file %s of field %s -> %v\n", header.Filename, field, err)
						continue
					}
					defer file.Close()

					id, err := uuid.NewUUID()
					if err != nil {
						fmt.Printf("Failed to generate UUID -> %v\n", err)
						continue
					}

					filePath := path.Join(folderPath, id.String()+filepath.Ext(header.Filename))
					dest, err := os.Create(filePath)
					if err != nil {
						fmt.Printf("Failed to create new file -> %v\n", err)
						continue
					}
					defer dest.Close()

					_, err = io.Copy(dest, file)
					if err != nil {
						fmt.Printf("Failed to copy file -> %v\n", err)
						continue
					}

					fmt.Printf("%s = Saved to: %s\n", field, filePath)
				}
			}
		} else {
			bytes, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Failed to read body -> %v\n", err)
				return
			}

			if len(bytes) > 0 {
				fmt.Println(string(bytes))
			}
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Printf("Server listens on port %s\n", strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(port)), nil))
}
