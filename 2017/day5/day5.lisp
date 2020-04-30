#!/usr/bin/sbcl --script

(require "asdf")

(defun get-input ()
	(map 'list (lambda (x) (parse-integer x)) (uiop:read-file-lines "input.txt")))

(defun jump-around (arr &optional part2)
	(let ((pos 0)
		(steps 0))
		(loop while (< pos (length arr)) do
			(let ((jump (aref arr pos)))
				(if (and part2 (>= jump 3))
					(setf (aref arr pos) (- jump 1))
					(setf (aref arr pos) (+ jump 1)))
				(setq pos (+ pos jump))
				(setq steps (+ steps 1))))
		steps))
				
(defun run ()
	(let ((instructions (get-input)))
		(let ((steps (jump-around (make-array (list-length instructions) :initial-contents instructions)))
			(steps2 (jump-around (make-array (list-length instructions) :initial-contents instructions) t)))
			(princ steps)
			(terpri)
			(princ steps2)
			(terpri))))

(run)
