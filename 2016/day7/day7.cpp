#include <iostream>
#include <string>
#include <fstream>
using namespace std;

int main() {
	int valid = 0;
	string ip_address;
	ifstream in("input.txt");
	while (getline(in, ip_address)) {

		bool in_brackets = false;
		bool found = false;
		for (int i = 0; i < ip_address.size() - 2; i++) {
			if (ip_address[i] == '[') {
				in_brackets = true;
			}
			else if (ip_address[i] == ']') {
				in_brackets = false;
			}
			else {
				if (
					ip_address[i] != ip_address[i + 1]
					&& ip_address[i] == ip_address[i + 3]
					&& ip_address[i + 1] == ip_address[i + 2]
				) {
					if (in_brackets == true) {
						found = false;
						break;
					}
					else {
						found = true;
					}
				}
			}
		}
		if (found == true) {
			valid++;
		}

        }
        in.close();
	cout << valid << endl;
	return 0;
}
