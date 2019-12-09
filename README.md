# gosembler
a assembler simulator written in go
## Assembler Standard
### use memory address per bit
```
00000000: 0
00000001: 1
00000002: 0
00000003: 1
00000004: 1
00000005: 0
00000006: 0
00000007: 0
00000008: 0
00000009: 0
0000000a: 0
0000000b: 0
0000000c: 0
0000000d: 0
0000000e: 0
0000000f: 0
00000010: 0
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
