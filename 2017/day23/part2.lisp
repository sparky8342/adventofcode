#!/usr/bin/sbcl --script

(defun is-prime(n)
	(if (= (rem n 2) 0)
		(return-from is-prime 0))

	(let ((i 3))
		(loop
			(if (= (rem n i) 0)
				(return-from is-prime 0))
			(if (> (* i i) n)
				(return-from is-prime 1))
			(setq i (+ i 2)))))

(let ((nonprimes 0))
	(loop for n from 109900 to 126900 by 17 do
		(if (= (is-prime n) 0)
			(setq nonprimes (+ nonprimes 1))))

	(princ nonprimes)
	(terpri))
