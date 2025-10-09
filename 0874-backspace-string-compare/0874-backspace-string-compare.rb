# @param {String} s
# @param {String} t
# @return {Boolean}
def backspace_compare(s, t)
    return backspace(s) == backspace(t)
end

def backspace(s)
    bs = 0
    s2 = ""
    s.chars.reverse_each do |c|
        if c == '#'
            bs += 1
            next
        end
        
        if bs > 0
            bs -= 1
            next
        end

        s2 << c
    end

    s2.reverse
end