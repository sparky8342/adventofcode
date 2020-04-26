#!/usr/bin/sbcl --script

(defun distance (n)
	(let ((x 1)
		(side 0))
		(loop while (< x n) do
			(setq side (+ side 2))
			(let ((L 1))
				(loop while (and (<= L 4) (< x n)) do
					(setq x (+ x side))
					(setq L (+ L 1)))))
		(let ((d (list (- x n) (- n (- x side)))))
			(- side (apply 'min d)))))

(let ((input 361527))
	(princ (distance input))
	(terpri))
