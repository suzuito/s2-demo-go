package inject

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/common-go/cgcp"
	"github.com/suzuito/s2-demo-go/setting"
	"golang.org/x/xerrors"
)

func NewImplement() (
	*setting.Env,
	*cgcp.GCPContextResourceGenerator,
	error,
) {
	env := setting.Env{}
	if err := envconfig.Process("", &env); err != nil {
		return nil, nil, xerrors.Errorf("Cannot process env : %w", err)
	}
	genGCP := cgcp.NewGCPContextResourceGenerator()
	genGCP.GCS()
	return &env, genGCP, nil
}
