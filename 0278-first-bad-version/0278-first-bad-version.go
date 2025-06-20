/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

/**
 * 111100000000000
 * l  m  r      
 *
 */

func firstBadVersion(n int) int {
    low := 1
    high := n
    for (low < high) {
        midd := (high + low) / 2
        if (isBadVersion(midd)) {
            high = midd
        } else {
            low = midd + 1
        }
    }

    return high
}