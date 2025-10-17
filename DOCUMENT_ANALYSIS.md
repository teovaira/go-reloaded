# Go-Reloaded: My Analysis Document

**Student:** [Your Name]  
**Project:** go-reloaded  
**Date:** October 15, 2025  
**Week 1 Assignment** - Problem Analysis (No Coding!)

---

## Part 1: What is this project about? (In my own words)

### The Problem
I need to build a command-line tool in Go that reads a text file, applies some transformation rules to it, and writes the result to another file.

Think of it like this: someone gives me a messy text with special commands inside it, and my program needs to clean it up and apply those commands.

**How it works:**
```
go run . input.txt output.txt
```

The program reads `input.txt`, transforms the text according to rules, and saves the result in `output.txt`.

### Why is this challenging?
The text has special "markers" like `(hex)`, `(up)`, `(cap)` that tell my program what to do. The hard part is:
1. Understanding what each marker means
2. Applying them in the right order
3. Making sure the output is **exactly** correct (100% match)
4. Handling weird cases (what if something is missing? what if there are too many quotes?)

### What I'm learning:
- Reading and writing files in Go
- String manipulation (changing text)
- Number conversions (hex, binary → decimal)
- Thinking about edge cases
- Planning before coding!

---

## Part 2: The Transformation Rules

Here are all the rules my program needs to follow. I'm writing them in my own words with examples.

### Rule 1: Hexadecimal to Decimal
**What it does:** When I see `(hex)` after a word, that word is a hexadecimal number and I need to convert it to decimal.

**Example:**
```
Input:  "1E (hex) files"
Output: "30 files"
```

**Why:** 1E in hex = 30 in decimal.

**Another example:**
```
Input:  "FF (hex) is the max"
Output: "255 is the max"
```

**Things to think about:**
- What if the hex has lowercase letters? (should work)
- What if it's not a valid hex number?
- What if `(hex)` is at the beginning with no word before it?

---

### Rule 2: Binary to Decimal
**What it does:** When I see `(bin)` after a word, that word is a binary number and I need to convert it to decimal.

**Example:**
```
Input:  "10 (bin) years"
Output: "2 years"
```

**Why:** 10 in binary = 2 in decimal.

**Another example:**
```
Input:  "1010 (bin) equals"
Output: "10 equals"
```

**Things to think about:**
- What if it has numbers other than 0 and 1?
- What if `(bin)` has no word before it?

---

### Rule 3: Uppercase
**What it does:** `(up)` makes the previous word uppercase. If there's a number like `(up, 3)`, it makes the previous 3 words uppercase.

**Example:**
```
Input:  "ready, set, go (up)!"
Output: "ready, set, GO!"
```

**With a number:**
```
Input:  "this is exciting (up, 2)"
Output: "this is SO EXCITING"
```

**Things to think about:**
- What if I say `(up, 5)` but there are only 2 words before it?
- What if the word is already uppercase?

---

### Rule 4: Lowercase
**What it does:** `(low)` makes the previous word lowercase. With `(low, N)` it makes N previous words lowercase.

**Example:**
```
Input:  "STOP SHOUTING (low)"
Output: "STOP shouting"
```

**With a number:**
```
Input:  "WHY ARE WE YELLING (low, 4)"
Output: "why are we yelling"
```

---

### Rule 5: Capitalize
**What it does:** `(cap)` capitalizes the first letter of the previous word. With `(cap, N)` it capitalizes N previous words.

**Example:**
```
Input:  "welcome to the brooklyn bridge (cap)"
Output: "welcome to the brooklyn Bridge"
```

**With a number (like titles):**
```
Input:  "the new york times (cap, 4)"
Output: "The New York Times"
```

**Things to think about:**
- If a word is all CAPS, does it become "Word"? (probably yes)

---

### Rule 6: Article Correction (a → an)
**What it does:** Automatically changes "a" to "an" when the next word starts with a vowel (a, e, i, o, u) or the letter 'h'.

**Example:**
```
Input:  "a apple a day"
Output: "an apple a day"
```

**Another example:**
```
Input:  "There it was. A amazing rock!"
Output: "There it was. An amazing rock!"
```

**With 'h':**
```
Input:  "a hour passed"
Output: "an hour passed"
```

**Things to think about:**
- What about capital 'A'? (should become "An")
- What if "a" is at the end with no word after? (leave it as "a")
- Edge case: "a university" technically sounds like "you-niversity" so should stay "a", but the spec says all vowels and 'h', so maybe it becomes "an"? Need to check!

---

### Rule 7: Punctuation Spacing
**What it does:** Punctuation marks (. , ! ? : ;) should be right next to the word before them (no space), and have one space after them.

