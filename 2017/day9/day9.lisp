#!/usr/bin/sbcl --script

(require "asdf")

(defun get-data ()
	(let* ((in (uiop:read-file-string "input.txt"))
		; coerce input into a list with a symbol per character, removing final carriage return
		(data (butlast (coerce in 'list))))
	data))

(defun get-score (data)
	(let ((score 0)
		(in-garbage 0)
		(after-exclam 0)
		(level 0)
		(garbage-count 0))
			(loop for chr in data do
				(if (= after-exclam 1)
					(setq after-exclam 0)
					(if (= in-garbage 1)
						(case chr
							(#\> (setq in-garbage 0))
							(#\! (setq after-exclam 1))
							(otherwise (setq garbage-count (+ garbage-count 1))))
						(case chr
							(#\< (setq in-garbage 1))
							(#\{
								(setq score (+ score 1 level))
								(setq level (+ level 1)))
							(#\} (setq level (- level 1)))))))
			(list score garbage-count)))

(let ((data (get-data)))
	(princ (get-score data))
	(terpri))
