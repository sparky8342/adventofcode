#!/usr/bin/sbcl --script

(require "asdf")

(defun get-input ()
	(map 'list (lambda (x) (parse-integer x)) (uiop:split-string (car (uiop:read-file-lines "input.txt")) :separator "	")))

(defun cycle (lst)
	(let* ((mx (apply 'max lst))
		(pos (position mx lst)))
			(setf (nth pos lst) 0)
			(loop while (> mx 0) do
				(setq pos (+ pos 1))
				(if (= pos (list-length lst))
					(setq pos 0))
				(setf (nth pos lst) (+ (nth pos lst) 1))
				(setq mx (- mx 1)))
			lst))

(defun get-key (lst)
	(format nil "~{~a~^ ~}" lst))

(defun reallocate (lst)
	(let ((states (make-hash-table :test #'equal))
		(key (get-key lst))
		(amount 0))
			(loop while (not (gethash key states)) do
		  		(setf (gethash key states) amount)
				(setq lst (cycle lst))
				(setq key (get-key lst))
				(setq amount (+ amount 1)))
			(list amount (- amount (gethash key states)))))

(defun run ()
	(let ((input (get-input)))
		(princ (reallocate input))
		(terpri)))

(run)
