package main

import (
	as "Assign2/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Student is a struct
type Student struct {
	Name     string      `json:"name"`
	Marks    as.Marks    `json:"marks"`
	School   as.School   `json:"school"`
	Address  as.Address  `json:"address"`
	Personal as.Personal `json:"personal"`
}

type ErrorStruct struct {
	Status string
	Error  string
}

// type Error struct {
// 	status  string
// 	message string
// }

var Students []Student /* this slice is used as main database for the api*/

func main() {
	log.Println("Starting...")
	http.HandleFunc("/getStudent", getStudent)
	http.HandleFunc("/createStudent", createStudent)
	http.HandleFunc("/updateStudent", updateStudent)
	http.ListenAndServe(":10000", nil)
}

/*
Purpose : to get the student details
Request :
   Body : nil
----------------------
   Header : nil
----------------------
   Params :
	  name :
	 dtype :
----------------------
	Request :
	On Success
	==========
	{
		name : "anyName"
		status : s
		error : nil
	}
	On Error
	=========
	{
		name : ""
		status : e
		error : error msg
	}
Authorization : SAI ARAVINDH
Date : 08-07-2024 */

func getStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting data ...")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Content-Type", "application/json")

	if r.Method == "GET" { // to ensure that this func should work only the method is fixed in GET

		var lresponse ErrorStruct
		lresponse.Status = "s"

		name := r.URL.Query().Get("name")
		dtype := r.URL.Query().Get("dtype")

		// if the database is empty it will throw a message
		if Students == nil {
			http.Error(w, "No Data Found : Empty Database ", http.StatusBadRequest)
			return
		}

		if name == "" && dtype != "" { // used to filter the data based on the type and not considering the name

			if dtype == "All" || dtype == "all" {
				Data, err := json.Marshal(Students)
				if err != nil {
					lresponse.Status = "e"
					lresponse.Error = err.Error()
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(Data)

			} else if dtype == "Name" || dtype == "name" { // to get the Names of all the students

				var NameData []string
				for _, student := range Students {
					NameData = append(NameData, student.Name)
				}
				Data, err := json.Marshal(NameData)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(Data)

			} else if dtype == "Marks" || dtype == "marks" { // to get marks of all the students

				type MarksData struct {
					Name  string   `json:"name"`
					Marks as.Marks `json:"marks"`
				}

				var markData MarksData
				var AllMarks []MarksData

				for i := 0; i < len(Students); i++ {

					markData.Name = Students[i].Name
					markData.Marks = Students[i].Marks
					AllMarks = append(AllMarks, markData)
				}

				Data, err := json.Marshal(AllMarks)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(Data)

			} else if dtype == "School" || dtype == "school" { // to get School Details of all the students

				type SchoolData struct {
					Name   string    `json:"name"`
					School as.School `json:"school"`
				}

				var schoolData SchoolData
				var AllSchool []SchoolData
				for i := 0; i < len(Students); i++ {

					schoolData.Name = Students[i].Name
					schoolData.School = Students[i].School
					AllSchool = append(AllSchool, schoolData)

				}

				Data, err := json.Marshal(AllSchool)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(Data)

			} else if dtype == "Address" || dtype == "address" { // to get Address Details of all the students

				type AddressData struct {
					Name    string     `json:"name"`
					Address as.Address `json:"address"`
				}

				var addressData AddressData
				var AllAddress []AddressData
				for i := 0; i < len(Students); i++ {

					addressData.Name = Students[i].Name
					addressData.Address = Students[i].Address
					AllAddress = append(AllAddress, addressData)

				}

				Data, err := json.Marshal(AllAddress)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(Data)

			} else if dtype == "Personal" || dtype == "personal" { // to get Personal Details of all the students

				type PersonalData struct {
					Name     string      `json:"name"`
					Personal as.Personal `json:"personal"`
				}

				var personalData PersonalData
				var AllPersonal []PersonalData
				for i := 0; i < len(Students); i++ {

					personalData.Name = Students[i].Name
					personalData.Personal = Students[i].Personal
					AllPersonal = append(AllPersonal, personalData)

				}

				Data, err := json.Marshal(AllPersonal)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(Data)

			} else {
				http.Error(w, "Invalid data type", http.StatusBadRequest)
				return
			}
			/**************************************************************************************************************/
		} else if name != "" && dtype != "" { // to get Details of specific student

			var NameVerify bool = false

			// this range will check the database , whether the entered name is present or not, if not it will throw a error msg
			for _, s := range Students {
				if s.Name == name {
					NameVerify = true
				}
			}

			if NameVerify { // if name is presented, this block is executed.

				for i := range Students {

					if Students[i].Name == name {
						if dtype == "All" || dtype == "all" {
							// var tempDetails Student
							tempDetails := Students[i] // using the temporary variable to to get the list of all the details of the selected student.
							Data, err := json.Marshal(tempDetails)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}
							w.Write(Data)
							return
						} else if dtype == "Marks" || dtype == "marks" {

							type MarksData struct {
								Name  string   `json:"name"`
								Marks as.Marks `json:"marks"`
							}

							var markData MarksData
							markData.Name = Students[i].Name
							markData.Marks = Students[i].Marks

							Data, err := json.Marshal(markData)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}
							w.Write(Data)
							return
						} else if dtype == "School" || dtype == "school" {

							type SchoolData struct {
								Name   string    `json:"name"`
								School as.School `json:"school"`
							}

							var schoolData SchoolData
							schoolData.Name = Students[i].Name
							schoolData.School = Students[i].School

							Data, err := json.Marshal(schoolData)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}
							w.Write(Data)
							return
						} else if dtype == "Address" || dtype == "address" {

							type AddressData struct {
								Name    string     `json:"name"`
								Address as.Address `json:"address"`
							}

							var addressData AddressData
							addressData.Name = Students[i].Name
							addressData.Address = Students[i].Address

							Data, err := json.Marshal(addressData)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}
							w.Write(Data)
							return
						} else if dtype == "Personal" || dtype == "personal" {
							type PersonalData struct {
								Name     string      `json:"name"`
								Personal as.Personal `json:"personal"`
							}

							var personalData PersonalData
							personalData.Name = Students[i].Name
							personalData.Personal = Students[i].Personal

							Data, err := json.Marshal(personalData)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}
							w.Write(Data)
							return
						} else {
							http.Error(w, "Invalid data type", http.StatusBadRequest)
							return
						}
					}
				}

			} else {
				http.Error(w, "StudentName is not Present in the database", http.StatusBadRequest)
				NameVerify = true
				return

			}

		} else {
			http.Error(w, "Invalid data type", http.StatusBadRequest)
			return
		}

	}

}

