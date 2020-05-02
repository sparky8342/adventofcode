#!/usr/bin/sbcl --script

(require "asdf")

(defstruct node
	name
	weight
	totalweight
	parent
	children
)

(defun read-input ()
	(let ((nodehash (make-hash-table :test #'equal))
		(in (open "input.txt" :if-does-not-exist nil)))
		(when in
			(loop for line = (read-line in nil)
				while line do
					(let ((data (map 'list (lambda (x) (read-from-string x)) (uiop:split-string line :separator " "))))
						; make node and store in hash
						(let ((name (car data))
							(weight (car (car (cdr data))))
							(children (cdr (cdr (cdr data)))))
								(let ((node (make-node
									:name name
									:weight weight
									:parent nil
									:children children)))
										(setf (gethash name nodehash) node))))))
		(close in)

		; set the parent and totalweight of each node, and attach the child nodes directly
		(loop for key being the hash-keys of nodehash collect key do
			(let ((node (gethash key nodehash))
				(childnodes ()))
				(loop for child in (node-children node) do
					(let ((childnode (gethash child nodehash)))
						(setf (node-parent childnode) (node-name node))
						(setq childnodes (append childnodes (list childnode)))))
				(setf (node-children node) childnodes)))

		(let ((parent (find-parent nodehash)))
		  	(set-total-weights parent)
			parent)))

(defun find-parent (nodehash)
	; get any key from the hash
	(let ((name (first (loop for key being the hash-keys of nodehash collect key))))
		; walk up to the top parent
		(let ((node (gethash name nodehash)))
			(loop while (node-parent node) do
				(setq node (gethash (node-parent node) nodehash)))
		node)))

(defun set-total-weights (node)
	(let ((weight (node-weight node)))
		(loop for childnode in (node-children node) do
			(setq weight (+ weight (set-total-weights childnode))))
		(setf (node-totalweight node) weight)
		weight))

(defun find-different-node (node)
	; store a hash of weights and how many seen
	(let ((childweights (make-hash-table :test #'equal)))
		(loop for childnode in (node-children node) do
			(let ((w (node-totalweight childnode)))
				(if (gethash w childweights)
					(setf (gethash w childweights) (+ (gethash w childweights) 1))
					(setf (gethash w childweights) 1))))

		(if (= (hash-table-count childweights) 1)
			; only one hash key means all weights were the same
			0
			; otherwise find the unique node
			(progn
				(let ((single nil)
					(multiple nil)
					(uniquenode nil))
						(maphash (lambda (key value)
							(when (= value 1)
								(setq single key)
								(loop for childnode in (node-children node) do
									(if (= (node-totalweight childnode) key)
										(setq uniquenode childnode))))
							(when (> value 1)
								(setq multiple key)))
							childweights)
						; recurse on the unique node
						; if it returns 0, then this is the level where
						; a node has to be changed, so print the needed weight
						(let ((diff (find-different-node uniquenode)))
							(when (= diff 0)
								(princ (+ (node-weight uniquenode) (- multiple single)))
								(terpri))))
				1))))

(let* ((parent (read-input))
	(*print-case* :downcase))
		(princ (node-name parent))
		(terpri)
		(find-different-node parent))
