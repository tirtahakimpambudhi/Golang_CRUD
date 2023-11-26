package test

import (
	"testing"
	"ttd/config"
	"ttd/database"
	"ttd/repository"
	"ttd/util"
)

func TestQueryPagig(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()
	dbs := repository.NewStudent(db)
	jurusan := "XI SIJA 2"
	page := 2
	limit  := 5 
	offset := (page - 1) *  limit
	students := dbs.FindAllPaging(jurusan,limit,offset)
	t.Log(students)
}

func TestConfig(t *testing.T){
	t.Log(config.Connect)
}

func TestFindByJurusan(t *testing.T){
	db := database.GetConnection()
	defer db.Close()
	dbs := repository.NewStudent(db)
	students := dbs.FindByJurusan("XI SIJA")
	t.Log(len(students))
	for _, s := range students {
		t.Log(s.ID)
	}
}


func TestTypesCategory(t *testing.T){
	db := database.GetConnection()
	defer db.Close()
	dbs := repository.NewStudent(db)
	types := util.TypesJurusan(dbs)

	for _, ts := range types {
		t.Log(ts)
	}
}