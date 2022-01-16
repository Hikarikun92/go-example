package user

import (
	"database/sql"
)

type Repository interface {
	FindAll() ([]*User, error)
	FindCredentialsByUsername(username string) (*Credentials, error)
}

type repositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) FindAll() ([]*User, error) {
	rows, err := r.db.Query("select id, username from `user`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.Id, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *repositoryImpl) FindCredentialsByUsername(username string) (*Credentials, error) {
	rows, err := r.db.Query("select u.id, u.username, c.password, r.user_id, r.roles from `user` u inner join "+
		"user_credentials c on c.user_id = u.id left join user_roles r on r.user_id = u.id where username = ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//Read all the rows first so that we can perform checks on them later
	type record struct {
		id         int
		username   string
		password   string
		roleUserId sql.NullInt32
		role       sql.NullString
	}
	var records []*record

	for rows.Next() {
		r := record{}
		err := rows.Scan(&r.id, &r.username, &r.password, &r.roleUserId, &r.role)
		if err != nil {
			return nil, err
		}
		records = append(records, &r)
	}

	//If there were no rows, this user doesn't exist
	if len(records) == 0 {
		return nil, nil
	}

	//Create the user and its credentials from the first row
	first := records[0]
	user := &User{Id: first.id, Username: first.username}

	//Go through the other rows for the user's roles
	var roles []string
	for _, rec := range records {
		if rec.roleUserId.Valid { //If userId is not valid, it means there were no roles
			roles = append(roles, rec.role.String)
		}
	}

	credentials := &Credentials{User: user, Password: first.password, Roles: roles}
	return credentials, nil
}
