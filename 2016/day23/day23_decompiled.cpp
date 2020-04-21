#include <iostream>
using namespace std;

int calc(int a) {
	int loop = a - 2;

	int b = a - 1;
	for (int i = 0; i < loop; i++) {
		a *= b;
		b--;
	}

	a = a + 73 * 95;
	return a;
}

int main() {
	// decompiled version of the program
	// from observation the above loop is a - 2
	// haven't seen other inputs, but I would guess that
	// the numbers 73 and 95 are the things that change

	// part 1
	int result = calc(7);
	cout << result << endl;
	// part 2
	result = calc(12);
	cout << result << endl;
	return 0;
}
