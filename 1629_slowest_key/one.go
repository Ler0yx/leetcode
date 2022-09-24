package slowestKey

func slowestKey(releaseTimes []int, keysPressed string) byte {

	maxDur := 0
	slowestKey := 0

	for d := 0; d < len(releaseTimes); d++ {

		dur := 0

		if d == 0 {
			dur = releaseTimes[d]

		} else {
			dur = releaseTimes[d] - releaseTimes[d-1]
		}

		if dur > maxDur {
			maxDur = dur
			slowestKey = d

		} else if dur == maxDur {

			if byte(keysPressed[d]) > byte(keysPressed[slowestKey]) {
				slowestKey = d
			}
		}
	}
	return keysPressed[slowestKey]
}
