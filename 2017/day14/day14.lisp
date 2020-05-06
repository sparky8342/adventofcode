#!/usr/bin/sbcl --script

(require "asdf")

(defun split-chars-ascii (lst)
	(let ((nums (map 'list (lambda (x) (char-code x)) (coerce lst 'list))))
		(setq nums (append nums '(17 31 73 47 23)))
		nums))

(defun create-nums ()
	(let ((lst (create-lst)))
		(let ((lst (circular lst)))
			lst)))

(defun create-lst ()
	(let ((nums ()))
		(dotimes (i 256 (nreverse nums))
			(push i nums))))

(defun circular (items)
	; from lisp cookbook
	"Modifies the last cdr of list ITEMS, returning a circular list"
	(setf (cdr (last items)) items)
	items)

(defun knot (lst lengths rounds)
	(let ((pos 0)
		(skip 0))
			(loop repeat rounds do
				(loop for len in lengths do
					(let ((part (subseq lst pos (+ pos len))))
						(let ((part (nreverse part)))
							(loop for i from pos to (+ pos (- len 1)) do
								(setf (nth i lst) (car part))
								(setq part (cdr part)))))
					(setq pos (mod (+ pos len skip) 256))
					(setq skip (+ skip 1))))
			lst))

(defun knot-hash (str)
	(let ((lengths (split-chars-ascii str))
		(nums (create-nums))
		(hash))
			(setq nums (knot nums lengths 64))
			(loop for i from 0 to 15 do
				(let ((x 0))
					(loop for j from 0 to 15 do
						(setq x (logxor x (nth (+ (* i 16) j) nums))))
					(let ((hex (string-downcase (write-to-string x :base 16))))
						(if (= 1 (length hex))
							(setq hash (concatenate 'string hash "0")))
						(setq hash (concatenate 'string hash hex)))))
			hash))
		
(defun to-bits (hex-str)
	(let ((bitstr))
		(loop for c across hex-str do
			(let ((n))
				(case c
					(#\a (setq n 10))
					(#\b (setq n 11))
					(#\c (setq n 12))
					(#\d (setq n 13))
					(#\e (setq n 14))
					(#\f (setq n 15))
					(otherwise (setq n (digit-char-p c))))
				(setq n (format nil "~B" n))
				(if (< (length n) 4)
					(loop for j from 1 to (- 4 (length n)) do
						(setq bitstr (concatenate 'string bitstr "0"))))
				(setq bitstr (concatenate 'string bitstr n))))
		bitstr))

(defun count-ones (bit-str)
  	(let ((amount 0))
		(loop for c across bit-str do
			(if (eq c #\1)
				(setq amount (+ amount 1))))
		amount))

(defun count-groups (grid)
	(let ((groups 0))
		(dotimes (y 128)
			(dotimes (x 128)
				(when (= (aref grid x y) 1)
					(remove-group grid x y)
					(setq groups (+ groups 1)))))
		groups))

(defun remove-group (grid x y)
	(let ((queue ()))
		(push (list x y) queue)
		(loop while (> (list-length queue) 0) do
			(let ((square (car queue)))
				(setq queue (cdr queue))
				(let ((sqx (car square))
					(sqy (car (cdr square))))
						(when (= (aref grid sqx sqy) 1)
							(setf (aref grid sqx sqy) 0)
							(if (> sqx 0)
								(push (list (- sqx 1) sqy) queue))
							(if (< sqx 127)
								(push (list (+ sqx 1) sqy) queue))
							(if (> sqy 0)
								(push (list sqx (- sqy 1)) queue))
							(if (< sqy 127)
								(push (list sqx (+ sqy 1)) queue))
							(setq queue (nreverse queue))))))))

(defun print-grid (grid)
	(dotimes (i 128)
		(dotimes (j 128)
			(princ (aref grid i j)))
				(terpri)))

(let ((input "amgozmfv"))
	(let ((squares 0)
		(grid (make-array '(128 128))))
			(loop for i from 0 to 127 do
				(let ((key (concatenate 'string input "-" (write-to-string i))))
					(let ((bit-str (to-bits (knot-hash key))))
						(setq squares (+ squares (count-ones bit-str)))
						(loop for j from 0 to 127 do
							(setf (aref grid i j) (digit-char-p (char bit-str j)))))))

			; part 1
			(princ squares)
			(terpri)
			; part 2
			(princ (count-groups grid))
			(terpri)))
