package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/ArianaFreitag/mmyope-website/pkg/models"
)

const keyServerAddr = "serverAddr"

// do this correctly, toggle local vs not, use secrets, etc
const (
	host     = "db-postgres-mmyope-do-user-22616298-0.g.db.ondigitalocean.com"
	port     = 25060
	user     = "doadmin"
	password = ""
	dbname   = "mmyope"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	ctx := r.Context()

	hasId := r.URL.Query().Has("id")
	id := r.URL.Query().Get("id")

	fmt.Printf("%s: got /getBook request. id(%t)=%s",
		ctx.Value(keyServerAddr), hasId, id)
	book, err := models.GetBook(id)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	json_book, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	response := fmt.Sprintf("%v", string(json_book))
	io.WriteString(w, response)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Root of website\n")

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// query db for all books
	ctx := r.Context()
	fmt.Printf("%s: got /getBooks request\n", ctx.Value(keyServerAddr))
	books, err := models.GetBooks()
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	json_books, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	response := fmt.Sprintf("%v", string(json_books))
	io.WriteString(w, response)
}

func getMagazines(w http.ResponseWriter, r *http.Request) {
}

func getMagazine(w http.ResponseWriter, r *http.Request) {
}
func getImage() {

}
func getImages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	enableCors(&w)

	hasPath := r.URL.Query().Has("path")
	path := r.URL.Query().Get("path")
	fmt.Printf("%s: got /getImage request. path(%t)=%s",
		ctx.Value(keyServerAddr), hasPath, path)

	buf, err := os.ReadFile("./images/" + path)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "image/jpg")
	w.Write(buf)
}

func Start() {
	var err error

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	models.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("static"))

	// mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/getBooks", getBooks)
	mux.HandleFunc("/getBook", getBook)
	mux.HandleFunc("/getImages", getImages)
	mux.Handle("/", fileserver)

	ctx := context.Background()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	fmt.Printf("Server listening on port %s\n", "8080")
	// http.ListenAndServe()

	err = server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}

}
