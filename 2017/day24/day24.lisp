#!/usr/bin/sbcl --script

(require "asdf")

(defstruct component
	id
	port1
	port2
)

(defun get-input()
	(let ((comphash (make-hash-table :test #'equal))
		(id 1))
		(with-open-file (stream "input.txt")
			(loop for line = (read-line stream nil)
				while line do
					(let* ((pins (uiop:split-string line :separator "/"))
						(port1 (parse-integer (nth 0 pins)))
						(port2 (parse-integer (nth 1 pins)))
						(comp1 (make-component :id id :port1 port1 :port2 port2))
						(comp2 (make-component :id id :port1 port2 :port2 port1)))

						(push comp1 (gethash port1 comphash))
						(push comp2 (gethash port2 comphash))

						(setq id (+ id 1)))))
		comphash))

(defun dfs(comphash visited port total)
	(let ((maxtotal total))
		(loop for component in (gethash port comphash) do
			(when (not (gethash (component-id component) visited))
				(setf (gethash (component-id component) visited) 1)
				(let ((newtotal (dfs comphash visited (component-port2 component) (+ total (component-port1 component) (component-port2 component)))))
					(if (> newtotal maxtotal)
						(setq maxtotal newtotal)))
				(remhash (component-id component) visited)))
		maxtotal))

(let ((comphash (get-input)))
	(princ comphash)
	(terpri)
	(let ((visited (make-hash-table :test #'equal)))
		(princ (dfs comphash visited 0 0))
		(terpri)))
