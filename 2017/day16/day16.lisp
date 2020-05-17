#!/usr/bin/sbcl --script

(require "asdf")

(defun get-input ()
	(uiop:split-string (car (uiop:read-file-lines "input.txt")) :separator ","))

(defun spin (lst arg)
	(let ((n (parse-integer arg)))
		(let ((start (subseq lst 0 (- (list-length lst) n)))
			(end (subseq lst (- (list-length lst) n))))
				(append end start))))

(defun exchange (lst arg)
	(let ((args (uiop:split-string arg :separator "/")))
		(let ((a (parse-integer (car args)))
			(b (parse-integer (car (cdr args)))))
				(rotatef (nth a lst) (nth b lst)))))
	
(defun partner (lst arg)
	(let ((args (uiop:split-string arg :separator "/")))
		(let ((a (read-from-string (car args)))
			(b (read-from-string (car (cdr args)))))
				(let ((a-pos (position a lst))
					(b-pos (position b lst)))
						(rotatef (nth a-pos lst) (nth b-pos lst))))))

(defun play-moves (lst moves)
	(loop for move in moves do
		;(print lst)
		(let ((move-type (subseq move 0 1))
			(arg (subseq move 1)))
				(if (equal move-type "s")
					(setf lst (spin lst arg)))
				(if (equal move-type "x")
					(exchange lst arg))
				(if (equal move-type "p")
					(partner lst arg))))
	lst)

(let ((p '(a b c d e f g h i j k l m n o p))
	(moves (get-input))
	(*print-case* :downcase))
		(setf p (play-moves p moves))
		(princ (format nil "~{~a~}" p))
		(terpri)
		(let ((i 1))
			(loop while (not (equal p '(a b c d e f g h i j k l m n o p))) do
				(setf p (play-moves p moves))
				(setq i (+ i 1)))
			(let ((n (mod 1000000000 i)))
				(loop repeat n do
					(setf p (play-moves p moves)))))
		(princ (format nil "~{~a~}" p))
		(terpri))
