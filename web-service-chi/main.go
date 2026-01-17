package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("hello world")

	currentId := 1
	storage := make(map[int]CrudItem)

	router := chi.NewRouter()

	router.Post("/crud-items", func(w http.ResponseWriter, r *http.Request) {
		var item CrudItem
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		item.Id = currentId
		storage[currentId] = item

		jsonItem, err := json.Marshal(item)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(jsonItem)
		currentId++
	})

	router.Get("/crud-items/", func(w http.ResponseWriter, r *http.Request) {
		result := make([]CrudItem, 0, len(storage))
		for _, item := range storage {
			result = append(result, item)
		}

		resultAsJson, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(resultAsJson)
	})

	router.Get("/crud-items/{id}", func(w http.ResponseWriter, r *http.Request) {
		isAsString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(isAsString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if _, ok := storage[id]; !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		resultJson, err := json.Marshal(storage[id])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(resultJson)
	})

	router.Put("/crud-items/{id}", func(w http.ResponseWriter, r *http.Request) {
		isAsString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(isAsString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if _, ok := storage[id]; !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var item CrudItem
		err = json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		item.Id = id
		storage[id] = item
		jsonItem, err := json.Marshal(item)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(jsonItem)
	})

	router.Delete("/crud-items/{id}", func(w http.ResponseWriter, r *http.Request) {
		isAsString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(isAsString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if _, ok := storage[id]; !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		delete(storage, id)
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(router)
}

type CrudItem struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	internal    string
}
