#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <set>
#include <queue>
using namespace std;

struct Floor {
	set <string> generators;
	set <string> microchips;
};

struct Building {
	vector <Floor> floors;
	int elevator;
	int steps;
};

Building copybuilding(Building building) {
	Building copy = Building{
		.floors = {},
		.elevator = building.elevator,
		.steps = building.steps
	};

	for ( auto floor : building.floors ) {
		set<string> gencopy (floor.generators);
		set<string> chipcopy (floor.microchips);
		copy.floors.push_back( Floor{ .generators = gencopy, .microchips = chipcopy } );
	}

	return copy;
}

bool endcondition(Building building) {
	if (building.floors[3].generators.size() == 5 && building.floors[3].microchips.size() == 5) {
		return true;
	}
	return false;
}

void printbuilding(Building building) {
	cout << "elevator: " << to_string(building.elevator) << endl;
	for ( auto floor : building.floors ) {
		set<string>::iterator it;
		cout << "generators: ";
		for ( it = floor.generators.begin(); it != floor.generators.end(); ++it ) {
			cout << *it << ",";
		}
		cout << endl;

		cout << "microchips: ";
		for ( it = floor.microchips.begin(); it != floor.microchips.end(); ++it ) {
			cout << *it << ",";
		}
		cout << endl << "--------------------" << endl;
	}
	cout << endl;
}
	
bool validbuilding(Building building) {
	for ( auto floor : building.floors ) {
		if ( floor.generators.size() == 0 ) {
			continue;
		}

		set<string>::iterator it;
		for ( it = floor.microchips.begin(); it != floor.microchips.end(); ++it ) {
			if ( floor.generators.count( *it ) == 0 ) {
				//cout << "not valid" << endl;
				return false;
			}
		}
	}
	//printbuilding(building);
	//cout << "valid" << endl;
	return true;
}

string serialisebuilding(Building building) {
	string str = "";
	for (int i = 0; i < 4; i++) {
		Floor floor = building.floors[i];
		str = str + to_string(i);
		str = str + ":g:";
		set<string>::iterator it;
		for ( it = floor.generators.begin(); it != floor.generators.end(); ++it ) {
			str = str + *it + ",";
		}
		str = str + "m:";
		for ( it = floor.microchips.begin(); it != floor.microchips.end(); ++it ) {
			str = str + *it + ",";
		}
	}
	str = str + ":" + to_string(building.elevator);
	return str;
}

void bfs(Building building) {
	queue <Building> moves;
	moves.push(building);

	set <string> visited;

	while (moves.size() > 0) {
		Building bl = moves.front();
		moves.pop();

		string serial = serialisebuilding(bl);
		if (visited.count(serial) == 1) {
			continue;
		}
		visited.insert(serial);

		if (!validbuilding(bl)) {
			continue;
		}

		if (endcondition(bl)) {
			cout << to_string(bl.steps) << endl;
			break;
		}

		int elevator = bl.elevator;
		Floor floor = bl.floors[elevator];

		vector<int> floors_moving_to;
		if (elevator > 0) {
			floors_moving_to.push_back(elevator - 1);
		}
		if (elevator < 3) {
			floors_moving_to.push_back(elevator + 1);
		}

		for ( auto destination_floor : floors_moving_to ) {
			// 1 or 2 things from generator set
			set<string>::iterator it;
			set<string>::iterator it2;
			for ( it = floor.generators.begin(); it != floor.generators.end(); ++it ) {
				for ( it2 = floor.generators.begin(); it2 != floor.generators.end(); ++it2 ) {
					Building buildingcopy = copybuilding(bl);
					buildingcopy.floors[elevator].generators.erase(*it);
					buildingcopy.floors[destination_floor].generators.insert(*it);
					if (*it != *it2) {
						buildingcopy.floors[elevator].generators.erase(*it2);
						buildingcopy.floors[destination_floor].generators.insert(*it2);
					}
					buildingcopy.elevator = destination_floor;
					buildingcopy.steps++;
					moves.push(buildingcopy);
				}
			}
				
			// 1 or 2 things from microchip set
			for ( it = floor.microchips.begin(); it != floor.microchips.end(); ++it ) {
				for ( it2 = floor.microchips.begin(); it2 != floor.microchips.end(); ++it2 ) {
					Building buildingcopy = copybuilding(bl);
					buildingcopy.floors[elevator].microchips.erase(*it);
					buildingcopy.floors[destination_floor].microchips.insert(*it);
					if (*it != *it2) {
						buildingcopy.floors[elevator].microchips.erase(*it2);
						buildingcopy.floors[destination_floor].microchips.insert(*it2);
					}
					buildingcopy.elevator = destination_floor;
					buildingcopy.steps++;
					moves.push(buildingcopy);
				}
			}
			
			// 1 from each set
			for ( it = floor.microchips.begin(); it != floor.microchips.end(); ++it ) {
				for ( it2 = floor.generators.begin(); it2 != floor.generators.end(); ++it2 ) {
					Building buildingcopy = copybuilding(bl);
					buildingcopy.floors[elevator].microchips.erase(*it);
					buildingcopy.floors[destination_floor].microchips.insert(*it);
					buildingcopy.floors[elevator].generators.erase(*it2);
					buildingcopy.floors[destination_floor].generators.insert(*it2);
					buildingcopy.elevator = destination_floor;
					buildingcopy.steps++;
					moves.push(buildingcopy);

				}
			}
		}
	}
}

int main() {
	vector <Floor> floors;

	// hardcode initial data for now
	Floor fl1 = Floor{
		.generators = { "polonium", "thulium", "promethium", "ruthenium", "cobalt" },
		.microchips = { "thulium", "ruthenium", "cobalt" }
	};
	//Floor fl1 = Floor{
	//	.generators = {},
	//	.microchips = { "hydrogen", "lithium" }
	//};
	floors.push_back(fl1);

	Floor fl2 = Floor{
		.generators = {},
		.microchips = { "polonium", "promethium" }
	};
	//Floor fl2 = Floor{
	//	.generators = { "hydrogen" },
	//	.microchips = {}
	//};
	floors.push_back(fl2);

	Floor fl3 = Floor{ .generators = {}, .microchips = {} };
	//Floor fl3 = Floor{
	//	.generators = { "lithium" },
	//	.microchips = {}
	//};
	floors.push_back(fl3);

	Floor fl4 = Floor{ .generators = {}, .microchips = {} };
	floors.push_back(fl4);

	Building bl = Building{
		.floors = floors,
		.elevator = 0,
		.steps = 0
	};

	bfs(bl);

	return 0;
}

/* 
example:

The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.

my input:

The first floor contains a polonium generator, a thulium generator, a thulium-compatible microchip, a promethium generator, a ruthenium generator, a ruthenium-compatible microchip, a cobalt generator, and a cobalt-compatible microchip.
The second floor contains a polonium-compatible microchip and a promethium-compatible microchip.
The third floor contains nothing relevant.
The fourth floor contains nothing relevant.
*/
