package deserialize

import (
	"errors"
	"fmt"

	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/atlarge-research/opendc-emulate-kubernetes/api/controlplane"
	apiEvents "github.com/atlarge-research/opendc-emulate-kubernetes/api/controlplane/events"
	"github.com/atlarge-research/opendc-emulate-kubernetes/api/scenario"
	anyMarshal "github.com/atlarge-research/opendc-emulate-kubernetes/pkg/any"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/scenario/events"
	"github.com/atlarge-research/opendc-emulate-kubernetes/pkg/scenario/normalization/translate"
)

type protoEventMap map[int32]*anypb.Any

type customFlagParser struct {
	// A parsed public scenario, only missing the custom flags events
	scenario *controlplane.PublicScenario
}

func (cfp *customFlagParser) parse(json string) error {
	// JSON object looks as such: tasks[x].{custom_flags, pod_configs[y].custom_flags}
	for i, task := range gjson.Get(json, "tasks").Array() {
		currentParsedTask := cfp.scenario.Tasks[i]

		// If the current task has custom flags
		if hasCustomFlags(task) {
			taskCustomFlags := make(protoEventMap)

			// Parse them
			if err := cfp.parseCustomFlags(task, &taskCustomFlags); err != nil {
				return err
			}
			currentParsedTask.NodeEvent = &controlplane.Task_CustomFlags{CustomFlags: &apiEvents.CustomFlags{CustomFlags: taskCustomFlags}}
		}

		// Iterate over the pod configs
		for j, podConfig := range task.Get("pod_configs").Array() {
			// If this pod config has custom flags
			if hasCustomFlags(podConfig) {
				podCustomFlags := make(protoEventMap)

				// Parse them
				if err := cfp.parseCustomFlags(podConfig, &podCustomFlags); err != nil {
					return err
				}
				currentParsedTask.PodConfigs[j].PodEvent = &controlplane.PodConfig_CustomFlags{CustomFlags: &apiEvents.CustomFlags{CustomFlags: podCustomFlags}}
			}
		}
	}

	return nil
}

func (cfp *customFlagParser) parseCustomFlags(objectWithCustomFlags gjson.Result, customFlags *protoEventMap) error {
	// Iterate over all set custom flags and fill the given map
	for k, v := range objectWithCustomFlags.Get("custom_flags").Map() {
		ef, anyValue, err := cfp.parseKey(k, v)
		if err != nil {
			return err
		}

		anyMarshalled, err := anyMarshal.Marshal(anyValue)
		if err != nil {
			return err
		}

		(*customFlags)[ef] = anyMarshalled
	}

	return nil
}

