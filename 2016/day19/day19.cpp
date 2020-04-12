#include <iostream>
#include <fstream>
#include <list>
using namespace std;

#define NUM_ELVES 3005290

list<int> init_elves() {
	int num_elves = NUM_ELVES;

	// use a double linked list
	list<int> elves;

	for (int i = 0; i < num_elves; i++) {
		elves.push_back(num_elves - i);
	}
	return elves;
}

int part1() {
	list<int> elves = init_elves();

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

	return *elves.begin();
}

int part2() {
	list<int> elves = init_elves();

	list<int>::iterator it;
	it = elves.end();
	it--;

	list<int>::iterator oppositeit = it;
	int steps = elves.size() / 2;
	for (int i = 1; i <= steps; i++) {
		oppositeit--;
	}

	while (elves.size() > 1) {
		list<int>::iterator eraseit = oppositeit;
		int moves = 1;
		if (elves.size() % 2 == 1) {
			moves++;
		}
		for (int i = 1; i <= moves; i++) {
			if (oppositeit == elves.begin()) {
				oppositeit = elves.end();
			}
			oppositeit--;
		}
		elves.erase(eraseit);

		if (it == elves.begin()) {
			it = elves.end();
		}
		it--;

		if (it == oppositeit) {
			if (oppositeit == elves.begin()) {
				oppositeit = elves.end();
			}
			oppositeit--;
		}
	}

	return *elves.begin();
}


int main() {
	int elf = part1();
	cout << elf << endl;
	elf = part2();
	cout << elf << endl;

	return 0;
}
