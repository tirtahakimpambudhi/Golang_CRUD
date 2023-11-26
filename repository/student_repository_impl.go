package repository

import (
  "database/sql"
  "ttd/config"
  "ttd/models"
  "log"
  "fmt"
  "errors"
  "strings"
  )
type StudentRepoImpl struct {
  DB *sql.DB
}

func (repo *StudentRepoImpl ) FindAll()  []models.Student {
  var students []models.Student
  query := fmt.Sprintf("SELECT id,nis,name,jurusan FROM %v",config.TBName)
  stmt, err := repo.DB.Prepare(query)
  if err != nil {
    log.Fatal(err.Error())
  }
  rows , err := stmt.Query()
  defer rows.Close()
  if err != nil {
    log.Fatal(err.Error())
  }
  
  for rows.Next(){
    row := models.Student{}
    rows.Scan(&row.ID,&row.NIS,&row.Name,&row.Jurusan)
    students = append(students,row)
  }
  
  return students
}

func (repo *StudentRepoImpl) FindAllPaging(jurusan string, limit, offset int) []models.Student {
  var students []models.Student
  var query string
  var err error

  if jurusan != "" {
      query = fmt.Sprintf("SELECT id, nis, name, jurusan FROM %v WHERE jurusan = ? LIMIT %v OFFSET %v", config.TBName, limit, offset)
  } else {
      query = fmt.Sprintf("SELECT id, nis, name, jurusan FROM %v LIMIT %v OFFSET %v", config.TBName, limit, offset)
  }

  stmt, err := repo.DB.Prepare(query)
  if err != nil {
      log.Fatal(err.Error())
  }
  defer stmt.Close()

  var rows *sql.Rows
  if jurusan != "" {
      rows, err = stmt.Query(jurusan)
  } else {
      rows, err = stmt.Query()
  }
  if err != nil {
      log.Fatal(err.Error())
  }
  defer rows.Close()

  for rows.Next() {
      row := models.Student{}
      rows.Scan(&row.ID, &row.NIS, &row.Name, &row.Jurusan)
      students = append(students, row)
  }

  return students
}
func (repo *StudentRepoImpl) FindSearchPaging(key string, limit, offset int) []models.Student {
  var students []models.Student
  var query string
  var err error
  key = "%" + key + "%"
  query = fmt.Sprintf("SELECT * FROM %v WHERE nis LIKE ? OR name LIKE ? OR jurusan LIKE ? LIMIT %v OFFSET %v",config.TBName,limit,offset)



  stmt, err := repo.DB.Prepare(query)
  if err != nil {
      log.Fatal(err.Error())
  }
  defer stmt.Close()

  var rows *sql.Rows
  rows, err = stmt.Query(key,key,key)
  
  if err != nil {
      log.Fatal(err.Error())
  }
  defer rows.Close()

  for rows.Next() {
      row := models.Student{}
      rows.Scan(&row.ID, &row.NIS, &row.Name, &row.Jurusan)
      students = append(students, row)
  }

  return students
}
func (repo *StudentRepoImpl) FindSearch(key string) []models.Student {
  var students []models.Student
  var query string
  var err error
  key = "%" + key + "%"
  query = fmt.Sprintf("SELECT * FROM %v WHERE nis LIKE ? OR name LIKE ? OR jurusan LIKE ?",config.TBName)



  stmt, err := repo.DB.Prepare(query)
  if err != nil {
      log.Fatal(err.Error())
  }
  defer stmt.Close()

  var rows *sql.Rows
  rows, err = stmt.Query(key,key,key)
  
  if err != nil {
      log.Fatal(err.Error())
  }
  defer rows.Close()

  for rows.Next() {
      row := models.Student{}
      rows.Scan(&row.ID, &row.NIS, &row.Name, &row.Jurusan)
      students = append(students, row)
  }

  return students
}
func (repo *StudentRepoImpl) FindByJurusan(jurusan string) []models.Student {
  var students []models.Student
  var query string
  var err error

  if jurusan != "" {
      query = fmt.Sprintf("SELECT id, nis, name, jurusan FROM %v WHERE jurusan = ? ", config.TBName)
  } else {
      query = fmt.Sprintf("SELECT id, nis, name, jurusan FROM %v ", config.TBName)
  }

  stmt, err := repo.DB.Prepare(query)
  if err != nil {
      log.Fatal(err.Error())
  }
  defer stmt.Close()

  var rows *sql.Rows
  if jurusan != "" {
      rows, err = stmt.Query(jurusan)
  } else {
      rows, err = stmt.Query()
  }
  if err != nil {
      log.Fatal(err.Error())
  }
  defer rows.Close()

  for rows.Next() {
      row := models.Student{}
      rows.Scan(&row.ID, &row.NIS, &row.Name, &row.Jurusan)
      students = append(students, row)
  }

  return students
}


