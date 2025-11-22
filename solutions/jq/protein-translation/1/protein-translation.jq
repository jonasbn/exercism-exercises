
def handle:
	[ label $out | foreach .[] as $codon (null; if $codon == "STOP" then break $out elif $codon == "Invalid codon" then ($codon|halt_error) else $codon end)]
;

def translate:
	if . == "AUG" then "Methionine"
	elif . == "UUU" or . == "UUC" then "Phenylalanine"
	elif . == "UUA" or . == "UUG" then "Leucine"
	elif . == "UCU" or . == "UCC" or . == "UCA" or . == "UCG" then "Serine"
	elif . == "UAU" or . == "UAC" then "Tyrosine"
	elif . == "UGU" or . == "UGC" then "Cysteine"
	elif . == "UGG" then "Tryptophan"
	elif . == "" then null
	elif . == "UAA" then "STOP"
	elif . == "UAG" then "STOP"
	elif . == "UGA" then "STOP"
	else "Invalid codon"
	end
;

def parse:
	. | [ foreach scan("\\w{1,3}", "g") as $codon (null; $codon ) | translate ]
;

.strand | parse | handle

