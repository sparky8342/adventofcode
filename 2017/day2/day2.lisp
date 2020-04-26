#!/usr/bin/sbcl --script

(require "asdf")

(defun convert-line (line)
	(map 'list (lambda (x) (parse-integer x)) (uiop:split-string line :separator "	")))
					
(defun min-max-difference (vals)
	(- (apply 'max vals) (apply 'min vals)))

(defun division (vals)
	(loop named loop-1
		for n in vals
		do (loop for n2 in vals
			do (if (> n n2)
				(let ((div (/ n n2)))
					(if (= div (floor div))
						(return-from loop-1 (values div))))))))

(let ((part1 0)
	(part2 0)
	(in (open "input.txt" :if-does-not-exist nil)))
	(when in
		(loop for line = (read-line in nil)
			while line do
				(let ((vals (convert-line line)))
					(setq part1 (+ part1 (min-max-difference vals)))
					(setq part2 (+ part2 (division vals))))))
	(close in)
	(princ part1)
	(terpri)
	(princ part2)
	(terpri))
