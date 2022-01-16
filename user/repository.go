package user

import "database/sql"

type Repository interface {
	FindAll() ([]*User, error)
	FindCredentialsByUsername(username string) *Credentials
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
		var (
			id       int
			username string
		)
		err := rows.Scan(&id, &username)
		if err != nil {
			return nil, err
		}
		users = append(users, &User{Id: id, Username: username})
	}

	return users, nil
}

func (*repositoryImpl) FindCredentialsByUsername(username string) *Credentials {
	switch username {
	case "user1":
		return &Credentials{User: &User{Id: 1, Username: "user1"}, Password: "pass1", Roles: []string{"ROLE_ADMIN", "ROLE_USER"}}
	case "user2":
		return &Credentials{User: &User{Id: 2, Username: "user2"}, Password: "pass2", Roles: []string{"ROLE_USER"}}
	default:
		return nil
	}
}
