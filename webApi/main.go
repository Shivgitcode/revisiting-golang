package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course-file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake db

var courses []Course

//middleware,helper-file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func main() {

	r := mux.NewRouter()

	// fmt.Println("Starting the server")

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})

	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCoureses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllers-file
// serve home route

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome ot ApI </h1>"))
}

func getAllCoureses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	fmt.Println(params)

	// loop through courses , find matching id and return the response
	// for _, course := range courses {
	// 	if course.CourseId == params["id"] {
	// 		json.NewEncoder(w).Encode(course)
	// 		return
	// 	}
	// }
	json.NewEncoder(w).Encode("No course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	//generate unique id,string
	//append course into courses

	course.CourseId = strconv.Itoa(rand.Intn(100) + 1)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}
