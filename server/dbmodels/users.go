package dbmodels

type User struct {
	tableName struct{} `pg:"users"`

	ID        string `pg:"id,pk"`
	FirstName string `pg:"first_name"`
	LastName  string `pg:"last_name"`
	Gender    string `pg:"gender"`
	Email     string `pg:"email"`
	Age       *int32 `pg:"age"`
	Password  string `pg:"password"`
}
