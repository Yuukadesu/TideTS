package metrics

import "github.com/hanami/tidets/core/queryengine/plan"

func planKindLabel(kind plan.Kind) string {
	switch kind {
	case plan.KindInsert:
		return "insert"
	case plan.KindSelect:
		return "select"
	case plan.KindCreateTimeseries:
		return "create_timeseries"
	case plan.KindShowDevices:
		return "show_devices"
	case plan.KindShowTimeseries:
		return "show_timeseries"
	case plan.KindDelete:
		return "delete"
	default:
		return "unknown"
	}
}
