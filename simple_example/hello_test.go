package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		names  []string
		result string
	}{
		{
			names:  []string{},
			result: "Hello, World!",
		},
		{
			names:  []string{"David", "Tom"},
			result: "Hello, David, Tom!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.names); s != st.result {
			t.Errorf("Wanted %s, got %s", st.result, s)
		}
	}
}