**Example:**
```
Input:  "Hello , world !"
Output: "Hello, world!"
```

**Special case - groups of punctuation:**
When there are groups like `...` or `!?` they stay together.

```
Input:  "Wait ... really ?"
Output: "Wait... really?"
```

**Another group example:**
```
Input:  "What !? No way"
Output: "What!? No way"
```

**Things to think about:**
- What counts as a "group"? (probably `...`, `!!`, `??`, `!?`, `?!`)
- What if there's punctuation at the very start?

---

### Rule 8: Quote Handling
**What it does:** Single quotes `'` come in pairs. Remove spaces between the quotes and the words inside them.

**Single word:**
```
Input:  "He said: ' hello '"
Output: "He said: 'hello'"
```

**Multiple words:**
```
Input:  "As he said: ' I am the best player '"
Output: "As he said: 'I am the best player'"
```

**Things to think about:**
- What if there's an odd number of quotes? (3 quotes total)
- What if quotes are empty? `' '`
- Do other rules (like punctuation) still apply inside quotes?

---

## Part 3: Architecture Comparison

Now I need to choose HOW to build this. There are two main approaches:

### Approach 1: The Pipeline ("Car Wash")

**How it works:**
The text goes through multiple "stations", each station does one job.

```
Input Text
   ↓
[Split into words/tokens]
   ↓
[Convert hex/bin numbers]
   ↓
[Apply uppercase/lowercase/capitalize]
   ↓
[Fix a/an articles]
   ↓
[Fix punctuation spacing]
   ↓
[Fix quotes]
   ↓
[Put everything back together]
   ↓
Output Text
```

**Good things:**
- ✅ **Simple to understand** - each function does ONE thing
- ✅ **Easy to test** - I can test each step separately
- ✅ **Easy to debug** - if something breaks, I know which step
- ✅ **Easy to change** - want to add a new rule? Just add a new step
- ✅ **Good for learning** - clear and organized

**Not so good things:**
- ❌ Goes through the text multiple times (slower)
- ❌ Uses more memory (keeps everything in memory)

**Example of a function:**
```
Function: ConvertHex
Input: List of words
Job: Find "(hex)", convert the word before it
Output: List of words with conversions done
```

---

### Approach 2: The FSM ("Conveyor Belt")

**How it works:**
Read the text character by character, and the program is in different "states". Depending on what I see, I change state and do actions.

**States might be:**
- Reading a word
- Reading a command like (hex)
- Reading punctuation
- Inside quotes
- Between words

**Good things:**
- ✅ **Fast** - only reads through text once
- ✅ **Memory efficient** - doesn't need to store everything
- ✅ **Professional** - this is how real parsers work

**Not so good things:**
- ❌ **Complex** - harder to understand
- ❌ **Harder to test** - states interact in complicated ways
- ❌ **Harder to debug** - when something breaks, it's hard to find why
- ❌ **Harder to change** - adding a new rule means understanding all states

---

### My Choice: Pipeline Architecture

**I'm choosing the Pipeline approach because:**

1. **I'm learning** - I want to understand the problem clearly, and Pipeline is more clear
2. **Easier to test** - I can write tests for each small function
3. **Easier to debug** - When something doesn't work, I can see which step is wrong
4. **Good enough** - The project files won't be huge, so speed isn't critical
5. **Team-friendly** - If I work with someone, we can each do different steps

**My plan:**
- Split text into pieces (tokens)
- Apply each rule one by one
- Put text back together
- Each rule is a separate function

**Trade-off I'm accepting:**
My solution won't be the fastest possible, but it will be correct and easy to understand. For this project, **correctness is more important than speed**.

---

## Part 4: Golden Test Set (Success Test Cases)

These are the tests I'll use to know if my program works correctly. I'm writing them in plain language (no code yet).

### Category A: Basic Audit Examples (from the project spec)

#### Test 1: Hex and Binary Conversion
**Input:**
```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```

**Expected Output:**
```
Simply add 66 and 2 and you will see the result is 68.
```

**Why this matters:** Tests both hex and binary conversion in one sentence.

---

#### Test 2: Article Correction
**Input:**
```
There is no greater agony than bearing a untold story inside you.
```

**Expected Output:**
```
There is no greater agony than bearing an untold story inside you.
```

**Why this matters:** Tests the a→an rule with vowels.

---

#### Test 3: Punctuation Spacing
**Input:**
```
Punctuation tests are ... kinda boring ,what do you think ?
```

**Expected Output:**
```
Punctuation tests are... kinda boring, what do you think?
```

