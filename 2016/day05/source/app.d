import std.stdio;
import std.uni;
import std.conv;
import std.digest.md;

void main()
{
    string id = "abbhdwsy";
    char[] code1;
    char[8] code2;
    int found2;
    int index = 1;

    while(code1.length < 8 || found2 < 8) {
        auto md5 = new MD5Digest();
        ubyte[] hash = md5.digest(id ~ to!string(index));
        string hashS = toHexString(hash[0..4]);

        if(hashS[0..5] == "00000") {
            if (code1.length < 8) {
                code1 ~= hashS[5..6];
                writeln("Part 1: ", index, " -- ", code1);
            }

            if (found2 >= 8) {
                continue;
            }
            try {
                int pos = to!int(hashS[5..6]);
                if (pos < 8 && code2[pos] == 0xFF) {
                    found2++;
                    code2[pos] = to!char(hashS[6..7]);
                    writeln("Part 2: ", index, " -- ", code2, " found: ", found2);
                }
            } catch (Exception e) {
                // Converting a character like "E" to int failed
            }
        }

        index++;
    }

    writeln("The first code is  ", toLower(code1));
    writeln("The second code is ", toLower(to!string(code2)));
}
