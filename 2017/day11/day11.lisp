#!/usr/bin/sbcl --script

(require "asdf")

(defun get-input ()
	(map 'list (lambda (x) (read-from-string x)) (uiop:split-string (car (uiop:read-file-lines "input.txt")) :separator ",")))

(defun trace-path (steps)
	(let ((x 0)
		(y 0)
		(z 0)
		(mx 0))
		(loop for move in steps do
			(case move
				('N  (setq y (+ y 1)) (setq z (- z 1)))
				('S  (setq y (- y 1)) (setq z (+ z 1)))
				('SE (setq y (- y 1)) (setq x (+ x 1)))
				('SW (setq z (+ z 1)) (setq x (- x 1)))
				('NE (setq z (- z 1)) (setq x (+ x 1)))
				('NW (setq y (+ y 1)) (setq x (- x 1))))
			(setq mx (max mx (distance x y z))))
		(list (distance x y z) mx)))

(defun distance (x y z)
	(/ (+ (abs x) (abs y) (abs z)) 2))

(let ((data (get-input)))
	(princ (trace-path data))
	(terpri))
