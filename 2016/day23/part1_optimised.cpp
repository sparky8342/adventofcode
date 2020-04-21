#include <iostream>
using namespace std;

int main() {
	int a = 7;

	int b = a - 1;
	for (int i = 0; i < 5; i++) {
		a *= b;
		b--;
	}

	a = a + 73 * 95;

	cout << a << endl;

	return 0;
}
