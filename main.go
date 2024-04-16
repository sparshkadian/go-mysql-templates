package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Task struct{
	TaskId 		int64 		`json:"task_id"`
	Task 		string 		`json:"task"`
	Username 	string 		`json:"username"`
}

var (
	db *sql.DB
	tmpl *template.Template
	fs http.Handler
	tasks []Task
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./static/index.html")
	} 

	if r.Method == "POST"{
		if err := r.ParseForm(); err != nil {
			fmt.Println("Error Parsing Form", err)
			return
		}
		task := r.FormValue("task")
		username := r.FormValue("username")

		_, queryErr := db.Query("INSERT INTO `go`.`tasks` (`task`, `username`) VALUES (?, ?)", task, username)
		if queryErr != nil {
			fmt.Println("Error Adding data to DB", queryErr)
			return
		}
	}

	fs.ServeHTTP(w, r)
}

func taskHandler(w http.ResponseWriter, r *http.Request){
	tmpl = template.Must(template.ParseFiles("./templates/tasks.html"))
	resultRow, queryErr := db.Query("SELECT * FROM `go`.`tasks`")
	if queryErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error Fetching data from DB", queryErr)
	}
	defer resultRow.Close()

	for resultRow.Next() {
		var task Task
		err := resultRow.Scan(&task.TaskId, & task.Task, &task.Username)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Printf("Error Fetching Data from DB %v\n", err)
			return 
		}
		tasks = append(tasks, task)
	}

	if len(tasks) == 0 {
		if err := tmpl.ExecuteTemplate(w, "tasks.html", nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Printf("Error Executing Template %v\n", err)
		}
	}

	if err := tmpl.ExecuteTemplate(w, "tasks.html", tasks); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error Executing Template %v\n", err)
	}
}

func main(){
	envErr := godotenv.Load()
	if envErr != nil {
    log.Fatalf("Error loading .env file: %v", envErr)
	}

	dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

	fs = http.FileServer(http.Dir("./static/"))

	var err error
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
    db, err = sql.Open("mysql", connString)
    if err != nil {
        log.Fatal(err)
    }
	
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging Database", err)
	} else {
		fmt.Println("DB Connection Successfull!")
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tasks", taskHandler)

	fmt.Println("Server Started at port: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error Starting Server", err)
	}
	
}