#!/usr/bin/sbcl --script

(require "asdf")

(defun convert-line (line)
	(map 'list (lambda (x) (read-from-string x)) (uiop:split-string line :separator " ")))

(defun convert-line-sort-chars (line)
	(map 'list (lambda (x) (sort-chars x)) (uiop:split-string line :separator " ")))

(defun sort-chars (str)
	(read-from-string (format nil "~{~a~}" (sort (coerce str 'list) #'char-lessp))))

(defun all-unique (lst)
	(if lst
		(if (member (car lst) (cdr lst))
			nil
			(all-unique (cdr lst)))
		t))

(defun count-valid (file)
	(let ((part1 0)
		(part2 0)
		(in (open file :if-does-not-exist nil)))
			(when in
				(loop for line = (read-line in nil)
					while line do
						(let ((vals (convert-line line))
							(vals2 (convert-line-sort-chars line)))
							(if (all-unique vals)
								(setq part1 (+ part1 1)))
							(if (all-unique vals2)
								(setq part2 (+ part2 1))))))
		(list part1 part2)))

(let ((ans (count-valid "input.txt")))
	(princ (car ans))
	(terpri)
	(princ (car (cdr ans)))
	(terpri))
