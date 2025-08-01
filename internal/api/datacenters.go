package api

import "strings"

type DataCenter struct {
	ID                       int32   `json:"id"`
	Name                     string  `json:"name"`
	VpsTariffs               []int32 `json:"vps_tariffs"`
	DriveTariffs             []int32 `json:"drive_tariffs"`
	AdditionalDriveTariffs   []int32 `json:"additional_drive_tariffs"`
	AdditionalServiceTariffs []int32 `json:"additional_service_tariffs"`
}

type DataCentersResponse struct {
	DataCenters []DataCenter `json:"datacenters"`
}

// GetDatacenterCode returns a unique code for the data center by splitting name by ':' and transliterating it.
func (dc *DataCenter) GetDatacenterCode() string {
	parts := strings.Split(dc.Name, ":")
	if len(parts) > 1 {
		return strings.ToUpper(transliterate(parts[0]))
	}
	return strings.ToUpper(transliterate(dc.Name))
}

// GetDatacenterCountryCode returns the country code for the data center based on its name (if name has country name as substring).
func (dc *DataCenter) GetDatacenterCountryCode() string {
	for country, code := range RuCountryToCode {
		if strings.Contains(dc.Name, country) {
			return code
		}
	}
	return ""
}

// Transliteration of Russian characters to Latin.
var rumap = map[rune]string{
	'а': "a", 'А': "A",
	'б': "b", 'Б': "B",
	'в': "v", 'В': "V",
	'г': "g", 'Г': "G",
	'д': "d", 'Д': "D",
	'е': "e", 'Е': "E",
	'ё': "yo", 'Ё': "YO",
	'ж': "zh", 'Ж': "ZH",
	'з': "z", 'З': "Z",
	'и': "i", 'И': "I",
	'й': "j", 'Й': "J",
	'к': "k", 'К': "K",
	'л': "l", 'Л': "L",
	'м': "m", 'М': "M",
	'н': "n", 'Н': "N",
	'о': "o", 'О': "O",
	'п': "p", 'П': "P",
	'р': "r", 'Р': "R",
	'с': "s", 'С': "S",
	'т': "t", 'Т': "T",
	'у': "u", 'У': "U",
	'ф': "f", 'Ф': "F",
	'х': "h", 'Х': "H",
	'ц': "c", 'Ц': "C",
	'ч': "ch", 'Ч': "CH",
	'ш': "sh", 'Ш': "SH",
	'щ': "sch", 'Щ': "SCH",
	'ъ': "", 'Ъ': "",
	'ы': "y", 'Ы': "Y",
	'ь': "", 'Ь': "",
	'э': "e", 'Э': "E",
	'ю': "ju", 'Ю': "JU",
	'я': "ja",
}

// transliterate converts a string from Russian to Latin characters.
func transliterate(s string) string {
	var result []rune
	for _, r := range s {
		if tr, ok := rumap[r]; ok {
			result = append(result, []rune(tr)...)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
