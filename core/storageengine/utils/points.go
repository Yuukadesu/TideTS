package utils

import (
	"sort"

	"github.com/hanami/tidets/core/storageengine/model"
)

// MergeSorted 归并两个按时间升序的点列；同时间戳保留 a。
func MergeSorted(a, b []model.Point) []model.Point {
	if len(a) == 0 {
		return append([]model.Point(nil), b...)
	}
	out := make([]model.Point, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i].Timestamp <= b[j].Timestamp {
			if len(out) == 0 || out[len(out)-1].Timestamp != a[i].Timestamp {
				out = append(out, a[i])
			}
			i++
		} else {
			if len(out) == 0 || out[len(out)-1].Timestamp != b[j].Timestamp {
				out = append(out, b[j])
			}
			j++
		}
	}
	for i < len(a) {
		if len(out) == 0 || out[len(out)-1].Timestamp != a[i].Timestamp {
			out = append(out, a[i])
		}
		i++
	}
	for j < len(b) {
		if len(out) == 0 || out[len(out)-1].Timestamp != b[j].Timestamp {
			out = append(out, b[j])
		}
		j++
	}
	return out
}

// MergeSortedPreferNewer 归并升序序列；时间戳相同时保留 b（较新来源）。
func MergeSortedPreferNewer(a, b []model.Point) []model.Point {
	if len(a) == 0 {
		return append([]model.Point(nil), b...)
	}
	if len(b) == 0 {
		return append([]model.Point(nil), a...)
	}
	out := make([]model.Point, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i].Timestamp < b[j].Timestamp {
			out = append(out, a[i])
			i++
		} else if a[i].Timestamp > b[j].Timestamp {
			out = append(out, b[j])
			j++
		} else {
			out = append(out, b[j])
			i++
			j++
		}
	}
	for i < len(a) {
		out = append(out, a[i])
		i++
	}
	for j < len(b) {
		out = append(out, b[j])
		j++
	}
	return out
}

// TruncatePoints 截断结果集；limit <= 0 时不截断。
func TruncatePoints(pts []model.Point, limit int) []model.Point {
	if limit <= 0 || len(pts) <= limit {
		return pts
	}
	return append([]model.Point(nil), pts[:limit]...)
}

// QueryPointsInRange 在有序点列上按时间范围切片。
func QueryPointsInRange(pts []model.Point, start, end int64) []model.Point {
	if len(pts) == 0 {
		return nil
	}
	if end < pts[0].Timestamp || start > pts[len(pts)-1].Timestamp {
		return nil
	}
	i := sort.Search(len(pts), func(i int) bool { return pts[i].Timestamp >= start })
	if i >= len(pts) {
		return nil
	}
	j := sort.Search(len(pts), func(j int) bool { return pts[j].Timestamp > end })
	return append([]model.Point(nil), pts[i:j]...)
}

// InsertSorted 向有序点列插入一点；同时间戳覆盖。
func InsertSorted(points []model.Point, p model.Point) []model.Point {
	i := sort.Search(len(points), func(i int) bool { return points[i].Timestamp >= p.Timestamp })
	if i < len(points) && points[i].Timestamp == p.Timestamp {
		dup := append([]model.Point(nil), points...)
		dup[i] = p
		return dup
	}
	return append(points[:i:i], append([]model.Point{p}, points[i:]...)...)
}
