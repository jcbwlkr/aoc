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

    auto md5 = new MD5Digest();
    while(found2 < 8 || code1.length < 8) {
        ubyte[] hash = md5.digest(id ~ to!string(index));
        string hashS = toHexString(hash[0..4]);

        index++;
        if(hashS[0..5] != "00000") {
            continue;
        }

        if (code1.length < 8) {
            code1 ~= hashS[5..6];
            writeln("Part 1: ", index, " -- ", code1);
        }

        if (found2 >= 8) {
            continue;
        }

        int pos;
        try {
            pos = to!int(hashS[5..6]);
        } catch (Exception e) {
            // Converting a character like "E" to int failed
            continue;
        }

        if (pos < 8 && code2[pos] == 0xFF) {
            found2++;
            code2[pos] = to!char(hashS[6..7]);
            writeln("Part 2: ", index, " -- ", code2, " found: ", found2);
        }
    }

    writeln("The first code is  ", toLower(code1));
    writeln("The second code is ", toLower(to!string(code2)));
}