/*
Purpose : to Create the student details
Request :
   Body :
    {
   		marks    :
		school   :
		address  :
		personal :
    }
----------------------
   Header :
   {
     name :
   }
----------------------
   Params : nil
----------------------
	Request :
	On Success
	==========
	{
		name : "anyName"
		status : s
		error : nil
	}
	On Error
	=========
	{
		name : ""
		status : e
		error : error msg
	}
Authorization : SAI ARAVINDH
Date : 08-07-2024 */

func createStudent(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Creating Student .........")
	// fmt.Fprintln(w, len(Students))
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {

		namekey := r.Header.Get("name")

		Studentcheck := 0
		defer r.Body.Close()
		var newStudent Student
		newStudent.Name = namekey

		err := json.NewDecoder(r.Body).Decode(&newStudent)
		if err != nil {
			log.Println("Error parsing request body:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(Students) == 0 {

			if newStudent.Name != "" {

				if newStudent.Marks.M10.English != 0 && newStudent.Marks.M10.Tamil != 0 && newStudent.Marks.M10.Maths != 0 && newStudent.Marks.M10.Science != 0 && newStudent.Marks.M10.Social != 0 {

					if newStudent.School.S10.SchoolName != "" && newStudent.School.S10.Address != "" && newStudent.School.S10.Place != "" && newStudent.School.S10.Pincode != "" && newStudent.School.S10.Type != "" {

						if newStudent.Address.Landmark != "" && newStudent.Address.Pincode != "" && newStudent.Address.StudentAddress != "" {

							if newStudent.Personal.FatherName != "" && newStudent.Personal.MotherName != "" && newStudent.Personal.Age != 0 && newStudent.Personal.Gender != "" {

								Students = append(Students, newStudent)
								fmt.Fprintf(w, "Student %v created", newStudent.Name)
								// fmt.Fprintln(w, Students)
								return
							}
						}
					}

				}

			}
		} else if len(Students) != 0 {

			for index := 0; index < len(Students); index++ {
				if Students[index].Name == namekey {
					Studentcheck++

				}

			}

			if Studentcheck == 0 {
				if newStudent.Name != "" {
					if newStudent.Marks.M10.English != 0 && newStudent.Marks.M10.Tamil != 0 && newStudent.Marks.M10.Maths != 0 && newStudent.Marks.M10.Science != 0 && newStudent.Marks.M10.Social != 0 {
						if newStudent.School.S10.SchoolName != "" && newStudent.School.S10.Address != "" && newStudent.School.S10.Place != "" && newStudent.School.S10.Pincode != "" && newStudent.School.S10.Type != "" {
							if newStudent.Address.Landmark != "" && newStudent.Address.Pincode != "" && newStudent.Address.StudentAddress != "" {
								if newStudent.Personal.FatherName != "" && newStudent.Personal.MotherName != "" && newStudent.Personal.Age != 0 && newStudent.Personal.Gender != "" {

									Students = append(Students, newStudent)
									fmt.Fprintf(w, "Student %v created", newStudent.Name)
									Studentcheck = 0
									// fmt.Fprintln(w, Students)
									return
								}
							}
						}

					}

				}
			} else {
				fmt.Fprintln(w, "Student", namekey, "is already Presented ...:)")
			}

		}

		fmt.Fprintln(w, "Student Not created")
		Studentcheck = 0
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	log.Println("Error reading request body:", err)
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// err = json.Unmarshal(body, &newStudent)
		// if err != nil {
		// 	log.Println("Error parsing request body:", err)
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// Students = append(Students, newStudent)
	}
}

func updateStudent(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Updating Student .........")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {

		defer r.Body.Close()
		studentCheck := 0
		name := r.URL.Query().Get("name")
		dtype := r.URL.Query().Get("dtype")

		for i := 0; i < len(Students); i++ {

			if Students[i].Name == name {

				if dtype == "Marks" || dtype == "marks" {

					var tempMarks as.Marks

					body, err := ioutil.ReadAll(r.Body)

					if err != nil {
						log.Println("Error reading request body:", err)
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}

					err = json.Unmarshal(body, &tempMarks)
					if err != nil {
						log.Println("Error parsing request body:", err)
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}

					if tempMarks.M10 == nil {
						fmt.Fprint(w, "Invalid Details entered.")
						return
					}

					if tempMarks.M10.English > 0 && tempMarks.M10.Tamil > 0 && tempMarks.M10.Maths > 0 && tempMarks.M10.Science > 0 && tempMarks.M10.Social > 0 {

						// (tempMarks.M12.English > 0 || tempMarks.M12.Tamil > 0 || tempMarks.M12.Biology > 0 || tempMarks.M12.Chemistry > 0 || tempMarks.M12.Maths > 0)

						Students[i].Marks = tempMarks
						message := Students[i].Name + " Marks is Updated Successfully"
						studentCheck = 0
						jsonData, err := json.Marshal(message)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						w.Write(jsonData)
						return
					} else {
						http.Error(w, "Student Name Already Presented", http.StatusBadRequest)
						return
					}

				} else if dtype == "name" || dtype == "Name" {

					oldname := Students[i].Name

					tempname := r.Header.Get("name")

					if tempname == "" {
						http.Error(w, "Data not entered in header .", http.StatusBadRequest)
						return
					}
					var updateNamesCheck bool = true

					for names := range Students {
						if tempname == Students[names].Name {
							updateNamesCheck = false
						}
					}

					if updateNamesCheck {
						Students[i].Name = tempname

						message := oldname + " changed to " + Students[i].Name + " **  Updated Successfully **"
						studentCheck = 0
						jsonData, err := json.Marshal(message)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						w.Write(jsonData)
						return
					} else {
						http.Error(w, "Student Name Already Presented", http.StatusBadRequest)
						return
					}

				} else if dtype == "School" || dtype == "school" {

					var tempSchool as.School

					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						log.Println("Error reading request body:", err)
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}

					err1 := json.Unmarshal(body, &tempSchool)
					fmt.Printf("Type of school is  %T", tempSchool)
					if err1 != nil {
						log.Println("Error parsing request body:", err1)
						http.Error(w, err1.Error(), http.StatusBadRequest)
						return
					}

					if tempSchool.S10 == nil {
						fmt.Fprint(w, "Invalid Details entered.")
						return
					}

					if tempSchool.S10.SchoolName != "" && tempSchool.S10.Place != "" &&
						tempSchool.S10.Address != "" && tempSchool.S10.Pincode != "" && tempSchool.S10.Type != "" {
						Students[i].School = tempSchool
						message := Students[i].Name + " School Details is Updated Successfully"
						studentCheck = 0
						jsonData, err := json.Marshal(message)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						w.Write(jsonData)
						return
					} else {
						http.Error(w, "Some Mandatory detais are not entered.", http.StatusBadRequest)
						return
					}

				} else if dtype == "Address" || dtype == "address" {

					var tempAddress as.Address
					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						log.Println("Error reading request body:", err)
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					err = json.Unmarshal(body, &tempAddress)
					if err != nil {
						log.Println("Error parsing request body:", err)
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}

					// if tempAddress.Landmark == nil {
					// 	fmt.Fprint(w, "Invalid Details entered.")
					// 	return
					// }

					if tempAddress.StudentAddress != "" && tempAddress.Landmark != "" && tempAddress.Pincode != "" {
						Students[i].Address = tempAddress
						message := Students[i].Name + " Address is Updated Successfully"
						studentCheck = 0
						jsonData, err := json.Marshal(message)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						w.Write(jsonData)
						return
					} else {
						http.Error(w, "Some Mandatory detais are not entered.", http.StatusBadRequest)
						return
					}

				} else if dtype == "Personal" || dtype == "personal" {

					var tempPersonal as.Personal
					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						log.Println("Error reading request body:", err)
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					err = json.Unmarshal(body, &tempPersonal)
					if err != nil {
						log.Println("Error parsing request body:", err)
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}

					if tempPersonal.FatherName != "" && tempPersonal.MotherName != "" && tempPersonal.Age > 0 && tempPersonal.Gender != "" {
						Students[i].Personal = tempPersonal
						message := Students[i].Name + " Personal Details is Updated Successfully"
						studentCheck = 0
						jsonData, err := json.Marshal(message)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						w.Write(jsonData)
						return
					} else {
						http.Error(w, "Some Mandatory detais are not entered.", http.StatusBadRequest)
						return
					}

				} else {

					http.Error(w, "Invalid data type", http.StatusBadRequest)
					studentCheck = 0
					return
				}

			} else {
				studentCheck++
			}
		}

		if studentCheck != 0 {
			http.Error(w, "Invalid data type", http.StatusBadRequest)
			studentCheck = 0
			return
		}

	}

}
