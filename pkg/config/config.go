package config

// name of the input / output files
var DataName = "peaceIndex"

// overrides any date found
var CustomYear = 0

// fallback value if target data column does not have the same value as dataName
var CustomValueName = "Score"

// option to print the output on the console
var PrintAllData = false

// country name harmonization: if the key is not found, it searches the values
var CountryVariations = map[string][]string{
	"bahamas":                      {"the bahamas"},
	"czechia":                      {"czech republic"},
	"cape verde":                   {"cabo verde"},
	"congo":                        {"republic of the congo"},
	"cote d'ivoire":                {"ivory coast"},
	"democratic republic of congo": {"dr congo", "democratic republic of the congo"},
	"eswatini":                     {"lesotho"},
	"gambia":                       {"the gambia"},
	"gibraltar":                    {"gibraltar (uk)"},
	"greenland":                    {"greenland (denmark)"},
	"korea (the democratic people's republic of)": {"north korea"},
	"korea (the republic of)":                     {"south korea"},
	"sao tome and principe":                       {"são tomé and príncipe"},
	"hong kong":                                   {"hong kong (china)"},
	"timor-leste":                                 {"east timor"},
	"united states":                               {"united states of america"},
}

// used to identify the potential date column
var PossibleDateColumns = []string{
	"Year",
	"Date",
}

// indexes the data in the SQL table for which this script was created
var AttributeList = []string{
	"area",
	"population",
	"hdi",
	"ldi",
	"peaceIndex",
}

// used for parsing dates if the date column contains more than just the year
var PossibleDateFormats = []string{
	"2 Jan 2006",
}
