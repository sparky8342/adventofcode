(require "asdf")

(let* ((in (uiop:read-file-string "input.txt"))
	(digits (coerce in 'list))
	(sum 0))
	
  	; copy the first digit to the end
	(setf digits (cons (car digits) digits))

	(loop for i from 0 to (- (length digits) 2)
        	do (if (eq (nth i digits) (nth (+ i 1) digits))
		     (setq sum (+ sum (digit-char-p (nth i digits))))))

	(princ sum))
