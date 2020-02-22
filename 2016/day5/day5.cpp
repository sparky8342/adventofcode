#include <iostream>
#include <fstream>
#include <openssl/md5.h>
#include <sstream>
#include <iomanip>
using namespace std;

int main() {
	string door_id;
	ifstream in("input.txt");
	getline(in, door_id);
	in.close();

	int index = 0;

	string password;
	string password2 = "--------";

	while (1) {
		string str = door_id + to_string(index);

		unsigned char result[MD5_DIGEST_LENGTH];
		MD5((unsigned char*)str.c_str(), str.size(), result);

		ostringstream sout;
		sout<<hex<<setfill('0');
		for(long long c: result)
		{
			    sout<<setw(2)<<(long long)c;
		}

		string out = sout.str();
		bool found = true;
		for (int i = 0; i < 5; i++) {
			if (out[i] != '0') {
				found = false;
				break;
			}
		}
		if (found == true) {
			if (password.size() < 8) {
				password = password + out[5];
			}
			int pos = out[5] - '0';
			if (pos >= 0 && pos <= 7 && password2[pos] == '-') {
				password2[pos] = out[6];
				bool done = true;
				for (int i = 0; i < 8; i++) {
					if (password2[i] == '-') {
						done = false;
						break;
					}
				}
				if (done == true) {
					break;
				}
			}
		}

		index++;
	}
	cout << password << endl << password2 << endl;
	return 0;
}
