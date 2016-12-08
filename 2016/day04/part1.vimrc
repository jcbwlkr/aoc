set nowrap
set shortmess=at
set cmdheight=8
set nohlsearch

" Macro for part one (o)
let @o=':%normal @r:%g!/\(\[[a-z]\{5\}]\)\1/d:%s/.\{14\}//:1,$-s/$/ +/:%s/\n//:.!bc -l'
let @r='dd:newpkdd@cdd:bd!Pj'
let @c='f[F-rf[ikmak:s/-//g:s/./&\r/gddkVgg:sortgv:!uniq -c' . "'" . 'ak:1,.!sort -k1,1rn -k2,2' . "'" . 'ak:6,.dk2dwk2dwk2dwk2dwk2dw5J:s/ //gi[A]jddpkJxkJx0'
