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
	for (int i = 0; i < 3; i++) {
		if (building.floors[i].generators.size() > 0 || building.floors[i].microchips.size() > 0) {
			return false;
		}
	}
	return true;
}

bool validbuilding(Building building) {
	for ( auto floor : building.floors ) {
		if ( floor.generators.size() == 0 ) {
			continue;
		}

		set<string>::iterator it;
		for ( it = floor.microchips.begin(); it != floor.microchips.end(); ++it ) {
			if ( floor.generators.count( *it ) == 0 ) {
				return false;
			}
		}
	}
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
				it2 = floor.generators.begin();
				while (*it != *it2) {
					it2++;
				}
				for ( ; it2 != floor.generators.end(); ++it2 ) {
					Building buildingcopy = copybuilding(bl);
					buildingcopy.floors[elevator].generators.erase(*it);
					buildingcopy.floors[destination_floor].generators.insert(*it);
					if (*it != *it2) {
						buildingcopy.floors[elevator].generators.erase(*it2);
						buildingcopy.floors[destination_floor].generators.insert(*it2);
					}
					if (validbuilding(buildingcopy)) {
						buildingcopy.elevator = destination_floor;
						buildingcopy.steps++;
						moves.push(buildingcopy);
					}
				}
			}
				
			// 1 or 2 things from microchip set
			for ( it = floor.microchips.begin(); it != floor.microchips.end(); ++it ) {
				it2 = floor.microchips.begin();
				while (*it != *it2) {
					it2++;
				}
				for ( ; it2 != floor.microchips.end(); ++it2 ) {
					Building buildingcopy = copybuilding(bl);
					buildingcopy.floors[elevator].microchips.erase(*it);
					buildingcopy.floors[destination_floor].microchips.insert(*it);
					if (*it != *it2) {
						buildingcopy.floors[elevator].microchips.erase(*it2);
						buildingcopy.floors[destination_floor].microchips.insert(*it2);
					}
					if (validbuilding(buildingcopy)) {
						buildingcopy.elevator = destination_floor;
						buildingcopy.steps++;
						moves.push(buildingcopy);
					}
				}
			}
			
			// 1 from each set
			if (floor.microchips.size() > 0 && floor.generators.size() > 0) {
				for ( it = floor.microchips.begin(); it != floor.microchips.end(); ++it ) {
					for ( it2 = floor.generators.begin(); it2 != floor.generators.end(); ++it2 ) {
						Building buildingcopy = copybuilding(bl);
						buildingcopy.floors[elevator].microchips.erase(*it);
						buildingcopy.floors[destination_floor].microchips.insert(*it);
						buildingcopy.floors[elevator].generators.erase(*it2);
						buildingcopy.floors[destination_floor].generators.insert(*it2);
						if (validbuilding(buildingcopy)) {
							buildingcopy.elevator = destination_floor;
							buildingcopy.steps++;
							moves.push(buildingcopy);
						}
					}
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
	floors.push_back(fl1);

	Floor fl2 = Floor{
		.generators = {},
		.microchips = { "polonium", "promethium" }
	};
	floors.push_back(fl2);

	Floor fl3 = Floor{ .generators = {}, .microchips = {} };
	floors.push_back(fl3);

	Floor fl4 = Floor{ .generators = {}, .microchips = {} };
	floors.push_back(fl4);

	Building bl = Building{
		.floors = floors,
		.elevator = 0,
		.steps = 0
	};

	// part 1
	bfs(bl);

	// part 2
	bl.floors[0].generators.insert("elerium");
	bl.floors[0].generators.insert("dilithium");
	bl.floors[0].microchips.insert("elerium");
	bl.floors[0].microchips.insert("dilithium");
	bfs(bl);

	return 0;
}

/* 
input:

The first floor contains a polonium generator, a thulium generator, a thulium-compatible microchip, a promethium generator, a ruthenium generator, a ruthenium-compatible microchip, a cobalt generator, and a cobalt-compatible microchip.
The second floor contains a polonium-compatible microchip and a promethium-compatible microchip.
The third floor contains nothing relevant.
The fourth floor contains nothing relevant.

part 2 additions (to first floor):

An elerium generator.
An elerium-compatible microchip.
A dilithium generator.
A dilithium-compatible microchip.
*/
