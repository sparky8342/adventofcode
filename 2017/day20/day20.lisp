#!/usr/bin/sbcl --script

(require "asdf")

(defstruct particle
	id
	x
	y
	z
	vx
	vy
	vz
	ax
	ay
	az
)

(defun get-input ()
	(let ((particles ())
		(id 0))
			(with-open-file (stream "input.txt")
				(loop for line = (read-line stream nil)
					while line do
						(setq line (remove-if (lambda (ch) (find ch "pva=<> ")) line))
						(let ((parts (map 'list (lambda (x) (parse-integer x)) (uiop:split-string line :separator ","))))
							(let ((pt (make-particle
								:id id
								:x (first parts)
								:y (second parts)
								:z (third parts)
								:vx (fourth parts)
								:vy (fifth parts)
								:vz (sixth parts)
								:ax (seventh parts)
								:ay (eighth parts)
								:az (ninth parts))))
									(push pt particles)))
						(setq id (+ id 1))))
		particles))

(defun move-particle (pt)
	(setf (particle-vx pt) (+ (particle-vx pt) (particle-ax pt)))
	(setf (particle-vy pt) (+ (particle-vy pt) (particle-ay pt)))
	(setf (particle-vz pt) (+ (particle-vz pt) (particle-az pt)))
	(setf (particle-x pt) (+ (particle-x pt) (particle-vx pt)))
	(setf (particle-y pt) (+ (particle-y pt) (particle-vy pt)))
	(setf (particle-z pt) (+ (particle-z pt) (particle-vz pt))))

(defun distance (pt)
	(+ (abs (particle-x pt)) (abs (particle-y pt)) (abs (particle-z pt))))

(defun move-particles (particles)
	(loop for pt in particles do
		(move-particle pt)))

(defun find-closest-id (particles)
	(let ((mindistance (distance (first particles)))
		(bestid (particle-id (first particles))))
		(loop for pt in particles do
			(let ((dist (distance pt)))
				(when (< dist mindistance)
					(setq mindistance dist)
					(setq bestid (particle-id pt)))))
	bestid))

(let ((particles (get-input)))
	(loop repeat 10000 do
		(move-particles particles))
	(princ (find-closest-id particles))
	(terpri))
