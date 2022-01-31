//go:build go1.18

package main

import (
	"testing"
)

func TestNonLocalSymbolLink(t *testing.T) {
	requireEqual := func(a, b string) {
		if a != b {
			t.Fatalf("%s != %s", a, b)
		}
	}
	requireEqual(
		nonLocalSymbolLink(
			"github.com/bradenaw/juniper",
			"github.com/bradenaw/juniper/container/tree",
			"github.com/bradenaw/juniper/iterator",
			"Iterator",
		),
		"../iterator.html#Iterator",
	)
	requireEqual(
		nonLocalSymbolLink(
			"github.com/bradenaw/juniper",
			"github.com/bradenaw/juniper/container/tree",
			"github.com/bradenaw/juniper/internal/fuzz",
			"Operations",
		),
		"../internal/fuzz.html#Operations",
	)
	requireEqual(
		nonLocalSymbolLink(
			"github.com/bradenaw/juniper",
			"github.com/bradenaw/juniper/xmath",
			"github.com/bradenaw/juniper/xmath/xrand",
			"Sample",
		),
		"xmath/xrand.html#Sample",
	)
	requireEqual(
		nonLocalSymbolLink(
			"github.com/bradenaw/juniper",
			"github.com/bradenaw/juniper/xmath/xrand",
			"github.com/bradenaw/juniper/xmath",
			"Min",
		),
		"../xmath.html#Min",
	)
}
