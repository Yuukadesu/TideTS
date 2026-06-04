package utils

import (
	"sort"

	"github.com/hanami/tidets/core/tsmodel"
)

// MergeSorted 归并两个按时间升序的点列；同时间戳保留 a。
func MergeSorted(a, b []tsmodel.Point) []tsmodel.Point {
	if len(a) == 0 {
		return append([]tsmodel.Point(nil), b...)
	}
	out := make([]tsmodel.Point, 0, len(a)+len(b))
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
func MergeSortedPreferNewer(a, b []tsmodel.Point) []tsmodel.Point {
	if len(a) == 0 {
		return append([]tsmodel.Point(nil), b...)
	}
	if len(b) == 0 {
		return append([]tsmodel.Point(nil), a...)
	}
	out := make([]tsmodel.Point, 0, len(a)+len(b))
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
func TruncatePoints(pts []tsmodel.Point, limit int) []tsmodel.Point {
	if limit <= 0 || len(pts) <= limit {
		return pts
	}
	return append([]tsmodel.Point(nil), pts[:limit]...)
}

// QueryPointsInRange 在有序点列上按时间范围切片。
func QueryPointsInRange(pts []tsmodel.Point, start, end int64) []tsmodel.Point {
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
	return append([]tsmodel.Point(nil), pts[i:j]...)
}

// InsertSorted 向有序点列插入一点；同时间戳覆盖。
func InsertSorted(points []tsmodel.Point, p tsmodel.Point) []tsmodel.Point {
	i := sort.Search(len(points), func(i int) bool { return points[i].Timestamp >= p.Timestamp })
	if i < len(points) && points[i].Timestamp == p.Timestamp {
		dup := append([]tsmodel.Point(nil), points...)
		dup[i] = p
		return dup
	}
	return append(points[:i:i], append([]tsmodel.Point{p}, points[i:]...)...)
}

// DeleteFromSorted 从有序点列删除指定时间戳。
func DeleteFromSorted(points []tsmodel.Point, ts int64) []tsmodel.Point {
	i := sort.Search(len(points), func(i int) bool { return points[i].Timestamp >= ts })
	if i >= len(points) || points[i].Timestamp != ts {
		return points
	}
	out := append([]tsmodel.Point(nil), points[:i]...)
	return append(out, points[i+1:]...)
}

// DeleteRangeFromSorted 从有序点列删除闭区间 [start, end] 内的点。
func DeleteRangeFromSorted(points []tsmodel.Point, start, end int64) []tsmodel.Point {
	if len(points) == 0 || start > end {
		return points
	}
	i := sort.Search(len(points), func(i int) bool { return points[i].Timestamp >= start })
	if i >= len(points) {
		return points
	}
	j := sort.Search(len(points), func(j int) bool { return points[j].Timestamp > end })
	if i >= j {
		return points
	}
	out := append([]tsmodel.Point(nil), points[:i]...)
	return append(out, points[j:]...)
}

// CountInRange 统计有序点列在 [start, end] 内的点数。
func CountInRange(points []tsmodel.Point, start, end int64) int {
	if len(points) == 0 || start > end {
		return 0
	}
	i := sort.Search(len(points), func(i int) bool { return points[i].Timestamp >= start })
	if i >= len(points) {
		return 0
	}
	j := sort.Search(len(points), func(j int) bool { return points[j].Timestamp > end })
	return j - i
}
