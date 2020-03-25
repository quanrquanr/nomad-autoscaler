package plugins

import (
	"os"
	"path"

	"github.com/hashicorp/nomad-autoscaler/apm"
	nomadapm "github.com/hashicorp/nomad-autoscaler/plugins/nomad/apm"
	nomadtarget "github.com/hashicorp/nomad-autoscaler/plugins/nomad/target"
	prometheus "github.com/hashicorp/nomad-autoscaler/plugins/prometheus/apm"
	targetvalue "github.com/hashicorp/nomad-autoscaler/plugins/target-value/strategy"
	"github.com/hashicorp/nomad-autoscaler/strategy"
	"github.com/hashicorp/nomad-autoscaler/target"
)

const (
	NomadAPM            = "nomad-apm"
	NomadTarget         = "nomad"
	PrometheusAPM       = "prometheus"
	TargetValueStrategy = "target-value"
)

func IsInternal(driver, pluginDir string) bool {
	// Use a plugin binary if one is available
	if _, err := os.Stat(path.Join(pluginDir, driver)); err == nil {
		return false
	}

	switch driver {
	case
		NomadAPM,
		NomadTarget,
		PrometheusAPM,
		TargetValueStrategy:
		return true
	}
	return false
}

func NewInternalAPM(driver string) apm.APM {
	switch driver {
	case NomadAPM:
		return &nomadapm.MetricsAPM{}
	case PrometheusAPM:
		return &prometheus.APM{}
	}
	return nil
}

func NewInternalStrategy(driver string) strategy.Strategy {
	switch driver {
	case TargetValueStrategy:
		return &targetvalue.Strategy{}
	}
	return nil
}

func NewInternalTarget(driver string) target.Target {
	switch driver {
	case NomadTarget:
		return &nomadtarget.NomadGroupCount{}
	}
	return nil
}
