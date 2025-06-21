class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        int alpha[26] = {0};
        for (int i = 0; i < magazine.size(); i++) {
            alpha[magazine[i] - 'a']++;
        }

        //for (int i = 0; i < 26; i++) {
        //    cout << i << " " << alpha[i] << endl; 
        //}

        for (int i = 0; i < ransomNote.size(); i++) {
            int idx = ransomNote[i] - 'a';
            if (alpha[idx] == 0) {
                return false;
            }

            alpha[idx]--;
        }

        return true;
    }
};