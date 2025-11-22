#include "collatz_conjecture.h"

int steps(int n) {
	if (n <= 0) {
		return -1;
	} else if (n == 1) {
		return 0;
	} else if (n % 2 == 0) {
		return 1 + steps(n / 2);
	} else {
		return 1 + steps(3 * n + 1);
	}
}
