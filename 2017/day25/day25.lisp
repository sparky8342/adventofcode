#!/usr/bin/sbcl --script

(require "asdf")

(defstruct newstate
	value
	move
	state
)

(defstruct tapeslot
	value
	prev
	next
)

(defun get-input()
	(with-open-file (stream "input.txt")
		(let* ((startstate (subseq (read-line stream nil) 15 16))
			(steps (parse-integer (car (cdr (cdr (cdr (cdr (cdr (uiop:split-string (read-line stream nil) :separator " "))))))))))
				(read-line stream nil)
				(let ((zerohash (make-hash-table :test #'equal))
					(onehash (make-hash-table :test #'equal)))

					(loop for line = (read-line stream nil)
						while line do
							(let ((state (subseq line 9 10)))
								(read-line stream nil)
								(let* ((zerovalue (parse-integer (subseq (read-line stream nil) 22 23)))
									(zeromove (char (read-line stream nil) 27))
									(zerostate (subseq (read-line stream nil) 26 27))
									(ignoreline (read-line stream nil))
									(onevalue (parse-integer (subseq (read-line stream nil) 22 23)))
									(onemove (char (read-line stream nil) 27))
									(onestate (subseq (read-line stream nil) 26 27))
									(ignoreline (read-line stream nil))
									(zero (make-newstate
										:value zerovalue
										:move zeromove
										:state zerostate))
									(one (make-newstate
										:value onevalue
										:move onemove
										:state onestate)))
										(setf (gethash state zerohash) zero)
										(setf (gethash state onehash) one))))
					(values startstate steps zerohash onehash)))))
								

(multiple-value-bind (startstate steps zerohash onehash) (get-input)
	(let ((state startstate)
		(ts (make-tapeslot
			:value 0))
		(nextstate 0))

			(dotimes (i steps)
				(if (eq (tapeslot-value ts) 0)
					(setf nextstate (gethash state zerohash))
					(setf nextstate (gethash state onehash)))
				(setf (tapeslot-value ts) (newstate-value nextstate))
				(setq state (newstate-state nextstate))
				(when (eq (newstate-move nextstate) #\l)
					(when (not (tapeslot-prev ts))
						(let ((new-ts (make-tapeslot :value 0)))
							(setf (tapeslot-prev ts) new-ts)
							(setf (tapeslot-next new-ts) ts)))
					(setq ts (tapeslot-prev ts)))
				(when (eq (newstate-move nextstate) #\r)
					(when (not (tapeslot-next ts))
						(let ((new-ts (make-tapeslot :value 0)))
							(setf (tapeslot-next ts) new-ts)
							(setf (tapeslot-prev new-ts) ts)))
					(setq ts (tapeslot-next ts))))

			(loop while (not (eq (tapeslot-prev ts) nil)) do
					(setq ts (tapeslot-prev ts)))

			(let ((cnt (tapeslot-value ts)))
				(loop while (not (eq (tapeslot-next ts) nil)) do
					(setq ts (tapeslot-next ts))
					(setq cnt (+ cnt (tapeslot-value ts))))
				(princ cnt)
				(terpri))))
