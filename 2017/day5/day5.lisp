#!/usr/bin/sbcl --script

(require "asdf")

(defun get-input ()
	(map 'list (lambda (x) (parse-integer x)) (uiop:read-file-lines "input.txt")))

(defun jump-around (lst &optional part2)
	(let ((pos 0)
		(steps 0))
		(loop while (< pos (list-length lst)) do
			(let ((jump (nth pos lst)))
				(if (and part2 (>= jump 3))
					(setf (nth pos lst) (- jump 1))
					(setf (nth pos lst) (+ jump 1)))
				(setq pos (+ pos jump))
				(setq steps (+ steps 1))))
		steps))
				
(defun run ()
	(let ((instructions (get-input)))
		(let ((steps (jump-around (copy-list instructions)))
			(steps2 (jump-around (copy-list instructions) t)))
			(princ steps)
			(terpri)
			(princ steps2)
			(terpri))))

(run)
