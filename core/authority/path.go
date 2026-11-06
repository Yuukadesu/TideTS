package authority

import (
	"strings"
	"unicode"

	commons "github.com/hanami/tidets/commons/errors"
)

func validateUsername(name string) error {
	if name == "" || strings.TrimSpace(name) != name {
		return commons.ErrAuthUsernameRequired
	}
	for _, r := range name {
		if unicode.IsSpace(r) {
			return commons.ErrAuthUsernameRequired
		}
	}
	return nil
}

// validateGrantPath 接受 root、root.sg1、root.sg1.*、root.sg1.**。
func validateGrantPath(path string) error {
	if path == "" {
		return commons.ErrAuthInvalidGrantPath
	}
	base := path
	switch {
	case strings.HasSuffix(path, ".**"):
		base = strings.TrimSuffix(path, ".**")
	case strings.HasSuffix(path, ".*"):
		base = strings.TrimSuffix(path, ".*")
	}
	if base == "" || strings.Contains(base, "*") {
		return commons.ErrAuthInvalidGrantPath
	}
	parts := strings.Split(base, ".")
	if len(parts) == 0 || parts[0] != "root" {
		return commons.ErrAuthInvalidGrantPath
	}
	for _, part := range parts {
		if part == "" {
			return commons.ErrAuthInvalidGrantPath
		}
		for _, r := range part {
			if unicode.IsSpace(r) {
				return commons.ErrAuthInvalidGrantPath
			}
		}
	}
	return nil
}
