#!/usr/bin/sbcl --script

(require "asdf")

(defun convert-line (line)
	(map 'list (lambda (x) (parse-integer x)) (uiop:split-string line :separator "	")))
					
(defun min-max-difference (vals)
	(- (apply 'max vals) (apply 'min vals)))

(let ((checksum 0)
	(in (open "input.txt" :if-does-not-exist nil)))
	(when in
		(loop for line = (read-line in nil)
			while line do
				(let ((vals (convert-line line)))
				(setq checksum (+ checksum (min-max-difference vals)))))
	(close in)
	(princ checksum)
	(terpri)))
