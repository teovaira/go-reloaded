#!/bin/bash
# 🧪 go-reloaded test runner

INPUT="sample.txt"
OUTPUT="result.txt"
EXPECTED="result_expected.txt"

echo "🔹 Running go-reloaded on sample.txt..."
go run .. "$INPUT" "$OUTPUT"

echo "🔍 Comparing result.txt with expected output..."
if diff -w "$OUTPUT" "$EXPECTED" > /dev/null; then
  echo "✅ Test passed: output matches expected result."
else
  echo "❌ Test failed: differences found!"
  diff -y --suppress-common-lines "$OUTPUT" "$EXPECTED"
fi


# // Make it executable with chmod +x tests/run_tests.sh.