#!/usr/bin/sbcl --script

(require "asdf")

(defun get-input ()
	(car (uiop:read-file-lines "input.txt")))

(defun split-comma (lst)
	(map 'list (lambda (x) (parse-integer x)) (uiop:split-string lst :separator ",")))

(defun split-chars-ascii (lst)
	(let ((nums (map 'list (lambda (x) (char-code x)) (coerce lst 'list))))
		(setq nums (append nums '(17 31 73 47 23)))
		nums))

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

(defun knot (lst lengths rounds)
	(let ((pos 0)
		(skip 0))
			(loop repeat rounds do
				(loop for len in lengths do
					(let ((part (subseq lst pos (+ pos len))))
						(let ((part (nreverse part)))
							(loop for i from pos to (+ pos (- len 1)) do
								(setf (nth i lst) (car part))
								(setq part (cdr part)))))
					(setq pos (mod (+ pos len skip) 256))
					(setq skip (+ skip 1))))
			lst))


(let ((data (get-input)))
	; part1
	(let ((lengths (split-comma data))
		(nums (create-nums)))
			;(setf *print-circle* t)
			(setq nums (knot nums lengths 1))
			(princ (* (car nums) (nth 1 nums)))
			(terpri))
	; part2
	(let ((lengths (split-chars-ascii data))
		(nums (create-nums)))
		(setq nums (knot nums lengths 64))
		(loop for i from 0 to 15 do
			(let ((x 0))
				(loop for j from 0 to 15 do
					(setq x (logxor x (nth (+ (* i 16) j) nums))))
				(let ((str (string-downcase (write-to-string x :base 16))))
					(if (= 1 (length str))
						(princ 0))
					(princ str))))
		(terpri)))
