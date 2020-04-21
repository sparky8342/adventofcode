#include <iostream>
using namespace std;

int main() {
	int a = 7;
	int b = 0;
	int c = 0;
	int d = 0;

	b = a;
	b--;
	for (int i = 0; i < 5; i++) {
		d = a;
		a = 0;
		do {
			c = b;
			do {
				a++;
				c--;
			} while (c != 0);

			d--;
		} while (d != 0);
		b--;
		c = b;
		d = c;
		do {
			d--;
			c++;
		} while (d != 0);
	}
	c = 95;
	do {
		d = 73;
		do {
			a++;
			d--;
		} while (d != 0);
		c--;
	} while (c != 0);

	cout << a << endl;

	return 0;
}
