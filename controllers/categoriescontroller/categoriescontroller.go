package categoriescontroller

import (
	"encoding/json"
	"fmt"
	"go-web-native/entities"
	"go-web-native/middleware"
	"go-web-native/models/categorymodel"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	middleware.Re(w)
	authorizationHeader := r.Header.Get("Authorization")

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	// fmt.Println("string token", tokenString, authorizationHeader)
	err := middleware.VerifyToken(tokenString)
	if err != nil {
		w.Write([]byte("Invalid token"))
		return
	}

	categories := categorymodel.GetAll()
	w.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(categories)
	w.Write(payload)
}

func Add(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	middleware.Re(w)
	authorizationHeader := r.Header.Get("Authorization")

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	// fmt.Println("string token", tokenString, authorizationHeader)
	err := middleware.VerifyToken(tokenString)
	if err != nil {
		w.Write([]byte("Invalid token"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var category entities.Category
	defer r.Body.Close()
	request, _ := io.ReadAll(r.Body)
	// fmt.Println(request)
	json.Unmarshal(request, &category)
	// fmt.Println(category)
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	if ok := categorymodel.Create(category); !ok {
		fmt.Println(ok)
	}
	// fmt.Println(category.Name)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Data created"))

}

func Detil(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	middleware.Re(w)
	authorizationHeader := r.Header.Get("Authorization")

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	err := middleware.VerifyToken(tokenString)
	if err != nil {
		w.Write([]byte("Invalid token"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// var categories entities.Category

	input := r.URL.String()
	// Metode 1: Menggunakan regular expression
	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllString(input, -1)
	// fmt.Println("Hasil dengan regexp:", strings.Join(numbers, ""))
	idString := strings.Join(numbers, "")
	id, _ := strconv.Atoi(idString)
	// fmt.Println(id)

	category := categorymodel.Detil(int(id))
	payload, _ := json.Marshal(category)
	w.Write(payload)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	middleware.Re(w)
	authorizationHeader := r.Header.Get("Authorization")

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	// fmt.Println("string token", tokenString, authorizationHeader)
	err := middleware.VerifyToken(tokenString)
	if err != nil {
		w.Write([]byte("Invalid token"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var category entities.Category

	defer r.Body.Close()
	request, _ := io.ReadAll(r.Body)
	json.Unmarshal(request, &category)
	category.UpdatedAt = time.Now()
	id, _ := strconv.Atoi(category.Id)

	if ok := categorymodel.Update(int(id), category); !ok {
		fmt.Println(ok)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Data created"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	middleware.Re(w)
	authorizationHeader := r.Header.Get("Authorization")

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	err := middleware.VerifyToken(tokenString)
	if err != nil {
		w.Write([]byte("Invalid token"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var category entities.Category
	defer r.Body.Close()
	request, _ := io.ReadAll(r.Body)
	json.Unmarshal(request, &category)
	category.UpdatedAt = time.Now()
	id, _ := strconv.Atoi(category.Id)
	fmt.Println(id)

	if ok := categorymodel.Delete(id); ok != nil {
		fmt.Println(ok)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to delete data"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data deleted"))
}
