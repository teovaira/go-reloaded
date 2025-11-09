#!/bin/bash
# 🧪 go-reloaded test runner

echo "🧪 Running go-reloaded tests..."
echo ""

# Test 1: sample.txt
echo "🔹 Test 1: Running go-reloaded on sample.txt..."
go run .. "sample.txt" "result.txt"

echo "🔍 Comparing result.txt with expected output..."
if diff -w "result.txt" "result_expected.txt" > /dev/null; then
  echo "✅ Test 1 passed: output matches expected result."
else
  echo "❌ Test 1 failed: differences found!"
  diff -y --suppress-common-lines "result.txt" "result_expected.txt"
fi

echo ""

# Test 2: golden.txt
echo "🔹 Test 2: Running go-reloaded on golden.txt..."
go run .. "golden.txt" "golden_result.txt"

echo "🔍 Comparing golden_result.txt with expected output..."
if diff -w "golden_result.txt" "golden_expected.txt" > /dev/null; then
  echo "✅ Test 2 passed: output matches expected result."
else
  echo "❌ Test 2 failed: differences found!"
  diff -y --suppress-common-lines "golden_result.txt" "golden_expected.txt"
fi

echo ""
echo "🏁 Test run complete!"

# Make it executable with chmod +x tests/run_tests.sh.