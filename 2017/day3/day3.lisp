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

(defun get-grid-key (x y)
	(concatenate 'string (write-to-string x) ":" (write-to-string y)))

(defun surround-total (grid x y)
	(let ((sum 0))
		(loop for dx in '(-1 0 1) do
			(loop for dy in '(-1 0 1) do
				(let ((key (get-grid-key (+ x dx) (+ y dy))))
					(if (gethash key grid)
						(setq sum (+ sum (gethash key grid)))))))
		(+ 0 sum)))

(defun spiral (n)
	(let ((grid (make-hash-table :test #'equal))
		(x 0)
		(y 0)
		(key (get-grid-key 0 0)))

		(setf (gethash key grid) 1)
		(loop while (> n (gethash key grid)) do
			(if (and (or (/= x y) (>= x 0)) (<= (abs x) (abs y)))
				(if (>= y 0)
					(setq x (+ x 1))
					(setq x (- x 1)))
				(if (>= x 0)
					(setq y (- y 1))
					(setq y (+ y 1))))
			(setq key (get-grid-key x y))
			(setf (gethash key grid) (surround-total grid x y)))
		(gethash key grid)))

(let ((input 361527))
	(princ (distance input))
	(terpri)
	(princ (spiral input))
	(terpri))
