#!/usr/bin/sbcl --script

(require "asdf")

(defstruct program
	instructions
	registers
	pos
	inbuffer
	sendcount
)

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

(defun run-instructions (prg otherprg)
	(let ((registers (program-registers prg))
		(instructions (program-instructions prg))
		(pos (program-pos prg))
		(done 0))
		(loop while (and (>= pos 0) (< pos (list-length instructions)) (= done 0)) do
			(let ((instruction (nth pos instructions)))
				(let ((op (first instruction))
					(arg1 (second instruction))
					(arg2 (third instruction)))
						; look up arg2 if it's a register
						(if (not (numberp arg2))
							(setq arg2 (get-reg arg2 registers)))
						; execute op
						(when (eq op 'snd)
							(setf (program-inbuffer otherprg) (append (program-inbuffer otherprg) (list (get-reg arg1 registers))))
							(setf (program-sendcount prg) (+ (program-sendcount prg) 1)))
						(if (eq op 'set)
							(setf (gethash arg1 registers) arg2))
						(if (eq op 'add)
							(setf (gethash arg1 registers) (+ (get-reg arg1 registers) arg2)))
						(if (eq op 'mul)
							(setf (gethash arg1 registers) (* (get-reg arg1 registers) arg2)))
						(if (eq op 'mod)
							(setf (gethash arg1 registers) (mod (get-reg arg1 registers) arg2)))
						(if (eq op 'rcv)
							(if (= (list-length (program-inbuffer prg)) 0)
								(setq done 1)
								(progn
									(setf (gethash arg1 registers) (car (program-inbuffer prg)))
									(setf (program-inbuffer prg) (cdr (program-inbuffer prg))))))
						(when (eq op 'jgz)
							(if (not (numberp arg1))
								(setq arg1 (get-reg arg1 registers)))
							(if (> arg1 0)
								(setq pos (+ pos arg2))
								(setq pos (+ pos 1))))
						(if (and (not (eq op 'jgz)) (= done 0))
							(setq pos (+ pos 1))))))
		(setf (program-pos prg) pos)))

(let ((ins (get-input)))
	; part 1
	(let (
		(prg (make-program
			:instructions ins
			:registers (make-hash-table :test #'equal)
			:pos 0
			:inbuffer ()
			:sendcount 0))
		(prg2 (make-program
			:instructions ins
			:registers (make-hash-table :test #'equal)
			:pos 0
			:inbuffer ()
			:sendcount 0)))

			(run-instructions prg prg2)
			(princ (car (last (program-inbuffer prg2))))
			(terpri))
	; part 2
	(let (
		(prg0 (make-program
			:instructions ins
			:registers (make-hash-table :test #'equal)
			:pos 0
			:inbuffer ()
			:sendcount 0))
		(prg1 (make-program
			:instructions ins
			:registers (make-hash-table :test #'equal)
			:pos 0
			:inbuffer ()
			:sendcount 0)))

			(setf (gethash 'p (program-registers prg1)) 1)

			(run-instructions prg0 prg1)
			(run-instructions prg1 prg0)
			(loop while (or (> (list-length (program-inbuffer prg0)) 0) (> (list-length (program-inbuffer prg1)) 0)) do
				(run-instructions prg0 prg1)
				(run-instructions prg1 prg0))
			(princ (program-sendcount prg1))
			(terpri)))
