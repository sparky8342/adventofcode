#!/usr/bin/sbcl --script

(require "asdf")

(defun get-data ()
	(with-open-file (stream "input.txt")
		(loop for line = (read-line stream nil)
			while line
				collect line)))

(defun get-reg (registers reg)
	(let ((val (gethash reg registers)))
		(if val val 0)))

(defun get-max (registers)
	(let ((max 0))
		(loop for val being each hash-value of registers do
			(if (> val max)
				(setq max val)))
		max))

(defun run-commands (data)
	(let ((registers (make-hash-table :test #'equal))
		(running-max 0))
		(loop for line in data do
			(let ((instruction (uiop:split-string line :separator " ")))
				(let ((set-reg (read-from-string (nth 0 instruction)))
					(inc-dec (read-from-string (nth 1 instruction)))
					(amount (parse-integer (nth 2 instruction)))
					(check-reg (read-from-string (nth 4 instruction)))
					(comparison (read-from-string (nth 5 instruction)))
					(num (parse-integer (nth 6 instruction)))
					(result 0))
						(case comparison
							('>	 (if (>  (get-reg registers check-reg) num) (setq result 1)))
							('<	 (if (<  (get-reg registers check-reg) num) (setq result 1)))
							('>= (if (>= (get-reg registers check-reg) num) (setq result 1)))
							('<= (if (<= (get-reg registers check-reg) num) (setq result 1)))
							('== (if (=  (get-reg registers check-reg) num) (setq result 1)))
							('!= (if (/= (get-reg registers check-reg) num) (setq result 1))))
						(when (= result 1)
							(let ((newval 0))
								(case inc-dec
									('inc (setq newval (+ (get-reg registers set-reg) amount)))
									('dec (setq newval (- (get-reg registers set-reg) amount))))
								(setf (gethash set-reg registers) newval)
								(if (> newval running-max)
									(setq running-max newval)))))))
		(list (get-max registers) running-max)))
	
(let ((data (get-data)))
	(let ((max (run-commands data)))
		(princ max)
		(terpri)))
