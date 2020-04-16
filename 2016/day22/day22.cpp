#include <iostream>
#include <fstream>
#include <regex>
using namespace std;

struct Node {
	int x;
	int y;
	int size;
	int used;
	int avail;
	int used_percent;
};

int main() {
	const regex node_regex("^/dev/grid/node-x([0-9]+)-y([0-9]+)[ ]+([0-9]+)T[ ]+([0-9]+)T[ ]+([0-9]+)T[ ]+([0-9]+)%$");
	vector<Node> nodes;

	string line;
	smatch match;
	ifstream in("input.txt");
	getline(in,line);
	getline(in,line);
	while(getline(in, line)) {
		regex_match(line, match, node_regex);
		Node node = Node {
			.x = stoi(match[1]),
			.y = stoi(match[2]),
			.size = stoi(match[3]),
			.used = stoi(match[4]),
			.avail = stoi(match[5]),
			.used_percent = stoi(match[6])
		};
		nodes.push_back(node);
	}
	in.close();

	int pairs = 0;
	for (int i = 0; i < nodes.size(); i++) {
		for (int j = 0; j < nodes.size(); j++) {
			if (i != j && nodes[i].used != 0 && nodes[i].used <= nodes[j].avail) {
			       pairs++;
		      	}	       
		}
	}

	cout << pairs << endl;

	return 0;
}
