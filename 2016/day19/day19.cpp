#include <iostream>
using namespace std;

#define NUM_ELVES 3005290

struct Elf {
	int id;
	Elf *next;
};

Elf* init_elves() {
	Elf* head = new Elf;
	head->id = 1;

	Elf* elf = head;

	for (int i = 2; i <= NUM_ELVES; i++) {
		Elf* next = new Elf;
		next->id = i;
		elf->next = next;
		elf = elf->next;
	}

	elf->next = head;
	return head;
}

int part1() {
	Elf* elf = init_elves();

	while (elf != elf->next) {
		elf->next = elf->next->next;
		elf = elf->next;
	}

	return elf->id;
}

int part2() {
	Elf* elf = init_elves();

	Elf* opposite = elf;
	for (int i = 1; i < NUM_ELVES / 2; i++) {
		opposite = opposite->next;
	}

	while (elf != elf->next) {
		opposite->next = opposite->next->next;
		elf = elf->next;

		opposite->next = opposite->next->next;
		opposite = opposite->next;
		elf = elf->next;
	}

	return elf->id;
}

int main() {
	int elf = part1();
	cout << elf << endl;
	elf = part2();
	cout << elf << endl;

	return 0;
}
