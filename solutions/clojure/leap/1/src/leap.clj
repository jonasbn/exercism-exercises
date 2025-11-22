(ns leap)

(defn leap-year? [year] ;; <- argslist goes here
  (if (= (mod year 4) 0)
    (if (and (= (mod year 100) 0) (not= (mod year 400) 0)) false true)))