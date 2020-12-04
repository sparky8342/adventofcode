#!/usr/bin/python3
import re

required_keys = {'byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid'}
min_max = {
	'byr' : (1920, 2002),
	'iyr' : (2010, 2020),
	'eyr' : (2020, 2030)
}
height_limits = {
	'in' : (59, 76),
	'cm' : (150, 193)
}
eye_colours = {'amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'}

def check_record(record):
	for key in required_keys:
		if not key in record:
			return False, False

	for key in min_max:
		mn, mx = min_max[key]
		val = int(record[key])
		if val < mn or val > mx:
			return True, False

	height_type = record['hgt'][-2:]
	if height_type not in height_limits:
		return True, False
	height = int(record['hgt'][:-2])
	mn, mx = height_limits[height_type]
	if height < mn or height > mx:
		return True, False

	if not re.search("^#[0-9a-f]{6}$", record['hcl']):
		return True, False

	if record['ecl'] not in eye_colours:
		return True, False

	if not re.search("^[0-9]{9}$", record['pid']):
		return True, False

	return True, True

present_count = 0
valid_count = 0
with open('input.txt') as f:
	record = {}
	for line in f:
		line = line.strip()

		if line:
			parts = line.split()
			for part in parts:
				key, value = part.split(":")
				record[key] = value
		else:
			present, valid = check_record(record)
			present_count += present
			valid_count += valid
			record = {}

	# check last one
	present, valid = check_record(record)
	present_count += present
	valid_count += valid

print(present_count)
print(valid_count)
