package auth

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(model, policy string) *Authorizer {
	enforcer, err := casbin.NewEnforcer(model, policy)
	if err != nil {
		panic(err)
	}
	return &Authorizer{
		enforcer: enforcer,
	}
}

type Authorizer struct {
	enforcer *casbin.Enforcer
}

func (a *Authorizer) Authorize(subject, object, action string) error {
	permitted, err := a.enforcer.Enforce(subject, object, action)
	if err != nil {
		return err
	}
	if !permitted {
		msg := fmt.Sprintf("%s not permitted to %s to %s", subject, object, action)
		st := status.New(codes.PermissionDenied, msg)
		return st.Err()
	}
	return nil
}
