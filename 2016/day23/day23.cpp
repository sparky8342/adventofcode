#include <iostream>
#include <string>
#include <fstream>
#include <sstream>
#include <vector>
#include <map>
using namespace std;

enum Ops { cpy, inc, decr, jnz, tgl };

struct Ins {
	Ops op;
	string arg1;
	string arg2;
};

int run_program(vector<Ins> &program, int a, int b, int c, int d) {
	map<string, int> registers = { { "a", a }, { "b", b }, { "c", c }, { "d", d } };

	int pos = 0;

	while (1) {
		if (pos >= program.size()) {
			break;
		}

		Ins ins = program[pos];

		if (ins.op == cpy) {
			if (registers.count(ins.arg2) == 1) {
				if (registers.count(ins.arg1) == 1) {
					registers[ins.arg2] = registers[ins.arg1];
				}
				else {
					registers[ins.arg2] = stoi(ins.arg1);
				}
			}
			pos++;
		}
		else if (ins.op == inc) {
			registers[ins.arg1]++;
			pos++;
		}
		else if (ins.op == decr) {
			registers[ins.arg1]--;
			pos++;
		}
		else if (ins.op == jnz) {
			bool jump = false;
			if (registers.count(ins.arg1) == 1) {
				if (registers[ins.arg1] != 0) {
					jump = true;
				}
			}
			else if (stoi(ins.arg1) != 0) {
				jump = true;
			}

			if (jump == true) {
				if (registers.count(ins.arg2) == 1) {
					pos += registers[ins.arg2];
				}
				else {
					pos += stoi(ins.arg2);
				}
			}
			else {
				pos++;
			}
		}
		else if (ins.op == tgl) {
			int toggle;
			if (registers.count(ins.arg1) == 1) {
				toggle = pos + registers[ins.arg1];
			}
			else {
				toggle = pos + stoi(ins.arg1);
			}

			if (toggle >= 0 || toggle < program.size()) {
				Ins toggle_ins = program[toggle];
				if (toggle_ins.op == cpy) {
					program[toggle].op = jnz;
				}
				else if (toggle_ins.op == inc) {
					program[toggle].op = decr;
				}
				else if (toggle_ins.op == decr) {
					program[toggle].op = inc;
				}
				else if (toggle_ins.op == jnz) {
					program[toggle].op = cpy;
				}
				else if (toggle_ins.op == tgl) {
					program[toggle].op = inc;
				}
			}

			pos++;
		}
	}

	return registers["a"];
}

int main() {
	vector<Ins> program;

	string line;
	ifstream in("input.txt");
	while (getline(in, line)) {
		Ins ins = Ins{};
		
		stringstream ss(line);
		string cmd;
		getline(ss, cmd, ' ');
		if (cmd == "cpy") { ins.op = cpy; }
		else if (cmd == "inc") { ins.op = inc; }
		else if (cmd == "dec") { ins.op = decr; }
		else if (cmd == "jnz") { ins.op = jnz; }
		else if (cmd == "tgl") { ins.op = tgl; }
		getline(ss, ins.arg1, ' ');
		getline(ss, ins.arg2, ' ');

		program.push_back(ins);
	}
	in.close();

	int register_a = run_program(program, 7, 0, 0, 0);
	cout << register_a << endl;

	return 0;
}
