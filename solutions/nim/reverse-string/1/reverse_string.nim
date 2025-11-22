import std/strutils

proc reverse*(s: string): string =

  var a = newSeq[char](s.len)
  
  var i = len(s)-1
  var j = 0

  while i >= 0:
    a[j] = s[i]
    dec(i)
    inc(j)

  return a.join("")
