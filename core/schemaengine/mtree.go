package schemaengine

import (
	"sort"
	"strings"
	"sync"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel"
)

// ChildPath 路径树中的直接子节点。
type ChildPath struct {
	Name     string
	FullPath string
	IsDevice bool
}

type mtreeNode struct {
	name     string
	children map[string]*mtreeNode
	isDevice bool
	series   map[string]Timeseries
}

func newMTreeNode(name string) *mtreeNode {
	return &mtreeNode{
		name:     name,
		children: make(map[string]*mtreeNode),
		series:   make(map[string]Timeseries),
	}
}

// MTree 统一 schema + metadata 路径树（对齐 IoTDB MTree 子集）。
type MTree struct {
	mu   sync.RWMutex
	root *mtreeNode
}

func newMTree() *MTree {
	return &MTree{root: newMTreeNode("root")}
}

func newMTreeFromSeries(items []Timeseries) *MTree {
	t := newMTree()
	for _, ts := range items {
		_ = t.put(ts)
	}
	return t
}

func splitDevicePath(path string) []string {
	path = strings.Trim(path, ".")
	if path == "" {
		return nil
	}
	return strings.Split(path, ".")
}

func validateDevicePath(path string) error {
	if path == "" {
		return commons.ErrMetadataPathRequired
	}
	if !strings.HasPrefix(path, "root") {
		return commons.ErrMetadataInvalidPath
	}
	return nil
}

func (t *MTree) put(ts Timeseries) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.putLocked(ts)
}

func (t *MTree) putLocked(ts Timeseries) error {
	node, err := t.ensureDeviceLocked(ts.DevicePath)
	if err != nil {
		return err
	}
	node.series[ts.Measurement] = ts
	return nil
}

func (t *MTree) ensureDeviceLocked(devicePath string) (*mtreeNode, error) {
	if err := validateDevicePath(devicePath); err != nil {
		return nil, err
	}
	parts := splitDevicePath(devicePath)
	if len(parts) == 0 || parts[0] != "root" {
		return nil, commons.ErrMetadataInvalidPath
	}
	cur := t.root
	for _, part := range parts[1:] {
		child, ok := cur.children[part]
		if !ok {
			child = newMTreeNode(part)
			cur.children[part] = child
		}
		cur = child
	}
	cur.isDevice = true
	return cur, nil
}

func (t *MTree) get(key tsmodel.SeriesKey) (Timeseries, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.getLocked(key)
}

func (t *MTree) getLocked(key tsmodel.SeriesKey) (Timeseries, bool) {
	node, ok := t.deviceNodeLocked(key.DevicePath)
	if !ok {
		return Timeseries{}, false
	}
	ts, ok := node.series[key.Measurement]
	return ts, ok
}

func (t *MTree) has(key tsmodel.SeriesKey) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.hasLocked(key)
}

func (t *MTree) hasLocked(key tsmodel.SeriesKey) bool {
	_, ok := t.getLocked(key)
	return ok
}

func (t *MTree) deviceNodeLocked(devicePath string) (*mtreeNode, bool) {
	if err := validateDevicePath(devicePath); err != nil {
		return nil, false
	}
	cur := t.root
	for _, part := range splitDevicePath(devicePath)[1:] {
		child, ok := cur.children[part]
		if !ok {
			return nil, false
		}
		cur = child
	}
	if !cur.isDevice {
		return nil, false
	}
	return cur, true
}

func (t *MTree) list(pathPrefix string) []Timeseries {
	t.mu.RLock()
	defer t.mu.RUnlock()
	out := make([]Timeseries, 0)
	t.walkSeries(t.root, func(ts Timeseries) {
		if pathPrefix == "" || strings.HasPrefix(ts.FullPath(), pathPrefix) {
			out = append(out, ts)
		}
	})
	sort.Slice(out, func(i, j int) bool {
		return out[i].FullPath() < out[j].FullPath()
	})
	return out
}

func (t *MTree) walkSeries(node *mtreeNode, fn func(Timeseries)) {
	for _, ts := range node.series {
		fn(ts)
	}
	for _, child := range node.children {
		t.walkSeries(child, fn)
	}
}

func (t *MTree) snapshotSeries() []Timeseries {
	return t.list("")
}

func (t *MTree) ListDevices(pattern string) []string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if pattern != "" && !strings.Contains(pattern, "*") {
		if err := validateDevicePath(strings.TrimSuffix(strings.TrimSuffix(pattern, "."), " ")); err != nil {
			return nil
		}
	}
	var out []string
	t.walkDevices(t.root, "root", &out)
	sort.Strings(out)
	if pattern == "" {
		return out
	}
	filtered := make([]string, 0, len(out))
	for _, p := range out {
		if tsmodel.MatchDevicePattern(p, pattern) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func (t *MTree) walkDevices(node *mtreeNode, fullPath string, out *[]string) {
	if node.isDevice {
		*out = append(*out, fullPath)
	}
	for name, child := range node.children {
		t.walkDevices(child, fullPath+"."+name, out)
	}
}

func (t *MTree) ListMeasurements(devicePath string) []Timeseries {
	t.mu.RLock()
	defer t.mu.RUnlock()
	node, ok := t.deviceNodeLocked(devicePath)
	if !ok {
		return nil
	}
	out := make([]Timeseries, 0, len(node.series))
	for _, ts := range node.series {
		out = append(out, ts)
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].Measurement < out[j].Measurement
	})
	return out
}

func (t *MTree) ChildPaths(prefix string) []ChildPath {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if prefix == "" {
		prefix = "root"
	}
	if err := validateDevicePath(prefix); err != nil {
		return nil
	}
	cur := t.root
	parts := splitDevicePath(prefix)
	for i := 1; i < len(parts); i++ {
		child, ok := cur.children[parts[i]]
		if !ok {
			return nil
		}
		cur = child
	}
	out := make([]ChildPath, 0, len(cur.children))
	for name, child := range cur.children {
		full := prefix + "." + name
		out = append(out, ChildPath{
			Name:     name,
			FullPath: full,
			IsDevice: child.isDevice,
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}
