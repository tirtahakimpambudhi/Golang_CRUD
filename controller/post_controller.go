package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"ttd/database"
	"ttd/models"
	"ttd/repository"
	"ttd/util"

	"github.com/gorilla/mux"
)


func AddStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	err := r.ParseForm()
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	body := r.PostForm
	err = util.ValidateForm(body,&student)
	if err != nil {
		log.Println("users doing bad request...")
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	log.Println("successfully entered input from users...")
	//Connection DB
	log.Println("Connection To Database...")
	db := database.GetConnection()
	log.Println("successfully to connection Database")
	defer db.Close()
	dbs := repository.NewStudent(db)
	//Insert DB
	student.Jurusan = strings.Replace(student.Jurusan,"_"," ",-1)
	err = dbs.CreateStudent(student)

	if err != nil {
		log.Println("failed to insert in database...")
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	log.Println("successfully to insert in database...")
	http.Redirect(w,r,fmt.Sprintf("/students/%v",strings.Replace(student.Jurusan," ","_", -1)),301)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	NIS := vars["NIS"]
	//Parse NIS
	log.Println("parsing params NIS...")
	parseNIS , err := strconv.Atoi(NIS)
	if err != nil {
		log.Println("users doing bad request...")
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	log.Println("Connection To Database...")
	db := database.GetConnection()
	log.Println("successfully to connection Database")
	defer db.Close()
	dbs := repository.NewStudent(db)

	err = dbs.DeleteStudent(parseNIS)
	if err != nil {
		log.Println("failed to delete in database")
		http.Error(w,err.Error(),http.StatusNotFound)
		return
	}
	log.Println("successfully to delete in database")
	w.WriteHeader(http.StatusOK)
}
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	vars := mux.Vars(r)
	NIS := vars["NIS"]
	//Parse NIS
	log.Println("parsing params NIS...")
	parseNIS , err := strconv.Atoi(NIS)
	if err != nil {
		log.Println("users doing bad request...")
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	body := r.PostForm
	err = util.ValidateForm(body,&student)
	if err != nil {
		log.Println("users doing bad request...")
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	log.Println("successfully entered input from users...")

	log.Println("Connection To Database...")
	db := database.GetConnection()
	log.Println("successfully to connection Database")
	defer db.Close()
	dbs := repository.NewStudent(db)
	student.Jurusan = strings.Replace(student.Jurusan,"_"," ",-1)
	err = dbs.UpdateStudent(parseNIS,student)
	if err != nil {
		log.Println("failed to update in database")
		http.Error(w,err.Error(),http.StatusNotFound)
		return
	}
	log.Println("successfully to update in database")
	http.Redirect(w,r,fmt.Sprintf("/students/%v",strings.Replace(student.Jurusan," ","_", -1)),301)
}
