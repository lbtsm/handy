package maphelper

import "reflect"

// CheckDesMapDiffFromSrc check diff des map by src
func CheckDesMapDiffFromSrc(src map[string]interface{}, des map[string]interface{}) map[string]map[string]interface{} {
	updateMap := map[string]map[string]interface{}{
		"add":    make(map[string]interface{}),
		"del":    make(map[string]interface{}),
		"modify": make(map[string]interface{}),
	}
	// get the deleted or modified field
	for sKey, sValue := range src {
		dValue, ok := des[sKey]
		if !ok {
			updateMap["del"][sKey] = struct{}{}
			continue
		}
		if !reflect.DeepEqual(sValue, dValue) {
			updateMap["modify"][sKey] = dValue
		}
	}

	// get the added field
	for dKey, dValue := range des {
		_, ok := src[dKey]
		if !ok {
			updateMap["add"][dKey] = dValue
		}
	}

	return updateMap
}
