#!/usr/bin/sbcl --script

(require "asdf")

(defun convert-line (line)
	(map 'list (lambda (x) (read-from-string x)) (uiop:split-string line :separator "  ")))

(defun all-unique (l)
	(if l
		(if (member (car l) (cdr l))
			nil
			(all-unique (cdr l)))
		t))

(defun count-valid (file)
	(let ((valid 0)
		(in (open file :if-does-not-exist nil)))
			(when in
				(loop for line = (read-line in nil)
					while line do
						(let ((vals (convert-line line)))
							(if (all-unique vals)
								(setq valid (+ valid 1))))))
		valid))

(princ (count-valid "input.txt"))
(terpri)
