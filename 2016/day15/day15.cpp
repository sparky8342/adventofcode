#include <iostream>
#include <fstream>
#include <regex>
#include <vector>
using namespace std;

struct Disc {
	int positions;
	int at;
};

int find_time (vector<Disc> &discs) {
	int time = 0;
	while (1) {
		bool ok = true;
		for (int i = 0; i < discs.size(); i++) {
			Disc disc = discs[i];
			if ((disc.at + time + i + 1) % disc.positions != 0) {
				ok = false;
				break;
			}
		}

		if (ok) {
			return time;
		}

		time++;
	}
}

int main() {
	const regex disc_regex("^Disc .[0-9] has ([0-9]+) positions; at time=0, it is at position ([0-9]+).$");
	vector <Disc> discs;

	string line;
	ifstream in("input.txt");
	while (getline(in, line)) {
		smatch match;
		if (regex_match(line, match, disc_regex)) {
			Disc disc = Disc{ .positions = stoi(match[1]), .at = stoi(match[2]) };
			discs.push_back(disc);
		}
	}
	in.close();

	// part 1
	int time = find_time(discs);
	cout << time << endl;

	// part 2
	discs.push_back(Disc{ .positions = 11, .at = 0});
	time = find_time(discs);
	cout << time << endl;

	return 0;
}

