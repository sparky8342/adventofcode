#!/usr/bin/python3
import re

def check_record(record):
	for key in ('byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid'):
		if not key in record:
			return False, False

	valid = True
	if int(record['byr']) < 1920 or int(record['byr']) > 2002:
		valid = False

	if int(record['iyr']) < 2010 or int(record['iyr']) > 2020:
		valid = False

	if int(record['eyr']) < 2020 or int(record['eyr']) > 2030:
		valid = False

	height_type = record['hgt'][-2:]

	if height_type == 'in':
		height = int(record['hgt'][:-2])
		if height < 59 or height > 76:
			valid = False
	elif height_type == 'cm':
		height = int(record['hgt'][:-2])
		if height < 150 or height > 193:
			valid = False
	else:
		valid = False

	if record['hcl'][0] == '#':
		if not re.search("^[0-9a-f]{6}$", record['hcl'][1:]):
			valid = False
	else:
		valid = False

	if record['ecl'] not in eye_colours:
		valid = False

	if not re.search("^[0-9]{9}$", record['pid']):
		valid = False

	return True, valid


present_count = 0
valid_count = 0
eye_colours = {'amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'}
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
			if present:
				present_count += 1
				if valid:
					valid_count += 1
			record = {}

	# check last one
	present, valid = check_record(record)
	if present:
		present_count += 1
		if valid:
			valid_count += 1
	
print(present_count)
print(valid_count)
