package solution

func FindFirstBad(version int) int {
	isBad := false
	badVersion := 0
	mid := version/2 + 1
	if mid == 1 {
		// NEED API isBadVersion
		// isBad := isBadVersion(mid)
		if isBad {
			badVersion = version
			return badVersion
		} else {
			return badVersion
		}
	} else {
		// NEED API isBadVersion
		// isBad := isBadVersion(mid)
		if isBad {
			badVersion = FindFirstBad(mid)
		} else {
			badVersion = FindFirstBad(version - mid)
		}
	}
	return badVersion
}
