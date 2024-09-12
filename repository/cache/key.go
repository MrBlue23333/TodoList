package cache

import "fmt"

// TaskViewKey Returns the corresponding key by task id
func TaskViewKey(id int64) string {
	return fmt.Sprintf("view:task:%d", id)
}
