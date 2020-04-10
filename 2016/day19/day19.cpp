#include <iostream>
#include <fstream>
#include <list>
using namespace std;

#define NUM_ELVES 3005290

int main() {
	int num_elves = NUM_ELVES;

	// use a double linked list
	list<int> elves;

	for (int i = 0; i < num_elves; i++) {
		elves.push_back(num_elves - i);
	}

	list<int>::iterator it;
       	it = elves.end();
	it--;

	while (elves.size() > 1) {
		if (it == elves.begin()) {
			it = elves.end();
		}
		it--;

		list<int>::iterator newit = it;

		if (newit == elves.begin()) {
			newit = elves.end();
		}
		newit--;

		elves.erase(it); // erasing invalidates this iterator

		it = newit;
	}

	cout << *elves.begin() << endl;

	return 0;
}
