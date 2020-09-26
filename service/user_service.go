package service

import (
	"database/sql" // package to encode and decode the json into struct and vice versa
	"mywebproj/models"

	// "go-postgres/models" // models package where User schema is defined
	// used to access the request and response object of the api
	// used to read the environment variable

	// package used to covert string into int type
	// used to get the params from the route

	// "github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq" // postgres golang driver
	genv "github.com/sakirsensoy/genv"
)

// // response format
// type response struct {
// 	ID      int64  `json:"id,omitempty"`
// 	Message string `json:"message,omitempty"`
// }

// create connection with postgres db
func CreateConnection() (*sql.DB, error) {
	// Open the connection
	url := genv.Key("DATABASE_URL").Default("").String()
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateUser create a user in the postgres db
// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	// set the header to content type x-www-form-urlencoded
// 	// Allow all origin to handle cors issue
// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	// create an empty user of type models.User
// 	var user models.User

// 	// decode the json request to user
// 	err := json.NewDecoder(r.Body).Decode(&user)

// 	if err != nil {
// 		log.Fatalf("Unable to decode the request body.  %v", err)
// 	}

// 	// call insert user function and pass the user
// 	insertID := insertUser(user)

// 	// format a response object
// 	res := response{
// 		ID:      insertID,
// 		Message: "User created successfully",
// 	}

// 	// send the response
// 	json.NewEncoder(w).Encode(res)
// }

// // GetUser will return a single user by its id
// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	// get the userid from the request params, key is "id"
// 	params := mux.Vars(r)

// 	// convert the id type from string to int
// 	id, err := strconv.Atoi(params["id"])

// 	if err != nil {
// 		log.Fatalf("Unable to convert the string into int.  %v", err)
// 	}

// 	// call the getUser function with user id to retrieve a single user
// 	user, err := getUser(int64(id))

// 	if err != nil {
// 		log.Fatalf("Unable to get user. %v", err)
// 	}

// 	// send the response
// 	json.NewEncoder(w).Encode(user)
// }

// // GetAllUser will return all the users
// func GetAllUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	// get all the users in the db
// 	users, err := getAllUsers()

// 	if err != nil {
// 		log.Fatalf("Unable to get all user. %v", err)
// 	}

// 	// send all the users as response
// 	json.NewEncoder(w).Encode(users)
// }

// // UpdateUser update user's detail in the postgres db
// func UpdateUser(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "PUT")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	// get the userid from the request params, key is "id"
// 	params := mux.Vars(r)

// 	// convert the id type from string to int
// 	id, err := strconv.Atoi(params["id"])

// 	if err != nil {
// 		log.Fatalf("Unable to convert the string into int.  %v", err)
// 	}

// 	// create an empty user of type models.User
// 	var user models.User

// 	// decode the json request to user
// 	err = json.NewDecoder(r.Body).Decode(&user)

// 	if err != nil {
// 		log.Fatalf("Unable to decode the request body.  %v", err)
// 	}

// 	// call update user to update the user
// 	updatedRows := updateUser(int64(id), user)

// 	// format the message string
// 	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

// 	// format the response message
// 	res := response{
// 		ID:      int64(id),
// 		Message: msg,
// 	}

// 	// send the response
// 	json.NewEncoder(w).Encode(res)
// }

// // DeleteUser delete user's detail in the postgres db
// func DeleteUser(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	// get the userid from the request params, key is "id"
// 	params := mux.Vars(r)

// 	// convert the id in string to int
// 	id, err := strconv.Atoi(params["id"])

// 	if err != nil {
// 		log.Fatalf("Unable to convert the string into int.  %v", err)
// 	}

// 	// call the deleteUser, convert the int to int64
// 	deletedRows := deleteUser(int64(id))

// 	// format the message string
// 	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

// 	// format the reponse message
// 	res := response{
// 		ID:      int64(id),
// 		Message: msg,
// 	}

// 	// send the response
// 	json.NewEncoder(w).Encode(res)
// }

// //------------------------- handler functions ----------------
// // insert one user in the DB
func InsertUser(user models.User) (int64, error) {
	db, err := CreateConnection()
	if err != nil {
		return 0, err
	}
	if db != nil {
		defer db.Close()
	}

	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`
	var id int64
	err = db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetUser(id int64) (*models.User, error) {
	// create the postgres db connection
	var err error
	if db, err := CreateConnection(); db != nil {
		defer db.Close()
		user := new(models.User)

		sqlStatement := `SELECT * FROM users WHERE userid=$1`
		row := db.QueryRow(sqlStatement, id)

		err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

		switch err {
		case nil:
			return user, nil
		case sql.ErrNoRows:
		default:
			return nil, err
		}
	}

	return nil, err
}

// // get one user from the DB by its userid
// func getAllUsers() ([]models.User, error) {
// 	// create the postgres db connection
// 	db := createConnection()

// 	// close the db connection
// 	defer db.Close()

// 	var users []models.User

// 	// create the select sql query
// 	sqlStatement := `SELECT * FROM users`

// 	// execute the sql statement
// 	rows, err := db.Query(sqlStatement)

// 	if err != nil {
// 		log.Fatalf("Unable to execute the query. %v", err)
// 	}

// 	// close the statement
// 	defer rows.Close()

// 	// iterate over the rows
// 	for rows.Next() {
// 		var user models.User

// 		// unmarshal the row object to user
// 		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

// 		if err != nil {
// 			log.Fatalf("Unable to scan the row. %v", err)
// 		}

// 		// append the user in the users slice
// 		users = append(users, user)

// 	}

// 	// return empty user on error
// 	return users, err
// }

// // update user in the DB
// func updateUser(id int64, user models.User) int64 {

// 	// create the postgres db connection
// 	db := createConnection()

// 	// close the db connection
// 	defer db.Close()

// 	// create the update sql query
// 	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`

// 	// execute the sql statement
// 	res, err := db.Exec(sqlStatement, id, user.Name, user.Location, user.Age)

// 	if err != nil {
// 		log.Fatalf("Unable to execute the query. %v", err)
// 	}

// 	// check how many rows affected
// 	rowsAffected, err := res.RowsAffected()

// 	if err != nil {
// 		log.Fatalf("Error while checking the affected rows. %v", err)
// 	}

// 	fmt.Printf("Total rows/record affected %v", rowsAffected)

// 	return rowsAffected
// }

// // delete user in the DB
// func deleteUser(id int64) int64 {

// 	// create the postgres db connection
// 	db := createConnection()

// 	// close the db connection
// 	defer db.Close()

// 	// create the delete sql query
// 	sqlStatement := `DELETE FROM users WHERE userid=$1`

// 	// execute the sql statement
// 	res, err := db.Exec(sqlStatement, id)

// 	if err != nil {
// 		log.Fatalf("Unable to execute the query. %v", err)
// 	}

// 	// check how many rows affected
// 	rowsAffected, err := res.RowsAffected()

// 	if err != nil {
// 		log.Fatalf("Error while checking the affected rows. %v", err)
// 	}

// 	fmt.Printf("Total rows/record affected %v", rowsAffected)

// 	return rowsAffected
// }
