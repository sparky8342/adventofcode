#!/usr/bin/sbcl --script

(let ((steps 303)
	(lst '(0))
	(pos 0))
		(loop for x from 1 to 2017 do
			(setq pos (mod (+ pos steps) x))
			(push x (cdr (nthcdr pos lst)))
			(setq pos (+ pos 1)))
		(setq pos (+ pos 1))
		(if (= pos 2019)
			(setq pos 0))
		(princ (nth pos lst))
		(terpri))
