package gochill

import json "github.com/bitly/go-simplejson"

//CombineMaps used to merge list of maps into one single map
//contain of key string and value interface
func CombineMaps(maps ...map[string]interface{}) map[string]interface{} {
	var m map[string]interface{}
	m = make(map[string]interface{})

	if len(maps) > 0 {
		for _, mapped := range maps {
			if len(mapped) > 0 {
				for k, v := range mapped {
					m[k] = v
				}
			}
		}
	}

	return m
}

//MapToJSON used to convert map to json byte
func MapToJSON(maps map[string]interface{}) []byte {
	/*
		I'm not using native golang json encoding because
		it's very hard to convert map to json because native
		golang use struct to marshal and unmarshal data, we
		doesn't need any struct here.

		I'm using third party library here from bitly/go-simplejson
		it's easier to convert map to json, just iterate the map, set
		and convert it to json.
	*/
	jsonEngine := json.New()
	if len(maps) > 0 {
		for k, v := range maps {
			jsonEngine.Set(k, v)
		}
	}

	jsonByte, err := jsonEngine.MarshalJSON()
	if err != nil {
		return nil
	}

	return jsonByte
}
