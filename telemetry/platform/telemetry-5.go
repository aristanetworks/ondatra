package platform

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Subcomponent_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Subcomponent{}
	md, ok := oc.Lookup(t, n, "Component_Subcomponent", goStruct, true, false)
	if ok {
		return convertComponent_Subcomponent_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Subcomponent_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Subcomponent_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Subcomponent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Subcomponent", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Subcomponent_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a ONCE subscription.
func (n *Component_Subcomponent_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Subcomponent_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Subcomponent_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_Subcomponent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Subcomponent", gs, queryPath, true, false)
		return convertComponent_Subcomponent_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Subcomponent_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Subcomponent_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Subcomponent_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/subcomponents/subcomponent/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/subcomponents/subcomponent/state/name to the batch object.
func (n *Component_Subcomponent_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Subcomponent_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Subcomponent_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Subcomponent_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/subcomponents/subcomponent/state/name to the batch object.
func (n *Component_Subcomponent_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Subcomponent_NamePath extracts the value of the leaf Name from its parent oc.Component_Subcomponent
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Subcomponent_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Subcomponent) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Name
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_TemperaturePath) Lookup(t testing.TB) *oc.QualifiedComponent_Temperature {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Temperature{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_TemperaturePath) Get(t testing.TB) *oc.Component_Temperature {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_TemperaturePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Temperature {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Temperature
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Temperature{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature with a ONCE subscription.
func (n *Component_TemperaturePathAny) Get(t testing.TB) []*oc.Component_Temperature {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Temperature
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_TemperaturePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Temperature {
	t.Helper()
	c := &oc.CollectionComponent_Temperature{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Temperature) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Temperature{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Temperature)))
		return false
	})
	return c
}

func watch_Component_TemperaturePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Temperature) bool) *oc.Component_TemperatureWatcher {
	t.Helper()
	w := &oc.Component_TemperatureWatcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Temperature{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Temperature)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_TemperaturePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Temperature) bool) *oc.Component_TemperatureWatcher {
	t.Helper()
	return watch_Component_TemperaturePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_TemperaturePath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Temperature) *oc.QualifiedComponent_Temperature {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Temperature) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature to the batch object.
