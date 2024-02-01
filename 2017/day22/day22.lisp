#!/usr/bin/sbcl --script

(require "asdf")

(defstruct node
	x
	y
)

(defconstant UP 1)
(defconstant RIGHT 2)
(defconstant DOWN 3)
(defconstant LEFT 4)

(defun get-input()
	(let ((nodehash (make-hash-table :test #'equalp))
		(y 0)
		(len 0))
		(with-open-file (stream "input.txt")
			(loop for line = (read-line stream nil)
				while line do
					(if (= len 0)
						(setq len (- (length line) 1)))
					(loop for x from 0 to len do
						(if (char= (aref line x) #\#)
							(let ((node (make-node
								:x x
								:y y)))
								(setf (gethash node nodehash) 0))))
					(setq y (+ y 1))))
		(values nodehash (/ len 2))))

(defun turn(dir turndir)
	(if (= turndir LEFT)
		(progn
			(if (= dir UP)
				(return-from turn LEFT))
			(if (= dir RIGHT)
				(return-from turn UP))
			(if (= dir DOWN)
				(return-from turn RIGHT))
			(if (= dir LEFT)
				(return-from turn DOWN))))
	(if (= turndir RIGHT)
		(progn
			(if (= dir UP)
				(return-from turn RIGHT))
			(if (= dir RIGHT)
				(return-from turn DOWN))
			(if (= dir DOWN)
				(return-from turn LEFT))
			(if (= dir LEFT)
				(return-from turn UP)))))

(defun move(carrier dir)
	(if (= dir UP)
		(setf (node-y carrier) (- (node-y carrier) 1)))
	(if (= dir RIGHT)
		(setf (node-x carrier) (+ (node-x carrier) 1)))
	(if (= dir DOWN)
		(setf (node-y carrier) (+ (node-y carrier) 1)))
	(if (= dir LEFT)
		(setf (node-x carrier) (- (node-x carrier) 1))))

(defun part1()
	(let ((dir UP)
		(infections 0)
		(nodehash 0)
		(middle 0))
		(setf (values nodehash middle) (get-input))
		(let ((carrier (make-node
			 :x middle
			 :y middle)))
		(loop for i from 1 to 10000 do
			(if (gethash carrier nodehash)
				(progn
					(remhash carrier nodehash)
					(setq dir (turn dir RIGHT))
					(move carrier dir))
				(progn
					(setf (gethash (make-node :x (node-x carrier) :y (node-y carrier)) nodehash) 0)
					(setq infections (+ infections 1))
					(setq dir (turn dir LEFT))
					(move carrier dir)))))
		(princ infections)
		(terpri)))

(part1)
