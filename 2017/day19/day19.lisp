#!/usr/bin/sbcl --script

(defstruct pos
	x
	y
)

(defstruct move
	dx
	dy
)

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
	(let ((start 0))
		(loop for x from 0 to 199 do
			(if (equal (aref grid 0 x) #\|)
				(setf start (make-pos
								:x x
								:y 0))))
		start))

(defun follow-path (grid pos)
	(let ((move (make-move :dx 0 :dy 1))
		(letters ())
		(steps 0))
			(loop while (not (equal (aref grid (pos-y pos) (pos-x pos)) #\Space)) do
				(setf (pos-x pos) (+ (pos-x pos) (move-dx move)))
				(setf (pos-y pos) (+ (pos-y pos) (move-dy move)))
				(setq steps (+ steps 1))
				(if (equal (aref grid (pos-y pos) (pos-x pos)) #\+)
					(if (= (move-dx move) 0)
						(progn
							(setf (move-dy move) 0)
							(if (equal (aref grid (pos-y pos) (+ (pos-x pos) 1)) #\Space)
								(setf (move-dx move) -1)
								(setf (move-dx move)  1)))
						(progn
							(setf (move-dx move) 0)
							(if (equal (aref grid (+ (pos-y pos) 1) (pos-x pos)) #\Space)
								(setf (move-dy move) -1)
								(setf (move-dy move)  1))))
					(progn
						(let ((chr (aref grid (pos-y pos) (pos-x pos))))
							(if (not (or (equal chr #\|) (equal chr #\-)))
								(push chr letters))))))
			(list (reverse letters) steps)))

(let* ((grid (get-input))
	(start (find-start grid))
	(result (follow-path grid start)))
		(princ (format nil "~{~a~}" (car result)))
		(terpri)
		(princ (car (cdr result)))
		(terpri))
