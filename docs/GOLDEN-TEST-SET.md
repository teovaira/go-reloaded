# Golden Test Set (Success Test Cases)

### **Student:** Theodore Vairaktaris
### **Project:** go-reloaded  
### **Date:** October 19, 2025

---

## What is this document?

These are the tests I'll use to know if my program works correctly. I've organized them by difficulty - starting with the official examples from the project spec, then moving to tricky edge cases I thought of, and finally two big complex tests with everything combined.

**My test coverage:**
- 4 tests from the official audit examples
- 7 tricky cases that could break my program
- 2 huge paragraphs with multiple rules at once

---

## Category A: Official Audit Examples (from the project spec)

These test cases come from the official project specification.

### Test 1: Hex and Binary Conversion

**Input:**
```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```

**Expected Output:**
```
Simply add 66 and 2 and you will see the result is 68.
```

**Why this matters:** Tests both hex and binary conversion in one sentence. 42 in hex = 66 decimal, and 10 in binary = 2 decimal.

---

### Test 2: Article Correction

**Input:**
```
There is no greater agony than bearing a untold story inside you.
```

**Expected Output:**
```
There is no greater agony than bearing an untold story inside you.
```

**Why this matters:** Tests the a→an rule with vowels. The word "untold" starts with 'u', so "a" becomes "an".

---

### Test 3: Punctuation Spacing

**Input:**
```
Punctuation tests are ... kinda boring ,what do you think ?
```

**Expected Output:**
```
Punctuation tests are... kinda boring, what do you think?
```

**Why this matters:** Tests punctuation spacing AND the group `...` staying together. Commas and question marks need to be right next to words with one space after.

---

### Test 4: Multiple Rules Combined (Long Paragraph)

**Input:**
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```

**Expected Output:**
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

**Why this matters:** This is the big one from the audit! It has:
- `(cap)` - capitalize one word
- `(up)` - uppercase one word  
- `(cap, 6)` - capitalize 6 words
- `(low, 3)` - lowercase 3 words
- Punctuation spacing throughout
- All rules working together perfectly!

---

## Category B: Tricky Cases (My Own Test Cases)

These are edge cases I came up with that could break my program if I'm not careful.

### Test 5: Quotes with Multiple Words and Punctuation

**Input:**
```
She said: ' hello , how are you ? '
```

**Expected Output:**
```
She said: 'hello, how are you?'
```

**Why this is tricky:** 
- Quotes need to be fixed (remove spaces inside)
- Punctuation inside quotes also needs fixing
- Tests if my punctuation rules still work inside quotes

---

### Test 6: Article Correction with 'h'

**Input:**
```
a honor to meet a hero
```

**Expected Output:**
```
an honor to meet an hero
```

**Why this is tricky:**
- In proper English: "an honor" is correct (silent h) but "a hero" is also correct (pronounced h)
- BUT the spec says ALL words starting with 'h' should get "an"
- So even though "an hero" sounds wrong, I'm following the spec strictly
- This tests that I follow the rules even when they're different from normal English

---

### Test 7: Count Exceeds Available Words

**Input:**
```
only two words (up, 10)
```

**Expected Output:**
```
ONLY TWO WORDS
```

**Why this is tricky:**
- I asked for 10 words to be uppercase but there are only 3 words before the command
- My program should handle this gracefully - just uppercase all available words
- Tests that my program doesn't crash when the count is too high

---

### Test 8: Multiple Conversions in One Line

**Input:**
```
Values: 1E (hex) and FF (hex) and A (hex)
```

**Expected Output:**
```
Values: 30 and 255 and 10
```

**Why this is tricky:**
- Three separate hex conversions in one sentence
- Tests if my program can handle multiple markers of the same type
- Each conversion should work independently

---

### Test 9: Case Transformation on Already Transformed Text

**Input:**
```
THIS WORD (up) again
```

**Expected Output:**
```
THIS WORD again
```

**Why this is tricky:**
- The word is already uppercase
- Should stay uppercase (applying uppercase to uppercase = still uppercase)
- Tests that my commands are idempotent (safe to apply multiple times)

---

### Test 10: Empty and Edge Cases

**Input 10a (empty file):**
```
[empty file]
```
**Expected:** Program runs without crashing, creates empty output

**Input 10b (only spaces):**
```
     
```
**Expected:** Program handles it gracefully (probably outputs spaces or empty)

**Input 10c (command at beginning with no previous word):**
```
(hex) word
```
**Expected:** Leaves it unchanged because there's no word before `(hex)` to convert

**Why this is tricky:**
- Edge cases that could crash my program if I don't handle them
- Empty files shouldn't cause errors
- Commands without targets should be left alone

---

### Test 11: Invalid Commands

**Input:**
```
test (hex) word and (low, -1) text
```

**Expected Output:**
```
test (hex) word and (low, -1) text
```

**Why this is tricky:**
- "test" is not a valid hex number (has letters other than A-F)
- `(low, -1)` has an invalid negative count
- My program should handle these gracefully - just leave them unchanged
- Tests robustness - program shouldn't crash on bad input

---

## Category C: The Big Complex Test

### Test 12: Everything at Once (My Ultimate Test)

**Input:**
```
here (cap) is a interesting text with 1A (hex) items and 11 (bin) more , all in ' a epic document (cap, 2) ' ... what do you think (up, 4) ?
```

**Expected Output:**
```
Here is an interesting text with 26 items and 3 more, all in 'an Epic Document'... WHAT DO YOU THINK?
```

**What's being tested:**

1. `(cap)` - capitalize "here" → "Here"
2. `a` → `an` before "interesting" (vowel)
3. `1A (hex)` → `26` (hex conversion)
4. `11 (bin)` → `3` (binary conversion)
5. Punctuation spacing (comma, ellipsis, question mark)
6. Quotes with multiple words inside
7. `(cap, 2)` inside quotes: "epic document" → "Epic Document"
8. `a` → `an` before "epic" (vowel, inside quotes)
9. `(up, 4)` - uppercase last 4 words: "WHAT DO YOU THINK"
10. All rules working together perfectly!

**Why this matters:**

This is my ultimate integration test! If this works correctly, I'm confident that:
- All my transformation rules work
- Rules don't conflict with each other
- My order of operations is correct
- I can handle complex real-world input

If I can pass this test, my program is solid! 🎯

---

**End of Test Cases Analysis**