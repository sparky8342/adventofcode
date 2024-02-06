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

(defun dfs(comphash visited port total len results)
	(if (> total (gethash "maxstrength" results))
		(setf (gethash "maxstrength" results) total))

	(let ((maxlen (gethash "maxlen" results)))
		(when (> len maxlen)
			(setf (gethash "maxlen" results) len)
			(setf (gethash "longeststrength" results) total))
		(if (= len maxlen)
			(if (> total (gethash "longeststrength" results))
				(setf (gethash "longeststrength" results) total))))

	(loop for component in (gethash port comphash) do
		(when (not (gethash (component-id component) visited))
			(setf (gethash (component-id component) visited) 1)
			(dfs comphash visited (component-port2 component) (+ total (component-port1 component) (component-port2 component)) (+ len 1) results)
			(remhash (component-id component) visited))))

(let ((comphash (get-input)))
	(let ((visited (make-hash-table :test #'equal))
		(results (make-hash-table :test #'equal)))
		(setf (gethash "maxstrength" results) 0)
		(setf (gethash "maxlen" results) 0)
		(setf (gethash "longeststrength" results) 0)
		(dfs comphash visited 0 0 0 results)
		(princ (gethash "maxstrength" results))
		(terpri)
		(princ (gethash "longeststrength" results))
		(terpri)))
