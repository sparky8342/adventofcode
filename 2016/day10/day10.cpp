#include <iostream>
#include <string>
#include <fstream>
#include <map>
#include <regex>
#include <vector>
using namespace std;

struct Bot {
	int low;
	string low_dest;
	int high;
	string high_dest;
	vector<int> vals;
};

struct Val {
	int bot;
	int val;
};

int main() {
	const regex bot_regex("^bot ([0-9]+) gives low to (bot|output) ([0-9]+) and high to (bot|output) ([0-9]+)$");
	const regex val_regex("^value ([0-9]+) goes to bot ([0-9]+)$");

	map<int, Bot> bots;

	vector <Val> values;
	string line;
	ifstream in("input.txt");
	while (getline(in, line)) {
		smatch match;
		if (regex_match(line, match, bot_regex)) {
			bots[stoi(match[1])] = Bot{
				.low_dest = match[2],
				.low = stoi(match[3]),
				.high_dest = match[4],
				.high = stoi(match[5])
			};
		}
		else if (regex_match(line, match, val_regex)) {
			values.push_back(Val{ .bot = stoi(match[2]), .val = stoi(match[1])});
		}
	}
	in.close();

	map <int, int> output;

	while (values.size() > 0) {
		Val val_struct = values[0];
		int bot_id = val_struct.bot;
		int val = val_struct.val;

		bots[bot_id].vals.push_back(val);

		if (bots[bot_id].vals.size() == 2) {
			sort(bots[bot_id].vals.begin(), bots[bot_id].vals.end());	
			if (bots[bot_id].vals[0] == 17 && bots[bot_id].vals[1] == 61) {
				// part 1
				cout << bot_id << endl;
			}

			if (bots[bot_id].low_dest == "bot") {
				values.push_back(Val{ .bot = bots[bot_id].low, val = bots[bot_id].vals[0] });
			}
			else {
				output[bots[bot_id].low] = bots[bot_id].vals[0];
			}				
			if (bots[bot_id].high_dest == "bot") {
				values.push_back(Val{ .bot = bots[bot_id].high, val = bots[bot_id].vals[1] });
			}
			else {
				output[bots[bot_id].high] = bots[bot_id].vals[1];
			}
			
			bots[bot_id].vals.clear();

		}

		values.erase(values.begin());
	}

	// part 2
	cout << output[0] * output[1] * output[2] << endl;
	return 0;
}
