#!/usr/bin/sbcl --script

(let ((steps 303)
	(lst '(0))
	(pos 0)
	(pos1 0))
		; part 1 (insert values, find value after 2017)
		(loop for x from 1 to 2017 do
			(setq pos (mod (+ pos steps) x))
			(push x (cdr (nthcdr pos lst)))
			(setq pos (+ pos 1)))
		(setq pos (+ pos 1))
		(if (= pos 2019)
			(setq pos 0))
		(princ (nth pos lst))
		(terpri)

		; part 2 (don't insert, just save the value when at position 1)
		(loop for x from 1 to 50000000 do
			(setq pos (mod (+ pos steps) x))
			(setq pos (+ pos 1))
			(if (= pos 1)
				(setq pos1 x)))
		(princ pos1)
		(terpri))
