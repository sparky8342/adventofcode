#!/usr/bin/sbcl --script

(require "asdf")

(defun get-data ()
	(let ((layers ()))
		(with-open-file (stream "input.txt")
			(loop for line = (read-line stream nil)
				while line do
					(let ((parts (uiop:split-string line :separator ":")))
						(let ((key (parse-integer (car parts)))
							(val (parse-integer (car (cdr parts)))))
								(push (list key val) layers)))))
		(nreverse layers)))

(defun find-severity (layers tm caught-exit)
	(let ((current-layer 0)
		(caught 0)
		(severity 0))
			(loop named loop1 for layer in layers do
				(let ((layer-no (first layer))
					(layer-length (car (cdr layer))))
					(when (< current-layer layer-no)
						(setq tm (+ tm (- layer-no current-layer)))
						(setq current-layer layer-no))
					
					; scanner reaches zero every ((length * 2) - 2) steps
					(when (= (mod tm (- (* layer-length 2) 2)) 0)
						(setq severity (+ severity (* current-layer layer-length)))
						(setq caught 1)
						(if (= caught-exit 1)
							(return-from loop1)))))
		(list severity caught)))

(let ((layers (get-data))
	(tm 0))
	(let ((rlst (find-severity layers tm 0)))
		(princ (first rlst))
		(terpri)
		(let ((caught 1))
			(loop while (= caught 1) do
				(setq tm (+ tm 1))
				(setq rlst (find-severity layers tm 1))
				(setq caught (car (cdr rlst))))
			(princ tm)
			(terpri))))
