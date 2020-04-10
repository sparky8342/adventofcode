#include <iostream>
#include <fstream>
#include <openssl/md5.h>
#include <sstream>
#include <iomanip>
#include <set>
#include <vector>
using namespace std;

#define INPUT_PASSCODE "gdjjyniy"

struct Space {
	int x;
	int y;       
	string path;
};

struct Move {
	int dx;
	int dy;
	string dir;
};

struct Path {
	int length;
	string path;
};

string get_hash(string str) {
	unsigned char result[MD5_DIGEST_LENGTH];
	MD5((unsigned char*)str.c_str(), str.size(), result);

	ostringstream sout;
	sout<<hex<<setfill('0');
	for(long long c: result)
	{
		    sout<<setw(2)<<(long long)c;
	}

	string out = sout.str();
	return out;
}

void search (Space space, Path &bestpath, Path &longestpath) {
	const vector<Move> moves = {
		Move{ .dx =  0, .dy = -1, .dir = "U" },
		Move{ .dx =  0, .dy =  1, .dir = "D" },
		Move{ .dx = -1, .dy =  0, .dir = "L" },
		Move{ .dx =  1, .dy =  0, .dir = "R" }
	};
	const set<char> openchars = { 'b', 'c', 'd', 'e', 'f' };

	// don't leave the map
	if (space.x < 0 || space.x > 3 || space.y < 0 || space.y > 3) {
		return;
	}

	if (space.x == 3 && space.y == 3) {
		// found the destination
		if (space.path.size() < bestpath.length) {
			bestpath.length = space.path.size();
			bestpath.path = space.path;
		}
		if (space.path.size() > longestpath.length) {
			longestpath.length = space.path.size();
		}
		return;
	}

	string hash = get_hash(space.path);

	// go to open spaces	
	for (int i = 0; i <= 3; i++) {
		if (openchars.count(hash[i]) == 1) {
			Move move = moves[i];
			Space newspace = { .x = space.x + move.dx, .y = space.y + move.dy, .path = space.path + move.dir };
			search(newspace, bestpath, longestpath);
		}
	}
}

int main() {
	string passcode = INPUT_PASSCODE;
	Space start = Space { .x = 0, .y = 0, .path = passcode};
	Path bestpath = { .length = 9999, .path = "" };
	Path longestpath = { .length = 0, .path = "" };
	search(start, bestpath, longestpath);	

	cout << bestpath.path.substr(passcode.size(), bestpath.path.size() - passcode.size()) << endl;
	cout << longestpath.length - passcode.size() << endl;

	return 0;
}
