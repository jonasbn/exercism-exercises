#include "reverse_string.h"
#include <string>

using namespace std;

namespace reverse_string {

string reverse_string(string s) {
    std::string o;

    for (int i=s.length()-1; i>=0; --i)
    {
        o.append(s.substr(i, 1));
    }

    return o;
}

}  // namespace reverse_string
