# Go-Reloaded: Project Analysis (Pipeline Model)

### **Student:** Theodore Vairaktaris
### **Project:** go-reloaded (Pipeline Implementation)
### **Date:** October 19, 2025  
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

**With a number:**
```
Input:  "the new york times (cap, 4)"
Output: "The New York Times"
```

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

## Edge Cases & Open Questions

Below are things I'm still figuring out. These questions apply no matter which architecture I choose - they're about the problem itself, not how I implement it.

### When Commands Ask for More Than Available

**The situation:** What if someone writes `(up, 10)` but there are only 3 words before it?

**My thinking:** I should just transform all the available words. So if they ask for 10 but I only have 3, I uppercase those 3 words.

**Why this makes sense:** It's called "graceful degradation" - when you can't do exactly what was asked, you do what you can instead of crashing. Better to do something useful than to give up completely!

**Example:**
```
Input:  "only two words (up, 10)"
Output: "ONLY TWO WORDS"
```

I transform all 3 words even though they asked for 10. Seems like the safest approach.

---

### When Hex or Binary Conversions Get Bad Input

**The situation:** What if someone writes `"XYZ (hex)"` where XYZ isn't valid hex? (Has letters other than A-F)

**My thinking:** I should probably just leave it as-is and skip the conversion. Don't crash, don't give an error, just... ignore it and move on.

**Why:** Better to skip a weird conversion than to crash the whole program. The user might have made a typo, or maybe they meant something else.

**Example:**
```
Input:  "XYZ (hex) is weird"
Output: "XYZ (hex) is weird"  (unchanged)
```

Same logic for binary - if someone writes `"102 (bin)"` (which has a 2, not valid binary), I'll just leave it alone.

---

### When There Are Odd Numbers of Quotes

**The situation:** What if there are 3 quotes total in the file? Quotes come in pairs, so what do I do with the third one?

**My thinking:** I'm not 100% sure yet... Maybe treat the last unpaired quote as an opening quote? Or maybe just leave it as a regular character?

**Example:**
```
Input:  "He said ' hello ' and then ' goodbye"
Output: "He said 'hello' and then 'goodbye"  (???)
```

I need to test this edge case and see what makes sense. Probably the third quote just stays as an opening quote without a closing one.

---

### The 'h' Rule - Should I Follow It Exactly?

**The situation:** The spec says ALL words starting with 'h' should get "an". But in real English, "a hero" is correct (you pronounce the h) while "an honor" is correct (silent h).

**My decision:** I'm going to follow the spec strictly and make ALL 'h' words use "an", even if it sounds weird in English.

**Why:** Rules are rules! The spec says "vowels and h", so I'll do exactly that. Even if "an hero" sounds wrong to my ears, that's what the spec asks for.

**Example:**
```
Input:  "a honor to meet a hero"
Output: "an honor to meet an hero"
```

Yes, "an hero" sounds weird, but I'm following the spec exactly. If the spec wanted phonetic rules, it would have said so!

---

### What If There's an Article at the End?

**The situation:** What if the text ends with "a" and there's no word after it?

**My thinking:** Just leave it as "a". Can't change it to "an" if there's nothing following it!

**Example:**
```
Input:  "I saw a"
Output: "I saw a"  (unchanged)
```

Simple - no next word means no transformation.

---

### Commands Inside Quotes

**The situation:** What if there's a command inside quotes? Like `' hello (up) world '`

**My guess:** Commands should probably still work inside quotes. The quotes just affect spacing (removing spaces around the quotes), not whether commands execute.

**Example:**
```
Input:  "He said ' hello (up) world '"
Output: "He said 'hello WORLD'"
```

The `(up)` still works, and the quotes still get fixed. Both rules apply!

**Status:** Need to verify this with testing, but it seems logical.

---

### Nested or Malformed Command Markers

**The situation:** What if someone writes something weird like `hello (up (cap) test` or `((hex))`?

**My thinking:** I'm not sure how to handle this... Maybe just process the first valid marker I find and treat the rest as regular text?

**Example:**
```
Input:  "word ((hex))"
Output: "word ((hex))"  (treat as invalid, leave unchanged)
```

Or maybe I try to find the first valid pattern and ignore the broken ones? This could be tricky. I'll need to think about this more when I'm implementing the command parser.

---

### How Do I Detect Punctuation Groups?

**The situation:** Groups like `...` should stay together, not become `. . .` with spaces.

**My thinking:** When I see multiple punctuation marks in a row (no spaces between them), I should keep them together as one unit.

**Valid groups:** `...`, `!!`, `??`, `!?`, `?!`

**But what about:** If someone writes `. . .` with spaces? That's probably three separate periods, not a group.

**Example:**
```
Input:  "Wait ... really"
Output: "Wait... really"  (group stays together)

Input:  "Wait . . . really"
Output: "Wait. . . really"  (three separate periods? or still a group?)
```

I need to decide: do I only group punctuation that's already together, or do I also merge spaced punctuation? Probably only group what's already together - don't try to be too smart.

---

### Is There a Maximum Count for Commands?

**The situation:** The spec doesn't say a limit... Someone could write `(cap, 1000)` or even `(cap, 999999)`.

**My thinking:** Do I need to keep ALL previous words in memory? That could use a lot of memory for long texts. Or should I set a reasonable maximum (like 100) and document it as a limitation?

**Or maybe:** Just use graceful degradation - if they ask for 1000 words but there are only 50, transform all 50. No artificial limits, just work with what's available.

**Decision:** I think graceful degradation is the answer. No maximum limit - just transform however many words are actually there. If someone writes `(cap, 1000)` in a 50-word document, I capitalize all 50 words. Simple!

---

### Negative or Zero Counts

**The situation:** What if someone writes `(up, -1)` or `(up, 0)`?

**My thinking:** These don't make sense. You can't uppercase -1 words or 0 words.

**Decision:** Treat these as invalid commands and leave them unchanged.

**Example:**
```
Input:  "hello (up, -1) world"
Output: "hello (up, -1) world"  (invalid, unchanged)
```

Don't crash, don't error - just ignore the weird command and move on.

---

### What If Capitalization Hits Already-Uppercase Words?

**The situation:** If I run `(cap)` on "HELLO", what happens?

**My thinking:** Capitalize means "first letter uppercase, rest lowercase". So "HELLO" becomes "Hello".

**Example:**
```
Input:  "HELLO (cap)"
Output: "Hello"
```

This makes sense - capitalize is a specific format, not just "make first letter uppercase". It's: first letter UP, everything else DOWN.

---

**End of Project Analysis**
