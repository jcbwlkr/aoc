use strict;
use warnings;
use v5.10;

my $file = 'input.txt';
open(my $fh, '<:encoding(UTF-8)', $file) or die "Could not open file '$file' $!";

my $possible1 = 0; # 993
my $possible2 = 0; # 1849

# Array of arrays to hold the triangle pieces from each column for part 2
my @buff = [[], [], []];

my $i = 0;
while (my $row = <$fh>) {
    # Trim whitespace
    $row =~ s/^\s+|\s+$//g;

    # Split out columns
    my ($a, $b, $c) = split(/\s+/, $row);

    # Immediately check validity for part 1
    $possible1+=valid($a, $b, $c);

    # Track state for part 2 then validate it every 3 rows
    $buff[0][$i%3] = $a;
    $buff[1][$i%3] = $b;
    $buff[2][$i%3] = $c;

    $i++;
    if ($i%3 == 0) {
        $possible2+=valid($buff[0][0], $buff[0][1], $buff[0][2]);
        $possible2+=valid($buff[1][0], $buff[1][1], $buff[1][2]);
        $possible2+=valid($buff[2][0], $buff[2][1], $buff[2][2]);
    }
}

say "Possible triangles (1): ${possible1}";
say "Possible triangles (2): ${possible2}";

################################################################################
# Subroutines
################################################################################

sub valid {
    my ($a, $b, $c) = @_;

    if ($a+$b > $c && $a+$c > $b && $c+$b>$a) {
        return 1;
    }
    return 0;
}
