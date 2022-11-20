package helpers

import "fmt"

func GetAggregatedTimeline(userId uint64) string {
	return fmt.Sprintf("{user_id: %v}", userId)
}
