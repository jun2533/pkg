package mapstr_test

import (
	"testing"

	"github.com/jun2533/pkg/mapstr"
)

func TestSetValueToMapStrByTagsWithTagName(t *testing.T) {
	type testStruct struct {
		FieldString *string `field:"field-string"`
		FieldInt    *int    `field:"field-int"`
	}

	tmpStr := "tmpstr"
	tmpStruct := testStruct{FieldString: &tmpStr}

	testcases := map[string]interface{}{
		"nil-case": nil,
		"struct":   tmpStruct,
	}

	// for key, caseItem := range testcases {
	// 	returnMapStr := mapstr.SetValueToMapStrByTagsWithTagName(caseItem, "field")
	// 	require.NotNil(t, returnMapStr)
	// 	t.Logf("return: key: %s %#v", key, returnMapStr)
	// }
	returnMapStr := mapstr.SetValueToMapStrByTagsWithTagName(testcases["struct"], "field")
	t.Logf("return: key: %#v", returnMapStr)
}