func (n *Component_TemperaturePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_TemperaturePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Temperature {
	t.Helper()
	c := &oc.CollectionComponent_Temperature{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Temperature) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_TemperaturePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Temperature) bool) *oc.Component_TemperatureWatcher {
	t.Helper()
	return watch_Component_TemperaturePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature to the batch object.
func (n *Component_TemperaturePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/alarm-severity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_AlarmSeverityPath) Lookup(t testing.TB) *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_AlarmSeverityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/alarm-severity with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_AlarmSeverityPath) Get(t testing.TB) oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/alarm-severity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_AlarmSeverityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_AlarmSeverityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/alarm-severity with a ONCE subscription.
func (n *Component_Temperature_AlarmSeverityPathAny) Get(t testing.TB) []oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/alarm-severity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AlarmSeverityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	c := &oc.CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_AlarmSeverityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool) *oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher {
	t.Helper()
	w := &oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_AlarmSeverityPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/alarm-severity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AlarmSeverityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool) *oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher {
	t.Helper()
	return watch_Component_Temperature_AlarmSeverityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/alarm-severity with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_AlarmSeverityPath) Await(t testing.TB, timeout time.Duration, val oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/alarm-severity failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/alarm-severity to the batch object.
func (n *Component_Temperature_AlarmSeverityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/alarm-severity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AlarmSeverityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	c := &oc.CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/alarm-severity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AlarmSeverityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool) *oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher {
	t.Helper()
	return watch_Component_Temperature_AlarmSeverityPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/alarm-severity to the batch object.
func (n *Component_Temperature_AlarmSeverityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_AlarmSeverityPath extracts the value of the leaf AlarmSeverity from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY.
func convertComponent_Temperature_AlarmSeverityPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	qv := &oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY{
		Metadata: md,
	}
	val := parent.AlarmSeverity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/alarm-status with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_AlarmStatusPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_AlarmStatusPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/alarm-status with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_AlarmStatusPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/alarm-status with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_AlarmStatusPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_AlarmStatusPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/alarm-status with a ONCE subscription.
func (n *Component_Temperature_AlarmStatusPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/alarm-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AlarmStatusPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_AlarmStatusPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_AlarmStatusPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/alarm-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AlarmStatusPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Component_Temperature_AlarmStatusPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/alarm-status with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_AlarmStatusPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/alarm-status failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/alarm-status to the batch object.
func (n *Component_Temperature_AlarmStatusPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/alarm-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AlarmStatusPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/alarm-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AlarmStatusPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Component_Temperature_AlarmStatusPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/alarm-status to the batch object.
func (n *Component_Temperature_AlarmStatusPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_AlarmStatusPath extracts the value of the leaf AlarmStatus from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertComponent_Temperature_AlarmStatusPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.AlarmStatus
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/alarm-threshold with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_AlarmThresholdPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_AlarmThresholdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/alarm-threshold with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_AlarmThresholdPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_AlarmThresholdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_AlarmThresholdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a ONCE subscription.
func (n *Component_Temperature_AlarmThresholdPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AlarmThresholdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_AlarmThresholdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_AlarmThresholdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AlarmThresholdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Component_Temperature_AlarmThresholdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_AlarmThresholdPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/alarm-threshold failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/alarm-threshold to the batch object.
func (n *Component_Temperature_AlarmThresholdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AlarmThresholdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AlarmThresholdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Component_Temperature_AlarmThresholdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/alarm-threshold to the batch object.
func (n *Component_Temperature_AlarmThresholdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_AlarmThresholdPath extracts the value of the leaf AlarmThreshold from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertComponent_Temperature_AlarmThresholdPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.AlarmThreshold
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/avg with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/avg with a ONCE subscription.
func (n *Component_Temperature_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_AvgPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/avg to the batch object.
func (n *Component_Temperature_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_AvgPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/avg to the batch object.
func (n *Component_Temperature_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_AvgPath extracts the value of the leaf Avg from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Temperature_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/instant with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/instant with a ONCE subscription.
func (n *Component_Temperature_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_InstantPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/instant to the batch object.
func (n *Component_Temperature_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_InstantPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/instant to the batch object.
func (n *Component_Temperature_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_InstantPath extracts the value of the leaf Instant from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Temperature_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Instant
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/interval with a ONCE subscription.
func (n *Component_Temperature_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_IntervalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Temperature_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/interval to the batch object.
func (n *Component_Temperature_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Temperature_IntervalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/interval to the batch object.
func (n *Component_Temperature_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Temperature_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Interval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/max with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/max with a ONCE subscription.
func (n *Component_Temperature_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_MaxPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/max to the batch object.
func (n *Component_Temperature_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_MaxPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/max to the batch object.
func (n *Component_Temperature_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_MaxPath extracts the value of the leaf Max from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Temperature_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Max
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/max-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/max-time with a ONCE subscription.
func (n *Component_Temperature_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_MaxTimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Temperature_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/max-time to the batch object.
func (n *Component_Temperature_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Temperature_MaxTimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/max-time to the batch object.
func (n *Component_Temperature_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Temperature_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/min with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/min with a ONCE subscription.
func (n *Component_Temperature_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_MinPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/min to the batch object.
func (n *Component_Temperature_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_MinPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/min to the batch object.
func (n *Component_Temperature_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_MinPath extracts the value of the leaf Min from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Temperature_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Min
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/min-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/min-time with a ONCE subscription.
func (n *Component_Temperature_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_MinTimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Temperature_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/min-time to the batch object.
func (n *Component_Temperature_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Temperature_MinTimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/min-time to the batch object.
func (n *Component_Temperature_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Temperature_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MinTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_TransceiverPath) Lookup(t testing.TB) *oc.QualifiedComponent_Transceiver {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Transceiver{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_TransceiverPath) Get(t testing.TB) *oc.Component_Transceiver {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_TransceiverPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Transceiver {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Transceiver
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Transceiver{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver with a ONCE subscription.
func (n *Component_TransceiverPathAny) Get(t testing.TB) []*oc.Component_Transceiver {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Transceiver
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_TransceiverPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Transceiver{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Transceiver)))
		return false
	})
	return c
}

func watch_Component_TransceiverPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver) bool) *oc.Component_TransceiverWatcher {
	t.Helper()
	w := &oc.Component_TransceiverWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Transceiver{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_TransceiverPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver) bool) *oc.Component_TransceiverWatcher {
	t.Helper()
	return watch_Component_TransceiverPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_TransceiverPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Transceiver) *oc.QualifiedComponent_Transceiver {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Transceiver) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver to the batch object.
func (n *Component_TransceiverPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_TransceiverPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_TransceiverPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver) bool) *oc.Component_TransceiverWatcher {
	t.Helper()
	return watch_Component_TransceiverPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver to the batch object.
func (n *Component_TransceiverPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_ChannelPath) Lookup(t testing.TB) *oc.QualifiedComponent_Transceiver_Channel {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Transceiver_Channel{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_ChannelPath) Get(t testing.TB) *oc.Component_Transceiver_Channel {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_ChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Transceiver_Channel {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Transceiver_Channel
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Transceiver_Channel{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel with a ONCE subscription.
func (n *Component_Transceiver_ChannelPathAny) Get(t testing.TB) []*oc.Component_Transceiver_Channel {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Transceiver_Channel
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_ChannelPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_Channel {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_Channel{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_Channel) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Transceiver_Channel{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Transceiver_Channel)))
		return false
	})
	return c
}

func watch_Component_Transceiver_ChannelPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_Channel) bool) *oc.Component_Transceiver_ChannelWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_ChannelWatcher{}
	gs := &oc.Component_Transceiver_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_Channel", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Transceiver_Channel{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_Channel)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_ChannelPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_Channel) bool) *oc.Component_Transceiver_ChannelWatcher {
	t.Helper()
	return watch_Component_Transceiver_ChannelPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/physical-channels/channel with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_ChannelPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Transceiver_Channel) *oc.QualifiedComponent_Transceiver_Channel {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Transceiver_Channel) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/physical-channels/channel failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel to the batch object.
func (n *Component_Transceiver_ChannelPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_ChannelPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_Channel {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_Channel{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_Channel) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_ChannelPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_Channel) bool) *oc.Component_Transceiver_ChannelWatcher {
	t.Helper()
	return watch_Component_Transceiver_ChannelPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel to the batch object.
func (n *Component_Transceiver_ChannelPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/associated-optical-channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_Channel_AssociatedOpticalChannelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/associated-optical-channel with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/associated-optical-channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_Channel_AssociatedOpticalChannelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/associated-optical-channel with a ONCE subscription.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/associated-optical-channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_Channel_AssociatedOpticalChannelPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_Transceiver_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_Channel", gs, queryPath, true, false)
		return convertComponent_Transceiver_Channel_AssociatedOpticalChannelPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/associated-optical-channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_AssociatedOpticalChannelPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/associated-optical-channel with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/associated-optical-channel failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/associated-optical-channel to the batch object.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/associated-optical-channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/associated-optical-channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_AssociatedOpticalChannelPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/associated-optical-channel to the batch object.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_Channel_AssociatedOpticalChannelPath extracts the value of the leaf AssociatedOpticalChannel from its parent oc.Component_Transceiver_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Transceiver_Channel_AssociatedOpticalChannelPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_Channel) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AssociatedOpticalChannel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_Channel_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/description with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_Channel_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/description with a ONCE subscription.
func (n *Component_Transceiver_Channel_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_DescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_Channel_DescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_Transceiver_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_Channel", gs, queryPath, true, false)
		return convertComponent_Transceiver_Channel_DescriptionPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_DescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_DescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_Channel_DescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/description to the batch object.
func (n *Component_Transceiver_Channel_DescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_DescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_DescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_DescriptionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/description to the batch object.
func (n *Component_Transceiver_Channel_DescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_Channel_DescriptionPath extracts the value of the leaf Description from its parent oc.Component_Transceiver_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Transceiver_Channel_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_Channel) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Description
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_IndexPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_Channel_IndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/index with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_IndexPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_IndexPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_Channel_IndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/index with a ONCE subscription.
func (n *Component_Transceiver_Channel_IndexPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_IndexPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_Channel_IndexPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Component_Transceiver_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_Channel", gs, queryPath, true, false)
		return convertComponent_Transceiver_Channel_IndexPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_IndexPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_IndexPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/index with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_Channel_IndexPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/index failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/index to the batch object.
func (n *Component_Transceiver_Channel_IndexPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_IndexPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_IndexPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_IndexPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/index to the batch object.
func (n *Component_Transceiver_Channel_IndexPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_Channel_IndexPath extracts the value of the leaf Index from its parent oc.Component_Transceiver_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertComponent_Transceiver_Channel_IndexPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_Channel) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Index
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_InputPowerPath) Lookup(t testing.TB) *oc.QualifiedComponent_Transceiver_Channel_InputPower {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel_InputPower", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Transceiver_Channel_InputPower{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_InputPowerPath) Get(t testing.TB) *oc.Component_Transceiver_Channel_InputPower {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_InputPowerPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Transceiver_Channel_InputPower {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Transceiver_Channel_InputPower
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel_InputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel_InputPower", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Transceiver_Channel_InputPower{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power with a ONCE subscription.
func (n *Component_Transceiver_Channel_InputPowerPathAny) Get(t testing.TB) []*oc.Component_Transceiver_Channel_InputPower {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Transceiver_Channel_InputPower
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_InputPowerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_Channel_InputPower {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_Channel_InputPower{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_Channel_InputPower) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Transceiver_Channel_InputPower{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Transceiver_Channel_InputPower)))
		return false
	})
	return c
}

func watch_Component_Transceiver_Channel_InputPowerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_Channel_InputPower) bool) *oc.Component_Transceiver_Channel_InputPowerWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_Channel_InputPowerWatcher{}
	gs := &oc.Component_Transceiver_Channel_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_Channel_InputPower", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Transceiver_Channel_InputPower{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_Channel_InputPower)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_InputPowerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_Channel_InputPower) bool) *oc.Component_Transceiver_Channel_InputPowerWatcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_InputPowerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_Channel_InputPowerPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Transceiver_Channel_InputPower) *oc.QualifiedComponent_Transceiver_Channel_InputPower {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Transceiver_Channel_InputPower) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power to the batch object.
func (n *Component_Transceiver_Channel_InputPowerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_InputPowerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_Channel_InputPower {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_Channel_InputPower{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_Channel_InputPower) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_InputPowerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_Channel_InputPower) bool) *oc.Component_Transceiver_Channel_InputPowerWatcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_InputPowerPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power to the batch object.
func (n *Component_Transceiver_Channel_InputPowerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_InputPower_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel_InputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_Channel_InputPower_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/avg with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_InputPower_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_InputPower_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel_InputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel_InputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_Channel_InputPower_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/avg with a ONCE subscription.
func (n *Component_Transceiver_Channel_InputPower_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_InputPower_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_Channel_InputPower_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_Channel_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_Channel_InputPower", gs, queryPath, true, false)
		return convertComponent_Transceiver_Channel_InputPower_AvgPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_InputPower_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_InputPower_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_Channel_InputPower_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/avg to the batch object.
func (n *Component_Transceiver_Channel_InputPower_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_InputPower_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_InputPower_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_InputPower_AvgPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/avg to the batch object.
func (n *Component_Transceiver_Channel_InputPower_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_Channel_InputPower_AvgPath extracts the value of the leaf Avg from its parent oc.Component_Transceiver_Channel_InputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_Channel_InputPower_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_Channel_InputPower) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_InputPower_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel_InputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_Channel_InputPower_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/instant with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_InputPower_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_InputPower_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel_InputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel_InputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_Channel_InputPower_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/instant with a ONCE subscription.
func (n *Component_Transceiver_Channel_InputPower_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_InputPower_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_Channel_InputPower_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_Channel_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_Channel_InputPower", gs, queryPath, true, false)
		return convertComponent_Transceiver_Channel_InputPower_InstantPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_InputPower_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_InputPower_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_Channel_InputPower_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/instant to the batch object.
func (n *Component_Transceiver_Channel_InputPower_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_InputPower_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_InputPower_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_InputPower_InstantPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/instant to the batch object.
func (n *Component_Transceiver_Channel_InputPower_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_Channel_InputPower_InstantPath extracts the value of the leaf Instant from its parent oc.Component_Transceiver_Channel_InputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_Channel_InputPower_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_Channel_InputPower) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Instant
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_InputPower_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel_InputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_Channel_InputPower_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_InputPower_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_InputPower_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel_InputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel_InputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_Channel_InputPower_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/interval with a ONCE subscription.
func (n *Component_Transceiver_Channel_InputPower_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_InputPower_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_Channel_InputPower_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_Channel_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_Channel_InputPower", gs, queryPath, true, false)
		return convertComponent_Transceiver_Channel_InputPower_IntervalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_InputPower_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_InputPower_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_Channel_InputPower_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/interval to the batch object.
func (n *Component_Transceiver_Channel_InputPower_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_InputPower_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_InputPower_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_InputPower_IntervalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/interval to the batch object.
func (n *Component_Transceiver_Channel_InputPower_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_Channel_InputPower_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_Transceiver_Channel_InputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_Channel_InputPower_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_Channel_InputPower) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Interval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_InputPower_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel_InputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_Channel_InputPower_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/max with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_InputPower_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_InputPower_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel_InputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel_InputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_Channel_InputPower_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/max with a ONCE subscription.
func (n *Component_Transceiver_Channel_InputPower_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_InputPower_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_Channel_InputPower_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_Channel_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_Channel_InputPower", gs, queryPath, true, false)
		return convertComponent_Transceiver_Channel_InputPower_MaxPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_InputPower_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_InputPower_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_Channel_InputPower_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/max to the batch object.
func (n *Component_Transceiver_Channel_InputPower_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_Channel_InputPower_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_Channel_InputPower_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_Channel_InputPower_MaxPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/physical-channels/channel/state/input-power/max to the batch object.
func (n *Component_Transceiver_Channel_InputPower_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_Channel_InputPower_MaxPath extracts the value of the leaf Max from its parent oc.Component_Transceiver_Channel_InputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_Channel_InputPower_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_Channel_InputPower) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Max
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