func (cfp *customFlagParser) parseKey(key string, value gjson.Result) (ef events.EventFlag, response interface{}, err error) {
	switch key {
	case "node_create_pod_response":
		ef = events.NodeCreatePodResponse
		response, err = getResponse(value)
	case "node_create_pod_response_percentage":
		ef = events.NodeCreatePodResponsePercentage
		response, err = getPercent(value)

	case "node_update_pod_response":
		ef = events.NodeUpdatePodResponse
		response, err = getResponse(value)
	case "node_update_pod_response_percentage":
		ef = events.NodeUpdatePodResponsePercentage
		response, err = getPercent(value)

	case "node_delete_pod_response":
		ef = events.NodeDeletePodResponse
		response, err = getResponse(value)
	case "node_delete_pod_response_percentage":
		ef = events.NodeDeletePodResponsePercentage
		response, err = getPercent(value)

	case "node_get_pod_response":
		ef = events.NodeGetPodResponse
		response, err = getResponse(value)
	case "node_get_pod_response_percentage":
		ef = events.NodeGetPodResponsePercentage
		response, err = getPercent(value)

	case "node_get_pod_status_response":
		ef = events.NodeGetPodStatusResponse
		response, err = getResponse(value)
	case "node_get_pod_status_response_percentage":
		ef = events.NodeGetPodStatusResponsePercentage
		response, err = getPercent(value)

	case "node_get_pods_response":
		ef = events.NodeGetPodsResponse
		response, err = getResponse(value)
	case "node_get_pods_response_percentage":
		ef = events.NodeGetPodsResponsePercentage
		response, err = getPercent(value)

	case "node_ping_response":
		ef = events.NodePingResponse
		response, err = getResponse(value)
	case "node_ping_response_percentage":
		ef = events.NodePingResponsePercentage
		response, err = getPercent(value)

	case "node_enable_resource_alteration":
		return events.NodeEnableResourceAlteration, value.Bool(), nil
	case "node_memory_usage":
		ef = events.NodeMemoryUsage
		response, err = getSize(value, "memory")
	case "node_cpu_usage":
		ef = events.NodeCPUUsage
		response, err = getIntMinZero(value)
	case "node_storage_usage":
		ef = events.NodeStorageUsage
		response, err = getSize(value, "storage")
	case "node_ephemeral_storage_usage":
		ef = events.NodeEphemeralStorageUsage
		response, err = getSize(value, "ephemeral storage")

	case "node_added_latency_enabled":
		return events.NodeAddedLatencyEnabled, value.Bool(), nil
	case "node_added_latency_msec":
		ef = events.NodeAddedLatencyMsec
		response, err = getIntMinZero(value)

	case "pod_create_pod_response":
		ef = events.PodCreatePodResponse
		response, err = getResponse(value)
	case "pod_create_pod_response_percentage":
		ef = events.PodCreatePodResponsePercentage
		response, err = getPercent(value)

	case "pod_update_pod_response":
		ef = events.PodUpdatePodResponse
		response, err = getResponse(value)
	case "pod_update_pod_response_percentage":
		ef = events.PodUpdatePodResponsePercentage
		response, err = getPercent(value)

	case "pod_delete_pod_response":
		ef = events.PodDeletePodResponse
		response, err = getResponse(value)
	case "pod_delete_pod_response_percentage":
		ef = events.PodDeletePodResponsePercentage
		response, err = getPercent(value)

	case "pod_get_pod_response":
		ef = events.PodGetPodResponse
		response, err = getResponse(value)
	case "pod_get_pod_response_percentage":
		ef = events.PodGetPodResponsePercentage
		response, err = getPercent(value)

	case "pod_get_pod_status_response":
		ef = events.PodGetPodStatusResponse
		response, err = getResponse(value)
	case "pod_get_pod_status_response_percentage":
		ef = events.PodGetPodStatusResponsePercentage
		response, err = getPercent(value)

	case "pod_update_pod_status":
		ef = events.PodUpdatePodStatus
		response, err = getPodStatus(value)
	case "pod_update_pod_status_percentage":
		ef = events.PodUpdatePodStatusPercentage
		response, err = getPercent(value)

	default:
		return 0, nil, fmt.Errorf("invalid custom flag key '%s'", key)
	}

	return ef, response, err
}

func hasCustomFlags(objectWithCustomFlags gjson.Result) bool {
	return objectWithCustomFlags.Get("custom_flags").Exists()
}

func getResponse(value gjson.Result) (scenario.Response, error) {
	if response, ok := scenario.Response_value[value.String()]; ok {
		return scenario.Response(response), nil
	}
	return 0, fmt.Errorf("invalid response '%v'", value.String())
}

func getPodStatus(value gjson.Result) (scenario.PodStatus, error) {
	if podStatus, ok := scenario.PodStatus_value[value.String()]; ok {
		return scenario.PodStatus(podStatus), nil
	}
	return 0, fmt.Errorf("invalid pod status '%v'", value.String())
}

func getSize(value gjson.Result, unitName string) (int64, error) {
	inBytes, err := translate.GetInBytes(value.String(), unitName)
	if err != nil {
		return 0, err
	}
	return inBytes, nil
}

func getPercent(value gjson.Result) (int64, error) {
	percent := value.Int()
	if percent < 0 || percent > 100 {
		return 0, errors.New("percentage should be between 0 and 100")
	}
	return percent, nil
}

func getIntMinZero(value gjson.Result) (int64, error) {
	valueInt := value.Int()
	if valueInt < 0 {
		return 0, errors.New("value should be at least 0")
	}
	return valueInt, nil
}