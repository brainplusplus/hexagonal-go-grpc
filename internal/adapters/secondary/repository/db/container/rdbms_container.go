package container

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"net"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // import file driver
	"github.com/testcontainers/testcontainers-go"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_USERNAME string = "postgres"
	DB_PASSWORD string = "postgres"
	DB_NAME     string = "hexagonal_go_grpc"
	DB_PORT     string = "15432"
	DB_DIR      string
)

var db *gorm.DB
var m *migrate.Migrate

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	DB_DIR = dir
	os.Chdir(dir)
	os.Chdir("../../../../../")
}

// StartContainer creates an instance of the mysql container type
func StartContainer() testcontainers.Container {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	//dbPort := nat.Port(fmt.Sprintf("%s/tcp", DB_PORT))
	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		Started: true,
		Reuse:   true,
		ContainerRequest: testcontainers.ContainerRequest{
			Name:         "hexagonal-go-grpc-testcontainer",
			Image:        "postgres:16",
			Hostname:     GetOutboundIP().String(),
			ExposedPorts: []string{DB_PORT + ":5432/tcp"},
			Env: map[string]string{
				"POSTGRES_USER":     DB_USERNAME,
				"POSTGRES_PASSWORD": DB_PASSWORD,
				"POSTGRES_DB":       DB_NAME,
			},
			WaitingFor: wait.ForListeningPort("5432/tcp"),
		},
	})
	if err != nil {
		panic(err)
	}

	return c
}

func StopContainer() {
	c := StartContainer()
	c.Terminate(context.Background())
}

func GetGorm() (*gorm.DB, error) {
	if db != nil {
		//clear first
		m.Down()
		//migrate again for fresh data
		m.Up()
		return db, nil
	}

	container := StartContainer()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT, "disable")
	gormdb, err := gorm.Open(gormPostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	d, err := gormdb.DB()
	if err != nil {
		return nil, err
	}

	d.SetMaxIdleConns(0)

	driver, err := postgres.WithInstance(d, &postgres.Config{DatabaseName: DB_NAME})
	if err != nil {
		return nil, err
	}

	m, err = migrate.NewWithDatabaseInstance("file://migrations/postgres", DB_NAME, driver)
	if err != nil {
		return nil, err
	}

	//clear first
	m.Down()
	//migrate again for fresh data
	m.Up()
	db = gormdb
	//time.Sleep(600 * time.Second)
	return gormdb, nil
}

// GetOutboundIP Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
