package constant

const (
	ES_INDEX_PREFIX = "evildoer_"
)

func GetESGlobalIndices() []string {
	return []string{ES_INDEX_PREFIX + "*"}
}