**Why this matters:** Tests punctuation spacing AND the group `...` staying together.

---

#### Test 4: Multiple Rules Combined (The Big One)
**Input:**
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```

**Expected Output:**
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

**Why this matters:** This is the ultimate test! It has:
- `(cap)` - capitalize one word
- `(up)` - uppercase one word
- `(cap, 6)` - capitalize 6 words
- `(low, 3)` - lowercase 3 words
- Punctuation spacing
- All rules working together!

---

### Category B: Tricky Cases (My Own Test Cases)

#### Test 5: Quote with Multiple Words and Punctuation
**Input:**
```
She said: ' hello , how are you ? '
```

**Expected Output:**
```
She said: 'hello, how are you?'
```

**Why this is tricky:** 
- Quotes need to be fixed (remove spaces)
- Punctuation inside quotes also needs fixing
- Tests if rules work inside quotes

---

#### Test 6: Article Before 'H'
**Input:**
```
We waited a hour for a bus.
```

**Expected Output:**
```
We waited an hour for an bus.
```
OR (if the spec is smart about it)
```
We waited an hour for a bus.
```

**Why this is tricky:**
- The spec says 'h' → an
- But "a bus" is correct English (b is consonant)
- Need to check: does spec mean ALL 'h' or just silent 'h'?

---

#### Test 7: Count Exceeds Available Words
**Input:**
```
only two words (up, 10)
```

**Expected Output:** (Need to test what happens!)
```
ONLY TWO WORDS
```
OR maybe error?

**Why this is tricky:**
- I asked for 10 words but only 3 exist
- What should happen? Transform all 3? Error? Need to test!

---

#### Test 8: Multiple Hex in One Line
**Input:**
```
Values: 1E (hex) and FF (hex) and A (hex)
```

**Expected Output:**
```
Values: 30 and 255 and 10
```

**Why this is tricky:**
- Multiple conversions in one line
- Tests if my program handles multiple markers

---

#### Test 9: Uppercase After Already Uppercase
**Input:**
```
THIS WORD (up) again
```

**Expected Output:**
```
THIS WORD again
```

**Why this is tricky:**
- Word is already uppercase
- Should stay uppercase (idempotent)

---

#### Test 10: Empty or Edge Cases
**Input 10a (empty file):**
```
[empty file]
```
**Expected:** Program runs without crashing, creates empty output

**Input 10b (only spaces):**
```
     
```
**Expected:** Program handles it gracefully

**Input 10c (marker at beginning):**
```
(hex) word
```
**Expected:** Probably leaves it unchanged (no word before)

**Why this is tricky:**
- Edge cases that could break my program
- Need graceful handling

---

### Category C: The Big Complex Test

#### Test 11: Kitchen Sink (Everything at Once)
**Input:**
```
here (cap) is a interesting text with 1A (hex) items and 11 (bin) more , all in ' a epic document (cap, 2) ' ... what do you think (up, 4) ?
```

**Expected Output:**
```
Here is an interesting text with 26 items and 3 more, all in 'an Epic Document'... WHAT DO YOU THINK?
```

**What's being tested:**
1. `(cap)` - capitalize "here"
2. `a` → `an` before "interesting"
3. `1A (hex)` → `26`
4. `11 (bin)` → `3`
5. Punctuation spacing (comma, ellipsis, question mark)
6. Quotes with multiple words
7. `(cap, 2)` inside quotes
8. `a` → `an` before "epic" inside quotes
9. `(up, 4)` - uppercase 4 words at end
10. All rules working together!

**Why this matters:**
This is my ultimate test. If this works, I'm confident my program is correct!

---

## My Questions / Things I'm Unsure About

1. **Order of operations:** If I have hex conversion and article correction, which happens first?
   - My guess: Probably hex/bin first (they remove markers), then case changes, then articles, then punctuation, then quotes

2. **What if (up, 10) but only 3 words exist?**
   - My guess: Transform all 3 available words (graceful)

3. **Invalid hex like "XYZ (hex)":**
   - My guess: Leave it unchanged

4. **Odd number of quotes (3 quotes total):**
   - My guess: This is an error, but need to check what to do

5. **Article + h rule:**
   - Spec says all 'h' → an, but "a university" is correct English
   - Need to clarify: strict spec or smart about pronunciation?

---

## My Next Steps

1. ✅ This analysis document
2. ⏭️ Ask instructor about my questions above
3. ⏭️ Create test files for all my test cases
4. ⏭️ Plan my functions (what does each function do?)
5. ⏭️ Start coding next week!

---

**End of Analysis**

I'm ready to start coding once I understand the problem fully!