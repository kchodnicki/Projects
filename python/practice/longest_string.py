def solution(p_string):
    prev_letter, max_letter = None, None
    curr_len = 0
    max_len = 0
    for curr_letter in p_string:
        print(prev_letter, curr_letter)
        if prev_letter == curr_letter:
            curr_len += 1
            if curr_len > max_len:
                max_len = curr_len
                max_letter = curr_letter
        else:
            curr_len = 1
        prev_letter = curr_letter
    return max_len, max_letter

if __name__ == '__main__':
    string = 'AAABBCDAABCCCCDDEEACDD'
    print(solution(string))