func (repo *StudentRepoImpl ) FindByNIS(NIS int) (models.Student,error) {
  count := 0
  student := models.Student{}
  
  query := fmt.Sprintf("SELECT COUNT(*),id,nis,name,jurusan FROM %v WHERE nis = ? LIMIT 1",config.TBName)
  stmt, err := repo.DB.Prepare(query)
  
  if err != nil {
    return student,err
  }
  
  stmt.QueryRow(NIS).Scan(&count,&student.ID, &student.NIS, &student.Name, &student.Jurusan)
  if count == 0 {
    return student,errors.New(fmt.Sprintf("student with NIS '%v' not found",NIS))
  }
  return student , nil
}


func (repo *StudentRepoImpl) CreateStudent(student models.Student) error {
  var count int
  err := repo.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE nis = ?",config.TBName), student.NIS).Scan(&count)
  if err != nil {
    return err
  }

  if count != 0 {
    return errors.New(fmt.Sprintf("student with NIS '%v' exists", student.NIS))
  }
  tx, err := repo.DB.Begin()
  if err != nil {
    return err
  }
  defer func() {
    if err != nil {
      tx.Rollback()
      return 
    }
    err = tx.Commit()
  }()
  
  query := fmt.Sprintf("INSERT INTO %v (nis, name, jurusan) VALUES (?, ?, ?);", config.TBName)
  stmt, err := tx.Prepare(query)
  if err != nil {
    return err
  }

  _, err = stmt.Exec(student.NIS, strings.ToUpper(student.Name), strings.ToUpper(student.Jurusan))
  if err != nil {
    return err
  }

  return nil
}
func (repo *StudentRepoImpl) UpdateStudent(NIS int,student models.Student) error {
  var count int
  err := repo.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE nis = ?",config.TBName),NIS).Scan(&count)
  if err != nil {
    return err
  }

  if count == 0 {
    return errors.New(fmt.Sprintf("student with NIS '%v' not exists", NIS))
  }
  tx, err := repo.DB.Begin()
  if err != nil {
    return err
  }
  defer func() {
    if err != nil {
      tx.Rollback()
      return 
    }
    err = tx.Commit()
  }()
  
  query := fmt.Sprintf("UPDATE %v SET nis = ? ,name = ?, jurusan = ? WHERE nis = ?;", config.TBName)
  stmt, err := tx.Prepare(query)
  if err != nil {
    return err
  }

  _, err = stmt.Exec(student.NIS, strings.ToUpper(student.Name), strings.ToUpper(student.Jurusan),NIS)
  if err != nil {
    return err
  }

  return nil
}
func (repo *StudentRepoImpl) DeleteStudent(NIS int) error {
  var count int
  err := repo.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE nis = ?",config.TBName),NIS).Scan(&count)
  if err != nil {
    return err
  }

  if count == 0 {
    return errors.New(fmt.Sprintf("student with NIS '%v' not exists", NIS))
  }
  tx, err := repo.DB.Begin()
  if err != nil {
    return err
  }
  defer func() {
    if err != nil {
      tx.Rollback()
      return 
    }
    err = tx.Commit()
  }()
  
  query := fmt.Sprintf("DELETE FROM %v WHERE nis = ?;", config.TBName)
  stmt, err := tx.Prepare(query)
  if err != nil {
    return err
  }

  _, err = stmt.Exec(NIS)
  if err != nil {
    return err
  }

  return nil
}

func NewStudent(db *sql.DB) StudentRepository {
  return &StudentRepoImpl{
    DB:db,
  }
}