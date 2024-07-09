package main

// import (
// 	as "Assign2/structs"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type Student struct {
// 	Name     string      `json:"name"`
// 	Marks    as.Marks    `json:"marks"`
// 	School   as.School   `json:"school"`
// 	Address  as.Address  `json:"address"`
// 	Personal as.Personal `json:"personal"`
// }

// var Students []Student

// func main() {
// 	log.Println("Starting...")
// 	http.HandleFunc("/getStudent", getStudent)
// 	http.HandleFunc("/createStudent", createStudent)
// 	http.HandleFunc("/updateStudent", updateStudent)
// 	http.ListenAndServe(":10000", nil)
// }

// func getStudent(w http.ResponseWriter, r *http.Request) {
// 	(w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(w).Header().Set("Access-Control-Allow-Credentials", "true")
// 	(w).Header().Set("Access-Control-Allow-Methods", "GET")
// 	(w).Header().Set("Content-Type", "application/json")

// 	name := r.URL.Query().Get("name")
// 	dtype := r.URL.Query().Get("dtype")

// 	if name == "" && dtype != "" {

// 		if dtype == "All" || dtype == "all" {
// 			for _, student := range Students {
// 				json.NewEncoder(w).Encode(student)

// 			}

// 		} else if dtype == "Marks" || dtype == "marks" {
// 			for _, student := range Students {
//				json.NewEncoder(w).Encode(student.Name)
// 				json.NewEncoder(w).Encode(student.Marks)
// 			}
// 		} else if dtype == "School" || dtype == "school" {
// 			for _, student := range Students {
// 				json.NewEncoder(w).Encode(student.School)
// 			}
// 		} else if dtype == "Address" || dtype == "address" {
// 			for _, student := range Students {
// 				json.NewEncoder(w).Encode(student.Address)
// 			}
// 		} else if dtype == "Personal" || dtype == "personal" {
// 			for _, student := range Students {
// 				json.NewEncoder(w).Encode(student.Personal)
// 			}
// 		} else {
// 			http.Error(w, "Invalid data type", http.StatusBadRequest)
// 			return
// 		}

// 	} else if name != "" && dtype != "" {

// 		if dtype == "All" || dtype == "all" {
// 			for _, student := range Students {
// 				if student.Name == name {
// 					json.NewEncoder(w).Encode(student)
// 					return
// 				}
// 			}
// 		} else if dtype == "Marks" || dtype == "marks" {
// 			for _, student := range Students {
// 				if student.Name == name {
// 					json.NewEncoder(w).Encode(student.Marks)
// 					return
// 				}
// 			}
// 		} else if dtype == "School" || dtype == "school" {
// 			for _, student := range Students {
// 				if student.Name == name {
// 					json.NewEncoder(w).Encode(student.School)
// 					return
// 				}
// 			}
// 		} else if dtype == "Address" || dtype == "address" {
// 			for _, student := range Students {
// 				if student.Name == name {
// 					json.NewEncoder(w).Encode(student.Address)
// 					return
// 				}
// 			}
// 		} else if dtype == "Personal" || dtype == "personal" {
// 			for _, student := range Students {
// 				if student.Name == name {
// 					json.NewEncoder(w).Encode(student.Personal)
// 					return
// 				}
// 			}
// 		} else {
// 			http.Error(w, "Invalid data type", http.StatusBadRequest)
// 			return
// 		}

// 	} else {
// 		http.Error(w, "Invalid data type", http.StatusBadRequest)
// 		return

// 	}

// }

// func createStudent(w http.ResponseWriter, r *http.Request) {

// 	fmt.Println("Creating Student .........")
// 	(w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(w).Header().Set("Access-Control-Allow-Credentials", "true")
// 	(w).Header().Set("Access-Control-Allow-Methods", "POST")
// 	(w).Header().Set("Content-Type", "application/json")

// 	defer r.Body.Close()
// 	var newStudent Student

// 	err := json.NewDecoder(r.Body).Decode(&newStudent)
// 	if err != nil {
// 		log.Println("Error parsing request body:", err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	Students = append(Students, newStudent)

// 	fmt.Fprintf(w, "Student %v created", newStudent.Name)
// 	fmt.Fprintln(w, Students)
// }

// func updateStudent(w http.ResponseWriter, r *http.Request) {

// 	fmt.Println("Updating Student .........")
// 	(w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(w).Header().Set("Access-Control-Allow-Credentials", "true")
// 	(w).Header().Set("Access-Control-Allow-Methods", "PUT")
// 	(w).Header().Set("Content-Type", "application/json")

// 	defer r.Body.Close()

// 	name := r.URL.Query().Get("name")
// 	dtype := r.URL.Query().Get("dtype")

// 	for i := 0; i < len(Students); i++ {

// 		if Students[i].Name == name {

// 			if dtype == "Marks" || dtype == "marks" {

// 				var tempMarks as.Marks
// 				err := json.NewDecoder(r.Body).Decode(&tempMarks)
// 				if err != nil {
// 					log.Println("Error parsing request body:", err)
// 					http.Error(w, err.Error(), http.StatusBadRequest)
// 					return
// 				}
// 				Students[i].Marks = tempMarks
// 				json.NewEncoder(w).Encode(Students[i].Name + " Marks is Updated Successfully")
// 				return

// 			} else if dtype == "School" || dtype == "school" {

// 				var tempSchool as.School
// 				err := json.NewDecoder(r.Body).Decode(&tempSchool)
// 				if err != nil {
// 					log.Println("Error parsing request body:", err)
// 					http.Error(w, err.Error(), http.StatusBadRequest)
// 					return
// 				}
// 				Students[i].School = tempSchool
// 				json.NewEncoder(w).Encode(Students[i].Name + " School Details is Updated Successfully")
// 				return

// 			} else if dtype == "Address" || dtype == "address" {

// 				var tempAddress as.Address
// 				err := json.NewDecoder(r.Body).Decode(&tempAddress)
// 				if err != nil {
// 					log.Println("Error parsing request body:", err)
// 					http.Error(w, err.Error(), http.StatusBadRequest)
// 					return
// 				}
// 				Students[i].Address = tempAddress
// 				json.NewEncoder(w).Encode(Students[i].Name + " Address is Updated Successfully")
// 				return

// 			} else if dtype == "Personal" || dtype == "personal" {

// 				var tempPersonal as.Personal
// 				err := json.NewDecoder(r.Body).Decode(&tempPersonal)
// 				if err != nil {
// 					log.Println("Error parsing request body:", err)
// 					http.Error(w, err.Error(), http.StatusBadRequest)
// 					return
// 				}
// 				Students[i].Personal = tempPersonal
// 				json.NewEncoder(w).Encode(Students[i].Name + " Personal Details is Updated Successfully")
// 				return

// 			} else {

// 				http.Error(w, "Invalid data type", http.StatusBadRequest)
// 				return
// 			}

// 		}
// 	}

// }
