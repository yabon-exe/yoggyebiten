package system_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/yabon-exe/yoggyebiten/game/system"
)

func Test_PopRandValue(t *testing.T) {

	vals := []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}
	limitedRandom := system.NewLimitedRandom[string](vals)

	results := []string{}
	for range vals {
		popV := limitedRandom.PopRandValue()
		results = append(results, popV)
	}
	sort.Strings(results)

	exps := []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}
	if !reflect.DeepEqual(exps, results) {
		t.Errorf("PopRandValue Error:\nvals=>[%s]\nress=>[%s]\n", strings.Join(exps, ","), strings.Join(results, ","))
	}
}

func Test_Reset(t *testing.T) {

	vals := []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}
	limitedRandom := system.NewLimitedRandom[string](vals)

	results := []string{}
	for range vals {
		popV := limitedRandom.PopRandValue()
		results = append(results, popV)
	}
	sort.Strings(results)

	exps := []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}
	if !reflect.DeepEqual(exps, results) {
		t.Errorf("PopRandValue Error:\nvals=>[%s]\nress=>[%s]\n", strings.Join(exps, ","), strings.Join(results, ","))
	}

	limitedRandom.Reset()

	results2 := []string{}
	for range vals {
		popV := limitedRandom.PopRandValue()
		results2 = append(results2, popV)
	}
	sort.Strings(results2)

	if !reflect.DeepEqual(exps, results2) {
		t.Errorf("PopRandValue Error:\nvals=>[%s]\nress=>[%s]\n", strings.Join(exps, ","), strings.Join(results2, ","))
	}

}
