package models

import "regexp"

func HydrateShallowJson(object string, manifest map[string]string) (string, error) {
	var err error
	replacementsAvailable := true
	for replacementsAvailable {
		begin := object
		for k, v := range manifest {
			object, err = replaceUUID(object, k, v)
			if err != nil {
				return "", err
			}
		}
		end := object
		if begin == end {
			break
		}
	}
	return object, nil
}

func replaceUUID(input, match, replace string) (string, error) {
    // Define the regex pattern to match the UUID within quoted values but not as a key
    pattern := `(?<!":\s*)"` + regexp.QuoteMeta(match) + `"`

    // Compile the regex
    re, err := regexp.Compile(pattern)
    if err != nil {
        return "", err
    }

    // Replace all matches with the new UUID
    result := re.ReplaceAllString(input, `"`+replace+`"`)
    return result, nil
}