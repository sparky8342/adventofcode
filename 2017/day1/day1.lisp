#!/usr/bin/sbcl --script

(require "asdf")

(let* ((in (uiop:read-file-string "input.txt"))
	; coerce input into a list, removing final carriage return
	(captcha (butlast (coerce in 'list)))
	(len (length captcha))
	(half_len (/ len 2))
	(sum_part1 0)
	(sum_part2 0))

	; repeat the list
	(setf captcha (append captcha captcha))

	(loop for i from 0 to (- len 1) do
		; part 1
		(if (eq (nth i captcha) (nth (+ i 1) captcha))
			(setq sum_part1 (+ sum_part1 (digit-char-p (nth i captcha)))))
		; part 2
		(if (eq (nth i captcha) (nth (+ i half_len) captcha))
			(setq sum_part2 (+ sum_part2 (digit-char-p (nth i captcha))))))

	(princ sum_part1)
	(terpri)
	(princ sum_part2)
	(terpri))
