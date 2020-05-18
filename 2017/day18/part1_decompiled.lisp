#!/usr/bin/sbcl --script

(let ((a (- (expt 2 31) 1))
	(p 826)
	(sound 0))
		(loop repeat 127 do
			(setq p (* p 8505))
			(setq p (mod p a))
			(setq p (* p 129749))
			(setq p (+ p 12345))
			(setq p (mod p a))
			(setq sound (mod p 10000)))
		(princ sound)
		(terpri))
