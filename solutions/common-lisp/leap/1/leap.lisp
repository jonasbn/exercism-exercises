(defpackage :leap
  (:use :cl)
  (:export :leap-year-p))
(in-package :leap)

(defun leap-year-p (year)
  (if (= (mod year 4) 0) (if (and (= (mod year 100) 0) (not (= (mod year 400) 0))) NIL T) NIL)
)
