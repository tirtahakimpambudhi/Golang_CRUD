package repository
import (
  "ttd/models"
  )
type StudentRepository interface {
  FindAll()([]models.Student)
  FindAllPaging(jurusan string,limit , offset int)([]models.Student)
  FindByNIS(NIS int)(models.Student,error)
  FindSearchPaging(key string,limit , offset int)([]models.Student)
  FindSearch(key string)([]models.Student)
  FindByJurusan(Jurusan string)([]models.Student)
  CreateStudent(student models.Student)(error)
  UpdateStudent(NIS int,student models.Student)(error)
  DeleteStudent(NIS int)(error)
}