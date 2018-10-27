package routes
import (
		"encoding/json"
		"github.com/gorilla/mux"
		"net/http"
		"fmt"
		_ "github.com/go-sql-driver/mysql"
        "database/sql"
		)

		type person_struct struct {
			Name string
			Age int
		}


func HealthCheck(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode("Still alive!")
}

func HandleQryMessage(w http.ResponseWriter, r *http.Request) {
    vars := r.URL.Query()
    message := vars.Get("msg")

    json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func HandleUrlMessage(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    message := vars["msg"]

    json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func HandlePost(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	p := person_struct{}
	err := decoder.Decode(&p)

	if err != nil {
		panic(err)
	}

	personJson, err := json.Marshal(p)

	if err != nil {
		panic(err)
	}
	createUser(p)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(p)
	w.Write(personJson)
}

func createUser(person person_struct) person_struct {
	db, err := sql.Open("mysql", "root:1234@/go?charset=utf8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT person SET name=?, age=?")
	checkErr(err)

	res, err := stmt.Exec(person.Name, person.Age)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	db.Close()
 return person
}

func SelectUsers() []person_struct {
	db, err := sql.Open("mysql", "root:1234@/go?charset=utf8")
	checkErr(err)

	// insert
	rows, err := db.Query("SELECT * FROM person")
	checkErr(err)
	var i = 0
	persons := make([]person_struct, 2)
	for rows.Next() {
		i++
		person := person_struct{}
		err = rows.Scan(&person.Name, &person.Age)
		checkErr(err)

		println(person.Name)
		persons = append(persons, person)
	}
	fmt.Printf("%v", persons)
	return persons
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleGet(w http.ResponseWriter, request *http.Request){
	personJson, err := json.Marshal(SelectUsers())
	checkErr(err)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(personJson)
}