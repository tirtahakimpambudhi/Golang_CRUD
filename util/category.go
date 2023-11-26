package util

import (
	"strings"
	"ttd/repository"
)

func TypesJurusan(dbs repository.StudentRepository) map[string]string {
	typesJurusan := make(map[string]string)

	students := dbs.FindAll()

	for _, student := range students {
		if typesJurusan[student.Jurusan] == "" {
			typesJurusan[student.Jurusan] = strings.Replace(student.Jurusan," ","_",-1)
		}
	}
	return typesJurusan
}