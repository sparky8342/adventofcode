#!/usr/bin/sbcl --script

(defun next (factor n)
	(setq n (* n factor))
	(setq n (mod n 2147483647))
	n)

(defun next-a (n)
	(next 16807 n))

(defun next-b (n)
	(next 48271 n))

(defun bot16 (n)
	(mod n 65536))

(let ((a 277)
	(b 349)
	(matches 0))
		(dotimes (i 40000000)
			(setq a (next-a a))
			(setq b (next-b b))
			(if (= (bot16 a) (bot16 b))
				(setq matches (+ matches 1))))
		(princ matches)
		(terpri))
