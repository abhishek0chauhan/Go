package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses - file
type Course struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  int     `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware - file
func (c *Course) IsEmpty() bool {
	return c.Id == "" && c.Name == ""
	// return c.Name == ""
}

func main() {
	fmt.Println("API - in golang")

	r := mux.NewRouter()

	//seedinng
	courses = append(courses, Course{Id: "1", Name: "ReactJS", Price: 299, Author: &Author{Fullname: "Abhishek", Website: "http://www.abhishek.com"}})
	courses = append(courses, Course{Id: "2", Name: "Python", Price: 199, Author: &Author{Fullname: "abc", Website: "http://www.abc.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to API in golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)
	fmt.Println(params)

	// loop throug courses, find matching id and return the response
	for _, course := range courses {
		if course.Id == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
		json.NewEncoder(w).Encode("No course found with given id")
		return
	}
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about = {}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	//generate unique id, string
	//append new course into courses

	rand.Seed(time.Now().UnixNano())
	course.Id = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update a course")
	w.Header().Set("Content-Type", "application/json")

	//first grab id from request

	params := mux.Vars(r)

	// loop get id, remove, update, add with my ID

	for index, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.Id = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//TODO send a response when id is not found

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete a course")
	w.Header().Set("Content-Type", "application/json")

	//first grab id from request

	params := mux.Vars(r)

	for index, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course is deleted")
			return
		}
	}
}
