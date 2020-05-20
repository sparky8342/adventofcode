#!/usr/bin/sbcl --script

(defun get-input ()
	(let ((grid (make-array '(201 201)))
		(y 0))
		(with-open-file (stream "input.txt")
			(loop for line = (read-line stream nil)
				while line do
					(loop for x from 0 to 200 do
						(setf (aref grid y x) (aref line x)))
					(setq y (+ y 1))))
		(loop for x from 0 to 200 do
			(setf (aref grid 200 x) #\Space))
		grid))

(defun find-start (grid)
	(let ((x 0))
		(loop for i from 0 to 199 do
			(if (equal (aref grid 0 i) #\|)
				(setq x i)))
		(list x 0)))

(defun follow-path (grid x y)
	(let ((dx 0)
		(dy 1)
		(letters ())
		(steps 0))
			(loop while (not (equal (aref grid y x) #\Space)) do
				(setq x (+ x dx))
				(setq y (+ y dy))
				(setq steps (+ steps 1))
				(if (equal (aref grid y x) #\+)
					(if (= dx 0)
						(progn
							(setq dy 0)
							(if (equal (aref grid y (+ x 1)) #\Space)
								(setq dx -1)
								(setq dx 1)))
						(progn
							(setf dx 0)
							(if (equal (aref grid (+ y 1) x) #\Space)
								(setq dy -1)
								(setq dy 1))))
					(progn
						(let ((chr (aref grid y x)))
							(if (not (or (equal chr #\|) (equal chr #\-)))
								(push chr letters))))))
			(list (reverse letters) steps)))

(let* ((grid (get-input))
	(start (find-start grid))
	(x (car start))
	(y (car (cdr start)))
	(result (follow-path grid x y)))
		(princ (format nil "~{~a~}" (car result)))
		(terpri)
		(princ (car (cdr result)))
		(terpri))
