package storage

// for integration test only
import (
	"context"
	"fmt"
	"sthl/config"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
)

// PostgresContainer represents the postgres container type used in the module
type PostgresContainer struct {
	testcontainers.Container
}

type postgresContainerOption func(req *testcontainers.ContainerRequest)

func WithWaitStrategy(strategies ...wait.Strategy) func(req *testcontainers.ContainerRequest) {
	return func(req *testcontainers.ContainerRequest) {
		req.WaitingFor = wait.ForAll(strategies...).WithDeadline(1 * time.Minute)
	}
}

func WithPort(port string) func(req *testcontainers.ContainerRequest) {
	return func(req *testcontainers.ContainerRequest) {
		req.ExposedPorts = append(req.ExposedPorts, port)
	}
}

func WithInitialDatabase(user string, password string, dbName string) func(req *testcontainers.ContainerRequest) {
	return func(req *testcontainers.ContainerRequest) {
		req.Env["POSTGRES_USER"] = user
		req.Env["POSTGRES_PASSWORD"] = password
		req.Env["POSTGRES_DB"] = dbName
	}
}

// startContainer creates an instance of the postgres container type
func startContainer(ctx context.Context, opts ...postgresContainerOption) (*PostgresContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:14-alpine",
		Env:          map[string]string{},
		ExposedPorts: []string{},
		Cmd:          []string{"postgres", "-c", "fsync=off"},
	}

	for _, opt := range opts {
		opt(&req)
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{Container: container}, nil
}

func NewPostgresTestContainer(l *zap.Logger, c *config.Config) (*PostgresContainer, error) {
	ctx := context.TODO()
	const dbname = "postgres"

	port, err := nat.NewPort("tcp", "5432")
	if err != nil {
		fmt.Println("port err:", err)
		return nil, err
	}

	container, err := startContainer(ctx,
		WithPort(port.Port()),
		WithInitialDatabase(c.GetDbUser(), c.GetDbPw(), dbname),
		WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		l.Info("fail to start container", zap.Error(err))
		return nil, err
	}

	containerPort, err := container.MappedPort(ctx, port)
	if err != nil {
		l.Info("fail to container.MappedPort", zap.Error(err))
		return nil, err
	}

	// host, err := container.Host(ctx)
	// if err != nil {
	// 	l.Info("fail to container.Host", zap.Error(err))
	// 	return nil, err
	// }

	// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, containerPort.Port(), user, password, dbname)
	// fmt.Println("connStr:", connStr)

	c.SetDbPort(containerPort.Int())
	return container, nil
}
