use strict;
use warnings;
use Map qw(furthest_path);

use Test::More tests => 5;

is(furthest_path('^WNE$'),3);
is(furthest_path('^ENWWW(NEEE|SSE(EE|N))$'),10);
is(furthest_path('^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$'),18);
is(furthest_path('^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$'),23);
is(furthest_path('^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$'),31);
