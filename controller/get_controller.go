package controller

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"ttd/database"
	"ttd/models"
	"ttd/repository"
	"ttd/util"
	"ttd/views"

	"github.com/gorilla/mux"
)


type Data struct {
	Title string
	Message	string
	Result interface{}
}

type result struct {
	Students []models.Student
	Page	Pagination
	TypesJurusan map[string]string
}

type resultSearch struct {
	Students []models.Student
	Page	Pagination
}

type Pagination struct {
	Next		int
	Current 	int
	Previous 	int
	TotalPage	int
	Data		int
}

func Home(w http.ResponseWriter, r *http.Request) {
	
	var students []models.Student
	var totalStudents int
	//Get Params Category
	vars := mux.Vars(r)
	category := vars["category"]
	 
	

	//DB CONNECTION
	db := database.GetConnection()
	defer db.Close()
	dbstudent := repository.NewStudent(db)
	typesJurusan := util.TypesJurusan(dbstudent)
	//Get Params Page
	page := r.URL.Query().Get("p")
	if page == "" || page == "0" {
		page = "1"
	} 
	//Parse Params to INT
	p, err := strconv.Atoi(page)
	limit := 10
	offset := (p - 1) * limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Query SQL
	if strings.ToUpper(category) == "ALL" {
		students = dbstudent.FindAllPaging("", limit, offset)
		totalStudents = len(dbstudent.FindAll()) 
	} else {
		jurusan := strings.Replace(category,"_"," ",-1)
		totalStudents = len(dbstudent.FindByJurusan(strings.ToUpper(jurusan)))
		students = dbstudent.FindAllPaging(jurusan, limit, offset)
	}

	if totalStudents == 0 {
		http.Error(w,"Jurusan Not Found",http.StatusNotFound)
		return
	}

	totalPage := int(math.Ceil(float64(totalStudents) / float64(limit)))

	views.Render(w, "index", Data{
		Title:   "Home Page",
		Message: "Welcome",
		Result: result{
			Page:      Pagination{
				Next: p + 1,
				Current: p,
				Previous: p - 1,
				TotalPage: totalPage,
				Data : limit ,
			},
			Students:  students,
			TypesJurusan: typesJurusan,
		},
	})
}


func Search(w http.ResponseWriter, r *http.Request) {
	log.Println("Parsing Params...")
	page := r.URL.Query().Get("p")
	key := r.URL.Query().Get("q")
	if page == "" || page == "0" {
		page = "1"
	} 

	p, err := strconv.Atoi(page)
	limit := 10
	offset := (p - 1) * limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	

	log.Println("Connection To Database...")
	db := database.GetConnection()
	log.Println("successfully to connection Database")
	defer db.Close()
	dbs := repository.NewStudent(db)
	students := dbs.FindSearchPaging(key,limit,offset)
	totalPage := int(math.Ceil(float64(len(dbs.FindSearch(key))) / float64(limit)))
	result := resultSearch{
		Page: Pagination{
			Next: p + 1,
			Current: p,
			Previous: p - 1,
			TotalPage: totalPage,
			Data: limit,
		},
		Students: students,
	}

	log.Println("Parsing result to JSON")
	resultJSON ,err := json.Marshal(result)

	if err != nil {
		log.Println("Failed Parsing JSON")
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	w.Header().Set("content-type","application/json")
	w.Write(resultJSON)
}


