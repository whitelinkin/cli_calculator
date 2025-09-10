package calculator

import "testing"

func TestHistory(t *testing.T) {
	ClearHistory()

	AddHistory("2 + 2", 4)
	AddHistory("10 - 3", 7)

	h := GetHistory()
	if len(h) != 2 {
		t.Errorf("Ожидалось 2 записи в истории, получили %d", len(h))
	}

	if h[0] != "2 + 2 = 4" {
		t.Errorf("Ожидалось '2 + 2 = 4', получили '%s'", h[0])
	}
}
