package tests

import (
	"context"
	"os"
	"testing"

	"github.com/DreamerVulpi/bracket/repo"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

const dbDsnEnvVar = "DB_DSN"

func clearConnect(t *testing.T, ctx context.Context, connStr string) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, connStr)
	require.NoError(t, err, "Can't create pool coonection for db")

	err = pool.Ping(ctx)
	require.NoError(t, err, "Can't connect to database CI. Check DSN and state container.")

	t.Log("Successfuly connected to PostgreSQL CI-service!")
	_, err = pool.Exec(ctx, "TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	require.NoError(t, err, "Can't clear table users before test")

	return pool
}

// REFACTOR:
func TestMethodsUser(t *testing.T) {
	TestUser_Add(t)
	TestUser_Edit(t)
	TestUser_Delete(t)
	TestUser_Get(t)
}

func TestUser_Get(t *testing.T) {
	ctx := context.Background()
	connStr := os.Getenv(dbDsnEnvVar)
	if connStr == "" {
		t.Fatalf("Value %s isn't install. Launch in CI!", dbDsnEnvVar)
	}

	pool := clearConnect(t, ctx, connStr)
	defer pool.Close()

	t.Run("Success_DeleteUser", func(t *testing.T) {
		u := &repo.User{
			Conn: pool,
		}

		testNickname := "new_player"
		testHash := "hashed_secret"

		gotID, err := u.Add(testNickname, testHash)

		require.NoError(t, err, "[User.Add] User.Add must will be done!")
		require.Greater(t, gotID, 0, "[User.Add] ID must be positive!")

		var storedNickname string

		err = pool.QueryRow(ctx, "SELECT nickname FROM users WHERE id = $1", gotID).
			Scan(&storedNickname)

		require.NoError(t, err, "[User.Add] Can't get user from database")

		require.Equal(t, testNickname, storedNickname, "[User.Add] Nicknames is different!")

		gotUser, err := u.Get(gotID)
		require.NoError(t, err, "[User.Get] User.Get must gave correct data!")

		require.Equal(t, testNickname, gotUser.Nickname, "[User.Get] Nicknames is different!")
	})
}

func TestUser_Delete(t *testing.T) {
	ctx := context.Background()
	connStr := os.Getenv(dbDsnEnvVar)
	if connStr == "" {
		t.Fatalf("Value %s isn't install. Launch in CI!", dbDsnEnvVar)
	}

	pool := clearConnect(t, ctx, connStr)
	defer pool.Close()

	t.Run("Success_DeleteUser", func(t *testing.T) {
		u := &repo.User{
			Conn: pool,
		}

		testNickname := "new_player"
		testHash := "hashed_secret"

		gotID, err := u.Add(testNickname, testHash)

		require.NoError(t, err, "[User.Add] User.Add must will be done!")
		require.Greater(t, gotID, 0, "[User.Add] ID must be positive!")

		var storedNickname string
		var storedHash string

		err = pool.QueryRow(ctx, "SELECT nickname, password_hash FROM users WHERE id = $1", gotID).
			Scan(&storedNickname, &storedHash)

		require.NoError(t, err, "[User.Add] Can't get user from database")

		require.Equal(t, testNickname, storedNickname, " [User.Add]Nicknames is different!")
		require.Equal(t, testHash, storedHash, "[User.Add] Hashes is different")

		err = u.Delete(gotID)
		require.NoError(t, err, "[User.Delete] User.Delete must delete correct without errors!")
	})
}

func TestUser_Edit(t *testing.T) {
	ctx := context.Background()
	connStr := os.Getenv(dbDsnEnvVar)
	if connStr == "" {
		t.Fatalf("Value %s isn't install. Launch in CI!", dbDsnEnvVar)
	}

	pool := clearConnect(t, ctx, connStr)
	defer pool.Close()

	t.Run("Success_EditUser", func(t *testing.T) {
		u := &repo.User{
			Conn: pool,
		}

		testNickname := "new_player"
		testHash := "hashed_secret"
		gotID, err := u.Add(testNickname, testHash)
		require.NoError(t, err, "[User.Add] must will be done!")
		require.Greater(t, gotID, 0, "[User.Add] ID must be positive!")

		gotUser, err := u.Get(1)
		require.NoError(t, err, "[User.Get] must will be in stock!")
		require.NotNil(t, gotUser, "[User.Get] must gave result with data!")

		editedUser := gotUser
		editedUser.Nickname = "omg"
		err = u.Edit(editedUser)
		require.NoError(t, err, "[User.Edit] must work correctly!")

		var storedNickname string

		err = pool.QueryRow(ctx, "SELECT nickname FROM users WHERE id = $1", gotUser.Id).
			Scan(&storedNickname)

		require.NoError(t, err, "[User.Get] Can't get user from database")
		require.Equal(t, "omg", storedNickname, "[User.Get] Nicknames is different!")
	})

}

func TestUser_Add(t *testing.T) {
	ctx := context.Background()
	connStr := os.Getenv(dbDsnEnvVar)
	if connStr == "" {
		t.Fatalf("Value %s isn't install. Launch in CI!", dbDsnEnvVar)
	}

	pool := clearConnect(t, ctx, connStr)
	defer pool.Close()

	t.Run("Success_AddNewUser", func(t *testing.T) {
		u := &repo.User{
			Conn: pool,
		}

		testNickname := "new_player"
		testHash := "hashed_secret"

		gotID, err := u.Add(testNickname, testHash)

		require.NoError(t, err, "[User.Add] User.Add must will be done!")
		require.Greater(t, gotID, 0, "[User.Add] ID must be positive!")

		var storedNickname string
		var storedHash string

		err = pool.QueryRow(ctx, "SELECT nickname, password_hash FROM users WHERE id = $1", gotID).
			Scan(&storedNickname, &storedHash)

		require.NoError(t, err, "[User.Add] Can't get user from database")

		require.Equal(t, testNickname, storedNickname, " [User.Add]Nicknames is different!")
		require.Equal(t, testHash, storedHash, "[User.Add] Hashes is different")
	})

	t.Run("Error_DuplicateNickname", func(t *testing.T) {
		_, err := pool.Exec(ctx, "TRUNCATE TABLE users RESTART IDENTITY CASCADE")
		require.NoError(t, err)

		u := &repo.User{Conn: pool}

		testNickname := "duplicate_nick"
		_, err = u.Add(testNickname, "first_hash")
		require.NoError(t, err)

		_, err = u.Add(testNickname, "second_hash")
		require.Error(t, err, "[User.Add] Was waited error for dublicate nickname")
	})
}
