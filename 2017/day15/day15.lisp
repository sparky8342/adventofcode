#!/usr/bin/sbcl --script

(defun next (factor n)
	(setq n (* n factor))
	(setq n (mod n 2147483647))
	n)

(defun next-a (n)
	(next 16807 n))

(defun next-b (n)
	(next 48271 n))

(defun next-a-2 (n)
	(setq n (next-a n))
	(loop while (> (mod n 4) 0) do
		(setq n (next-a n)))
	n)

(defun next-b-2 (n)
	(setq n (next-b n))
	(loop while (> (mod n 8) 0) do
		(setq n (next-b n)))
	n)

(defun bot16 (n)
	(mod n 65536))

(defun count-matches (a b gen-a gen-b amount)
	(let ((matches 0))
		(loop repeat amount do
			(setq a (funcall gen-a a))
			(setq b (funcall gen-b b))
			(if (= (bot16 a) (bot16 b))
				(setq matches (+ matches 1))))
		matches))

(let ((a 277)
	(b 349))
		(princ (count-matches a b (lambda (x) (next-a x)) (lambda (x) (next-b x)) 40000000))
		(terpri)
		(princ (count-matches a b (lambda (x) (next-a-2 x)) (lambda (x) (next-b-2 x)) 5000000))
		(terpri))
