let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/src/maze
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +69 ~/src/maze/model/mazeGen/maze-gen.go
badd +23 main.go
badd +6 ~/src/maze/model/mazeGen/mazeGenUtil.go
badd +25 model/mazeGen/test/mazeGenUtil_test.go
badd +25 term://~/src/maze//74494:/usr/bin/zsh
badd +38 model/mazeGen/mazeGen_test/mazeGenUtil_test.go
badd +9 tanker.md
badd +1 term://~/src/maze//25406:/usr/bin/zsh
argglobal
%argdel
tabnew +setlocal\ bufhidden=wipe
tabrewind
edit ~/src/maze/model/mazeGen/mazeGenUtil.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
wincmd _ | wincmd |
split
1wincmd k
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 105 + 105) / 211)
exe '2resize ' . ((&lines * 25 + 27) / 55)
exe 'vert 2resize ' . ((&columns * 105 + 105) / 211)
exe '3resize ' . ((&lines * 26 + 27) / 55)
exe 'vert 3resize ' . ((&columns * 105 + 105) / 211)
argglobal
balt model/mazeGen/test/mazeGenUtil_test.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 6 - ((5 * winheight(0) + 26) / 52)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 6
normal! 041|
wincmd w
argglobal
if bufexists(fnamemodify("model/mazeGen/mazeGen_test/mazeGenUtil_test.go", ":p")) | buffer model/mazeGen/mazeGen_test/mazeGenUtil_test.go | else | edit model/mazeGen/mazeGen_test/mazeGenUtil_test.go | endif
if &buftype ==# 'terminal'
  silent file model/mazeGen/mazeGen_test/mazeGenUtil_test.go
endif
balt term://~/src/maze//74494:/usr/bin/zsh
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 40 - ((7 * winheight(0) + 12) / 25)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 40
normal! 017|
wincmd w
argglobal
if bufexists(fnamemodify("term://~/src/maze//25406:/usr/bin/zsh", ":p")) | buffer term://~/src/maze//25406:/usr/bin/zsh | else | edit term://~/src/maze//25406:/usr/bin/zsh | endif
if &buftype ==# 'terminal'
  silent file term://~/src/maze//25406:/usr/bin/zsh
endif
balt model/mazeGen/mazeGen_test/mazeGenUtil_test.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 59 - ((25 * winheight(0) + 13) / 26)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 59
normal! 023|
wincmd w
exe 'vert 1resize ' . ((&columns * 105 + 105) / 211)
exe '2resize ' . ((&lines * 25 + 27) / 55)
exe 'vert 2resize ' . ((&columns * 105 + 105) / 211)
exe '3resize ' . ((&lines * 26 + 27) / 55)
exe 'vert 3resize ' . ((&columns * 105 + 105) / 211)
tabnext
edit tanker.md
argglobal
balt ~/src/maze/model/mazeGen/mazeGenUtil.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 9 - ((8 * winheight(0) + 26) / 52)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 9
normal! 0
tabnext 2
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
nohlsearch
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
