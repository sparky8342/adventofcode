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

(defun get-input ()
	(let ((nodehash (make-hash-table :test #'equalp))
		(y 0))
		(with-open-file (stream "input.txt")
			(loop for line = (read-line stream nil)
				while line do
					(loop for x from 0 to (- (length line) 1) do
						(if (char= (aref line x) #\#)
							(let ((node (make-node
								:x x
								:y y)))
								(setf (gethash node nodehash) 0))))
					(setq y (+ y 1))))
		nodehash))

(let ((nodehash (get-input))
	(carrier (make-node
		 :x 12
		 :y 12))
	(dir UP)
	(infections 0))
	(loop for i from 1 to 10000 do
		(if (gethash carrier nodehash)
			(progn
				(remhash carrier nodehash)
				(let ((newdir 0))
					(when (= dir UP)
						(setf newdir RIGHT)
						(setf (node-x carrier) (+ (node-x carrier) 1)))
					(when (= dir RIGHT)
						(setf newdir DOWN)
						(setf (node-y carrier) (+ (node-y carrier) 1)))
					(when (= dir DOWN)
						(setf newdir LEFT)
						(setf (node-x carrier) (- (node-x carrier) 1)))
					(when (= dir LEFT)
					 	(setf newdir UP)
						(setf (node-y carrier) (- (node-y carrier) 1)))
					(setq dir newdir)))
			(progn
				(setf (gethash (make-node :x (node-x carrier) :y (node-y carrier)) nodehash) 0)
				(setq infections (+ infections 1))
				(let ((newdir 0))
					(when (= dir UP)
						(setf newdir LEFT)
						(setf (node-x carrier) (- (node-x carrier) 1)))
					(when (= dir RIGHT)
						(setf newdir UP)
						(setf (node-y carrier) (- (node-y carrier) 1)))
					(when (= dir DOWN)
						(setf newdir RIGHT)
						(setf (node-x carrier) (+ (node-x carrier) 1)))
					(when (= dir LEFT)
					 	(setf newdir DOWN)
						(setf (node-y carrier) (+ (node-y carrier) 1)))
					(setq dir newdir)))))

	(princ infections)
	(terpri))
