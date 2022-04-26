package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {
	var num int
	fmt.Print("num:")
	fmt.Scanln(&num)
	hap(1, 2)

	if num == 1 {
		os.Remove("sqlite-database.db") //파일 중복 확인

		log.Println("Creating sqlite-database.db...")
		file, err := os.Create("sqlite-database.db") // sqlite 파일 생성
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("sqlite-database.db created")

		//파일 열기
		sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
		defer sqliteDatabase.Close() //
		createTable(sqliteDatabase)  // Create Database Tables

		//데이터 입력
		insertStudent(sqliteDatabase, "0001", "Liana Kim", "Bachelor")
		insertStudent(sqliteDatabase, "0002", "Glen Rangel", "Bachelor")
		insertStudent(sqliteDatabase, "0003", "Martin Martins", "Master")
		insertStudent(sqliteDatabase, "0004", "Alayna Armitage", "PHD")
		insertStudent(sqliteDatabase, "0005", "Marni Benson", "Bachelor")
		insertStudent(sqliteDatabase, "0006", "Derrick Griffiths", "Master")
		insertStudent(sqliteDatabase, "0007", "Leigh Daly", "Bachelor")
		insertStudent(sqliteDatabase, "0008", "Marni Benson", "PHD")
		insertStudent(sqliteDatabase, "0009", "Klay Correa", "Bachelor")
	} else {
		// DISPLAY INSERTED RECORDS
		sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
		displayStudents(sqliteDatabase)
	}

}

func createTable(db *sql.DB) { // *sql.DB 받기
	createStudentTableSQL := `CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
	  );`

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")
}

//*sql.DB 값 , 컬럼값 받기(code, name, program)
func insertStudent(db *sql.DB, code string, name string, program string) {
	log.Println("Inserting student record ...")
	//insert 구문 작성
	insertStudentSQL := "INSERT INTO student(code, name, program) VALUES (?, ?, ?)"
	//insert 구문 prepare 대기시켜놓기
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// SQLinjection 방지
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, program)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayStudents(db *sql.DB) {
	row, err := db.Query("SELECT idStudent,code, name, program FROM student ORDER BY idStudent")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var code string
		var name string
		var program string
		//select 구분 리턴값을 스캔해서 포인터에 넣기
		row.Scan(&id, &code, &program, &name) //컬럼 순서에 맞춰줘야 함
		//출력위치 설정
		log.Println("Student: ", code, " ", program, " ", name, " ", id)

	}

}
