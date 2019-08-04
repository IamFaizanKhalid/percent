package percent

import (
	"testing"
)

func TestPercent(t *testing.T) {
	pcent, all := 25, 2000
	part := Percent(pcent, all)

	if part != 500.0 {
		t.Fatalf("%d is Wrong number for %d percent!", int(part), pcent)
	}
}

func TestPercentFloat(t *testing.T) {
	pcent, all := 25.0, 2000
	part := PercentFloat(pcent, all)

	if part != 500.0 {
		t.Fatalf("%d is Wrong number for %f percent!", int(part), pcent)
	}
}

func TestPercentOf(t *testing.T) {
	part, all := 20, 50
	pcent := PercentOf(part, all)

	if pcent != 40 {
		t.Fatalf("Wrong percent!")
	}
}

func TestChange(t *testing.T) {
	before, after := 25, 200
	pcent := Change(before, after)

	if int(pcent) != 700 {
		t.Fatalf("%f is Wrong percent!", pcent)
	}
}

func TestChangeFloat(t *testing.T) {
	before, after := 5.00, 100.00
	pcent := ChangeFloat(before, after)

	if int(pcent) != 1900 {
		t.Fatalf("%f is Wrong percent!", pcent)
	}
}
