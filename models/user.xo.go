package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// User represents a row from 'test.users'.
type User struct {
	ID        string    `json:"id"`         // ID
	Name      string    `json:"name"`       // ユーザ名
	CreatedAt time.Time `json:"created_at"` // created_at
	UpdatedAt time.Time `json:"updated_at"` // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [User] exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted returns true when the [User] has been marked for deletion
// from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the [User] to the database.
func (u *User) Insert(ctx context.Context, db DB) error {
	switch {
	case u._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case u._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO test.users (` +
		`id, name, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, u.ID, u.Name, u.CreatedAt, u.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID, u.Name, u.CreatedAt, u.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Update updates a [User] in the database.
func (u *User) Update(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case u._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE test.users SET ` +
		`name = ?, created_at = ?, updated_at = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, u.Name, u.CreatedAt, u.UpdatedAt, u.ID)
	if _, err := db.ExecContext(ctx, sqlstr, u.Name, u.CreatedAt, u.UpdatedAt, u.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [User] to the database.
func (u *User) Save(ctx context.Context, db DB) error {
	if u.Exists() {
		return u.Update(ctx, db)
	}
	return u.Insert(ctx, db)
}

// Upsert performs an upsert for [User].
func (u *User) Upsert(ctx context.Context, db DB) error {
	switch {
	case u._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO test.users (` +
		`id, name, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`id = VALUES(id), name = VALUES(name), created_at = VALUES(created_at), updated_at = VALUES(updated_at)`
	// run
	logf(sqlstr, u.ID, u.Name, u.CreatedAt, u.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID, u.Name, u.CreatedAt, u.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Delete deletes the [User] from the database.
func (u *User) Delete(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return nil
	case u._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM test.users ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, u.ID)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	u._deleted = true
	return nil
}

// UserByID retrieves a row from 'test.users' as a [User].
//
// Generated from index 'users_id_pkey'.
func UserByID(ctx context.Context, db DB, id string) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, created_at, updated_at ` +
		`FROM test.users ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&u.ID, &u.Name, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}
