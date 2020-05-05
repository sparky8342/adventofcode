#!/usr/bin/sbcl --script

(require "asdf")

(defun get-data ()
	(let ((ids (make-hash-table)))
		(with-open-file (stream "input.txt")
			(loop for line = (read-line stream nil)
				while line do
					(setq line (remove-if (lambda (ch) (find ch ",")) line))
					(let ((parts (uiop:split-string line :separator " ")))
						(let ((key (parse-integer (car parts)))
							(vals (map 'list (lambda (x) (parse-integer x)) (cdr (cdr parts)))))
								(setf (gethash key ids) vals)))))
		ids))
				  
(defun traverse (key datahash &optional visited)
	(if (not visited)
		(setq visited (make-hash-table)))

	(when (not (gethash key visited))
		(setf (gethash key visited) 1)
		(loop for val in (gethash key datahash) do
			(traverse val datahash visited)))
	visited)

(let ((data (get-data)))
	(let ((visited (traverse 0 data)))
		; part 1
		(princ (hash-table-count visited))
		(terpri)
		; part2
		; delete ids seen already
		(loop for key being the hash-keys of visited do
			(remhash key data))
		(let ((groups 1))
			(loop while (> (hash-table-count data) 0) do
				; fetch a key, find all in the same group, delete until empty
				(let ((key (first (loop for key being the hash-keys of data collect key))))
					(setq visited (traverse key data)))
				(setq groups (+ groups 1))
				(loop for key being the hash-keys of visited do
					(remhash key data)))
			(princ groups)
			(terpri))))
