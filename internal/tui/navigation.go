package tui

import (
	"github.com/mikanfactory/yakumo/internal/model"
)

// NextSelectable returns the next selectable index after current, or current if none.
func NextSelectable(items []model.NavigableItem, current int) int {
	for i := current + 1; i < len(items); i++ {
		if items[i].Selectable {
			return i
		}
	}
	return current
}

// PrevSelectable returns the previous selectable index before current, or current if none.
func PrevSelectable(items []model.NavigableItem, current int) int {
	for i := current - 1; i >= 0; i-- {
		if items[i].Selectable {
			return i
		}
	}
	return current
}

// FirstSelectable returns the index of the first selectable item, or 0.
func FirstSelectable(items []model.NavigableItem) int {
	for i, item := range items {
		if item.Selectable {
			return i
		}
	}
	return 0
}

// adjustScroll returns a scroll offset that keeps cursor inside the viewport.
func adjustScroll(cursor, scrollOff, viewportHeight, totalItems int) int {
	if totalItems <= viewportHeight {
		return 0
	}
	if cursor < scrollOff {
		return cursor
	}
	if cursor >= scrollOff+viewportHeight {
		return cursor - viewportHeight + 1
	}
	return scrollOff
}
