package suite

import (
	"context"
	"ssos/internal/config"
	"testing"

	"github.com/Pavel910862/protos/protos/gen/go/sso"
)

type Suite struct {
	*testing.T                //для вызова методов *testing.T внутри Suite
	Cfg        *config.Config //конфигурация приложения
	AuthClient sso.AuthClient //клиент для взаимодействия с gRPC-сервером
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()   //при фэйле теста правильно формировался стек вызов и функция не была финальная
	t.Parallel() // тесты параллельно

}
