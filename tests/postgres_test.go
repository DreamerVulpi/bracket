package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/lib/pq"
)

type TestLogger struct {
	T *testing.T
}

const (
	img        = "postgres:16-alpine"
	dbUser     = "user"
	dbPassword = "1234"
	dbName     = "test_db"
)

func (l *TestLogger) AcceptLog(log testcontainers.Log) {
	l.T.Helper()
	l.T.Logf("[PG LOG] %s", string(log.Content))
}

func TestPostgresContainer(t *testing.T) {
	testCtx := context.Background()

	postgresContainer, err := postgres.Run(
		testCtx, img,
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.WithDatabase(dbName),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2)),
	)
	require.NoError(t, err)

	// Гарантируем, что контейнер будет остановлен после завершения теста
	t.Cleanup(func() {
		if err := postgresContainer.Terminate(testCtx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	// 2. Получение DSN для подключения
	connStr, err := postgresContainer.ConnectionString(testCtx, "sslmode=disable")
	require.NoError(t, err)

	// Проверка: DSN должен содержать имя пользователя, пароль, хост, порт и имя БД
	t.Logf("PostgreSQL DSN: %s", connStr)

	// 3. Подключение к базе данных
	var closeErr error
	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)
	defer func() {
		if closeErr == nil {
			closeErr = db.Close()
		}
	}()

	// Проверка соединения
	err = db.Ping()
	require.NoError(t, err)
	t.Log("Successfully connected to PostgreSQL container!")

	// 4. Выполнение простого SQL-запроса (пример использования)
	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	require.NoError(t, err)

	fmt.Printf("PostgreSQL Version: %s\n", version)
	require.NoError(t, closeErr)
}
