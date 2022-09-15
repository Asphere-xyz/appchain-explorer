package gateway

func checkHasMore(limit uint32, length int) (int, bool) {
	hasMore := length > int(limit)
	if length <= int(limit) {
		return length, hasMore
	}
	return int(limit), hasMore
}
