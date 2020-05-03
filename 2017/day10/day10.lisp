#!/usr/bin/sbcl --script

(require "asdf")

(defun get-input ()
      (map 'list (lambda (x) (parse-integer x)) (uiop:split-string (car (uiop:read-file-lines "input.txt")) :separator ",")))

(defun create-nums ()
	(let ((lst (create-lst)))
		(let ((lst (circular lst)))
			lst)))

(defun create-lst ()
	(let ((nums ()))
		(dotimes (i 256 (nreverse nums))
			(push i nums))))

(defun circular (items)
	; from lisp cookbook
	"Modifies the last cdr of list ITEMS, returning a circular list"
	(setf (cdr (last items)) items)
	items)

(defun knot (lst lengths)
	(let ((pos 0)
		(skip 0))
			(loop for len in lengths do
				(let ((part (subseq lst pos (+ pos len))))
					(let ((part (nreverse part)))
						(loop for i from pos to (+ pos (- len 1)) do
							(setf (nth i lst) (car part))
							(setq part (cdr part)))))
				(setq pos (mod (+ pos len skip) 256))
				(setq skip (+ skip 1)))
			lst))

(let ((lengths (get-input))
	(nums (create-nums)))
		;(setf *print-circle* t)
		(setq nums (knot nums lengths))
		(princ (* (car nums) (nth 1 nums)))
		(terpri))
