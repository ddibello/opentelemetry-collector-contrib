// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func TestDefaultMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	mb := NewMetricsBuilder(DefaultMetricsSettings(), component.BuildInfo{}, WithStartTime(start))
	enabledMetrics := make(map[string]bool)

	enabledMetrics["active_directory.ds.bind.rate"] = true
	mb.RecordActiveDirectoryDsBindRateDataPoint(ts, 1, AttributeBindType(1))

	enabledMetrics["active_directory.ds.ldap.bind.last_successful.time"] = true
	mb.RecordActiveDirectoryDsLdapBindLastSuccessfulTimeDataPoint(ts, 1)

	enabledMetrics["active_directory.ds.ldap.bind.rate"] = true
	mb.RecordActiveDirectoryDsLdapBindRateDataPoint(ts, 1)

	enabledMetrics["active_directory.ds.ldap.client.session.count"] = true
	mb.RecordActiveDirectoryDsLdapClientSessionCountDataPoint(ts, 1)

	enabledMetrics["active_directory.ds.ldap.search.rate"] = true
	mb.RecordActiveDirectoryDsLdapSearchRateDataPoint(ts, 1)

	enabledMetrics["active_directory.ds.name_cache.hit_rate"] = true
	mb.RecordActiveDirectoryDsNameCacheHitRateDataPoint(ts, 1)

	enabledMetrics["active_directory.ds.notification.queued"] = true
	mb.RecordActiveDirectoryDsNotificationQueuedDataPoint(ts, 1)

	enabledMetrics["active_directory.ds.operation.rate"] = true
	mb.RecordActiveDirectoryDsOperationRateDataPoint(ts, 1, AttributeOperationType(1))

	enabledMetrics["active_directory.ds.replication.network.io"] = true
	mb.RecordActiveDirectoryDsReplicationNetworkIoDataPoint(ts, 1, AttributeDirection(1), AttributeNetworkDataType(1))

	enabledMetrics["active_directory.ds.replication.object.rate"] = true
	mb.RecordActiveDirectoryDsReplicationObjectRateDataPoint(ts, 1, AttributeDirection(1))

	enabledMetrics["active_directory.ds.replication.operation.pending"] = true
	mb.RecordActiveDirectoryDsReplicationOperationPendingDataPoint(ts, 1)

	enabledMetrics["active_directory.ds.replication.property.rate"] = true
	mb.RecordActiveDirectoryDsReplicationPropertyRateDataPoint(ts, 1, AttributeDirection(1))

	enabledMetrics["active_directory.ds.replication.sync.object.pending"] = true
	mb.RecordActiveDirectoryDsReplicationSyncObjectPendingDataPoint(ts, 1)

	enabledMetrics["active_directory.ds.replication.sync.request.count"] = true
	mb.RecordActiveDirectoryDsReplicationSyncRequestCountDataPoint(ts, 1, AttributeSyncResult(1))

	enabledMetrics["active_directory.ds.replication.value.rate"] = true
	mb.RecordActiveDirectoryDsReplicationValueRateDataPoint(ts, 1, AttributeDirection(1), AttributeValueType(1))

	enabledMetrics["active_directory.ds.security_descriptor_propagations_event.queued"] = true
	mb.RecordActiveDirectoryDsSecurityDescriptorPropagationsEventQueuedDataPoint(ts, 1)

	enabledMetrics["active_directory.ds.suboperation.rate"] = true
	mb.RecordActiveDirectoryDsSuboperationRateDataPoint(ts, 1, AttributeSuboperationType(1))

	enabledMetrics["active_directory.ds.thread.count"] = true
	mb.RecordActiveDirectoryDsThreadCountDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	sm := metrics.ResourceMetrics().At(0).ScopeMetrics()
	assert.Equal(t, 1, sm.Len())
	ms := sm.At(0).Metrics()
	assert.Equal(t, len(enabledMetrics), ms.Len())
	seenMetrics := make(map[string]bool)
	for i := 0; i < ms.Len(); i++ {
		assert.True(t, enabledMetrics[ms.At(i).Name()])
		seenMetrics[ms.At(i).Name()] = true
	}
	assert.Equal(t, len(enabledMetrics), len(seenMetrics))
}

func TestAllMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	settings := MetricsSettings{
		ActiveDirectoryDsBindRate:                                  MetricSettings{Enabled: true},
		ActiveDirectoryDsLdapBindLastSuccessfulTime:                MetricSettings{Enabled: true},
		ActiveDirectoryDsLdapBindRate:                              MetricSettings{Enabled: true},
		ActiveDirectoryDsLdapClientSessionCount:                    MetricSettings{Enabled: true},
		ActiveDirectoryDsLdapSearchRate:                            MetricSettings{Enabled: true},
		ActiveDirectoryDsNameCacheHitRate:                          MetricSettings{Enabled: true},
		ActiveDirectoryDsNotificationQueued:                        MetricSettings{Enabled: true},
		ActiveDirectoryDsOperationRate:                             MetricSettings{Enabled: true},
		ActiveDirectoryDsReplicationNetworkIo:                      MetricSettings{Enabled: true},
		ActiveDirectoryDsReplicationObjectRate:                     MetricSettings{Enabled: true},
		ActiveDirectoryDsReplicationOperationPending:               MetricSettings{Enabled: true},
		ActiveDirectoryDsReplicationPropertyRate:                   MetricSettings{Enabled: true},
		ActiveDirectoryDsReplicationSyncObjectPending:              MetricSettings{Enabled: true},
		ActiveDirectoryDsReplicationSyncRequestCount:               MetricSettings{Enabled: true},
		ActiveDirectoryDsReplicationValueRate:                      MetricSettings{Enabled: true},
		ActiveDirectoryDsSecurityDescriptorPropagationsEventQueued: MetricSettings{Enabled: true},
		ActiveDirectoryDsSuboperationRate:                          MetricSettings{Enabled: true},
		ActiveDirectoryDsThreadCount:                               MetricSettings{Enabled: true},
	}
	mb := NewMetricsBuilder(settings, component.BuildInfo{}, WithStartTime(start))

	mb.RecordActiveDirectoryDsBindRateDataPoint(ts, 1, AttributeBindType(1))
	mb.RecordActiveDirectoryDsLdapBindLastSuccessfulTimeDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsLdapBindRateDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsLdapClientSessionCountDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsLdapSearchRateDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsNameCacheHitRateDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsNotificationQueuedDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsOperationRateDataPoint(ts, 1, AttributeOperationType(1))
	mb.RecordActiveDirectoryDsReplicationNetworkIoDataPoint(ts, 1, AttributeDirection(1), AttributeNetworkDataType(1))
	mb.RecordActiveDirectoryDsReplicationObjectRateDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordActiveDirectoryDsReplicationOperationPendingDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsReplicationPropertyRateDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordActiveDirectoryDsReplicationSyncObjectPendingDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsReplicationSyncRequestCountDataPoint(ts, 1, AttributeSyncResult(1))
	mb.RecordActiveDirectoryDsReplicationValueRateDataPoint(ts, 1, AttributeDirection(1), AttributeValueType(1))
	mb.RecordActiveDirectoryDsSecurityDescriptorPropagationsEventQueuedDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsSuboperationRateDataPoint(ts, 1, AttributeSuboperationType(1))
	mb.RecordActiveDirectoryDsThreadCountDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	assert.Equal(t, attrCount, rm.Resource().Attributes().Len())

	assert.Equal(t, 1, rm.ScopeMetrics().Len())
	ms := rm.ScopeMetrics().At(0).Metrics()
	allMetricsCount := reflect.TypeOf(MetricsSettings{}).NumField()
	assert.Equal(t, allMetricsCount, ms.Len())
	validatedMetrics := make(map[string]struct{})
	for i := 0; i < ms.Len(); i++ {
		switch ms.At(i).Name() {
		case "active_directory.ds.bind.rate":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of binds per second serviced by this domain controller.", ms.At(i).Description())
			assert.Equal(t, "{binds}/s", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "server", attrVal.Str())
			validatedMetrics["active_directory.ds.bind.rate"] = struct{}{}
		case "active_directory.ds.ldap.bind.last_successful.time":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The amount of time taken for the last successful LDAP bind.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["active_directory.ds.ldap.bind.last_successful.time"] = struct{}{}
		case "active_directory.ds.ldap.bind.rate":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of successful LDAP binds per second.", ms.At(i).Description())
			assert.Equal(t, "{binds}/s", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["active_directory.ds.ldap.bind.rate"] = struct{}{}
		case "active_directory.ds.ldap.client.session.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of connected LDAP client sessions.", ms.At(i).Description())
			assert.Equal(t, "{sessions}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["active_directory.ds.ldap.client.session.count"] = struct{}{}
		case "active_directory.ds.ldap.search.rate":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of LDAP searches per second.", ms.At(i).Description())
			assert.Equal(t, "{searches}/s", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["active_directory.ds.ldap.search.rate"] = struct{}{}
		case "active_directory.ds.name_cache.hit_rate":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The percentage of directory object name component lookups that are satisfied by the Directory System Agent's name cache.", ms.At(i).Description())
			assert.Equal(t, "%", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["active_directory.ds.name_cache.hit_rate"] = struct{}{}
		case "active_directory.ds.notification.queued":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of pending update notifications that have been queued to push to clients.", ms.At(i).Description())
			assert.Equal(t, "{notifications}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["active_directory.ds.notification.queued"] = struct{}{}
		case "active_directory.ds.operation.rate":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of operations performed per second.", ms.At(i).Description())
			assert.Equal(t, "{operations}/s", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "read", attrVal.Str())
			validatedMetrics["active_directory.ds.operation.rate"] = struct{}{}
		case "active_directory.ds.replication.network.io":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of network data transmitted by the Directory Replication Agent.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("direction")
			assert.True(t, ok)
			assert.Equal(t, "sent", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "compressed", attrVal.Str())
			validatedMetrics["active_directory.ds.replication.network.io"] = struct{}{}
		case "active_directory.ds.replication.object.rate":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of objects transmitted by the Directory Replication Agent per second.", ms.At(i).Description())
			assert.Equal(t, "{objects}/s", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("direction")
			assert.True(t, ok)
			assert.Equal(t, "sent", attrVal.Str())
			validatedMetrics["active_directory.ds.replication.object.rate"] = struct{}{}
		case "active_directory.ds.replication.operation.pending":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of pending replication operations for the Directory Replication Agent.", ms.At(i).Description())
			assert.Equal(t, "{operations}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["active_directory.ds.replication.operation.pending"] = struct{}{}
		case "active_directory.ds.replication.property.rate":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of properties transmitted by the Directory Replication Agent per second.", ms.At(i).Description())
			assert.Equal(t, "{properties}/s", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("direction")
			assert.True(t, ok)
			assert.Equal(t, "sent", attrVal.Str())
			validatedMetrics["active_directory.ds.replication.property.rate"] = struct{}{}
		case "active_directory.ds.replication.sync.object.pending":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of objects remaining until the full sync completes for the Directory Replication Agent.", ms.At(i).Description())
			assert.Equal(t, "{objects}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["active_directory.ds.replication.sync.object.pending"] = struct{}{}
		case "active_directory.ds.replication.sync.request.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of sync requests made by the Directory Replication Agent.", ms.At(i).Description())
			assert.Equal(t, "{requests}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("result")
			assert.True(t, ok)
			assert.Equal(t, "success", attrVal.Str())
			validatedMetrics["active_directory.ds.replication.sync.request.count"] = struct{}{}
		case "active_directory.ds.replication.value.rate":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of values transmitted by the Directory Replication Agent per second.", ms.At(i).Description())
			assert.Equal(t, "{values}/s", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("direction")
			assert.True(t, ok)
			assert.Equal(t, "sent", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "distingushed_names", attrVal.Str())
			validatedMetrics["active_directory.ds.replication.value.rate"] = struct{}{}
		case "active_directory.ds.security_descriptor_propagations_event.queued":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of security descriptor propagation events that are queued for processing.", ms.At(i).Description())
			assert.Equal(t, "{events}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["active_directory.ds.security_descriptor_propagations_event.queued"] = struct{}{}
		case "active_directory.ds.suboperation.rate":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The rate of sub-operations performed.", ms.At(i).Description())
			assert.Equal(t, "{suboperations}/s", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "security_descriptor_propagations_event", attrVal.Str())
			validatedMetrics["active_directory.ds.suboperation.rate"] = struct{}{}
		case "active_directory.ds.thread.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of threads in use by the directory service.", ms.At(i).Description())
			assert.Equal(t, "{threads}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["active_directory.ds.thread.count"] = struct{}{}
		}
	}
	assert.Equal(t, allMetricsCount, len(validatedMetrics))
}

func TestNoMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	settings := MetricsSettings{
		ActiveDirectoryDsBindRate:                                  MetricSettings{Enabled: false},
		ActiveDirectoryDsLdapBindLastSuccessfulTime:                MetricSettings{Enabled: false},
		ActiveDirectoryDsLdapBindRate:                              MetricSettings{Enabled: false},
		ActiveDirectoryDsLdapClientSessionCount:                    MetricSettings{Enabled: false},
		ActiveDirectoryDsLdapSearchRate:                            MetricSettings{Enabled: false},
		ActiveDirectoryDsNameCacheHitRate:                          MetricSettings{Enabled: false},
		ActiveDirectoryDsNotificationQueued:                        MetricSettings{Enabled: false},
		ActiveDirectoryDsOperationRate:                             MetricSettings{Enabled: false},
		ActiveDirectoryDsReplicationNetworkIo:                      MetricSettings{Enabled: false},
		ActiveDirectoryDsReplicationObjectRate:                     MetricSettings{Enabled: false},
		ActiveDirectoryDsReplicationOperationPending:               MetricSettings{Enabled: false},
		ActiveDirectoryDsReplicationPropertyRate:                   MetricSettings{Enabled: false},
		ActiveDirectoryDsReplicationSyncObjectPending:              MetricSettings{Enabled: false},
		ActiveDirectoryDsReplicationSyncRequestCount:               MetricSettings{Enabled: false},
		ActiveDirectoryDsReplicationValueRate:                      MetricSettings{Enabled: false},
		ActiveDirectoryDsSecurityDescriptorPropagationsEventQueued: MetricSettings{Enabled: false},
		ActiveDirectoryDsSuboperationRate:                          MetricSettings{Enabled: false},
		ActiveDirectoryDsThreadCount:                               MetricSettings{Enabled: false},
	}
	mb := NewMetricsBuilder(settings, component.BuildInfo{}, WithStartTime(start))
	mb.RecordActiveDirectoryDsBindRateDataPoint(ts, 1, AttributeBindType(1))
	mb.RecordActiveDirectoryDsLdapBindLastSuccessfulTimeDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsLdapBindRateDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsLdapClientSessionCountDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsLdapSearchRateDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsNameCacheHitRateDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsNotificationQueuedDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsOperationRateDataPoint(ts, 1, AttributeOperationType(1))
	mb.RecordActiveDirectoryDsReplicationNetworkIoDataPoint(ts, 1, AttributeDirection(1), AttributeNetworkDataType(1))
	mb.RecordActiveDirectoryDsReplicationObjectRateDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordActiveDirectoryDsReplicationOperationPendingDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsReplicationPropertyRateDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordActiveDirectoryDsReplicationSyncObjectPendingDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsReplicationSyncRequestCountDataPoint(ts, 1, AttributeSyncResult(1))
	mb.RecordActiveDirectoryDsReplicationValueRateDataPoint(ts, 1, AttributeDirection(1), AttributeValueType(1))
	mb.RecordActiveDirectoryDsSecurityDescriptorPropagationsEventQueuedDataPoint(ts, 1)
	mb.RecordActiveDirectoryDsSuboperationRateDataPoint(ts, 1, AttributeSuboperationType(1))
	mb.RecordActiveDirectoryDsThreadCountDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 0, metrics.ResourceMetrics().Len())
}