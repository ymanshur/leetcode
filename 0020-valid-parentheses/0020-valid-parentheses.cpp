class Solution {
public:
    bool isValid(string s) {
        // Iterate each character on s
        // If it is open bracket, push it into a stack,
        // Else pop an open braket from stak, check if the close bracket match

        stack<char> openbs; // should contains open brackets

        for (int i = 0; i < s.size(); i++) {
            if (s[i] == '(' ||
                s[i] == '{' ||
                s[i] == '[') {
                    openbs.push(s[i]);
                    continue;
                }
            
            if (openbs.empty()) {
                return false;
            }
            
            char openb = openbs.top();

            switch (openb) {
                case '(':
                    if (s[i] != ')') {
                        return false;
                    }
                    break;
                case '{':
                    if (s[i] != '}') {
                        return false;
                    }
                    break;
                case '[':
                    if (s[i] != ']') {
                        return false;
                    }
                    break;
            }

            openbs.pop();
        }

        return openbs.empty();
    }
};