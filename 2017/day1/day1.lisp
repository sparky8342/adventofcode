#!/usr/bin/sbcl --script

(require "asdf")

(let* ((in (uiop:read-file-string "input.txt"))
	; coerce input into a list, removing final carriage return
	(digits (butlast (coerce in 'list)))
	(len (length digits))
	(half_len (/ len 2))
	(sum_part1 0)
	(sum_part2 0))

	; repeat the list
	(setf digits (append digits digits))

	(loop for i from 0 to (- len 1) do
		; part 1
		(if (eq (nth i digits) (nth (+ i 1) digits))
			(setq sum_part1 (+ sum_part1 (digit-char-p (nth i digits)))))
		; part 2
		(if (eq (nth i digits) (nth (+ i half_len) digits))
			(setq sum_part2 (+ sum_part2 (digit-char-p (nth i digits))))))

	(princ sum_part1)
	(terpri)
	(princ sum_part2)
	(terpri))
