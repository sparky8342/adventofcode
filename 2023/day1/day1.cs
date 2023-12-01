using System;
using System.IO;
using System.Collections.Generic;

public class AOC_2023_DAY1 {
	static public void Main() {

		string[] lines = File.ReadAllLines("input.txt");

		int total = 0;
		foreach (string line in lines) {
			List<int> digits = new List<int>();
			foreach (char ch in line) {
				if (ch >= '0' && ch <= '9') {
					digits.Add(ch - '0');
				}
			}
			total += digits[0] * 10 + digits[digits.Count - 1];

   		}

		Console.WriteLine(total);

		var words = new Dictionary<string, int>{
			["one"] = 1,
			["two"] = 2,
			["three"] = 3,
			["four"] = 4,
			["five"] = 5,
			["six"] = 6,
			["seven"] = 7,
			["eight"] = 8,
			["nine"] = 9
		};

		total = 0;
		foreach (string line in lines) {
			int first = -1;
			int last = -1;

			bool found = false;
			for (int i = 0; i < line.Length; i++) {
				if (line[i] >= '0' && line[i] <= '9') {
					first = line[i] - '0';
					break;
				} else {
					for (int j = 3; j <= 5 && i + j < line.Length; j++) {
						string sub = line.Substring(i, j);
						if (words.ContainsKey(sub)) {
							first = words[sub];
							found = true;
							break;
						}
					}
					if (found) {
						break;
					}
				}
			}

			found = false;
			for (int i = line.Length - 1; i >= 0; i--) {
				if (line[i] >= '0' && line[i] <= '9') {
					last = line[i] - '0';
					break;
				} else {
					for (int j = 3; j <= 5 && i + j <= line.Length; j++) {
						string sub = line.Substring(i, j);
						if (words.ContainsKey(sub)) {
							last = words[sub];
							found = true;
							break;
						}
					}
					if (found) {
						break;
					}
				}
			}

			total += first * 10 + last;
		}

		Console.WriteLine(total);
	}
}
