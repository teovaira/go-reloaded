# Go-Reloaded: Project Analysis (Pipeline Model)

### **Student:** Theodore Vairaktaris
### **Project:** go-reloaded (Pipeline Implementation)
### **Date:** October 18, 2025  
### **Week 1 Assignment** - Problem Analysis (No Coding!)

---

## Part 1: Problem Description

### The Problem

I need to build a command-line tool in Go that reads a text file, applies some commands and transformation rules to it, and writes the modified result to another file.

**How it works:**
```
go run . input.txt output.txt
```

The program reads `input.txt`, transforms the text according to rules, and saves the result in `output.txt`.

---

## Part 2: Transformation Rules

Below are all the rules my program needs to follow. 

### Rule 1: Hexadecimal to Decimal

**What it does:** When I see the command `(hex)` after a word, that word is a hexadecimal number and I need to convert it to decimal.

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

**Matters to think about:**
- What if the hex has lowercase letters? (should work)
- What if it's not a valid hex number? (probably leave it as it is or error handling?)
- What if `(hex)` is at the beginning with no word before it? (leave it as it is or erase command?)

---

### Rule 2: Binary to Decimal

**What it does:** When I see the command `(bin)` after a word, that word is a binary number and I need to convert it to decimal.

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

**Matters to think about:**
- What if it has numbers other than 0 and 1? (probably leave it as it is or error handling?)
- What if `(bin)` has no word before it? (leave it as it is or erase command?)

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
Output: "this IS EXCITING"
```

**Matters to think about:**
- What if I say `(up, 5)` but there are only 2 words before it? (alter all of them or error handling?)
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

**Matters to think about:**
- What if I say `(low, 5)` but there are only 2 words before it? (alter all of them or error handling?)
- What if the word is already lowercase?

---

### Rule 5: Capitalize

**What it does:** `(cap)` capitalizes the first letter of the previous word. With `(cap, N)` it capitalizes N previous words.

**Example:**
```
Input:  "welcome to the brooklyn bridge (cap)"
Output: "welcome to the brooklyn Bridge"
```

**With a number:**
```
Input:  "the new york times (cap, 4)"
Output: "The New York Times"
```

**Matters to think about:**
- If a word is already uppercase, does it become "Word"? (probably yes)
- If there are less words before the command than the number suggests? (transform all of them or error handling)

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

**Matters to think about:**
- What about capital 'A'? (should become "An")
- What if "a" is at the end with no word after? (leave it as "a")
- Edge case: "a university" according to english grammar rules is correct, but the spec says all vowels and 'h', so maybe it becomes "an"? Need to verify!

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

**Example:**
```
Input:  "Wait ... really ?"
Output: "Wait... really?"
```

**Another group example:**
```
Input:  "What !? No way"
Output: "What!? No way"
```

**Matters to think about:**
- What counts as a "group"? (probably `...`, `!!`, `??`, `!?`, `?!`) Need to verify!
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

**Matters to think about:**
- What if there's an odd number of quotes? (3 quotes total)
- What if quotes are empty? `' '`
- Do other rules (like punctuation) still apply inside quotes? (most probably yes!)

---

## Part 3: Architecture Comparison

Now I need to choose HOW to build this. There are two models to choose between:

### Model 1: The Pipeline ("Car Wash")

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

### Model 2: The FSM ("Conveyor Belt")

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

**I'm choosing the Pipeline model because:**

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

## Things I'm Still Figuring Out

There are some things I'm not 100% sure about yet. I'll need to figure these out when I start coding:

1. **What order should I apply the rules?**
   - My thinking: probably hex/bin first (since they remove the markers), then do case changes (up/low/cap), then fix articles (a/an), then punctuation, and finally quotes. But I might need to adjust this if things don't work right.

2. **What if someone asks for more words than exist?**
   - Like `(up, 10)` but there are only 3 words before it
   - My plan: just transform all the available words and don't crash
   - Seems like the safest approach

3. **What if the hex/bin conversion gets invalid input?**
   - Like "XYZ (hex)" where XYZ isn't valid hex
   - My plan: probably just leave it as-is and don't convert
   - Better to skip it than to crash

4. **What if there are an odd number of quotes?**
   - Like 3 quotes total in the file
   - I'm not sure yet... maybe treat the last one as a regular character? Or maybe it's an error?
   - Need to test this edge case

5. **The 'h' rule for articles - should I follow it exactly?**
   - The spec says all words starting with 'h' get "an" 
   - But in real English, "a hero" is correct (pronounced h) while "an honor" is correct (silent h)
   - I think I'll follow the spec strictly and make ALL 'h' words use "an", even if it sounds weird
   - Rules are rules!

---

**End of Project Analysis**