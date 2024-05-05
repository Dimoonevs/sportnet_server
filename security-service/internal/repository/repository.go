package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Dimoonevs/SportsApp/auth-service/pkg/data"
	proto "github.com/Dimoonevs/SportsApp/auth-service/proto/security"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	Register(ctx context.Context, req *data.User) (int32, error)
	UsernameAndEmailExist(ctx context.Context, username, email string) (bool, int32, error)
	GetUserPassword(ctx context.Context, username, email string) (string, error)
	GetCoachTimeZone(ctx context.Context, req *proto.GetCoachTimeZoneRequest) (*string, error)
	GetUsername(ctx context.Context, email string) (string, error)
	SaveCode(ctx context.Context, id int32, code string) error
	ConfirmEmailByUserId(ctx context.Context, code string, id int32) error
	GetUserById(ctx context.Context, id int32) (*data.User, error)
	ConfirmResetPassword(ctx context.Context, code, password string, id int32) error
	GetByUsernameOrEmail(ctx context.Context, email, username string) (*data.User, error)
	UpdateCoach(ctx context.Context, req *proto.CoachData) error
	ChangeEmail(ctx context.Context, req *proto.ChangeEmailReq) error
}

type RepositoryPostgres struct {
	DB *pgxpool.Pool
}

func (r *RepositoryPostgres) Register(ctx context.Context, req *data.User) (int32, error) {
	var idUser int32
	conn, err := r.DB.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()
	err = conn.QueryRow(ctx, "INSERT INTO users (first_name, last_name, username, date_of_birth, status, email, password, created_at, updated_at, time_zone, user_active) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW(), $8, false) RETURNING id",
		req.FirstName, req.LastName, req.Username, req.DateOfBirth, req.Status, req.Email, req.Password, req.TimeZone).Scan(&idUser)
	if err != nil {
		return 0, err
	}
	return idUser, nil
}

func (r *RepositoryPostgres) UsernameAndEmailExist(ctx context.Context, username, email string) (bool, int32, error) {
	// var count int
	var exist bool
	var id int32
	conn, err := r.DB.Acquire(ctx)
	if err != nil {
		return false, 0, err
	}
	defer conn.Release()
	if username != "" && email != "" {
		err = conn.QueryRow(ctx, "SELECT EXISTS(SELECT * FROM users WHERE username = $1 AND email = $2) AS exist", username, email).Scan(&exist)
		if exist {
			err = conn.QueryRow(ctx, "SELECT id FROM users WHERE username = $1 AND email = $2", username, email).Scan(&id)
			if err != nil {
				return false, 0, err
			}
			return true, id, nil
		}
		if !exist {
			return false, 0, nil
		}
	}
	if username == "" {
		err = conn.QueryRow(ctx, "SELECT EXISTS(SELECT * FROM users WHERE email = $1) AS exist", email).Scan(&exist)
		if exist {
			err = conn.QueryRow(ctx, "SELECT id FROM users WHERE email = $1", email).Scan(&id)
			if err != nil {
				return false, 0, err
			}
			return true, id, nil
		}
		if !exist {
			return false, 0, nil
		}
	}
	if email == "" {
		err = conn.QueryRow(ctx, "SELECT EXISTS(SELECT * FROM users WHERE username = $1) AS exist", username).Scan(&exist)
		if exist {
			err = conn.QueryRow(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&id)
			if err != nil {
				return false, 0, err
			}
			return true, id, nil
		}
		if !exist {
			return false, 0, nil
		}
	}
	return false, 0, nil
}

