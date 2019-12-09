# gosembler
a assembler simulator written in go
## Assembler Standard

### Assembler Syntax

Line syntax: ```([label]:) (([instuction] ([parameter]) (,[parameter])) (;[comment])```

Label syntax: ```.||[character]([alphanumeric]...)```

Instructions: 
```
DB
MOV
ADD
SUB
INC
MUL
DIV
AND
OR
XOR
NOT
SHL
SHR
JMP
JC
JNC
JZ
JNZ
JA
JNA
CALL
RET
PUSH
POP
HLT
```

parameter: 
```
reg
regAddress
raw
addressValue
number
string
```

### opt code syntax

#### 8 bit

``[byte specifing instuction id]``

#### 16 bit

``[2 bytes specifing instuction id]``

#### 32 bit

``[byte specifing instuction type][byte specifing instuction groupe][byte specifing instuction code][byte specifing instuction format]``

example:
```23D1```

instuction type = ``2``
instuction groupe = ``3``
instuction code = ``D``
instuction format = ``1``

#### 64 bit

``[2 bytes specifing instuction type][2 bytes specifing instuction groupe][2 bytes specifing instuction code][2 bytes specifing instuction format]``

example:
```02030D01```

instuction type = ``02``
instuction groupe = ``03``
instuction code = ``0D``
instuction format = ``01``
