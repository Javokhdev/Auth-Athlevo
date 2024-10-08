package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	pb "auth-athlevo/genproto/auth"

	_ "github.com/lib/pq"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetProfile(req *pb.GetByIdReq) (*pb.UserRepeated, error) {
	var users []*pb.UserRes
	query := `SELECT id, username, gym_id, phone_number, email, full_name, date_of_birth, role, gender FROM users WHERE 1=1`

	var args []interface{}
	i := 1

	if req.Id != "" {
		query += fmt.Sprintf(" AND id = $%d", i)
		args = append(args, req.Id)
		i++
	}
	if req.Username != "" {
		query += fmt.Sprintf(" AND username ILIKE $%d", i)
		args = append(args, "%"+req.Username+"%")
		i++
	}
	if req.GymId != "" {
		query += fmt.Sprintf(" AND gym_id = $%d", i)
		args = append(args, req.GymId)
		i++
	}
	if req.PhoneNumber != "" {
		query += fmt.Sprintf(" AND phone_number ILIKE $%d", i)
		args = append(args, "%"+req.PhoneNumber+"%")
		i++
	}

	if req.FullName != "" {
		query += fmt.Sprintf(" AND full_name ILIKE $%d", i)
		args = append(args, "%"+req.FullName+"%")
		i++
	}
	if req.Email != "" {
		query += fmt.Sprintf(" AND email ILIKE $%d", i)
		args = append(args, "%"+req.Email+"%")
		i++
	}
	if req.Gender != "" {
		query += fmt.Sprintf(" AND gender = $%d", i)
		args = append(args, req.Gender)
		i++
	}
	// Add other filters as needed (e.g., FaceId, PhoneNumber, DateOfBirth, Role)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var res pb.UserRes
		var date string
		err = rows.Scan(
			&res.Id,
			&res.Username,
			&res.GymId,
			&res.PhoneNumber,
			&res.Email,
			&res.FullName,
			&date,
			&res.Role,
			&res.Gender,
		)
		if err != nil {
			return nil, err
		}
		res.DateOfBirth = date[:10]
		users = append(users, &res)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.UserRepeated{Users: users}, nil
}

func (r *UserRepo) EditProfile(req *pb.UserRes) (*pb.UserRes, error) {
	res := &pb.UserRes{}

	query := `UPDATE users SET updated_at = NOW()`

	var arg []interface{}
	var conditions []string

	if req.Username != "" && req.Username != "string" {
		arg = append(arg, req.Username)
		conditions = append(conditions, fmt.Sprintf("username = $%d", len(arg)))
	}

	if req.PhoneNumber != "" && req.PhoneNumber != "string" {
		arg = append(arg, req.PhoneNumber)
		conditions = append(conditions, fmt.Sprintf("phone_number = $%d", len(arg)))
	}

	if req.GymId != "" && req.GymId != "string" {
		arg = append(arg, req.GymId)
		conditions = append(conditions, fmt.Sprintf("gym_id = $%d", len(arg)))
	}

	if req.Email != "" && req.Email != "string" {
		arg = append(arg, req.Email)
		conditions = append(conditions, fmt.Sprintf("email = $%d", len(arg)))
	}

	if req.FullName != "" && req.FullName != "string" {
		arg = append(arg, req.FullName)
		conditions = append(conditions, fmt.Sprintf("full_name = $%d", len(arg)))
	}

	if req.DateOfBirth != "" && req.DateOfBirth != "string" {
		arg = append(arg, req.DateOfBirth)
		conditions = append(conditions, fmt.Sprintf("date_of_birth = $%d", len(arg)))
	}

	if req.Gender != "" && req.Gender != "string" {
		arg = append(arg, req.Gender)
		conditions = append(conditions, fmt.Sprintf("gender = $%d", len(arg)))
	}

	if len(conditions) > 0 {
		query += ", " + strings.Join(conditions, ", ")
	}

	query += fmt.Sprintf(" WHERE id = $%d", len(arg)+1)
	arg = append(arg, req.Id)

	_, err := r.db.Exec(query, arg...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *UserRepo) ChangePassword(req *pb.ChangePasswordReq) (*pb.ChangePasswordRes, error) {
	res := &pb.ChangePasswordRes{Message: "Password changed successfully"}

	query := `SELECT password FROM users WHERE id = $1`
	var password string
	err := r.db.QueryRow(query, req.Id).Scan(&password)
	if err != nil {
		return nil, err
	}

	if password != req.CurrentPassword {
		return nil, fmt.Errorf("invalid current password")
	}

	query = `UPDATE users SET updated_at = NOW(), password = $1 WHERE id = $2`
	_, err = r.db.Exec(query, req.NewPassword, req.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *UserRepo) GetSetting(req *pb.GetById) (*pb.Setting, error) {
	res := &pb.Setting{}

	query := `SELECT privacy_level, notification, language, theme FROM settings WHERE user_id = $1`
	err := r.db.QueryRow(query, req.Id).
		Scan(
			&res.PrivacyLevel,
			&res.Notification,
			&res.Language,
			&res.Theme,
		)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *UserRepo) EditSetting(req *pb.SettingReq) (*pb.SettingRes, error) {
	res := &pb.SettingRes{Message: "Settings updated successfully"}

	query := `UPDATE settings SET updated_at = NOW()`

	var arg []interface{}
	var conditions []string

	if req.PrivacyLevel != "" && req.PrivacyLevel != "string" {
		arg = append(arg, req.PrivacyLevel)
		conditions = append(conditions, fmt.Sprintf("privacy_level = $%d", len(arg)))
	}

	if req.Notification != "" && req.Notification != "string" {
		arg = append(arg, req.Notification)
		conditions = append(conditions, fmt.Sprintf("notification = $%d", len(arg)))
	}

	if req.Language != "" && req.Language != "string" {
		arg = append(arg, req.Language)
		conditions = append(conditions, fmt.Sprintf("language = $%d", len(arg)))
	}

	if req.Theme != "" && req.Theme != "string" {
		arg = append(arg, req.Theme)
		conditions = append(conditions, fmt.Sprintf("theme = $%d", len(arg)))
	}

	if len(conditions) > 0 {
		query += ", " + strings.Join(conditions, ", ")
	}

	query += fmt.Sprintf(" WHERE user_id = $%d", len(arg)+1)
	arg = append(arg, req.Id)
	_, err := r.db.Exec(query, arg...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *UserRepo) DeleteUser(req *pb.GetById) (*pb.DeleteRes, error) {
	res := &pb.DeleteRes{Message: "User deleted successfully"}

	query := `UPDATE users SET deleted_at = $1 WHERE id = $2`
	_, err := r.db.Exec(query, time.Now().Unix(), req.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