func (r *RepositoryPostgres) GetUserPassword(ctx context.Context, username, email string) (string, error) {
	var password string
	conn, err := r.DB.Acquire(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Release()
	err = conn.QueryRow(ctx, "SELECT password FROM users WHERE username = $1 OR email = $2", username, email).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (r *RepositoryPostgres) GetCoachTimeZone(ctx context.Context, req *proto.GetCoachTimeZoneRequest) (*string, error) {
	conn, err := r.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, "SELECT time_zone FROM users WHERE id = $1", req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var timeZone string
	for rows.Next() {
		if err := rows.Scan(&timeZone); err != nil {
			return nil, err
		}
	}
	return &timeZone, nil
}

func (r *RepositoryPostgres) GetUsername(ctx context.Context, email string) (string, error) {
	var username string
	conn, err := r.DB.Acquire(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Release()
	err = conn.QueryRow(ctx, "SELECT username FROM users WHERE email = $1", email).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (r *RepositoryPostgres) SaveCode(ctx context.Context, id int32, code string) error {
	conn, err := r.DB.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	exporationDate := time.Now().Add(30 * time.Minute)
	_, err = conn.Exec(ctx, "UPDATE users SET user_active = false, updated_at = now() WHERE id = $1 ", id)
	if err != nil {
		return err
	}
	_, err = conn.Exec(ctx, "INSERT INTO users_confirms (user_id, code, expiration_time) VALUES ($1, $2, $3) ON CONFLICT (user_id) DO UPDATE SET code = $2, expiration_time = $3", id, code, exporationDate)
	if err != nil {
		return err
	}
	return nil
}
func (r *RepositoryPostgres) ConfirmEmailByUserId(ctx context.Context, code string, id int32) error {
	conn, err := r.DB.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	var reqPayload struct {
		code           string
		exporationDate time.Time
	}

	err = conn.QueryRow(ctx, "SELECT code, expiration_time FROM users_confirms WHERE user_id = $1", id).Scan(&reqPayload.code, &reqPayload.exporationDate)
	if err != nil {
		return err
	}
	if reqPayload.exporationDate.Before(time.Now()) {
		return errors.New("code expired")
	}
	if reqPayload.code != code {
		return errors.New("wrong code")
	}
	_, err = conn.Exec(ctx, "UPDATE users SET user_active = true, updated_at = now() WHERE id = (SELECT user_id FROM users_confirms WHERE code = $1)", code)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryPostgres) GetUserById(ctx context.Context, id int32) (*data.User, error) {
	var user data.User
	conn, err := r.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	err = conn.QueryRow(ctx, "SELECT id, first_name, last_name, username, date_of_birth, status, email, password, time_zone, user_active, created_at, updated_at FROM users WHERE id = $1", id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.DateOfBirth, &user.Status, &user.Email, &user.Password, &user.TimeZone, &user.Active, &user.Created_at, &user.Updated_at)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RepositoryPostgres) ConfirmResetPassword(ctx context.Context, code, password string, id int32) error {
	conn, err := r.DB.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	var reqPayload struct {
		code           string
		exporationDate time.Time
	}
	err = conn.QueryRow(ctx, "SELECT code, expiration_time FROM users_confirms WHERE user_id = $1", id).Scan(&reqPayload.code, &reqPayload.exporationDate)
	if err != nil {
		return err
	}
	if reqPayload.exporationDate.Before(time.Now()) {
		return errors.New("code expired")
	}
	if reqPayload.code != code {
		return errors.New("wrong code")
	}
	_, err = conn.Exec(ctx, "UPDATE users SET password = $1, updated_at = now() WHERE id = $2", password, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryPostgres) GetByUsernameOrEmail(ctx context.Context, email, username string) (*data.User, error) {
	var user data.User
	conn, err := r.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	err = conn.QueryRow(ctx, "SELECT id, first_name, last_name, username, date_of_birth, status, email, password, time_zone, user_active, created_at, updated_at FROM users WHERE username = $1 OR emai = $2", username, email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.DateOfBirth, &user.Status, &user.Email, &user.Password, &user.TimeZone, &user.Active, &user.Created_at, &user.Updated_at)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RepositoryPostgres) UpdateCoach(ctx context.Context, req *proto.CoachData) error {
	conn, err := r.DB.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "UPDATE users SET first_name = $1, last_name = $2, username = $3, date_of_birth = $4, updated_at = now() WHERE id = $5",
		req.FirstName,
		req.LastName,
		req.Username,
		req.DateOfBirth,
		req.Id,
	)
	if err != nil {
		return err
	}
	return nil
}
func (r *RepositoryPostgres) ChangeEmail(ctx context.Context, req *proto.ChangeEmailReq) error {
	conn, err := r.DB.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "UPDATE users SET email = $1, user_active = false, updated_at = now() WHERE id = $2",
		req.Email,
		req.Id,
	)
	if err != nil {
		return err
	}
	return nil
}
