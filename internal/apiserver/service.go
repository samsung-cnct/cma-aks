package apiserver

import (
	"github.com/juju/loggo"
	"github.com/samsung-cnct/cma-aks/pkg/util"
)

var (
	logger loggo.Logger
)

type Server struct{}

func SetLogger() {
	logger = util.GetModuleLogger("internal.cluster-manager-api", loggo.INFO)
}
