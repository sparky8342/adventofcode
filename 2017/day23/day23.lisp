#!/usr/bin/sbcl --script

(require "asdf")

(defun get-input ()
	(let ((commands ()))
		(with-open-file (stream "input.txt")
			(loop for line = (read-line stream nil)
				while line do
					(let ((parts (map 'list (lambda (x) (read-from-string x)) (uiop:split-string line :separator " "))))
						(push parts commands))))
		(reverse commands)))

(defun get-reg (reg registers)
	(let ((val (gethash reg registers)))
		(when (not val)
			(setf (gethash reg registers) 0)
			(setq val 0))
	val))

(defun run-instructions (instructions)
	(let ((registers (make-hash-table :test #'equal))
		(pos 0)
		(mulcount 0))
		(loop while (and (>= pos 0) (< pos (list-length instructions))) do
			(let ((instruction (nth pos instructions)))
				(let ((op (first instruction))
					(arg1 (second instruction))
					(arg2 (third instruction)))
						; look up arg2 if it's a register
						(if (not (numberp arg2))
							(setq arg2 (get-reg arg2 registers)))
						; execute op
						(if (eq op 'set)
							(setf (gethash arg1 registers) arg2))
						(if (eq op 'sub)
							(setf (gethash arg1 registers) (- (get-reg arg1 registers) arg2)))
						(when (eq op 'mul)
							(setf (gethash arg1 registers) (* (get-reg arg1 registers) arg2))
							(setq mulcount (+ mulcount 1)))
						(when (eq op 'jnz)
							(if (not (numberp arg1))
								(setq arg1 (get-reg arg1 registers)))
							(if (= arg1 0)
								(setq pos (+ pos 1))
								(setq pos (+ pos arg2))))
						(if (not (eq op 'jnz))
                                                        (setq pos (+ pos 1))))))
		mulcount))

(let ((instructions (get-input)))
	; part 1
	(princ (run-instructions instructions))
	(terpri))
