package suite

import (
	"context"
	"net"
	"ssos/internal/config"
	"strconv"
	"testing"

	"github.com/Pavel910862/protos/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T                //для вызова методов *testing.T внутри Suite
	Cfg        *config.Config //конфигурация приложения
	AuthClient sso.AuthClient //клиент для взаимодействия с gRPC-сервером
}

const (
	grpcHost = "localhost"
)

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()   //при фэйле теста правильно формировался стек вызов и функция не была финальная
	t.Parallel() // тесты параллельно

	// перед этим накатить тестовые миграции
	// сначала запускаем сервак go run cmd/sso/main.go --config=./сonfig/local.yaml (suite.go cfg := config.MustLoadByPath("./сonfig/local.yaml"))
	// потом меняем cfg := config.MustLoadByPath("C:/Users/Алёна Валерьевна/Desktop/sso/сonfig/local.yaml" или "../сonfig/local.yaml") и запускаем тест
	cfg := config.MustLoadByPath("../сonfig/local.yaml") //C:/Users/Алёна Валерьевна/Desktop/sso/сonfig/local.yaml //

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.DialContext(context.Background(),
		grpcAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials())) // Используем insecure-коннект для тестов
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		AuthClient: sso.NewAuthClient(cc),
	}
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}
