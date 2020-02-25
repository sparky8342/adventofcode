#include <iostream>
#include <string>
#include <fstream>
#include <regex>
using namespace std;

int main() {
	string line;
	ifstream in("input.txt");
	getline(in, line);
	in.close();
	
	string new_line;
	const regex reg("(.*?)\\(([0-9]+)x([0-9]+)\\)(.*)");
	smatch match;

	while (regex_match(line, match, reg)) {
		string prefix = match[1];
		int no_chars = stoi(match[2]);
		int no_repeat = stoi(match[3]);
		string remaining = match[4];

		new_line = new_line + prefix;

		string section;
		section = remaining.substr(0, no_chars);
		for (int i = 0; i < no_repeat; i++) {
			new_line = new_line + section;
		}		

		line = remaining.substr(no_chars, remaining.size() - no_chars);
	}

	cout << new_line.size() << endl;
	return 0;
}
