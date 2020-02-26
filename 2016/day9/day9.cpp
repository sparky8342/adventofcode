#include <iostream>
#include <string>
#include <fstream>
#include <regex>
using namespace std;

long decompress(string line, bool recurse) {
	const regex reg("(.*?)\\(([0-9]+)x([0-9]+)\\)(.*)");
	smatch match;
	long count = 0;

	while (regex_match(line, match, reg)) {
		string prefix = match[1];
		int no_chars = stoi(match[2]);
		int no_repeat = stoi(match[3]);
		string remaining = match[4];

		count += prefix.size();

		string section = remaining.substr(0, no_chars);
		if (recurse) {
			count += decompress(section, true) * no_repeat;
		}
		else {
			count += section.size() * no_repeat;
		}

		line = remaining.substr(no_chars, remaining.size() - no_chars);
	}

	count += line.size();
	return count;
}

int main() {
	string line;
	ifstream in("input.txt");
	getline(in, line);
	in.close();

	long count = decompress(line, false);
	cout << count << endl;

	count = decompress(line, true);
	cout << count << endl;

	return 0;
}
