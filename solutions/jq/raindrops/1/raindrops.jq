
[if .number % 3 == 0 then "Pling" else empty end, 
if .number % 5 == 0 then "Plang" else empty end,
if .number % 7 == 0 then "Plong" else empty end,
if (.number % 3 != 0 and .number % 5 != 0 and .number % 7 != 0) then .number else empty end] | join("")

